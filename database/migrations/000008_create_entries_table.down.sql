ALTER TABLE "entries" REMOVE FOREIGN KEY ("invoice_id");

ALTER TABLE "entries" REMOVE FOREIGN KEY ("category_id");

ALTER TABLE "entries" REMOVE FOREIGN KEY ("account_id");

DROP TABLE IF EXISTS "entries";

DROP TYPE IF EXISTS "entry_type";
