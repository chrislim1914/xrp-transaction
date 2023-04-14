CREATE TABLE IF NOT EXISTS balances (
  "id" serial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "total" numeric default 0,
  "available" numeric default 0,
  "hold" numeric default 0,
  "created_at" timestamp default now(),
  "updated_at" timestamp without time zone,
  "deleted_at" timestamp without time zone
);
CREATE INDEX uuid_balances_index ON balances(uuid);