package supplement

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"math"
	"strings"
	"time"
)

func CastBetweenIntAndFloat() {
	//math.MaxInt64     9223372036854775807
	//math.MaxInt32     2147483647
	//time.Now().Unix() 1678440318
	//below math.MaxInt32 is ok to cast between float64 and int
	//not equal at:     16777217

	fmt.Println("math.MaxInt64", math.MaxInt64)
	fmt.Println("math.MaxInt32", math.MaxInt32)
	fmt.Println("time.Now().Unix()", time.Now().Unix())
	var f float64 = 1.0
	var i int64 = 1
	for {
		if i != int64(f) {
			fmt.Println("not equal at:", i)
			break
		}
		if i >= math.MaxInt32 {
			fmt.Println("below math.MaxInt32 is ok to cast between float64 and int")
			break
		}
		f += 1.0
		i += 1
	}

	var f32 float32 = 1.0
	i = 1
	for {
		if i != int64(f32) {
			fmt.Println("not equal at:", i)
			break
		}
		if i >= math.MaxInt32 {
			fmt.Println("below math.MaxInt32 is ok to cast between float32 and int")
			break
		}
		f32 += 1.0
		i += 1
	}
}

type FooStruct struct {
	Big int64 `json:"big"`
}

type FooStruct2 struct {
	Big int64 `json:"big,string"` //并不是所有的类型都支持 string 这个 tag
}

func JsonMarshal() {
	//{"big":42}
	//{"big":9223372036854775807}
	//{"big":"42"}
	//{"big":"9223372036854775807"}
	bytes_, err := json.Marshal(FooStruct{42})
	if err != nil {
		return
	}
	fmt.Println(string(bytes_))
	bytes_, err = json.Marshal(FooStruct{int64(math.MaxInt64)})
	if err != nil {
		return
	}
	fmt.Println(string(bytes_))
	bytes_, err = json.Marshal(FooStruct2{42})
	if err != nil {
		return
	}
	fmt.Println(string(bytes_))
	bytes_, err = json.Marshal(FooStruct2{int64(math.MaxInt64)})
	if err != nil {
		return
	}
	fmt.Println(string(bytes_))
}

func JsonUnmarshalWithStructSchema() { //无损
	//&{42}
	//&{9223372036854775807}
	//&{0}
	//&{0}
	//===============
	//&{0}
	//&{0}
	//&{42}
	//&{9223372036854775807}
	strings_ := []string{
		`{"big":42}`,
		`{"big":9223372036854775807}`,
		`{"big":"42"}`,
		`{"big":"9223372036854775807"}`,
	}
	for _, str := range strings_ {
		target := &FooStruct{}
		_ = json.Unmarshal([]byte(str), target)
		fmt.Println(target)
	}
	fmt.Println("===============")
	for _, str := range strings_ {
		target2 := &FooStruct2{}
		_ = json.Unmarshal([]byte(str), target2)
		fmt.Println(target2)
	}
}

func JsonUnmarshalWithoutStructSchema() { //有损
	//map[string]interface {}{
	//  "big": 42.000000,
	//}
	//0 false
	//42 true
	//===============
	//map[string]interface {}{
	//  "big": 9223372036854775808.000000,
	//}
	//0 false
	//9.223372036854776e+18 true
	//===============
	//map[string]interface {}{
	//  "big": "42",
	//}
	//0 false
	//0 false
	//===============
	//map[string]interface {}{
	//  "big": "9223372036854775807",
	//}
	//0 false
	//0 false
	//===============
	strings_ := []string{
		`{"big":42}`,
		`{"big":9223372036854775807}`,
		`{"big":"42"}`,
		`{"big":"9223372036854775807"}`,
	}
	for _, str := range strings_ {
		map_ := make(map[string]interface{})
		_ = json.Unmarshal([]byte(str), &map_)
		pp.Println(map_)
		int64Value, ok := map_["big"].(int64)
		fmt.Println(int64Value, ok)
		float64Value, ok := map_["big"].(float64)
		fmt.Println(float64Value, ok)
		fmt.Println("===============")
	}
}

func JsonUnmarshalWithoutStructSchema2() { //无损
	//map[string]interface {}{
	//  "big": "42",
	//}
	//42 nil
	//===============
	//map[string]interface {}{
	//  "big": "9223372036854775807",
	//}
	//9223372036854775807 nil
	//===============
	//map[string]interface {}{
	//  "big": "42",
	//}
	//panic: interface conversion: interface {} is string, not json.Number

	strings_ := []string{
		`{"big":42}`,
		`{"big":9223372036854775807}`,
		`{"big":"42"}`,
		`{"big":"9223372036854775807"}`,
	}
	for _, str := range strings_ {
		map_ := make(map[string]interface{})
		decoder := json.NewDecoder(strings.NewReader(str))

		// json.UseNumber causes the Decoder to unmarshal a number into an interface{} as a
		// json.Number instead of as a float64.
		decoder.UseNumber()

		decoder.Decode(&map_)

		pp.Println(map_)
		jsonNum := map_["big"].(json.Number)
		pp.Println(jsonNum.Int64())
		fmt.Println("===============")
	}
}
