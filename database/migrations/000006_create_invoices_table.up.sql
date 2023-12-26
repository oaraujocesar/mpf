CREATE TABLE "invoices" (
  "id" bigserial PRIMARY KEY,
  "amount" float NOT NULL DEFAULT 0,
  "account_id" bigserial NOT NULL,
  "close_at" timestamptz NOT NULL DEFAULT (now()),
  "card_id" bigserial NOT NULL,
  "due_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "invoices" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("card_id") REFERENCES "cards" ("id");
