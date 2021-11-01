package ch12

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func PrintUnboundMethods(x interface{}) {
	var v reflect.Value = reflect.ValueOf(x)
	var t reflect.Type = v.Type()
	fmt.Printf("type: %s\n", t)
	fmt.Printf("value: %s\n", v.String())
	fmt.Printf("method name: method signature\n")
	fmt.Printf("-----------------------------\n")
	for i := 0; i < t.NumMethod(); i++ {

		var methodDeclaration reflect.Method = t.Method(i)
		var methodName string = methodDeclaration.Name
		var methodType reflect.Type = methodDeclaration.Type

		fmt.Printf("%s: %s\n", methodName, methodType)
	}
	fmt.Println()
}

func PrintBoundMethods(x interface{}) {
	var v reflect.Value = reflect.ValueOf(x)
	var t reflect.Type = v.Type()
	fmt.Printf("type: %s\n", t)
	fmt.Printf("value: %s\n", v.String())
	fmt.Printf("method name: method signature: bound method signature\n")
	fmt.Printf("-----------------------------------------------------\n")
	for i := 0; i < t.NumMethod(); i++ {

		var methodDeclaration reflect.Method = t.Method(i)
		var methodName string = methodDeclaration.Name
		var methodTypeFromMethodDeclaration reflect.Type = methodDeclaration.Type

		var methodInstance reflect.Value = v.Method(i)
		var methodTypeFromMethodInstance reflect.Type = methodInstance.Type()

		fmt.Printf("%s: %s: %s\n", methodName, methodTypeFromMethodDeclaration, methodTypeFromMethodInstance)
	}
	fmt.Println()
}

type MyInt int

func (p *MyInt) Increment() {
	*p += 1
}
func (p *MyInt) String() string {
	return strconv.Itoa(int(*p))
}

//使用reflect.Value.Call调用方法值
func InvokeBoundMethods(x interface{}, whichMethod string, params ...interface{}) {

	var methodFound = false
	var errorMethodNotFound string = fmt.Sprintf("InvokeBoundMethods: method %q not found", whichMethod)
	if whichMethod == "" {
		panic(errorMethodNotFound)
	}

	var v reflect.Value = reflect.ValueOf(x)
	var t reflect.Type = v.Type()
	for i := 0; i < t.NumMethod(); i++ {

		var methodDeclaration reflect.Method = t.Method(i)
		var methodName string = methodDeclaration.Name

		if methodName == whichMethod {
			methodFound = true

			var paramValues []reflect.Value
			for _, param := range params {
				paramValues = append(paramValues, reflect.ValueOf(param))
			}

			var methodInstance reflect.Value = v.Method(i)
			methodInstance.Call(paramValues)
			break
		}
		if !methodFound {
			panic(errorMethodNotFound)
		}
	}
}

func ReflectMethods() {
	PrintUnboundMethods(time.Hour)             //const Hour Duration = 60 * Minute
	PrintUnboundMethods(new(strings.Replacer)) //pointer to a strings.Replacer struct instance
	fmt.Println("============================================")

	PrintBoundMethods(time.Hour)
	PrintBoundMethods(new(strings.Replacer))
	fmt.Println("============================================")

	p := new(MyInt) //已经实例化一个int变量了, 该变量为int的零值0
	PrintBoundMethods(p)
	fmt.Printf("init value: %v\n", *p)
	InvokeBoundMethods(p, "Increment")
	fmt.Printf("later value: %v\n", *p)
	fmt.Println()
}
