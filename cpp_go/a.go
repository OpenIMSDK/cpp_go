package main
/*
#cgo CFLAGS: -I .
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
typedef void (*callback)(void *,int);
extern void c_callback (void *,int);
extern callback _cb;
*/
import "C"
import (
	"sync"
	"unsafe"
	"time"
	"fmt"
	"encoding/json"
)

var mutex  sync.Mutex


type HelloWorld  interface {
	  OnSuccessCallback(result string)
}

func doSomething(worker HelloWorld, input string){
	go func() {
		fmt.Println("func doSomething start...")
		time.Sleep(time.Duration(2)*time.Second)
		fmt.Println("func doSomething end...")
		result := "do something successful"
		worker.OnSuccessCallback(result)
	}()
}

type SomeHelloWorld struct {
	cb C.callback
	input string
	passBackData string
}

type CallbackOutput struct {
	Data     string   `json:"data"`
	Output   string `json:"output"`
}


func(t *SomeHelloWorld) OnSuccessCallback(result string){
	var callbackOutput CallbackOutput
	callbackOutput.Data = t.passBackData
	callbackOutput.Output = result

	jsonStr, err := json.Marshal(callbackOutput)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("json: ", string(jsonStr))

	var cmsg *C.char = C.CString(string(jsonStr))
	var ivalue C.int = C.int(len(jsonStr))
	defer C.free(unsafe.Pointer(cmsg))

	mutex.Lock()
	defer mutex.Unlock()
	C._cb = t.cb
	C.c_callback(unsafe.Pointer(cmsg), ivalue)
}

//export  doSomethingCallback
func doSomethingCallback(p C.callback, input *C.char, data *C.char){
	var one SomeHelloWorld
	one.cb = p
	one.passBackData = C.GoString(data)
	one.input = C.GoString(input)
	fmt.Println("one: ", one)
	doSomething(&one, one.input)
}

func main() {
}



