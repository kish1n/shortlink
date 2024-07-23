-- +migrate Up

CREATE TABLE links (
     original TEXT NOT NULL,
     shortened TEXT PRIMARY KEY
);

-- +migrate Down

drop table links;
