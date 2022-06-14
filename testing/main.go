package main

import "fmt"


type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){
	fmt.Println("1", address)
	fmt.Println("2", *address)
	address.Country = "Indonesia"
	*address = Address{"Utah", "Texas", "USA"}
	address.Country = "malaysia"
}

func changeName(name *string){
	fmt.Println("3", name)
	fmt.Println("4", *name)
	*name = "wp"
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamatPointer *Address = &Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(alamatPointer)

	a := "Jagad"
	alamatB := &a
	changeName(alamatB)
	fmt.Println(*alamatPointer)
	fmt.Println(a)
}

// type Person struct {
// 	name      string
// 	grade     int
// 	breakfast []string
// }

// func main() {
// 	var data map[string]interface{}

// 	data = map[string]interface{}{
// 		"name":      "ethan hunt",
// 		"grade":     2,
// 		"breakfast": []string{"apple", "manggo", "banana"},
// 	}

// 	ShowData(data)
// 	// fmt.Println(reflect.TypeOf(data.grade))
// }

// func ShowData(data map[string]interface{}) {
// 	// fmt.Println(reflect.TypeOf(data["name"]))
// 	res := Person{
// 		name:  data["name"].(string),
// 		grade: data["grade"].(int),
// 		breakfast:  data["breakfast"].([]string),
// 	}
// 	fmt.Println(res)
// }
