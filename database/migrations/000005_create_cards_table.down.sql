ALTER TABLE "cards" DROP CONSTRAINT "fk_family_id";

ALTER TABLE "cards" DROP CONSTRAINT "fk_user_id";

DROP TABLE IF EXISTS "cards";
