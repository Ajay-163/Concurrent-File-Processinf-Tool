package main

import (
	"encoding/json"
	"fmt"
)

func PrintReport(report Report) {

	data, err := json.MarshalIndent(
		report,
		"",
		"  ",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
