CREATE TYPE user_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- system user
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin user
);

CREATE TYPE app_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- client
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin
);

CREATE TYPE "order_type" AS ENUM (
    '4c60818f-29be-45df-9860-c7f2fe94d91d', -- order
    '4d990284-425b-4602-ab65-ebd8436b3ebd' -- return
);

CREATE TYPE "product_type" AS ENUM (
  '8b0bf29c-58e8-4310-8bb1-a1b9771f9c47',
  '2b98f424-91c9-46cc-abd7-c888208807da',
  'a19a514e-41c9-4666-a01a-e3f9c0255609'
);

CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY,
    "user_type_id" user_type NOT NULL,
    "first_name" VARCHAR(250) NOT NULL,
    "last_name" VARCHAR(250) NOT NULL,
    "phone_number" VARCHAR(30) NOT NULL,
    "image" TEXT,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX "user_deleted_at_idx" ON "user"("deleted_at");

INSERT INTO "user" (
    "id",
    "first_name",
    "last_name",
    "phone_number",
    "user_type_id"
) VALUES (
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    'admin',
    'admin',
    '99894172774',
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'
);


CREATE TABLE IF NOT EXISTS "product_sale_type" (
    "name" VARCHAR(10) PRIMARY KEY, 
    "description" JSONB NOT NULL
);

INSERT INTO "product_sale_type" ("name", "description") VALUES 
    ('LIFO',  '{"en":"", "uz":"", "ru":""}'),
    ('LILO',  '{"en":"", "uz":"", "ru":""}'),
    ('FIFO',  '{"en":"", "uz":"", "ru":""}'),
    ('FILO',  '{"en":"", "uz":"", "ru":""}');



CREATE TABLE IF NOT EXISTS "company" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX company_deleted_at_idx ON "company"("deleted_at");

CREATE TABLE IF NOT EXISTS "company_user" (
    "user_id" UUID,
    "company_id" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY("user_id", "company_id", "deleted_at")
);

CREATE INDEX "company_user_deleted_at_idx" ON "company_user"("deleted_at");


CREATE TABLE IF NOT EXISTS "shop" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "company_id" UUID NOT NULL,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX shop_deleted_at_idx ON "shop"("deleted_at");

CREATE TABLE IF NOT EXISTS "measurement_unit" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "is_deletable" BOOLEAN NOT  NULL DEFAULT TRUE,
    "short_name" VARCHAR NOT NULL,
    "long_name" VARCHAR NOT NULL,
    "precision" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID,
    UNIQUE ("short_name", "company_id", "deleted_at")
);
CREATE INDEX measurement_unit_deleted_at_idx ON "measurement_unit"("deleted_at");



CREATE TABLE "product" (
  "id" UUID PRIMARY KEY,
  "sale_type" VARCHAR(10) NOT NULL DEFAULT 'FIFO',
  "is_marking" BOOLEAN NOT NULL DEFAULT FALSE,
  "sku" VARCHAR NOT NULL,
  "measurement_unit_id" UUID,
  "mxik_code" VARCHAR,
  "name" varchar NOT NULL DEFAULT '',
  "company_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "updated_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  "type" product_type NOT NULL DEFAULT '8b0bf29c-58e8-4310-8bb1-a1b9771f9c47',
  UNIQUE ("sku", "company_id", "deleted_at")
);
CREATE INDEX "product_deleted_at_idx" ON "product"("deleted_at");

INSERT INTO "product" (
    "id","name", "sku", "company_id", "created_by")
VALUES (
    'af1cdaf8-3b7f-4a36-957b-cfbd55fcb48e',
    'apple',
    '123456789', 
    '262a4f1a-ecd0-488c-96b2-f90900fdb0a4', 
    '4f133155-f782-40e7-956b-c275281868ca');

CREATE TABLE IF NOT EXISTS "product_barcode" (
    "barcode" VARCHAR(300) NOT NULL,
    "product_id" UUID,
    PRIMARY KEY ("barcode", "product_id")
);

CREATE TABLE IF NOT EXISTS "shop_price" (
    "supply_price" NUMERIC NOT NULL DEFAULT 0,
    "min_price" NUMERIC NOT NULL DEFAULT 0,
    "max_price" NUMERIC NOT NULL DEFAULT 0,
    "retail_price" NUMERIC NOT NULL DEFAULT 0,
    "whole_sale_price" NUMERIC NOT NULL DEFAULT 0,
    "shop_id" UUID,
    "product_id" UUID,
    PRIMARY KEY("product_id", "shop_id")
);

CREATE TABLE IF NOT EXISTS "measurement_values" (
    "shop_id" UUID,
    "is_aviable" BOOLEAN NOT NULL DEFAULT TRUE,
    "in_stock" NUMERIC NOT NULL DEFAULT 0,
    "product_id" UUID,
    PRIMARY KEY("product_id", "shop_id")
);

CREATE TABLE "product_retail_price" (
    "product_id" UUID NOT NULL,
    "retail_price" NUMERIC NOT NULL,
    "wholesale_price" NUMERIC NOT NULL,
    "min_price" NUMERIC NOT NULL,
    "max_price" NUMERIC NOT NULL,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX "product_retail_price_deleted_at_idx" ON "product_retail_price"("deleted_at");

-- import status
CREATE TABLE "import_status" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE("name")
);

INSERT INTO "import_status" ("id","name","translation") VALUES 
('0b123839-be70-4372-866f-fce1c59f6ed0','NEW','{"en":"", "uz":"", "ru":""}'),
('0b123839-be70-4372-866f-fce1c59f6ed1','FINISHED','{"en":"", "uz":"", "ru":""}');


-- import status insert

CREATE TABLE IF NOT EXISTS "import" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status_id" UUID NOT NULL REFERENCES "import_status"("id"),
    "name" VARCHAR NOT NULL,
    "total_price" NUMERIC NOT NULL DEFAULT 0,
    "shop_id" UUID,
    "supplier" VARCHAR NOT NULL DEFAULT '',
    "quantity" NUMERIC NOT NULL DEFAULT 0,
    "completed_by" UUID,
    "completed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "type" VARCHAR NOT NULL DEFAULT '',
    "company_id" UUID,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS "import_deleted_at_idx" ON  "import"("deleted_at");



CREATE TABLE IF NOT EXISTS "import_item"(
    "id" UUID PRIMARY KEY,
    "product_id" UUID,
    "import_id" UUID,
    "wholesale_price" NUMERIC,
    "max_price" NUMERIC,
    "supply_price" NUMERIC NOT NULL DEFAULT 0,
    "sale_price" NUMERIC NOT NULL DEFAULT 0,
    "is_validated" boolean,
    "name" varchar,
    "line" varchar,
    "article" varchar,
    "barcode" varchar,
    "sku" varchar NOT NULL DEFAULT '',
    "created_at" timestamp,
    "quantity" NUMERIC NOT NULL DEFAULT 0,
    "sold_count" NUMERIC NOT NULL DEFAULT 0,
    measurement_unit_id UUID
);


CREATE TABLE IF NOT EXISTS "transfer_status" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE("name")
);

CREATE TABLE IF NOT EXISTS "transfer" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status_id" UUID NOT NULL REFERENCES "transfer_status"("id"),
    "name" VARCHAR NOT NULL,
    "departure_shop_id" UUID,
    "arrival_shop_id" UUID,
    "company_id" UUID,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS "transfer_deleted_at_idx" ON  "transfer"("deleted_at");

CREATE TABLE IF NOT EXISTS "supplier_order_status" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE("name")
);

CREATE TABLE IF NOT EXISTS "supplier_settlement" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE("name")
);

CREATE TABLE IF NOT EXISTS "transfer_item"(
    "id" UUID PRIMARY KEY,
    "product_id" UUID,
    "transfer_id" UUID NOT  NULL REFERENCES "transfer"("id") ON DELETE CASCADE,
    "expected_amount" NUMERIC NOT NULL,
    "arrived_amount" NUMERIC NOT NULL
);

INSERT INTO "supplier_settlement" (
    "id",
    "name",
    "translation"
) VALUES
(
    '2082d491-5eee-469c-903f-b6d50f0a794c',
    'Commision',
    '{"en":"Commision", "uz":"Komissiya", "ru":"Комиссия"}'
);

INSERT INTO "supplier_order_status" (
    "id",
    "name",
    "translation"
) VALUES 
(
    'b4e0db85-3a80-414f-abe2-421c6334b9d0',
    'Pending',
    '{"en":"Pending", "uz":"Kutilmoqda", "ru":"В ожидании"}'

),
(
    'f963a201-b8dd-4f52-b4a4-070041a0ed42',
    'New',
    '{"en":"New", "uz":"Yangi", "ru":"Новый"}'

),
(
    '973f4626-d0a3-4693-95a6-eb65e87b3a76',
    'Ordered',
    '{"en":"Ordered", "uz":"Buyurtma qilindi", "ru":"Заказан"}'
),
(
    '042781b6-9ed1-4cec-94ec-4ed4009117a3',
    'Canceled',
    '{"en":"Canceled", "uz":"Bekor qilingan", "ru":"Отменено"}'
),
(
    'caba1cc0-ba0d-4a03-8b78-c81c6730cca6',
    'Recieved',
    '{"en":"Recieved", "uz":"Qabul qilingan", "ru":"Получено"}'
);

CREATE TABLE IF NOT EXISTS "supplier" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "phone_number" VARCHAR NOT NULL,
    "company_id" UUID NOT NULL,
    "supplier_company_name" VARCHAR NOT NULL,
    "external_id" VARCHAR NOT NULL,
    "agreement_number" VARCHAR,
    "legal_address" VARCHAR,
    "bank_account" VARCHAR,
    "bank_name" VARCHAR,
    "tin" VARCHAR,
    "ibt" VARCHAR,
    "supplier_company_legal_name" VARCHAR,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    UNIQUE("company_id", "external_id", "deleted_at"),
    UNIQUE("company_id", "phone_number")
);
CREATE INDEX IF NOT EXISTS "supplier_deleted_at_idx" ON  "supplier"("deleted_at");

CREATE TABLE IF NOT EXISTS "supplier_file" (
    "id" UUID PRIMARY KEY,
    "supplier_id" UUID NOT NULL REFERENCES "supplier"("id") ON DELETE CASCADE,
    "file_name" TEXT NOT NULL,
    "name" VARCHAR NOT NULL,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    UNIQUE("supplier_id", "file_name", "deleted_at")
);
CREATE INDEX IF NOT EXISTS "supplier_file_deleted_at_idx" ON  "supplier_file"("deleted_at");
CREATE TABLE IF NOT EXISTS "supplier_order" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status_id" UUID NOT NULL REFERENCES "supplier_order_status"("id") DEFAULT 'f963a201-b8dd-4f52-b4a4-070041a0ed42',
    "supplier_id" UUID NOT NULL REFERENCES "supplier"("id") ON DELETE CASCADE,
    "invoice_number" VARCHAR,
    "shop_id" UUID NOT NULL,
    "company_id" UUID NOT NULL,
    "note" TEXT,
    "total_sold_amount" NUMERIC NOT NULL DEFAULT 0,
    "expected_date" TIMESTAMP NOT NULL,
    "created_by" UUID NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    UNIQUE("company_id", "external_id", "deleted_at")
);
CREATE INDEX IF NOT EXISTS "supplier_order_deleted_at_idx" ON  "supplier_order"("deleted_at");
CREATE TABLE IF NOT EXISTS "supplier_order_item"(
    "id" UUID PRIMARY KEY,
    "supplier_order_id" UUID NOT NULL REFERENCES "supplier_order"("id"),
    "product_id" UUID,
    "discount" NUMERIC,
    "expected_amount" NUMERIC NOT NULL,
    "arrived_amount" NUMERIC NOT NULL DEFAULT 0,
    "cost" NUMERIC NOT NULL DEFAULT 0,
    "sold_amount" NUMERIC NOT NULL DEFAULT 0,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    UNIQUE("supplier_order_id", "product_id", "deleted_at")
);

CREATE INDEX IF NOT EXISTS "supplier_order_item_deleted_at_idx" ON  "supplier_order_item"("deleted_at");

CREATE TABLE IF NOT EXISTS "write_off_status" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE("name")
);

CREATE TABLE IF NOT EXISTS "write_off" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status_id" UUID NOT NULL REFERENCES "write_off_status"("id"),
    "name" VARCHAR NOT NULL,
    "shop_id" UUID,
    "company_id" UUID,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS "write_off_deleted_at_idx" ON  "write_off"("deleted_at");

CREATE TABLE IF NOT EXISTS "write_off_item" (
    "id" UUID PRIMARY KEY,
    "product_id" UUID,
    "write_off_id" UUID NOT  NULL REFERENCES "write_off"("id") ON DELETE CASCADE,
    "expected_amount" NUMERIC NOT NULL,
    "arrived_amount" NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS "repricing" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status_id" UUID NOT NULL REFERENCES "transfer_status"("id"),
    "name" VARCHAR NOT NULL,
    "shop_id" UUID,
    "company_id" UUID,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS "repricing_deleted_at_idx" ON  "repricing"("deleted_at");

CREATE TABLE IF NOT EXISTS "repricing_item" (
    "id" UUID PRIMARY KEY,
    "product_id" UUID,
    "repricing_id" UUID NOT  NULL REFERENCES "repricing"("id") ON DELETE CASCADE,
    "old_price_id" NUMERIC NOT NULL,
    "arrived_amount" NUMERIC NOT NULL
);


CREATE TABLE IF NOT EXISTS "cashbox" (
  "id" UUID PRIMARY KEY,
  "company_id" UUID NOT NULL,
  "shop_id" UUID,
  "title" VARCHAR(200) NOT NULL,
  "created_by" UUID,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("company_id", "title", "deleted_at")
);
CREATE INDEX IF NOT EXISTS "cashbox_deleted_at_idx" ON "cashbox"("deleted_at");


CREATE TABLE IF NOT EXISTS "order_status" (
  "id" UUID PRIMARY KEY,
  "name" varchar NOT NULL,
  "translation" JSONB NOT NULL,
  UNIQUE ("name")
);

INSERT INTO "order_status"(id, name, translation)
VALUES
('7069e210-7d2e-4a12-9160-3ef82f18ef4d', 'draft', '{"uz": "Qoralama", "ru": "Черновик", "en": "Draft"}'),
('d3bde6a2-532c-4f08-811f-0385e804c885', 'payed', '{"uz": "T''olangan", "ru": "Оплаченный", "en": "Payed"}'),
('15c0d291-45e5-4077-89d7-9b365e65cfed', 'postpone', '{"uz": "Kechiktirish", "ru": "Oткладывать", "en": "Postpone"}');


CREATE TABLE IF NOT EXISTS "order" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status" UUID NOT NULL DEFAULT '7069e210-7d2e-4a12-9160-3ef82f18ef4d' REFERENCES "order_status"("id"),
    "total_price" NUMERIC DEFAULT 0 NOT NULL,
    "total_discount_price" NUMERIC NOT NULL DEFAULT 0,
    "product_sort_count" INT NOT NULL DEFAULT 0,
    "shop_id" UUID,
    "cashbox_id" UUID,
    "company_id" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" UUID ,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID 
);
CREATE INDEX IF NOT EXISTS "order_deleted_at_idx" ON  "order"("deleted_at");

CREATE TABLE IF NOT EXISTS "order_item" (
    "id" UUID PRIMARY KEY,
    "price" NUMERIC NOT NULL,
    "value" NUMERIC NOT NULL,
    "total_price" NUMERIC NOT NULL,
    "order_id" UUID NOT NULL REFERENCES "order"("id") ON DELETE CASCADE,
    "product_id" UUID NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID,
    UNIQUE ("order_id", "product_id", "deleted_at")
);
CREATE INDEX IF NOT EXISTS "order_item_deleted_at_idx" ON  "order_item"("deleted_at");


CREATE OR REPLACE FUNCTION create_supplier_external_id()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
  external_id INT := 0;

BEGIN
    SELECT
      COUNT(*) AS total
    FROM
      "supplier"
    WHERE
      "company_id" = NEW."company_id"
    INTO
      external_id;

    external_id := external_id + 1;

    NEW."external_id"=RIGHT(CONCAT('000000', external_id), 6);

    RETURN NEW;
END;
$$;

-- triggers

CREATE OR REPLACE TRIGGER create_supplier_external_id
    BEFORE INSERT ON "supplier"
    FOR EACH ROW
    EXECUTE PROCEDURE create_supplier_external_id();

-- functions

CREATE OR REPLACE FUNCTION create_supplier_order_external_id()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
  external_id INT := 0;

BEGIN
    SELECT
      COUNT(*) AS total
    FROM
      "supplier_order"
    WHERE
      "company_id" = NEW."company_id"
    INTO
      external_id;

    external_id := external_id + 1;

    NEW."external_id"=RIGHT(CONCAT('000000', external_id), 6);

    RETURN NEW;
END;
$$;

-- triggers

CREATE OR REPLACE TRIGGER create_supplier_order_external_id
    BEFORE INSERT ON "supplier_order"
    FOR EACH ROW
    EXECUTE PROCEDURE create_supplier_order_external_id();

-- PRODUCT HISTORY Trigger vs Function
CREATE TABLE "product_action" (
  "id" uuid PRIMARY KEY,
  "name" VARCHAR NOT NULL
);

INSERT INTO "product_action"("id", "name")
VALUES
('7069e210-7d2e-4a12-9160-3ef82f10ef4a', 'import'),
('7069e210-7d2e-4a12-9160-3ef82f10ef4b', 'transfer'),
('7069e210-7d2e-4a12-9160-3ef82f10ef4c', 'order'),
('7069e210-7d2e-4a12-9160-3ef82f10ef4d', 'suplier_order'),
('7069e210-7d2e-4a12-9160-3ef82f10ef4e', 'writeoff');


CREATE TABLE IF NOT EXISTS "product_history" (
    "product_id" uuid NOT NULL REFERENCES "product" ("id"),
    "quantity" int,
    "shop_id" uuid,
    "product_action_id" uuid REFERENCES "product_action"("id"),
    "source_id" uuid NOT NULL, -- import_id -- transfer_id -- supplier_order_id
    "external_id" VARCHAR NOT NULL DEFAULT '',
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "created_date" timestamp DEFAULT CURRENT_DATE,
    UNIQUE ("product_id", "product_action_id", "source_id")
);
CREATE INDEX IF NOT EXISTS "product_history_created_date" ON "product_history"("created_date");
CREATE INDEX IF NOT EXISTS "product_history_product_id" ON "product_history"("product_id");

CREATE TABLE "product_movement_status" (
  "id" uuid PRIMARY KEY,
  "name" VARCHAR NOT NULL
);

INSERT INTO "product_movement_status"("id", "name")
VALUES
('1069e210-7d2e-4a12-9160-3ef82f10ef0a', 'finished'),
('1069e210-7d2e-4a12-9160-3ef82f10ef0b', 'confirmed'),
('1069e210-7d2e-4a12-9160-3ef82f10ef0c', 'transfer_in'),
('1069e210-7d2e-4a12-9160-3ef82f10ef0d', 'transfer_out');

CREATE TABLE IF NOT EXISTS "product_movement" (
  "product_id" uuid NOT NULL REFERENCES "product" ("id"),
  "quantity" int,
  "shop_id" uuid,
  "type" uuid REFERENCES "product_action"("id"),
  "status" uuid REFERENCES "product_movement_status"("id"),
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "created_date" timestamp DEFAULT CURRENT_DATE,
  "order_item_id" uuid REFERENCES "order_item" ("id"),
  "import_item_id" uuid REFERENCES "import_item" ("id"),
  "transfer_item_id" uuid REFERENCES "transfer_item" ("id"),
  "supplier_order_item_id" uuid REFERENCES "supplier_order_item" ("id")
);
CREATE INDEX IF NOT EXISTS "product_movement_created_date" ON "product_movement"("created_date");
CREATE INDEX IF NOT EXISTS "product_movement_product_id" ON "product_movement"("product_id");

CREATE TABLE IF NOT EXISTS "product_queue" (
    "product_id" uuid NOT NULL REFERENCES "product" ("id"),
    "remainder_count" int,
    "shop_id" uuid,
    "source_id" uuid NOT NULL,-- import_item_id -- transfer_item_id -- supplier_order_id
    "product_action_id" uuid REFERENCES "product_action"("id"),
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

-- functions
CREATE OR REPLACE FUNCTION order_total_amount()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN

  UPDATE
		"supplier_order"
	SET
		total_sold_amount = subquery.sold_amount
	FROM (
    SELECT COALESCE(SUM(sold_amount),0) AS sold_amount
    FROM "supplier_order_item" soi
    WHERE soi.supplier_order_id = NEW."supplier_order_id" AND soi."deleted_at" = 0
  ) AS subquery
	WHERE
		id = NEW."supplier_order_id" AND deleted_at = 0;

  RETURN NEW;
END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER order_total_amount
    AFTER INSERT OR UPDATE ON "supplier_order_item"
    FOR EACH ROW
    EXECUTE PROCEDURE order_total_amount();