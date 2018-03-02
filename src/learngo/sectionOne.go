// sectionOne
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello SectionOne")
	i := int(1)
	fmt.Println(unsafe.Sizeof(i))
	j := int32(1)
	fmt.Println(unsafe.Sizeof(j))
	k := int32(500000000)
	fmt.Println(unsafe.Sizeof(k))
	var t string = "this test"
	fmt.Println(unsafe.Sizeof(t))
}
