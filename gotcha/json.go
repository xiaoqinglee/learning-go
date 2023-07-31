package gotcha

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp/v3"
)

type S struct {
	A int
}

func TestMarshal() {
	//{"A":42} <nil>
	//{"A":42} <nil>
	//{"A":42} <nil>

	src := &S{A: 42}

	bytes_, err := json.Marshal(*src)
	fmt.Println(string(bytes_), err)

	bytes_, err = json.Marshal(src) // 符合直觉
	fmt.Println(string(bytes_), err)

	bytes_, err = json.Marshal(&src)
	fmt.Println(string(bytes_), err)
}

func TestUnmarshal() {
	//<nil> &{42}
	//<nil> &{42}
	//json: Unmarshal(nil *gotcha.S) <nil>

	target1 := &S{} //符合直觉
	e := json.Unmarshal([]byte(`{"a": 42}`), target1)
	fmt.Println(e, target1)

	var target2 *S
	e = json.Unmarshal([]byte(`{"a": 42}`), &target2)
	fmt.Println(e, target2)

	var target3 *S
	e = json.Unmarshal([]byte(`{"a": 42}`), target3)
	fmt.Println(e, target3)
}

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

type GoodsItemAttrs struct {
	// OpenChest  default value is -1
	// https://stackoverflow.com/questions/39160807/default-value-golang-struct-using-encoding-json
	OpenChest int32 `json:"openChest"`
}

func TestJsonFieldDefaultValue() {
	for _, input := range []string{
		`{}`,
		`{"openChest":0}`,
		`{"openChest":1}`,
		`{"openChest":-1}`,
	} {
		inputAttrs := input
		fmt.Println("================================================")
		computedAttrs := &GoodsItemAttrs{OpenChest: -1}
		err := json.Unmarshal([]byte(inputAttrs), computedAttrs)
		if err != nil {
			panic(err)
		}
		bytes_, err := json.Marshal(computedAttrs)
		if err != nil {
			panic(err)
		}
		reMarshalled := string(bytes_)
		pp.Println("&GoodsItemAttrs{OpenChest: -1}", input, computedAttrs)
		pp.Println("&GoodsItemAttrs{OpenChest: -1}", input, reMarshalled)
		fmt.Println("------------------------------------------------")
		computedAttrs = &GoodsItemAttrs{}
		err = json.Unmarshal([]byte(inputAttrs), computedAttrs)
		if err != nil {
			panic(err)
		}
		bytes_, err = json.Marshal(computedAttrs)
		if err != nil {
			panic(err)
		}
		reMarshalled = string(bytes_)
		pp.Println("&GoodsItemAttrs{}", input, computedAttrs)
		pp.Println("&GoodsItemAttrs{}", input, reMarshalled)
	}
}

func TestUnmarshalMap() {
	//nil &gotcha.generalMap{
	//  Fields: map[string]interface {}{
	//    "a": "42",
	//  },
	//}
	//nil &gotcha.stringKV{
	//  Fields: map[string]string{
	//    "a": "42",
	//  },
	//}

	type generalMap struct {
		Fields map[string]any
	}
	type stringKV struct {
		Fields map[string]string
	}

	target1 := &generalMap{}
	e := json.Unmarshal([]byte(`{"fields": {"a": "42"}}`), target1)
	pp.Println(e, target1)

	target2 := &stringKV{}
	e = json.Unmarshal([]byte(`{"fields": {"a": "42"}}`), &target2)
	pp.Println(e, target2)

}

func TestNestFieldsMarshalUnMarshal() {

	//{"SmallValue":"s","BigValue":"b"} <nil>
	//<nil> &{{s} b}
	//"============================================="
	//{"Small":{"SmallValue":"s"},"BigValue":"b"} <nil>
	//<nil> &{{s} b}

	type Small struct {
		SmallValue string
	}
	type BigUsingNestFields struct {
		Small
		BigValue string
	}

	src := &BigUsingNestFields{Small{"s"}, "b"}

	bytes_, err := json.Marshal(src)
	fmt.Println(string(bytes_), err)

	dest := &BigUsingNestFields{}
	e := json.Unmarshal(bytes_, dest)
	fmt.Println(e, dest)

	pp.Println("=============================================")

	type BigNotUsingNestFields struct {
		Small    Small
		BigValue string
	}

	src2 := &BigNotUsingNestFields{Small{"s"}, "b"}

	bytes_2, err := json.Marshal(src2)
	fmt.Println(string(bytes_2), err)

	dest2 := &BigNotUsingNestFields{}
	e2 := json.Unmarshal(bytes_2, dest2)
	fmt.Println(e2, dest2)
}

func TestWhenFieldIsMissing() {
	//"TestWhenDestFieldIsMissing"
	//&{42} <nil>
	//"TestWhenSrcFieldIsMissing"
	//&{42 0} <nil>

	pp.Println("TestWhenDestFieldIsMissing")
	type Dest struct {
		A int
	}
	target := &Dest{}
	e := json.Unmarshal([]byte(`{"a": 42, "b": 42}`), target)
	fmt.Println(target, e)

	pp.Println("TestWhenSrcFieldIsMissing")
	type Dest2 struct {
		A int
		B int
	}
	target2 := &Dest2{}
	e2 := json.Unmarshal([]byte(`{"a": 42}`), target2)
	fmt.Println(target2, e2)
}
