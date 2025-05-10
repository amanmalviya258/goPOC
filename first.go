//static typed
//strongly typed
//better code compeltion
// go is compiled
// fast compile time
// built in concurrency
// simplicity
//garbage collection

package main

import "unicode/utf8"

import "fmt"

func main() {
	var intNum int = 32767
	intNum = intNum + intNum
	fmt.Println(intNum)

	var floatNum float64 = 234545.344432462634
	fmt.Println(floatNum)

	var typeCastingExapmple = floatNum + float64(intNum)
	fmt.Println(typeCastingExapmple)

	var myString string = "newhelloworld" ///`With backquotes you can arrange the string there only`
	fmt.Print(myString)

	var myString1 string = "Hello" + " " + "World"
	fmt.Println(myString1)

	fmt.Println(len(myString1))
	fmt.Println(len(myString)) // len does not reflect the true lenght of the it returns the number of bytes.

	fmt.Print(utf8.RuneCountInString(myString)) //this reflects the true length of the string

	var myBoolean bool = true;
	fmt.Println(myBoolean);

	var trail = "helloji";  //here we did not explicitly write string after the naming of the var
	fmt.Println(trail);
	
    var intNum1 int;
	fmt.Println(intNum1);//0    ///intitial value in golang is 0; java gives error doing this.; undefined in js ;python also gives error doing this


}
