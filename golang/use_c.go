// use_c.go
package main

//#include <stdlib.h>
//#include <stdio.h>
import "C"
import "fmt"

func main() {
	fmt.Println("Hello World! Wiht C!")

	r := C.rand()
	fmt.Println(r)

}
