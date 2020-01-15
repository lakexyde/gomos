package gomos

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lakexyde/gomos/types"
)

// SchemaProperties defines the schema properties
type SchemaProperties map[string]FieldOptions

// Schema is a Schema type for MongoDB database operations
type Schema struct {
	data       bson.M
	Properties SchemaProperties
	TimeStamps bool
	empty      bool
}

// New creates a new instance of the Schema
func (s Schema) New() *Schema {
	return &Schema{
		Properties: s.Properties,
		data:       bson.M{},
		TimeStamps: s.TimeStamps,
	}
}

// Add adds fields to the schema data
func (s *Schema) Add(val map[string]interface{}) {
	for k, v := range val {
		s.data[k] = v
	}

	s.buildData()
}

// Create populates the schema data with default values
func (s Schema) Create() *Schema {

	if len(s.data) == 0 {
		s = Schema{
			Properties: s.Properties,
			data:       bson.M{},
			TimeStamps: s.TimeStamps,
			empty:      true,
		}
	}

	s.buildData()
	return &s
}

// Data returns the Schema data
func (s *Schema) Data() (bson.M, error) {
	var err error
	if !s.empty {
		return s.data, err

	}

	s.data["_id"] = primitive.NewObjectID()

	// if not, range over the field options to look for required fields
	// and return an error
	for k, v := range s.Properties {
		if _, found := s.data[k]; !found {
			if v.Required {
				err = errors.New(k + " is missing or invalid")
				break
			}

			// set default values
			if v.Default != nil {
				s.data[k] = v.Default
			}
		}
	}

	return s.data, err
}

func (s *Schema) buildData() {
	for k, v := range s.data {
		// Delete keys that are not defined
		if _, found := s.Properties[k]; !found {
			delete(s.data, k)
			continue
		}

		opt := s.Properties[k]

		// Skip if types don't match
		if valid := opt.compareType(v); !valid {
			delete(s.data, k)
			continue
		}

		// Validate fields
		if valid := opt.validateFields(v); !valid {
			delete(s.data, k)
			continue
		}

		// Check if ObjectID string types are valid and convert to ObjectID
		if opt.fieldType == types.ObjectIDType && reflect.TypeOf(v).Kind() == reflect.String {
			t, err := primitive.ObjectIDFromHex(v.(string))
			if err != nil {
				delete(s.data, k)
				continue
			} else {
				s.data[k] = t
			}

		}

		// Trim value if string
		if opt.fieldType == types.StringType && opt.Trim {
			v = strings.Trim(v.(string), "")
		}

		// Convert to upper case
		if opt.fieldType == types.StringType && opt.UpperCase {
			v = strings.ToUpper(v.(string))
		}

		// Convert to lower case
		if opt.fieldType == types.StringType && opt.LowerCase {
			v = strings.ToLower(v.(string))
		}

		// updated the field
		s.data[k] = v
	}

	// set the updated time to now if Timestamps is set to true
	if s.TimeStamps {
		now := time.Now()
		s.data["updatedAt"] = now
		if s.empty {
			s.data["createdAt"] = now
		}
	}
}
