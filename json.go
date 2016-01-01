package main

import (
	"encoding/json"
	"fmt"

	log "github.com/cihub/seelog"
)

func testJson() {
	jsonStr, err := json.Marshal(student{Name: "joe1", Age: 20, human: human{Sex: 1}})
	if err != nil {
		log.Criticalf("marshal error")
	}
	fmt.Println(string(jsonStr))

	var std student
	if json.Unmarshal(jsonStr, &std) != nil {
		log.Criticalf("unmarshal error")
	}
	fmt.Println(std.Name)
}
