package validateJSON

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/fatih/structs"
)

//ValidateJSON : check if a json []byte is compatible with a given struct
// usage :
//var someStruct myPackage.myStruct
//err := ValidateJSON(JSONstring.Bytes(), someStruct)
//or err := ValidateJSON(JSONstring.Bytes(), &someStruct)
//returns nil if compatible
func Check(a []byte, input interface{}) error {
	expectedType := reflect.TypeOf(input)

	if expectedType.Kind() == reflect.Ptr {
		expectedType = expectedType.Elem()
	}
	// log.Print(expectedType.Name())

	var jsonMap map[string]interface{}
	json.Unmarshal(a, &jsonMap)

	s := structs.New(reflect.New(expectedType).Elem().Interface())
	structMap := s.Map()

	if len(structMap) != len(jsonMap) {
		return errors.New("length is different")
	}

	for k, _ := range structMap {
		if _, ok := jsonMap[k]; !ok {
			return errors.New("field " + k + " not found in JSON")
		}
	}

	return nil
}
