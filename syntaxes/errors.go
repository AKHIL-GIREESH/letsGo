package main

import "fmt"

type DivideError struct{
  divisor int
}

func (de DivideError) Error() string{
  return fmt.Sprintf("Division by zero %v",de.divisor)
}


func main(){
  var num int = 0;

  if num == 0 {
    fmt.println(DivideError{divisor:num})
  }
}
