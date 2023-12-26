CREATE TABLE "members" (
  "id" bigserial PRIMARY KEY,
  "family_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "members" ADD FOREIGN KEY ("family_id") REFERENCES "families" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
