package db

import (
	"database/sql"
	"log"
	"os"
	"path"

	"github.com/apereiroc/go-memo/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbFile := path.Join(homeDir, ".config", "memo", "caca.db")

	db, err := openDB(dbFile)
	if err != nil {
		return nil, err
	}
	return db, err
}

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)

	return db, nil
}

func LoadGroups(db *sql.DB) ([]models.Group, error) {
	// no database, no groups
	// valid scenario
	if db == nil {
		return []models.Group{}, nil
	}

	// i use left join to also get possible groups with no commands
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
