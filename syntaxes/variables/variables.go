package main

import "fmt"

func main(){
  
  //Method 1
  var a = "initial"
  
  //Method 2
  var b,c int = 1,2

  //Method 3 ( gets zeroed value which is 0 in the case of int )
  var d int

  //Method 4 ( only available inside functions )
  e := "string"

  //Constants
  const f int = 10

  fmt.Println(a,b,c,d,e,f+10)
}
