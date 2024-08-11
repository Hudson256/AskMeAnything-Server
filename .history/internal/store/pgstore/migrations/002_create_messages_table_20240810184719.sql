CREATE TABLE IF NOT EXISTS messages (
    "id"        uuid    PRIMARY KEY     NOT NULL    DEFAULT gen_random_uuid(),
    "room_id"   uuid                    NOT NULL,
    "message"   VARCHAR(255)
);

---- create above / drop below ----

DROP TABLE IF EXISTS messages;