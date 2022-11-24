package errors

import (
	"fmt"
)

func ExampleNew() {
	err := NewWithCode(unknownCoder.Code(), "whoops")
	fmt.Println(err)
	// fmt.Printf("%+v", err)

	// Output: An internal server error occurred
}

func ExampleNew_printf() {
	err := NewWithCode(10001, "whoops 1")
	fmt.Printf("%+v", err)

	// Example output:
	// whoops 1
	// github.com/panda/errors_test.ExampleNew_printf
	//         /home/dfc/src/github.com/panda/errors/example_test.go:17
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/panda/errors/_test/_testmain.go:106
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
}

func ExampleWithCode() {
	var err error
	err = NewWithCode(ConfigurationNotValid, "this is an error message")
	fmt.Println(err)
	fmt.Println()

	err = WrapC(err, ErrInvalidJSON)
	fmt.Println(codes[err.(*WithCode).code].String())
	fmt.Println(err)

	// Output:
	// ConfigurationNotValid error
	//
	// Data is not valid JSON
	// Data is not valid JSON
}

func ExampleWithCode_code() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	fmt.Println(err.(*WithCode).code)
	// Output: 1003
}

func ExampledefaultCoder_HTTPStatus() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	fmt.Println(codes[err.(*WithCode).code].HTTPStatus())
	// Output: 500
}

func ExampleCoder_HTTPStatus() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	coder := ParseCoder(err)
	fmt.Println(coder.HTTPStatus())
	// Output: 500
}

func ExampleString() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	fmt.Println(codes[err.(*WithCode).code].String())
	// Output: Load configuration file failed
}
