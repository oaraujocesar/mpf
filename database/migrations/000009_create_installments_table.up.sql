CREATE TABLE
    "installments" (
        "id" bigserial PRIMARY KEY,
        "amount" float NOT NULL,
        "entry_id" bigserial NOT NULL,
        "payday" timestamptz DEFAULT (now()),
        "paid_at" timestamptz DEFAULT (now()),
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
    );

ALTER TABLE "installments"
ADD
    CONSTRAINT "fk_entry_id" FOREIGN KEY ("entry_id") REFERENCES "entries" ("id");
