package gomos

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lakexyde/gomos/types"
	"github.com/lakexyde/gomos/util"
)

// FieldOptions is a Schema field options
type FieldOptions struct {
	fieldType interface{}
	Required  bool
	Min       int32
	Max       int32
	Default   interface{}
	Trim      bool
	Email     bool
	UpperCase bool
	LowerCase bool
}

// compareType compares a value type with the set fieldType
func (opt FieldOptions) compareType(val interface{}) bool {
	if opt.fieldType == types.ObjectIDType && (reflect.TypeOf(val).Kind() == reflect.String || reflect.TypeOf(val) == reflect.TypeOf(primitive.ObjectID{})) {
		return true
	}

	if opt.fieldType == types.StringType && reflect.TypeOf(val).Kind() == reflect.String {
		return true
	}

	if opt.fieldType == types.NumberType && reflect.TypeOf(val).Kind() == reflect.Float64 {
		return true
	}

	if opt.fieldType == types.IntType && reflect.TypeOf(val).Kind() == reflect.Float64 {
		return true
	}

	if opt.fieldType == types.FloatType && reflect.TypeOf(val).Kind() == reflect.Float64 {
		return true
	}

	if opt.fieldType == types.ArrayType && reflect.TypeOf(val).Kind() == reflect.Array {
		return true
	}

	if opt.fieldType == types.StringArray && reflect.TypeOf(val) == reflect.TypeOf([]string{}) {
		return true
	}

	if opt.fieldType == types.NumberArray && reflect.TypeOf(val) == reflect.TypeOf([]float64{}) {
		return true
	}

	if opt.fieldType == types.MapType && reflect.TypeOf(val).Kind() == reflect.Map {
		return true
	}

	if opt.fieldType == types.BooleanType && reflect.TypeOf(val).Kind() == reflect.Bool {
		return true
	}

	if opt.fieldType == types.TimeType && reflect.TypeOf(val) == reflect.TypeOf(time.Time{}) {
		return true
	}

	return false
}

// validateFields validates a value based on field validation params
func (opt FieldOptions) validateFields(val interface{}) (valid bool) {
	valid = true

	// Validate email
	if opt.fieldType == types.StringType && opt.Email {
		valid = util.ValidateEmail(val.(string))
	}

	return
}
