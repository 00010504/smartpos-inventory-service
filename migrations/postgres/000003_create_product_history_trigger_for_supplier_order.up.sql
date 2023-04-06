-- insert product history after creating import

CREATE OR REPLACE FUNCTION CREATE_PRODUCT_HISTORY_SUPPLIER_ORDER_TRIGGER_FUNCTION
() RETURNS TRIGGER LANGUAGE PLPGSQL AS 
	$$ BEGIN
	INSERT INTO
	    "product_history" (
	        "product_id",
	        "source_id",
	        "quantity",
	        "shop_id",
	        "product_action_id",
	        "external_id"
	    )
	SELECT
	    soi.product_id as product_id,
	    so.id as source_id,
	    soi.expected_amount as quantity,
	    so.shop_id as shop_id,
	    pa.id as product_action_id,
	    so.external_id as external_id
	FROM "supplier_order" as so
	    LEFT JOIN "supplier_order_item" as soi ON so.id = soi.supplier_order_id
	    JOIN "product_action" as pa ON pa.name = 'suplier_order'
	WHERE
	    so.id = NEW."id"
	    AND so.status_id = 'caba1cc0-ba0d-4a03-8b78-c81c6730cca6' ON CONFLICT (
	        "product_id",
	        "product_action_id",
	        "source_id"
	    )
	DO NOTHING;
	RETURN NULL;
END; 

$$;

CREATE OR REPLACE TRIGGER "PRODUCT_HISTORY_FOR_SUPPLIER_ORDER" 
	AFTER
	UPDATE OR
	INSERT ON "supplier_order" FOR EACH ROW
	EXECUTE
	    FUNCTION create_product_history_supplier_order_trigger_function();

-- insert product_movement after creating import
CREATE OR REPLACE FUNCTION  create_product_movement_supplier_order_trigger_function() 
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO "product_movement"
    (
        "product_id",
        "supplier_order_item_id",
        "quantity",
        "shop_id",
        "type",
        "status"
    )
    SELECT
        soi.product_id as product_id,
        soi.id as supplier_order_item_id,
        soi.expected_amount as quantity,
        so.shop_id as shop_id,
        pa.id as "type",
        pms.id as status
    FROM "supplier_order" as so LEFT JOIN "supplier_order_item" as soi
    ON so.id = soi.supplier_order_id
    JOIN "product_action" as pa ON pa.name = 'suplier_order'
    JOIN "product_movement_status" as pms ON pms.name = 'finished'
    WHERE so.id = NEW."id"  AND so.status_id = 'caba1cc0-ba0d-4a03-8b78-c81c6730cca6';
    RETURN NULL;
END;
$$;

CREATE OR REPLACE TRIGGER "product_movement_supplier_order"
AFTER UPDATE OR INSERT ON "supplier_order"
  FOR EACH ROW
EXECUTE FUNCTION create_product_movement_supplier_order_trigger_function();


-- insert product_queue after creating import
CREATE OR REPLACE FUNCTION  create_product_queue_for_supplier_order_trigger_function() 
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
        soi.product_id as product_id,
        soi.id as source_id,
        soi.expected_amount as remainder_count,
        so.shop_id as shop_id,
        pa.id as "product_action_id"
    FROM "supplier_order" as so LEFT JOIN "supplier_order_item" as soi
    ON so.id = soi.supplier_order_id
    JOIN "product_action" as pa ON pa.name = 'suplier_order'
    WHERE so.id = NEW."id"  AND so.status_id = 'caba1cc0-ba0d-4a03-8b78-c81c6730cca6';
    RETURN NULL;
END;
$$;

CREATE OR REPLACE TRIGGER "product_queue_supplier_order"
AFTER UPDATE OR INSERT ON "supplier_order"
  FOR EACH ROW
EXECUTE FUNCTION create_product_queue_for_supplier_order_trigger_function();