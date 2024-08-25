CREATE TABLE "likes" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "post_id" bigint,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamp
);

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");