-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  role TEXT NOT NULL,
  verified BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT NOW(),

  UNIQUE(user_id),
  UNIQUE(email)
);

CREATE TABLE profiles (
  user_id TEXT PRIMARY KEY,
  username TEXT NOT NULL,
  bio TEXT,
  picture TEXT,
  profile_link TEXT,
  following INTEGER,
  followers INTEGER,

  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO users (user_id, name, email, role, verified) VALUES ('hdiuebeufh82338d2', 'Habeeb', 'habeeb@gmail.com', 'user', true);
INSERT INTO profiles (user_id, username, bio, picture, profile_link, following, followers) VALUES ('hdiuebeufh82338d2', '@habeeb_amoo', 'Software Developer', 'b', '', 0, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABEL profiles;
DROP TABLE users;
-- +goose StatementEnd
