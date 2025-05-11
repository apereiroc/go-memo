package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/apereiroc/go-memo/debug"
	"github.com/apereiroc/go-memo/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		debug.Error(err)
		log.Panic(err)
	}

	dbFile := path.Join(homeDir, ".config", "memo", "commands.db")

	db, err := openDB(dbFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openDB(path string) (*sql.DB, error) {
	// check if path exists
	fileInfo, pathErr := os.Stat(path)
	if pathErr != nil && !os.IsNotExist(pathErr) {
		return nil, pathErr
	}

	if pathErr == nil && fileInfo.IsDir() {
		err := fmt.Errorf("%s is a directory", path)
		debug.Error(err)
		return nil, err
	}

	// open the DB
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		debug.Error(err)
		return nil, err
	}
	db.SetMaxOpenConns(1)

	// create the DB if it doesn't exist
	if os.IsNotExist(pathErr) {
		debug.Warnf("file %s did not exist; new DB file created", path)

		// create directory
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			mkdirErr := fmt.Errorf("failed to create directories %s: %w", dir, err)
			debug.Error(mkdirErr)
			return nil, mkdirErr
		}

		// touch the DB to trigger file creation, if needed
		if err := db.Ping(); err != nil {
			debug.Error(err)
			return nil, err
		}
	}

	return db, nil
}

func isEmpty(db *sql.DB) (bool, error) {
	if db == nil {
		return true, nil
	}

	// adapted from https://stackoverflow.com/questions/44098235/how-could-i-check-does-my-sqlite3-database-is-empty
	sqlStatement := `
	SELECT
		COUNT(name)
	FROM
		sqlite_master;
	`
	row := db.QueryRow(sqlStatement)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func LoadGroups(db *sql.DB) ([]models.Group, error) {
	// check if the database is empty
	empty, err := isEmpty(db)
	if err != nil {
		debug.Error(err)
		return nil, err
	}

	if empty {
		return []models.Group{}, nil
	}

	// i use left join to get also possible groups with no commands
	sqlStatement := `
	SELECT
		g.id, g.name,
		c.command, c.description
	FROM
		groups g
	LEFT JOIN
		commands c
	ON
		g.id = c.group_id
	ORDER BY
		g.id
	`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		debug.Error(err)
		return nil, err
	}
	defer rows.Close()

	groupsMap := make(map[int]*models.Group)
	for rows.Next() {
		var (
			groupID   int
			groupName string
			cmdText   sql.NullString
			descText  sql.NullString
		)

		err = rows.Scan(&groupID, &groupName, &cmdText, &descText)
		if err != nil {
			debug.Error(err)
			return nil, err
		}

		group, exists := groupsMap[groupID]
		if !exists {
			group = &models.Group{
				Name: groupName,
			}
			groupsMap[groupID] = group
		}

		if cmdText.Valid {
			cmd := models.Command{
				Cmd:         cmdText.String,
				Description: descText.String,
			}
			group.Cmds = append(group.Cmds, cmd)
		}
	}

	// flatten map into array
	groups := make([]models.Group, 0, len(groupsMap))
	for _, g := range groupsMap {
		groups = append(groups, *g)
	}

	return groups, nil
}
