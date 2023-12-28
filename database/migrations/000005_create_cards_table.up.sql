CREATE TABLE
    "cards" (
        "user_id" bigserial NOT NULL,
        "id" bigserial PRIMARY KEY,
        "name" varchar NOT NULL,
        "limit" float NOT NULL,
        "due_date" integer NOT NULL,
        "family_id" bigserial,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
    );

ALTER TABLE "cards"
ADD
    CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "cards"
ADD
    CONSTRAINT "fk_family_id" FOREIGN KEY ("family_id") REFERENCES "families" ("id");
