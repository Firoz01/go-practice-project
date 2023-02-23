CREATE TABLE IF NOT EXISTS "employee" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "salary" INT,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);