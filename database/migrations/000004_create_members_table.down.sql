ALTER TABLE "members" DROP CONSTRAINT "fk_user_id";

ALTER TABLE "members" DROP CONSTRAINT "fk_family_id";

DROP TABLE IF EXISTS "members";
