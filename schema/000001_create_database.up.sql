CREATE TABLE users
(
    "id"         int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "name"       text      NOT NULL,
    "last_name"  text      NOT NULL,
    "password"   text      NOT NULL UNIQUE,
    "email"      text      NOT NULL UNIQUE,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "enabled"    smallint  NOT NULL DEFAULT 0
);

-- UPDATES while user is updated --
CREATE OR REPLACE FUNCTION update_updated_at_on_update()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
end;
$$ language 'plpgsql';

CREATE TRIGGER update_user_update_at
    AFTER UPDATE
    ON
        users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_on_update();
-- END FUNCTION AND TRIGGER TO UPDATE --

CREATE TABLE verification_token
(
    "id"         int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "user_id"    int       NOT NULL,
    "token"      text      NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "expires_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '1 hour'),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE refresh_token
(
    "id"         int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "token"      text      NOT NULL,
    "expires_at" timestamp NOT NULL
)