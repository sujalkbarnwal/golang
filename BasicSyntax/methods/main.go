package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	sujal := User{"Sujal", "sujaldev@programmer.net", true, 19}
	fmt.Println(sujal)
	fmt.Printf("sujal details are: %+v\n", sujal)
	fmt.Printf("Name is %v and email is %v\n", sujal.Name, sujal.Email)

	sujal.GetStatus()
	sujal.NewMail()

	fmt.Printf("Name is %v and email is %v\n", sujal.Name, sujal.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
