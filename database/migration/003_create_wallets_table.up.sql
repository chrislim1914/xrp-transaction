CREATE TABLE IF NOT EXISTS wallets (
  "id" serial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "address" varchar NULL,
  "destination_tag" varchar NULL,
  "created_at" timestamp default now(),
  "updated_at" timestamp without time zone,
  "deleted_at" timestamp without time zone
);
CREATE INDEX uuid_wallets_index ON wallets(uuid);