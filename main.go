package main

import (
	"bytes"
	"encoding/json"
	"github.com/k0kubun/pp/v3"
	"net/http"
	"reflect"
)

type Student struct {
	Fname                string `json:"fname"`
	Lname                string `json:"lname"`
	City                 string `json:"city"`
	Mobile               int64  `json:"mobile"`
	privateFieldForDebug int64  `json:"privateFieldForDebug"`
}

func EndpointHandler(w http.ResponseWriter, r *http.Request) {
	//initPlayer (only playerId)

	//ExtractStudentFromPlayer()

	//CommitPlayerDiff()
}

func ExtractStudentFromPlayer(player any) *Student {
	//mget from redis, optionally populate redis

	//mark subset available

	return nil
}

type Field struct {
	Name       string
	DataType   reflect.Type
	MarshalTag string

	RV                   *reflect.Value
	MarshalValueOriginal []byte
}

func Load(fromBlob []byte, fromMap map[string]json.RawMessage, handle any) (trackingInfo []*Field) {
	rv := reflect.ValueOf(handle).Elem()
	rt := rv.Type()
	fields := reflect.VisibleFields(rt)
	for _, field := range fields {
		if !field.IsExported() {
			continue
		}
		elem := &Field{
			Name:       field.Name,
			DataType:   field.Type,
			MarshalTag: field.Tag.Get("json"),

			RV: &rv,

			MarshalValueOriginal: nil,
		}
		trackingInfo = append(trackingInfo, elem)
	}

	if len(fromBlob) > 0 {
		err := json.Unmarshal(fromBlob, handle)
		if err != nil {
			panic(err)
		}
		for _, elem := range trackingInfo {
			typedValue := elem.RV.FieldByName(elem.Name).Addr().Elem().Interface()
			fieldBlob, err := json.Marshal(typedValue)
			if err != nil {
				panic(err)
			}
			elem.MarshalValueOriginal = fieldBlob
		}
	} else if len(fromMap) > 0 {
		for _, elem := range trackingInfo {
			if fieldBlob, ok := fromMap[elem.MarshalTag]; ok {
				tempTypedValue := reflect.New(elem.DataType)
				err := json.Unmarshal(fieldBlob, tempTypedValue.Interface())
				if err != nil {
					panic(err)
				}
				elem.RV.FieldByName(elem.Name).Set(tempTypedValue.Elem())
				elem.MarshalValueOriginal = fieldBlob
			} else {
				panic("invalid map: 字段缺失")
			}
		}
	} else {
		panic("invalid input")
	}

	return trackingInfo
}

func operateFoo(handle *Student) {
	handle.Mobile = 999999
}

func Dump(trackingInfo []*Field) (diffMap, completeMap map[string]json.RawMessage) {
	diffMap = make(map[string]json.RawMessage)
	completeMap = make(map[string]json.RawMessage)
	for _, elem := range trackingInfo {
		typedValue := elem.RV.FieldByName(elem.Name).Addr().Elem().Interface()
		fieldBlob, err := json.Marshal(typedValue)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(elem.MarshalValueOriginal, fieldBlob) {
			diffMap[elem.MarshalTag] = fieldBlob
		}
		completeMap[elem.MarshalTag] = fieldBlob
	}
	return diffMap, completeMap
}

func main() {
	tracked := &Student{}
	fromBlob := []byte(`{"fname":"f","lname":"l","city":"ny","mobile":77777}`)
	//fromMap := map[string]json.RawMessage{
	//	"city":   []byte(`"f"`),
	//	"fname":  []byte(`"l"`),
	//	"lname":  []byte(`"ny"`),
	//	"mobile": []byte(`7777`),
	//}
	traced := Load(fromBlob, nil, tracked)
	pp.Println(tracked)
	elem := traced[len(traced)-1]
	pp.Println("mobile:", elem.RV.FieldByName(elem.Name).Addr().Elem().Interface())
	tracked.Mobile = 999999
	pp.Println(tracked)
	pp.Println("mobile:", elem.RV.FieldByName(elem.Name).Addr().Elem().Interface())
	diff, complete := Dump(traced)
	pp.Println(string(diff["mobile"]))
	pp.Println(diff)
	pp.Println(complete)
}
