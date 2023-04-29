package main

import (
	"fmt"
	"os"

	"github.com/wittwitt/go.dev/libs/json_json_iterator/json"
)

func main() {
	data, err := os.ReadFile("sectorInfoWaitSeed.json")
	fmt.Println(err)
	type Cs struct{}
	cs := &Cs{}
	json.Unmarshal(data, cs)
	fmt.Println("go")
}
