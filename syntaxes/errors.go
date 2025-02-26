package main

import "fmt"

type DivideError struct{
  divisor int
}

func (de DivideError) Error() string{
  return fmt.Sprintf("Division by zero %v",de.divisor)
}

//func (int num) error {
//  if num == 0 {
//    return DivideError{divisor:num}
//  }

//  return nil
//}


func main(){
  var num int = 0;

  if num == 0 {
    err:=DivideError{divisor:num}
    fmt.Println(err.Error())
  }
}
