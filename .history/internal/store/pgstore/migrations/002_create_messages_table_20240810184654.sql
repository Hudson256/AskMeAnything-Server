CREATE TABLE IF NOT EXISTS messages (
    "id"        uuid    PRIMARY KEY     NOT NULL    DEFAULT gen_random_uuid(),
    "room_id"   uuid                    NOT NULL,
    messages
);

---- create above / drop below ----

DROP TABLE IF EXISTS messages;