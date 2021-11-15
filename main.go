package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := ""
	fmt.Print("create your password:")
	fmt.Scan(&password)
	/*fmt.Printf("your password is: %v", password)*/
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(password)
	//fmt.Println(encrypted)
	/*password hash, to save in BD*/
	//fmt.Println(string(encrypted))
	BD_hash(string(encrypted))
	file_to_hash()
}
func BD_hash(encrypted string) {
	f, err := os.Create("hashes2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(encrypted))
	fmt.Fprintln(f, string(encrypted))
	f.Close()
	fmt.Println("SAVED IN BD")
}
func file_to_hash() {
	file_data, err := ioutil.ReadFile("./hashes2.txt")
	if err != nil {
		fmt.Println("ERROR!!!!!!!!!!!!!!!!!")
	}
	//fmt.Println(string(file_data))
	fmt.Println("LEIDO!!!")
	logg(file_data, err)
}
func logg(file_data []byte, err error) {
	loginPass := ""
	fmt.Print("login with your password:")
	fmt.Scan(&loginPass)
	//fmt.Printf("your password is: %v", loginPass)*/
	err = bcrypt.CompareHashAndPassword(file_data, []byte(loginPass))
	if err != nil {
		fmt.Println("LOGIN ERROR!!!!!!!!!!!")
		return
	}

	fmt.Println("YOU ARE SUCCESSFULLY LOG :)")
}
