CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "uuid" uuid UNIQUE NOT NULL,
  "token" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz NULL
);

CREATE TABLE "page_linked_lists" (
   "id" bigserial PRIMARY KEY,
   "uuid" uuid UNIQUE NOT NULL,
   "user_id" bigint NOT NULL,
   "type" smallint NOT NULL,
   "head" uuid NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT(now()),
   "updated_at" timestamptz NOT NULL DEFAULT(now()),
   "deleted_at" timestamptz NULL
);
CREATE TABLE "page_nodes" (
   "id" bigserial PRIMARY KEY,
   "uuid" uuid UNIQUE NOT NULL,
   "article_ids" int[] NOT NULL,
   "previous" uuid,
   "next" uuid,
   "created_at" timestamptz NOT NULL DEFAULT(now()),
   "updated_at" timestamptz NOT NULL DEFAULT(now()),
   "deleted_at" timestamptz NULL
);
CREATE TABLE "articles" (
    "id" bigserial PRIMARY KEY,
    "uuid" uuid UNIQUE NOT NULL,
    "title" varchar NOT NULL,
    "content" varchar NOT NULL,
    "sort" smallint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now()),
    "updated_at" timestamptz NOT NULL DEFAULT(now()),
    "deleted_at" timestamptz NULL
);

--index--
CREATE INDEX ON "users" ("uuid");
CREATE INDEX ON "users" ("token");
CREATE INDEX ON "users" ("email");
CREATE INDEX ON "page_linked_lists" ("user_id", "type");
CREATE INDEX ON "page_nodes" ("uuid");
CREATE INDEX ON "page_nodes" USING GIN ("article_ids");

--foreign keys--
ALTER TABLE "page_linked_lists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");