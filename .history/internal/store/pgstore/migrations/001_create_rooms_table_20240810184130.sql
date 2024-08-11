CREATE TABLE IF NOT EXISTS rooms(
    "id"    uuid    PRIMARY KEY     NOT NULL gen_random_uuid(),
);

---- create above / drop below ----

DROP TABLE IF EXISTS rooms: