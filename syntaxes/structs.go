package main

import "fmt"

type platform struct{
  name string
  kind string
}

type authenticationInfo struct{
  username string
  password string
  platform
}

func (user authenticationInfo) getBasicAuth() {
  fmt.Println("Auth Details 👇\n"+ user.username + ":" + user.password + "\nPlatform :" + user.platform.name + "," + user.kind );  
}


func main(){
  
  user1 := authenticationInfo{
    username: "John Doe",
    password: "password",
    platform:platform{
      name: "yt",
      kind: "brainrot",
    },
  }

  user1.getBasicAuth()

  
}


