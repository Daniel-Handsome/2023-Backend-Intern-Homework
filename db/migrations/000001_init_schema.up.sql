CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "token" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" varchar
);

CREATE TABLE "articles" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "title" varchar NOT NULL,
  "next_page_key" varchar
);

ALTER TABLE "articles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
