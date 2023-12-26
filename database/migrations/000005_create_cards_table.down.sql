ALTER TABLE "cards" REMOVE FOREIGN KEY ("family_id");

ALTER TABLE "cards" REMOVE FOREIGN KEY ("user_id");

DROP TABLE IF EXISTS "cards";
