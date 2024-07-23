-- +migrate Up

CREATE TABLE link (
     original TEXT NOT NULL,
     shortened TEXT PRIMARY KEY
);

-- +migrate Down

drop table link;
