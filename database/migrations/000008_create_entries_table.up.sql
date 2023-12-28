CREATE TYPE "entry_type" AS ENUM ('INCOME', 'EXPENSE');

CREATE TABLE
    "entries" (
        "id" bigserial PRIMARY KEY,
        "title" varchar NOT NULL,
        "amount" float NOT NULL,
        "account_id" bigserial,
        "installments" integer,
        "type" entry_type NOT NULL,
        "category_id" bigserial NOT NULL,
        "invoice_id" bigserial,
        "payday" timestamptz DEFAULT (now()),
        "paid_at" timestamptz DEFAULT (now()),
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
    );

ALTER TABLE "entries"
ADD
    CONSTRAINT "fk_account_id" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "entries"
ADD
    CONSTRAINT "fk_category_id" FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "entries"
ADD
    CONSTRAINT "fk_invoice_id" FOREIGN KEY ("invoice_id") REFERENCES "invoices" ("id");
