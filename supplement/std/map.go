package std

import (
	"fmt"
	"sort"
)

//Sort a map by key or value

func SortMapKeys() {
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

//Get Map keys and values

//func GetMapKeyOrValues() {
//
//	keys := make([]keyType, 0, len(myMap))
//	values := make([]valueType, 0, len(myMap))
//
//	for k, v := range myMap {
//		keys = append(keys, k)
//		values = append(values, v)
//	}
//}
