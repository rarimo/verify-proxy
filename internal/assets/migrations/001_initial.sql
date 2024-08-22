-- +migrate Up

CREATE TABLE verify_requests(
    id            BYTEA        PRIMARY KEY    NOT NULL,
    created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at    TIMESTAMP WITHOUT TIME ZONE,
    request_data BYTEA,
    callback_data BYTEA,
    status        VARCHAR(255)                NOT NULL
);

-- +migrate Down

DROP TABLE verify_requests;
