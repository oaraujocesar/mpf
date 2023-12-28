ALTER TABLE "invoices" DROP CONSTRAINT "fk_card_id";

ALTER TABLE "invoices" DROP CONSTRAINT "fk_account_id";

DROP TABLE IF EXISTS "invoices";
