package validateJSON

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

//ValidateJSON : check if a json []byte is compatible with a given struct
// usage :
//var someStruct myPackage.myStruct
//err := ValidateJSON(JSONstring.Bytes(), someStruct)
//or err := ValidateJSON(JSONstring.Bytes(), &someStruct)
//returns nil if compatible
func Check(inBytes []byte, inInterface interface{}) error {
	var inputMap, targetMap map[string]interface{}

	json.Unmarshal(inBytes, &inputMap)

	expectedType := reflect.TypeOf(inInterface)
	if expectedType.Kind() == reflect.Ptr {
		expectedType = expectedType.Elem()
	}
	s := reflect.New(expectedType).Elem().Interface()
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error:", err)
	}
	json.Unmarshal(b, &targetMap)

	if len(targetMap) != len(inputMap) {
		return errors.New("length is different")
	}

	for k, _ := range targetMap {
		if _, ok := inputMap[k]; !ok {
			return errors.New("field " + k + " not found in JSON")
		}
	}

	return nil
}
