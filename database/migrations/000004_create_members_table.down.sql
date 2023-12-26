ALTER TABLE "members" REMOVE FOREIGN KEY ("user_id");

ALTER TABLE "members" REMOVE FOREIGN KEY ("family_id");

DROP TABLE IF EXISTS "members";
