// c_fun.go
package main

/*
 #include <stdio.h>

int sum(int a, int b) {
	return a + b;
}

void hello() { // hello, world
	printf("Hello, world!\n");
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Hello World! cgo! ")
	var a, b int = 1, 2
	r := C.sum(C.int(a), C.int(b))
	fmt.Println(r)

	C.hello() // Hellow, world
}
