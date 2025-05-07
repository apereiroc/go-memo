-- create groups table
CREATE TABLE IF NOT EXISTS groups (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE NOT NULL
);

-- create commands table
CREATE TABLE IF NOT EXISTS commands (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  group_id INTEGER NOT NULL,
  command TEXT UNIQUE NOT NULL,
  description TEXT UNIQUE NOT NULL,
  FOREIGN KEY(group_id) REFERENCES groups(id)
);

-- insert some groups
INSERT INTO groups (name) VALUES
  ('C++'),
  ('Git'),
  ('Go'),
  ('Docker');

-- insert some commands 
INSERT INTO commands (group_id, command, description) VALUES
  (1, "g++ --help", "Display the compiler's help message"),
  (1, "g++ -o main main.cpp -I<include-path> -L<lib-path>", "Compile the source program, looking for possible headers in `include-path` and possible libraries in `lib-path`"),
  (2, "git add <file>", "Add file contents to the index"),
  (2, "git commit -m <message>", "Record changes to the respository"),
  (3, "find . -type f -name '*.go' -exec sed -i '' 's/OLDNAME/NEWNAME/g' {} +", "Rename a go module from 'OLDNAME' to 'NEWNAME'"),
  (4, "docker buildx build -f Dockerfile -t build .", "Build image from Dockerfile with tag build. Dockerfile is at ."),
  (4, "docker buildx build --platform=linux/amd64 -f Dockerfile -t dotfiles-dev . && docker run --platform=linux/amd64 --rm -it dotfiles-dev", "Build image for testing my dotfiles. Hard to remember ...");
