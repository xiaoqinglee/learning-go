package ch5

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp/v3"
	"strings"
)

func core() error {
	return fmt.Errorf("CoreError: %s", "Detail: xxx") //返回 *errors.errorString
}

func midLayer(e error) error {
	if e != nil {
		return fmt.Errorf("MidLayerError: %w", e) //返回 *fmt.wrapError
	}
	return nil
}

func topLayer(e error) error {
	if e != nil {
		return fmt.Errorf("TopLayerError: %w", e)
	}
	return nil
}

func TestGo13Errors() {
	e1 := core()
	e2 := midLayer(e1)
	e3 := topLayer(e2)

	fmt.Println("spew.Dump --------------------")
	spew.Dump(e1)
	spew.Dump(e2)
	spew.Dump(e3)

	fmt.Println("测试 != nil --------------------")
	//eInstance != nil 判断是否发生了错误
	fmt.Println(e1 != nil) //true

	fmt.Println("测试 == --------------------")
	//eInstance1 == eInstance2 判断是否是同一个error实例
	e4 := e1
	fmt.Println(e4 == e1)     //true
	fmt.Println(core() == e1) //false

	fmt.Println("测试 Unwrap --------------------")
	//解包装
	fmt.Println(errors.Unwrap(e3) == e2)  //true
	fmt.Println(errors.Unwrap(e2) == e1)  //true
	fmt.Println(errors.Unwrap(e1) == nil) //true

	fmt.Println("测试 Is --------------------")
	// func Is(err, target error) bool  判断err实例是否是target实例wrap 0次到多次的结果 (注意参数targe应该是comparable的)
	fmt.Println(errors.Is(e1, e1))     //true
	fmt.Println(errors.Is(e2, e1))     //true
	fmt.Println(errors.Is(e3, e1))     //true
	fmt.Println(errors.Is(e3, e2))     //true
	fmt.Println(errors.Is(e3, core())) //false
}
func coreS(detail string) error {
	return fmt.Errorf("coreS: %s", detail)
}

func midLayerS(e1, e2 error) error {
	return fmt.Errorf("midLayerS: %w and %w", e1, e2) // 返回 *fmt.wrapErrors, 结构体中的 errs []error 字段会过滤掉 nil error
}

func topLayerS(e1, e2 error) error {
	return errors.Join(e1, e2) // 返回 *errors.joinError, 结构体中的 errs []error 字段会过滤掉 nil error
}

//Wrapping multiple errors
//
//Go 1.20 expands support for error wrapping to permit an error to wrap multiple other errors.
//
//An error e can wrap more than one error by providing an Unwrap method that returns a []error.
//
//The errors.Is and errors.As functions have been updated to inspect multiply wrapped errors.
//
//The fmt.Errorf function now supports multiple occurrences of the %w format verb, which will cause it to return an error that wraps all of those error operands.
//
//The new function errors.Join returns an error wrapping a list of errors.

func TestGo20Errors() {
	coreE1 := coreS("coreE1")
	coreE2 := coreS("coreE2")
	midLayerSE := midLayerS(coreE1, coreE2)
	coreE3 := core()
	midLayerE := midLayer(coreE3)
	topSE := topLayerS(midLayerSE, midLayerE)

	midLayerSContainNilE := midLayerS(coreE1, nil)
	topContainNilE := topLayerS(midLayerSE, nil)

	fmt.Println("pp.Println --------------------")
	//"midLayerSE:" &fmt.wrapErrors{
	//  msg:  "midLayerS: coreS: coreE1 and coreS: coreE2",
	//  errs: []error{
	//    &errors.errorString{
	//      s: "coreS: coreE1",
	//    },
	//    &errors.errorString{
	//      s: "coreS: coreE2",
	//    },
	//  },
	//}
	//"midLayerE:" &fmt.wrapError{
	//  msg: "MidLayerError: CoreError: Detail: xxx",
	//  err: &errors.errorString{
	//    s: "CoreError: Detail: xxx",
	//  },
	//}
	//"topSE:" &errors.joinError{
	//  errs: []error{
	//    &fmt.wrapErrors{
	//      msg:  "midLayerS: coreS: coreE1 and coreS: coreE2",
	//      errs: []error{
	//        &errors.errorString{
	//          s: "coreS: coreE1",
	//        },
	//        &errors.errorString{
	//          s: "coreS: coreE2",
	//        },
	//      },
	//    },
	//    &fmt.wrapError{
	//      msg: "MidLayerError: CoreError: Detail: xxx",
	//      err: &errors.errorString{
	//        s: "CoreError: Detail: xxx",
	//      },
	//    },
	//  },
	//}
	//"midLayerSContainNilE:" &fmt.wrapErrors{
	//  msg:  "midLayerS: coreS: coreE1 and %!w(<nil>)",
	//  errs: []error{
	//    &errors.errorString{
	//      s: "coreS: coreE1",
	//    },
	//  },
	//}
	//"topContainNilE:" &errors.joinError{
	//  errs: []error{
	//    &fmt.wrapErrors{
	//      msg:  "midLayerS: coreS: coreE1 and coreS: coreE2",
	//      errs: []error{
	//        &errors.errorString{
	//          s: "coreS: coreE1",
	//        },
	//        &errors.errorString{
	//          s: "coreS: coreE2",
	//        },
	//      },
	//    },
	//  },
	//}

	pp.Println("midLayerSE:", midLayerSE)
	pp.Println("midLayerE:", midLayerE)
	pp.Println("topSE:", topSE)

	pp.Println("midLayerSContainNilE:", midLayerSContainNilE)
	pp.Println("topContainNilE:", topContainNilE)

	fmt.Println("测试 Unwrap --------------------")
	//解包装
	fmt.Println(errors.Unwrap(midLayerE) == coreE3) //true
	fmt.Println(errors.Unwrap(midLayerSE) == nil)   //true
	fmt.Println(errors.Unwrap(topSE) == nil)        //true

	fmt.Println(errors.Unwrap(midLayerSContainNilE) == nil) //true
	fmt.Println(errors.Unwrap(topContainNilE) == nil)       //true

	fmt.Println("测试 Is --------------------")
	// func Is(err, target error) bool  判断err实例是否是target实例wrap 0次到多次的结果 (注意参数targe应该是comparable的)
	fmt.Println(errors.Is(topSE, midLayerSE)) //true
	fmt.Println(errors.Is(topSE, midLayerE))  //true
	fmt.Println(errors.Is(topSE, coreE1))     //true
	fmt.Println(errors.Is(topSE, coreE1))     //true

	fmt.Println(errors.Is(midLayerSE, coreE1)) //true
	fmt.Println(errors.Is(midLayerSE, coreE1)) //true
}

//`
//MyError
//    +-- IOError
//    |    +-- ConnectionError
//    |    |    +-- ConnectionAbortedError
//    |    |    +-- ConnectionRefusedError
//    |    |    +-- ConnectionResetError
//    |    +-- FileError
//    |         +-- FileExistsError
//    |         +-- FileNotFoundError
//    +-- ValueError
//         +-- UnicodeError
//              +-- UnicodeDecodeError
//              +-- UnicodeEncodeError
//`

type eComponent struct { //参考fmt/errors.go
	msg   string
	cause error //可能是个nil, 如果是nil,那么说明当前error是最内层error //interface本质上是动态指针, 多层error的嵌套就发生在这里
}

func (e *eComponent) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.cause)
	} else {
		return e.msg
	}
}

func (e *eComponent) Unwrap() error {
	return e.cause
}

type MyError struct { //我们可以在各个Error结构体中添加各自特有的字段
	eComponent
}
type IOError struct {
	eComponent
}
type ConnectionError struct {
	eComponent
}
type ConnectionAbortedError struct {
	eComponent
}
type ConnectionRefusedError struct {
	eComponent
}
type ConnectionResetError struct {
	eComponent
}
type FileError struct {
	eComponent
}
type FileExistsError struct {
	eComponent
}
type FileNotFoundError struct {
	eComponent
}
type ValueError struct {
	eComponent
}
type UnicodeError struct {
	eComponent
}
type UnicodeDecodeError struct {
	eComponent
}
type UnicodeEncodeError struct {
	eComponent
}

func NewMyError(wrappedError error, msg string) error {
	switch wrappedError.(type) {
	case *IOError:
	case *ValueError:
	default:
		panic("Invalid Input")
	}
	return &MyError{eComponent: eComponent{msg: msg, cause: wrappedError}}
}

func NewIOError(wrappedError error, msg string) error {
	switch wrappedError.(type) {
	case *ConnectionError:
	case *FileError:
	default:
		panic("Invalid Input")
	}
	return &IOError{eComponent: eComponent{msg: msg, cause: wrappedError}}
}
func NewConnectionError(wrappedError error, msg string) error {
	switch wrappedError.(type) {
	case *ConnectionAbortedError:
	case *ConnectionRefusedError:
	case *ConnectionResetError:
	default:
		panic("Invalid Input")
	}
	return &ConnectionError{eComponent: eComponent{msg: msg, cause: wrappedError}}
}

func NewConnectionAbortedError(msg string) error {
	//两个不同的实例==测试应当返回false

	//  //(*string)(0xc0001042f0)("")
	//	//(*string)(0xc000104300)("")
	//	//(bool) false
	//	s1 := strings.Clone("")
	//	s2 := strings.Clone("")
	//	spew.Dump(&s1)
	//	spew.Dump(&s2)
	//	spew.Dump(&s1 == &s2)

	return &ConnectionAbortedError{eComponent: eComponent{msg: strings.Clone(msg)}}
}
func NewConnectionRefusedError(msg string) error {
	return &ConnectionRefusedError{eComponent: eComponent{msg: strings.Clone(msg)}}
}
func NewConnectionResetError(msg string) error {
	return &ConnectionResetError{eComponent: eComponent{msg: strings.Clone(msg)}}
}

func TestGo13Errors2() {

	e1 := NewConnectionAbortedError("ConnectionAbortedError: Detail: xxx")
	e2 := NewConnectionError(e1, "ConnectionError")
	e3 := NewIOError(e2, "IOError")

	fmt.Println("spew.Dump --------------------")
	spew.Dump(e1)
	spew.Dump(e2)
	spew.Dump(e3)

	fmt.Println("测试 != nil --------------------")
	//eInstance != nil 判断是否发生了错误
	fmt.Println(e1 != nil) //true

	fmt.Println("测试 == --------------------")
	//eInstance1 == eInstance2 判断是否是同一个error实例
	e4 := e1
	fmt.Println(e4 == e1)                                                               //true
	fmt.Println(NewConnectionAbortedError("ConnectionAbortedError: Detail: xxx") == e1) //false

	fmt.Println("测试 Unwrap --------------------")
	//解包装
	fmt.Println(errors.Unwrap(e3) == e2)  //true
	fmt.Println(errors.Unwrap(e2) == e1)  //true
	fmt.Println(errors.Unwrap(e1) == nil) //true

	fmt.Println("测试 Is --------------------")
	// func Is(err, target error) bool  判断err实例是否是target实例wrap 0次到多次的结果 (注意参数target应该是comparable的)
	fmt.Println(errors.Is(e1, e1))                                                               //true
	fmt.Println(errors.Is(e2, e1))                                                               //true
	fmt.Println(errors.Is(e3, e1))                                                               //true
	fmt.Println(errors.Is(e3, e2))                                                               //true
	fmt.Println(errors.Is(e3, NewConnectionAbortedError("ConnectionAbortedError: detail: xxx"))) //false

	fmt.Println("测试 As --------------------")
	//func As(err error, target any) bool 判断err实例是否是某个类型的实例wrap 0次到多次的结果, 并将该实例地址写入指定区域
	//在go1.13前只能使用switch e.(type) 来对错误进行类型判断, 能力较弱
	var e11 *ConnectionAbortedError
	var e22 *ConnectionError
	var e33 *IOError

	if errors.As(e3, &e33) {
		spew.Dump(e3)
		spew.Dump(e33)
		fmt.Println(e33 == e3) // true, 因为 e33 *IOError 和 e3 error 动态类型和动态值都相等
	}
	if errors.As(e3, &e22) {
		spew.Dump(e2)
		spew.Dump(e22)
		fmt.Println(e22 == e2)
	}
	if errors.As(e3, &e11) {
		spew.Dump(e1)
		spew.Dump(e11)
		fmt.Println(e11 == e1)
	}
}

// go 1.13 标准库错误处理
//http://books.studygolang.com/gobyexample/errors/
//https://polarisxu.studygolang.com/posts/go/translation/cockroachdb-errors-std-api/
//https://www.flysnow.org/2019/09/06/go1.13-error-wrapping.html
//https://www.kevinwu0904.top/blogs/golang-error/
//https://pkg.go.dev/errors
//https://go.dev/blog/go1.13-errors
//https://tonybai.com/2019/10/18/errors-handling-in-go-1-13/
//https://segmentfault.com/a/1190000020398774
