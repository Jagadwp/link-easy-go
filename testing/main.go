package main

import (
	"fmt"
)

type Person struct {
	name      string
	grade     int
	breakfast []string
}

func main() {
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	ShowData(data)
	// fmt.Println(reflect.TypeOf(data.grade))
}

func ShowData(data map[string]interface{}) {
	// fmt.Println(reflect.TypeOf(data["name"]))
	res := Person{
		name:  data["name"].(string),
		grade: data["grade"].(int),
		breakfast:  data["breakfast"].([]string),
	}
	fmt.Println(res)
}
