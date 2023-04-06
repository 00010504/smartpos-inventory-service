package config

import "github.com/golang/protobuf/jsonpb"

const (
	DateTimeFormat  = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
	ConsumerGroupID = "invan_inventory_service"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	// product type
	SimpleProductTypeID  = "8b0bf29c-58e8-4310-8bb1-a1b9771f9c47"
	ServiceProductTypeID = "2b98f424-91c9-46cc-abd7-c888208807da"
	SetProductTypeID     = "a19a514e-41c9-4666-a01a-e3f9c0255609"

	// custom field type
	BooleanCFTypeID = "8b0bf29c-58e8-4310-8bb1-a1b9771f9c47"
	StringCFTypeID  = "2b98f424-91c9-46cc-abd7-c888208807da"
	NumberCFTypeID  = "a19a514e-41c9-4666-a01a-e3f9c0255609"

	OrderedStatus = "973f4626-d0a3-4693-95a6-eb65e87b3a76"
	ReceivedStatus = "caba1cc0-ba0d-4a03-8b78-c81c6730cca6"

	ElasticProductIndex    = "products"
	ElasticImportItemIndex = "ImportItems_"
	ElasticOrderItemIndex  = "supplier_order_"
)

var (
	JSONPBMarshaler = jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
)
