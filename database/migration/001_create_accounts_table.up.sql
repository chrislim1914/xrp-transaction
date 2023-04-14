CREATE TABLE IF NOT EXISTS accounts (
  "id" serial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "account_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "api_key" varchar NOT NULL,
  "api_secret" varchar NOT NULL, 
  "created_at" timestamp default now(),
  "updated_at" timestamp without time zone,
  "deleted_at" timestamp without time zone
);
CREATE INDEX uuid_index ON accounts(uuid);
CREATE INDEX email_index ON accounts(email);