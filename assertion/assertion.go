package assertion

import (
    "testing"
    "fmt"
    "../color"
    "strings"
    "runtime/debug"
)

type testFunction func(*Assertion)
type callbackFunction func()

type Assertion struct {
    t *testing.T
}

func (lib *Assertion) Fail(message interface{}, args ... string) {
    msg := fmt.Sprint(message, args)
    fmt.Println(color.CreateText(">>>> StackTrace").Color("Cyan"))
    var stacktrace string
    stacktrace = fmt.Sprint(string(debug.Stack()))
    stacktrace = strings.SplitAfterN(stacktrace, "\n", 11)[10]
    fmt.Println(stacktrace)
    fmt.Println(color.CreateText("<<<< End Of StackTrace").Color("Cyan"))
    lib.t.Errorf(msg)
    panic("Fail")
}

func (lib *Assertion) equalAssertion(objectExpected interface{}, objectActual interface{}, args ... string) {
    if len(args) == 1 {
        message := color.CreateText(args[0]).Color("Red")
        lib.Fail(message)
    } else {
        message := color.CreateText("\nExpected: {0}\nActual  : {1}", objectExpected, objectActual).Color("Red")
        lib.Fail(message)
    }
}

func (lib *Assertion) AssertEqual(objectExpected interface{}, objectActual interface{}, args ... string) {
    if objectExpected == objectActual {
    } else {
        lib.equalAssertion(objectExpected, objectActual, args ...)
    }
}
func (lib *Assertion) AssertNotEqual(objectExpected interface{}, objectActual interface{}, args ... string) {
    if objectExpected == objectActual {
        lib.equalAssertion(objectExpected, objectActual, args ...)
    } else {
    }
}
func (lib *Assertion) AssertIsNil(arg interface{}) {
    if arg == nil {

    } else {
        lib.Fail("Argument is not nil")
    }
}
func (lib *Assertion) AssertIsNotNil(arg interface{}) {
    if arg == nil {
        lib.Fail("Argument is nil")
    } else {
    }
}

func (lib *Assertion) AssertPanic(callback callbackFunction, message string) {
    defer func() {
        if r := recover(); r != nil {
            lib.AssertIsNotNil(r)
            errMessage := fmt.Sprint(r)
            lib.AssertEqual(message, errMessage)
        } else {
            lib.Fail("Expect panic ", message)
        }
    }()
    callback()
}


func Init(t *testing.T, testFn testFunction) *Assertion {
    defer func() {
        if r := recover(); r != nil {
            if r == "Fail" {
                fmt.Println("Failed!!!")
            } else {
                fmt.Println("Failed!!! 2")
                fmt.Println(r)
                msg := fmt.Sprint("Exception occurs:\n", r)
                fmt.Println(color.CreateText(msg).Color("Red"))
            }
        }
    }()
    lib := &Assertion{t: t}
    testFn(lib)
    return lib
}