package constants

const (
	FAILED_GET  = "Failed fetching data"
	SUCCESS_GET = "Successfully fetching data"

	FAILED_INSERT  = "Failed creating new data"
	SUCCESS_INSERT = "Successfully creating new data"

	FAILED_UPDATE  = "Failed updating data"
	SUCCESS_UPDATE = "Successfully updating data"

	FAILED_DELETE    = "Failed deleting data"
	SUCCESS_DELETE   = "Successfully deleting data"
	NO_RECORD_DELETE = "No data found to delete"

	DUPLICATE_RECORD = "Data already exists"

	TRUE_VALUE  = true
	FALSE_VALUE = false

	EMPTY_VALUE     = ""
	EMPTY_VALUE_INT = 0

	YES_VALUE = "YES"
	NO_VALUE  = "NO"
)

var (
	// Empty Object Struct
	EMPTY_ARRAY_INTERFACE  = []interface{}{}
	EMPTY_SINGLE_INTERFACE = map[string]interface{}{}
)
