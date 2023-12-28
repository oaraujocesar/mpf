ALTER TABLE "entries" DROP CONSTRAINT "fk_invoice_id";

ALTER TABLE "entries" DROP CONSTRAINT "fk_category_id";

ALTER TABLE "entries" DROP CONSTRAINT "fk_account_id";

DROP TABLE IF EXISTS "entries";

DROP TYPE IF EXISTS "entry_type";
