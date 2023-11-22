//pointers

//pointers make sure that the actual value of the variable from it's address gets passed!
package main

import "fmt"

func main(){
	//fmt.Println("This is a class on pointers");
    //var ptr *int; 
	//fmt.Printf("The value of the ptr is", ptr);
	no := 23;
	var ptr = &no;
	fmt.Println("The memory address of the pointer is ", ptr);
	fmt.Println("The actual value of the pointer is ", *ptr);

	//this increments the actual value of the pointer by 2
	*ptr = *ptr +2;
	fmt.Println("The value of the pointer now is", *ptr);
}


