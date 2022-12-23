package main

import (
  "fmt"
)

func main() {

  /* bool data type */
  var is_string bool = true
  n := 1 == 1
  m := 1 == 2
  var unin bool    // uninitialize bool varible always store false.
  fmt.Printf("%v, %T\n", is_string, is_string)

  /* Int data type */
  /* signed integer
    1. int8
    2. int16
    3. int32
    4. int64
  */

  /* unsigned integer
    1. uint8
    2. uint16
    3. uint32
  */

  /* Float data type */
  /*
    1. float32
    2. float64
  */

  /* complex data type */
  /*
    1. complex64
    2. complex128
  */
  
  var com complex64 = 1 + 2i
  var com1 complex128 = complex(3, 4)
  fmt.Printf("%v, %T",com, com)
  fmt.Printf("%v, %T", real(com), real(com))
  fmt.Printf("%v, %T",imag(com), imag(com))

  /* text data type */
  /*
    1. string (utf8 character)
    2. 
  */

  str1 := "I from India."
  str2 := "My name is gotam."
  fmt.Printf("%v, %T", str1, str1)
  fmt.Printf("%v, %T", str2+str1, str2+str1)


  



  


}
