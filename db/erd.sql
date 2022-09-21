-- postgreSQL

DROP TABLE player_missions;
DROP TABLE players;
DROP TABLE missions;

ALTER SEQUENCE missions_id_seq RESTART WITH 1;

CREATE OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$ 
BEGIN 
  NEW.updated_at = now();
  RETURN NEW;
END; $$ language 'plpgsql';

CREATE TABLE players (
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT NOT NULL CHECK (char_length(name) <= 256),
	email TEXT NOT NULL CHECK (char_length(email) <= 256),
	password TEXT NOT NULL CHECK (char_length(password) <= 256),
	gold_amount NUMERIC(20,6) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by TEXT NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by TEXT
);

CREATE TRIGGER players BEFORE
UPDATE
	ON players FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

 

CREATE TABLE missions (
	id SERIAL PRIMARY KEY NOT NULL,
	title TEXT NOT NULL CHECK (char_length(title) <= 256),
	description TEXT NOT NULL CHECK (char_length(description) <= 256),
	gold_bounty NUMERIC(20,6) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by TEXT NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by TEXT
);

CREATE TRIGGER missions BEFORE
UPDATE
	ON missions FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE player_missions (
	id SERIAL PRIMARY KEY NOT NULL,
	player_id INTEGER NOT NULL REFERENCES players(id),
	mission_id INTEGER NOT NULL REFERENCES missions(id),
	status TEXT NOT NULL CHECK (char_length(status) <= 256),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by TEXT NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by TEXT
);

CREATE TRIGGER player_missions BEFORE
UPDATE
	ON player_missions FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
 
