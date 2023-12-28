CREATE TABLE
    "accounts" (
        "id" bigserial PRIMARY KEY,
        "name" varchar NOT NULL,
        "balance" float NOT NULL DEFAULT 0,
        "user_id" bigserial NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
    );

ALTER TABLE "accounts"
ADD
    CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id");
