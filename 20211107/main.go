package main

//#include <stdio.h>
// int PrintHello(void){
//     printf("Hello World!\n");
//     return 0;
// }
import "C"
import "fmt"

func main() {
	fmt.Println("golang output")
	C.PrintHello()
}
