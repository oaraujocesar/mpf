CREATE TABLE
    "families" (
        "id" bigserial PRIMARY KEY,
        "name" varchar NOT NULL,
        "user_id" bigserial NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
    );

ALTER TABLE "families"
ADD
    CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id");
