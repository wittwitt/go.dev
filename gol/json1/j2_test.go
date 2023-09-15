package json1

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestG(t *testing.T) {

	p1odec := map[string]interface{}{}

	p1odec["xxx"] = "dsafsdf"

	p1odec["ss"] = []byte("abcd")

	data, _ := json.Marshal(&p1odec)

	var p1odec2 map[string]json.RawMessage
	json.Unmarshal(data, &p1odec2)

	var cc []byte
	json.Unmarshal(p1odec2["ss"], &cc)

	fmt.Println(string(cc))


	
}
