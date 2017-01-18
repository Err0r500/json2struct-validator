# json2struct-validator

ValidateJSON : check if a json string is compatible with a given struct

usage :
```go
var JSONstring := "{field1:value}"
var someStruct myPackage.myStruct
err := Check(JSONstring.Bytes(), someStruct)

//or
err := Check(JSONstring.Bytes(), &someStruct)
```
returns nil if check OK
See the test file for an example
