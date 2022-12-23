package main

import (
  "fmt"
  "strconv"
)

/* At package Level */
l := 34  // we can not declare it. Gives Error

/* we can also crrate variable in block at a time */

var (
  Name string = "Gotam Gorabh"
  SchoolName string = IIIT
  Age int = 23
)

/*
   lower case at package level - varible are scoped to the packsge
   upper case at package level - variable are globally visible
*/

func main() {

  /* format one */
  var i int       // Declaration
  i = 30          // initialisation

  /* Format two */
  var j int = 1947    // Declaration + initialisation

/* Format three*/
  Age := 23

  /* Important */
  var Id int = 173
  Id := 200       // This gives error because we redeclaring it twice.
  fmt.Printf("%v, %T", Age, Age)

  /* covertion variable type */
  var k int =23
  var l float32
  var str string
  l = float32(k)
  str = strconv.Itoa(k)

}
