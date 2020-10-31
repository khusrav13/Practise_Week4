package main

import (
	"fmt"
)

type User struct {
	id       int64
	name     string
	username string
	surname  string
	login    string
	password string
	sex      string
	age      int64
	userInfo string
}




///1.Напишите функцию - done
//Меняет Значение Userinfo - done
//На новое значение - done

func (person *User) ChangeUserInfo(text string) {
	person.userInfo = text
	fmt.Println(person)
}



func main() {


	  NewUser := User{
		 id:       1,
		 name:     "Jones",
		 username: "EolX",
		 surname:  "Bones",
		 login:    "Jonesbones",
		 password: "qwerty1",
		 sex:      "male",
		 age:      23,
		 userInfo: "From CA",
	  }
	  fmt.Println(NewUser)

	  NewUser.userInfo = "From Dushanbe"
	  fmt.Println(NewUser)

	  //
	 //var a int64 = 10
	 //a := int64(10)

	NewUser.ChangeUserInfo("From GR")
	  fmt.Println(NewUser)

}
