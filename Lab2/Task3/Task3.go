package main

import (
	"fmt"
	"sync"
)
type Authenticator struct {
	username string
}
var instance *Authenticator

var once sync.Once

func GetInstance() *Authenticator {
	once.Do(func ()  {
		fmt.Println("Створюжмо новий екземпляр Authenticator")
		instance = &Authenticator{}
	})
	return instance
}
func (a *Authenticator) SetUsername(username string){
	a.username = username
}

func (a *Authenticator) GetUsername()string {
	return a.username
}
func main() {
	auth1 := GetInstance()
	auth1.SetUsername("Admin")
	
	auth2 := GetInstance()

	fmt.Println("ім'я користувача", auth2.GetUsername())
	if auth1 == auth2 {
		fmt.Println("auth1 і auth2 — це один і той самий екземпляр")
	}else {
		fmt.Println("Помилка: створено більше одного екземпляра")
	}
}