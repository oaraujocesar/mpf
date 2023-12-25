CREATE TYPE "entry_type" AS ENUM (
  'INCOME',
  'EXPENSE'
);


CREATE TABLE "families" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "members" (
  "id" bigserial PRIMARY KEY,
  "family_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "balance" float NOT NULL DEFAULT 0,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "cards" (
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

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "entries" (
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

CREATE TABLE "installments" (
  "id" bigserial PRIMARY KEY,
  "amount" float NOT NULL,
  "entry_id" bigserial NOT NULL,
  "payday" timestamptz DEFAULT (now()),
  "paid_at" timestamptz DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "families" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("family_id") REFERENCES "families" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "cards" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "cards" ADD FOREIGN KEY ("family_id") REFERENCES "families" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("card_id") REFERENCES "cards" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("invoice_id") REFERENCES "invoices" ("id");

ALTER TABLE "installments" ADD FOREIGN KEY ("entry_id") REFERENCES "entries" ("id");
