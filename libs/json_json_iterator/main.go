package main

import (
	"fmt"
	"os"

	"gol/demo/json_json_iterator/json"
)

func main() {
	data, err := os.ReadFile("sectorInfoWaitSeed.json")
	fmt.Println(err)
	type Cs struct{}
	cs := &Cs{}
	json.Unmarshal(data, cs)
	fmt.Println("go")
}
