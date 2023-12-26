ALTER TABLE "accounts" REMOVE FOREIGN KEY ("user_id");

DROP TABLE IF EXISTS "accounts";
