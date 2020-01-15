package types

// SchemaType is the Schema Type
type SchemaType int

const (
	// StringType defines a string type
	StringType SchemaType = iota
	// NumberType defines a number
	NumberType
	// IntType defines an interger type
	IntType
	// FloatType defines a float type
	FloatType
	// ObjectIDType defines and ObjectId Type
	ObjectIDType
	// ArrayType defines an Array type
	ArrayType
	// StringArray represents an array of string
	StringArray
	// NumberArray represents an array of number
	NumberArray
	// MapType defines a map type
	MapType
	// BooleanType defines a bool type
	BooleanType
	// TimeType define the time type
	TimeType
)
