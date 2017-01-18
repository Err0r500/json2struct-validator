package validateJSON

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Human struct {
	Age    int    `json:"age" structs:"age"`
	Gender string `json:"gender" structs:"gender"`
}
type Person struct {
	Human `structs:",flatten"`
	Name  string `json:"name" structs:"name"`
}
type Friend struct {
	Human     `structs:",flatten"`
	FirstName string `json:"firstname" structs:"firstName"`
}

func TestLib(t *testing.T) {
	human := Human{Age: 30, Gender: "female"}
	humanJSON, _ := json.Marshal(human)

	person := Person{Human: Human{Age: 30, Gender: "female"}, Name: "my Name"}
	personJSON, _ := json.Marshal(person)

	friend := Friend{Human: Human{Age: 30, Gender: "female"}, FirstName: "Elodie"}
	friendJSON, _ := json.Marshal(friend)

	t.Run("A=1", func(t *testing.T) { human2Human(t, humanJSON, human) })
	t.Run("A=2", func(t *testing.T) { person2Person(t, personJSON, person) })
	t.Run("A=3", func(t *testing.T) { human2Person(t, humanJSON, person) })
	t.Run("A=4", func(t *testing.T) { friend2Person(t, friendJSON, person) })
}

func human2Human(t *testing.T, a []byte, b Human) {
	if err := Check(a, reflect.TypeOf(b)); err != nil {
		t.Error(err)
	}
}
func person2Person(t *testing.T, a []byte, b Person) {
	if err := Check(a, reflect.TypeOf(b)); err != nil {
		t.Error(err)
	}
}

func human2Person(t *testing.T, a []byte, b Person) {
	if err := Check(a, reflect.TypeOf(b)); err == nil {
		//should throw an error to pass the test
		t.Error(err)
	}
}
func friend2Person(t *testing.T, a []byte, b Person) {
	if err := Check(a, reflect.TypeOf(b)); err == nil {
		//should throw an error to pass the test
		t.Error(err)
	}
}
