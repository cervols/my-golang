package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")

	paprika := User{"Tania", "paprika@go.dev", true, 18}
	fmt.Println(paprika)	// {Tania paprika@go.dev true 18}
	fmt.Printf("paprika details are: %+v\n", paprika)	// paprika details are: {Name:Tania Email:paprika@go.dev Status:true Age:18}
	fmt.Printf("Name is %v and email is %v\n", paprika.Name, paprika.Email)	// Name is Tania and email is paprika@go.dev

	paprika.GetStatus()	// Is user active: true
	paprika.NewMail()	// Email of this user is: test@go.dev
	fmt.Printf("Name is %v and email is %v\n", paprika.Name, paprika.Email)	// Name is Tania and email is paprika@go.dev
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

func(u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
