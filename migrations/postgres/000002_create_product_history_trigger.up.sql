-- insert product history after creating import
CREATE OR REPLACE FUNCTION  create_product_history_trigger_function() 
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO "product_history"
    (
        "product_id",
        "source_id",
        "quantity",
        "shop_id",
        "product_action_id",
        "external_id"
    )
    SELECT
        ii.product_id as product_id,
        i.id as source_id,
        ii.quantity as quantity,
        i.shop_id as shop_id,
        pa.id as product_action_id,
        i.external_id as external_id
    FROM "import" as i LEFT JOIN "import_item" as ii
    ON i.id = ii.import_id
    JOIN "product_action" as pa ON pa.name = 'import'
    WHERE i.id = NEW."id" AND i.status_id = '0b123839-be70-4372-866f-fce1c59f6ed1'
    ON CONFLICT ("product_id", "product_action_id", "source_id") DO NOTHING;
    RETURN NULL;
END;
$$;

CREATE OR REPLACE TRIGGER "product_history"
AFTER UPDATE OR INSERT ON "import"
  FOR EACH ROW
EXECUTE FUNCTION create_product_history_trigger_function();

-- insert product_movement after creating import
CREATE OR REPLACE FUNCTION  create_product_movement_trigger_function() 
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO "product_movement"
    (
        "product_id",
        "import_item_id",
        "quantity",
        "shop_id",
        "type",
        "status"
    )
    SELECT
        ii.product_id as product_id,
        ii.id as import_item_id,
        ii.quantity as quantity,
        i.shop_id as shop_id,
        pa.id as "type",
        pms.id as status
    FROM "import" as i LEFT JOIN "import_item" as ii
    ON i.id = ii.import_id
    JOIN "product_action" as pa ON pa.name = 'import'
    JOIN "product_movement_status" as pms ON pms.name = 'finished'
    WHERE i.id = NEW."id"  AND i.status_id = '0b123839-be70-4372-866f-fce1c59f6ed1';
    RETURN NULL;
END;
$$;

CREATE OR REPLACE TRIGGER "product_movement"
AFTER UPDATE OR INSERT ON "import"
  FOR EACH ROW
EXECUTE FUNCTION create_product_movement_trigger_function();

-- insert product history after creating order
CREATE OR REPLACE FUNCTION  create_product_history_trigger_function_for_order() 
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO "product_history"
    (
        "product_id",
        "source_id",
        "quantity",
        "shop_id",
        "product_action_id",
        "external_id"
    )
    SELECT
        oi.product_id as product_id,
        o.id as source_id,
        oi.value as quantity,
        o.shop_id as shop_id,
        pa.id as product_action_id,
        o.external_id as external_id
    FROM "order_item" as oi JOIN "order" as o
    ON  oi.order_id = o.id
    JOIN "product_action" as pa ON pa.name = 'order'
    WHERE oi.id = NEW."id"
    ON CONFLICT ("product_id", "product_action_id", "source_id") DO NOTHING;
    RETURN NULL;
END;
$$;

CREATE OR REPLACE TRIGGER "product_history_order"
AFTER INSERT ON "order_item"
  FOR EACH ROW
EXECUTE FUNCTION create_product_history_trigger_function_for_order();

-- insert product_queue after creating import
CREATE OR REPLACE FUNCTION  create_product_queue_for_import_trigger_function() 
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO "product_queue"
    (
        "product_id",
        "source_id",
        "remainder_count",
        "shop_id",
        "product_action_id"
    )
    SELECT
        ii.product_id as product_id,
        ii.id as source_id,
        ii.quantity as remainder_count,
        i.shop_id as shop_id,
        pa.id as "product_action_id"
    FROM "import" as i LEFT JOIN "import_item" as ii
    ON i.id = ii.import_id
    JOIN "product_action" as pa ON pa.name = 'import'
    WHERE i.id = NEW."id"  AND i.status_id = '0b123839-be70-4372-866f-fce1c59f6ed1';
    RETURN NULL;
END;
$$;

CREATE OR REPLACE TRIGGER "product_queue"
AFTER UPDATE OR INSERT ON "import"
  FOR EACH ROW
EXECUTE FUNCTION create_product_queue_for_import_trigger_function();
-- 
-- 
-- 
-- 
-- insert product history after creating order
CREATE OR REPLACE FUNCTION  create_product_movement_trigger_function_for_order() 
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
DECLARE 
   stock integer := 0;
   remainder_count integer := 0;
   remainder integer := 0;
   pqsource_id UUID ;
   product_action_id UUID ;
   product_id UUID ;
   sale_type VARCHAR := '';
    pm_product_id UUID;
    pm_order_item_id UUID;
    pm_shop_id UUID;
    pm_product_action_id UUID;
    pm_product_movement_id UUID;
BEGIN
    SELECT oi.value, p.sale_type FROM "order_item" as oi
    JOIN "product" as p ON p.id = oi.product_id
    INTO stock, sale_type WHERE oi.id = NEW."id";

    WHILE stock > 0  LOOP
        IF sale_type = 'FIFO' THEN
            SELECT pq."remainder_count", pq."source_id", pq."product_action_id" 
            FROM "product_queue" as pq
            JOIN "order_item" as oi ON oi.product_id =pq.product_id
            INTO remainder_count,  pqsource_id, product_action_id
            WHERE oi.id = NEW."id" AND pq.remainder_count > 0
            ORDER BY pq.created_at DESC limit 1;
        END IF;

        IF stock <= remainder_count THEN
            remainder = remainder_count - stock;
            stock = 0;
        ELSE 
            remainder = 0;
            stock = stock - remainder_count;
        END IF;

        IF product_action_id = '7069e210-7d2e-4a12-9160-3ef82f10ef4a' -- "import"
        THEN

            SELECT
                oi.product_id as product_id,
                oi.id as order_item_id,
                o.shop_id as shop_id,
                pa.id as product_action_id,
                pms.id as product_movement_id
            FROM "order_item" as oi 
            JOIN "order" as o
            ON  oi.order_id = o.id
            JOIN "product_action" as pa ON pa.name = 'order'
            JOIN "product_movement_status" as pms ON pms.name = 'finished'
            INTO    pm_product_id,
                    pm_order_item_id,
                    pm_shop_id,
                    pm_product_action_id,
                    pm_product_movement_id
            WHERE oi.id = NEW."id";

            INSERT INTO "product_movement"
            (
                "product_id",
                "order_item_id",
                "shop_id",
                "type",
                "status",
                "quantity",
                "import_item_id"
            )
            VALUES(pm_product_id,
            pm_order_item_id,
            pm_shop_id,
            pm_product_action_id,
            pm_product_movement_id,
            remainder_count - remainder,
            pqsource_id);
        END IF;

        -- IF product_action_id = '7069e210-7d2e-4a12-9160-3ef82f10ef4d' -- "supplier_order"
        -- THEN

        --     SELECT
        --         oi.product_id as product_id,
        --         oi.id as order_item_id,
        --         o.shop_id as shop_id,
        --         pa.id as product_action_id,
        --         pms.id as product_movement_id
        --     FROM "order_item" as oi 
        --     JOIN "order" as o
        --     ON  oi.order_id = o.id
        --     JOIN "product_action" as pa ON pa.name = 'order'
        --     JOIN "product_movement_status" as pms ON pms.name = 'finished'
        --     INTO    pm_product_id,
        --             pm_order_item_id,
        --             pm_shop_id,
        --             pm_product_action_id,
        --             pm_product_movement_id
        --     WHERE oi.id = NEW."id";

        --     INSERT INTO "product_movement"
        --     (
        --         "product_id",
        --         "order_item_id",
        --         "shop_id",
        --         "type",
        --         "status",
        --         "quantity",
        --         "supplier_order_item_id"
        --     )
        --     VALUES(pm_product_id,
        --     pm_order_item_id,
        --     pm_shop_id,
        --     pm_product_action_id,
        --     pm_product_movement_id,
        --     remainder_count - remainder,
        --     pqsource_id);
        -- END IF;

        UPDATE product_queue as pq SET "remainder_count" = remainder
        WHERE pq."product_id" = pm_product_id and pq."source_id" = pqsource_id ;
    END LOOP;
     
    RETURN NULL;
END;
$$;


CREATE OR REPLACE TRIGGER "product_movement_order_item"
AFTER INSERT ON "order_item"
  FOR EACH ROW
EXECUTE FUNCTION create_product_movement_trigger_function_for_order();

-- insert into "order"(id, external_id, total_price, product_sort_count, shop_id) 
-- values('351a3c19-d8fa-49f6-8ddc-6b71bf98cfd7','9999999990', 3000, 1, '6f5f6cf8-8caa-4384-bd5c-7961111599ca');

-- insert into "order_item"("id", "price","value", "total_price","order_id","product_id") 
-- values('351a3c19-d8fa-49f6-8ddc-6b71bf98cfe7',3000,1, 3000,'351a3c19-d8fa-49f6-8ddc-6b71bf98cfd7', 'c9ec7b54-5ce8-4fc1-aad2-7e3d3cd9c9d0');

SELECT pq."remainder_count", pq."source_id", pq."product_action_id" 
            FROM "product_queue" as pq
            JOIN "order_item" as oi ON oi.product_id =pq.product_id
            WHERE oi.id = 'f7841443-05e2-43e6-b8a7-60cfa7626f08' AND pq.remainder_count > 0
            ORDER BY pq.created_at DESC limit 1;