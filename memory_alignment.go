package main

import (
	"fmt"
	"unsafe"
)

/**
var a bool          // 1字节
var b int16         // 2字节
var c int32         // 4字节
var d int64         // 8字节
var e int32         // 4字节
var f int64         // 8字节
var g int           // 8字节
var h string        // 16字节
var i float32       // 4字节
var j float64       // 8字节
var k interface{}   // 16字节
var l time.Time     // 24字节，结构体字节数不稳定
var m time.Timer    // 80字节，结构体字节不稳定
var n time.Duration // 8字节
var o []byte        // 24字节
var p uint8         // 1字节
**/

type User1 struct {
	Age    uint8 // 1字节
	Hunger int64 // 8字节
	Happy  bool  // 1字节
}

type User2 struct {
	Hunger int64 // 8字节
	Age    uint8 // 1字节
	Happy  bool  // 1字节
}

var user1 User1
var user2 User2

func main() {
	fmt.Printf("Size of %T struct: %d bytes\n", user1, unsafe.Sizeof(user1))

	fmt.Printf("Size of %T struct: %d bytes\n", user2, unsafe.Sizeof(user2))
}
