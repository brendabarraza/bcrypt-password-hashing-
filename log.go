package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := ""
  
	fmt.Print("new password:")
	fmt.Scan(&password)
	/*fmt.Printf("your password is: %v", password)*/

	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	
  if err != nil {
		fmt.Println(err)
	}
	
  fmt.Println(password)
	fmt.Println(encrypted)

	/*password hash, to save in BD*/
	fmt.Println(string(encrypted))

	loginPass := ""
	 
  fmt.Print("login with your password:")
	fmt.Scan(&loginPass)
	/*fmt.Printf("your password is: %v", loginPass)*/

	err = bcrypt.CompareHashAndPassword(encrypted, []byte(loginPass))
	
  if err != nil {
		fmt.Println("LOGIN ERROR!!!!!!!!!!!")
		return
	}
  fmt.Println("you have successfully logged in :)")

}
