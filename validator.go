package validateJSON

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/fatih/structs"
)

//ValidateJSON : check if a json object is compatible with a given struct
// usage :
//var someStruct myPackage.myStruct
//err := ValidateJSON(JSONstring.Bytes(), reflect.TypeOf(someStruct)
//returns nil if compatible
func Check(a []byte, expectedType reflect.Type) error {

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
