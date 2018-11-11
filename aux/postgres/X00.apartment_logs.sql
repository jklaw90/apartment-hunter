BEGIN;
CREATE TABLE "public"."apartment_events" ( 
	"id" UUid NOT NULL PRIMARY KEY,
	"aggregate_id" UUid NOT NULL,
	"version" Integer NOT NULL,
	"timestamp" Timestamp Without Time Zone NOT NULL,
	"event_type" Character Varying( 254 ) COLLATE "pg_catalog"."default" NOT NULL,
	"event" jsonb NOT NULL,
	CONSTRAINT "order_aggregate_id_version" UNIQUE( "aggregate_id", "version" ) );
 ;
CREATE INDEX "index_aggregate_id" ON "public"."apartment_events" USING btree( "aggregate_id" Asc NULLS Last );
COMMIT;

--BEGIN;
--CREATE TABLE "public"."apartments" ( 
--	"id" UUid NOT NULL PRIMARY KEY,
--	"version" Integer NOT NULL,
--	"data" jsonb NOT NULL,
-- ;
--COMMIT;
