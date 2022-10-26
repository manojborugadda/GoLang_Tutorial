package main

import (
	"fmt"
)

func main() {

	var Stef = 3 < 5
	var Alex = 30 != 30

	var Nadal bool = true // way of declaring the boolean variables
	var Novak = false     //showing the use case of " Type Inference"
	var Meddy bool        //gives the ZERO value which is false

	//short circuiting example
	var b1 = 45 < 23 && 23 == 23   // second operand wont be evaluated since first evaluates to false 
	var b2 = 8*10 == 80 || 23 < 2  // second operand wont be evaluated since first evaluates to true

	//using short variable declaration
	a := 23
	b := 24

	fmt.Println(b1, b2)
	fmt.Println("a is 23 and b is 24 the result of a == b gives ---> ", a == b)
	fmt.Printf("%v,%T\n", Stef, Stef)
	fmt.Printf("%v,%T\n", Alex, Alex)
	fmt.Printf("%v , %T\n", Nadal, Nadal)
	fmt.Printf("%v, %T\n", Novak, Novak)
	fmt.Printf("%v , %T", Meddy, Meddy)
}


/* output of the above program would be
false true
a is 23 and b is 24 the result of a == b gives --->  false
true,bool
false,bool
true , bool
false, bool
false , bool

*/
