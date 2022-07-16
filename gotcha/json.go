package gotcha

//	Why does json.Marshal produce empty structs in the JSON text output?
//
//	type Person struct {
//		name string
//		age  int
//	}
//
//	p := Person{"Alice", 22}
//	jsonData, _ := json.Marshal(p)
//	fmt.Println(string(jsonData))
//
//	{}
//
//	Answer
//
//	Only exported fields of a Go struct will be present in the JSON output.
//
//	type Person struct {
//		Name string // Changed to capital N
//		Age  int    // Changed to capital A
//	}
//
//	p := Person{"Alice", 22}
//
//	jsonData, _ := json.Marshal(p)
//	fmt.Println(string(jsonData))
//
//	{"Name":"Alice","Age":22}
//
//	You can specify the JSON field name explicitly with a json: tag.
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	p := Person{"Alice", 22}
//
//	jsonData, _ := json.Marshal(p)
//	fmt.Println(string(jsonData))
//
//	{"name":"Alice","age":22}

//https://yourbasic.org/golang/gotcha-json-marshal-empty/
