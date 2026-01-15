package main

import "fmt"

func main() {
	hitesh := User{"Hitest", "hitesh@gmail.com", true, 23}
	fmt.Println(hitesh)
	fmt.Printf("hitesh status is %v", hitesh.GetStatus())
	fmt.Printf("New mail is %v",hitesh.NewMail())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() bool {
	return u.Status
}
func (u User) NewMail() string{
	u.Email = "test@gmail.com"
	return u.Email
}