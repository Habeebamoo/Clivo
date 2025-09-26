-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
  article_id TEXT NOT NULL,
  author_id TEXT NOT NULL,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  picture TEXT,
  read_time TEXT,
  created_at TIMESTAMP DEFAULT NOW(),

  UNIQUE (article_id),
  UNIQUE (author_id) REFERENCES users(user_id)
);

CREATE TABLE tags (
  article_id TEXT NOT NULL,
  tag TEXT NOT NULL,

  FOREIGN KEY (article_id) REFERENCES articles(article_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tags;
DROP TABLE articles;
-- +goose StatementEnd


