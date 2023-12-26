ALTER TABLE "invoices" REMOVE FOREIGN KEY ("card_id");

ALTER TABLE "invoices" REMOVE FOREIGN KEY ("account_id");

DROP TABLE IF EXISTS "invoices";
