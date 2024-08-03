CREATE TABLE "follows" (
  "following_user_id" bigint,
  "followed_user_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "username" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "user_id" bigint,
  "status" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint,
  "text" varchar NOT NULL,
  "user_id" bigint,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "comments" ("post_id");

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
