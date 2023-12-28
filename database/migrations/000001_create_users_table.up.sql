CREATE TABLE
    "users" (
        "id" bigserial PRIMARY KEY,
        "name" varchar NOT NULL,
        "email" varchar NOT NULL,
        "password" varchar NOT NULL,
        "avatar" varchar,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
    );

CREATE INDEX "idx_users_email" ON "users" ("email");
