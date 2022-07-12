package ch7

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func UsingByteCounter() {
	bc := new(ByteCounter)
	fmt.Printf("%d\n", *bc)
	bc.Write([]byte("Hello"))
	fmt.Printf("%d\n", *bc)
	*bc = 0
	fmt.Printf("%d\n", *bc)
	fmt.Fprintf(bc, "Hello%s", "World")
	fmt.Printf("%d\n", *bc)
}

func InterfaceAssignment() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second
	fmt.Printf("w: %v\n", w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = new(bytes.Buffer)

	w = rwc
	//rwc = w

	var any interface{}
	any = true
	any = 1.0
	any = 42
	any = "hello"
	any = map[string]int{"one": 1}
	fmt.Printf("any: %v\n", any)
	fmt.Println()

	//接口的动态类型和动态值 %T可打印动态类型 %#v打印动态值, 两者都为nil接口变量才等于nil
	var writer1 io.Writer
	fmt.Printf("writer1: %T %#[1]v\n", writer1)
	fmt.Printf("writer1 == nil: %t\n", writer1 == nil)
	writer1 = (*bytes.Buffer)(nil)
	fmt.Printf("writer1: %T %#[1]v\n", writer1)
	fmt.Printf("writer1 == nil: %t\n", writer1 == nil)
	writer1 = new(bytes.Buffer)
	fmt.Printf("writer1: %T %#[1]v\n", writer1)
	fmt.Printf("writer1 == nil: %t\n", writer1 == nil)
	writer1 = os.Stdout
	fmt.Printf("writer1: %T %#[1]v\n", writer1)
	fmt.Printf("writer1 == nil: %t\n", writer1 == nil)

	//error是个接口
	var err error = fmt.Errorf("new error %s %s", "hello", "world")
	fmt.Printf(`error content is "%s"`, err.Error())

}

func TypeAssertion() {
	//类型断言
	var w io.Writer
	w = os.Stdout
	var ok bool

	//1.类型断言一个返回值

	//(1)断言这个变量就是这个具体类型
	f := w.(*os.File)
	fmt.Printf("%#v\n", f)
	//b := w.(*bytes.Buffer) //panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	//fmt.Printf("%s\n", b)

	//(2)断言成一个方法更多的接口类型
	rw := w.(io.ReadWriter)
	fmt.Printf("%#v\n", rw)
	//w = new(ByteCounter)
	//rw = w.(io.ReadWriter) //panic: interface conversion: *ch7.ByteCounter is not io.ReadWriter: missing method Read
	//fmt.Printf("%s\n", rw)
	fmt.Println()

	//2.类型断言两个返回值

	//(1)断言这个变量就是这个具体类型
	f, ok = w.(*os.File)
	fmt.Printf("%#v %t\n", f, ok)
	b, ok := w.(*bytes.Buffer)
	fmt.Printf("%#v %t\n", b, ok)

	//(2)断言成一个方法更多的接口类型
	rw, ok = w.(io.ReadWriter)
	fmt.Printf("%#v %t\n", rw, ok)
	w = new(ByteCounter)
	rw, ok = w.(io.ReadWriter)
	fmt.Printf("%#v %t\n", rw, ok)

	////使用惯例
	//if f, ok := w.(*os.File); ok{
	//	//use f
	//}
	//if w, ok := w.(*os.File); ok{ //内部w屏蔽外部w
	//	//use w
	//}
	//短变量声明语句":="与隐式词法块相关陷阱见 <<Go程序设计语言>> P23 P23 P36 P161
	//简单来讲
	//1.短变量声明语句不需要声明所有":="左边的变量, 如果一个变量在同一个词法块中已经声明, 对于这个变量, 短变量声明相当于赋值
	//2.除了花括号{}创建的显式词法块, if switch for 可以创建隐式词法块.
}

func sqlQuoteString(str string) string {
	//省略
	return ""
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil: //单一类型分支中x是case所对应的类型的变量
		return "NULL"
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x)
	case int, uint:
		return fmt.Sprintf("%d", x) //非单一类型分支x是switch操作数,这里是interface{}
	default:
		panic(fmt.Sprintf("unexpected type %T: %#[1]v", x))
	}
}

func TypeSwitch() {
	fmt.Printf("%#v\n", sqlQuote(nil))
	fmt.Printf("%#v\n", sqlQuote("hello"))
	fmt.Printf("%#v\n", sqlQuote(15))
	fmt.Printf("%#v\n", sqlQuote(false))
	//fmt.Printf("%#v", sqlQuote(15.6)) //panic: unexpected type float64: 15.6
}
