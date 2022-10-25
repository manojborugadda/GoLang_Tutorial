/*
we know Go is Statically typed language but it provides the facility to REMOVE the data types while declaring the variables.This is generally called as TYPE INFERENCE.
*/

//Declaring Multiple variables using "var" keyword

package main
import "fmt"
func main() {
  
  var a , b , c  = 10 , 30 , 40   // declaring multiple variables
  var d , e , f = 10, true , 234

var(                                 ///Another way of declaring the  multiple variables
  number = "007"
  name  = "James Bond "
  salary = 200.90
  personType = true
)

    fmt.Println(number, name, salary, personType)
    fmt.Printf("%v \n", a)
    fmt.Printf("%v \n", b)
    fmt.Printf("%v \n", c)
    fmt.Printf("%v \n", d)
    fmt.Printf("%v \n", e)
    fmt.Printf("%v \n", f)
}

/*
output:
007 James Bond 200.90 true
10
30
40
10
true
234
*/
