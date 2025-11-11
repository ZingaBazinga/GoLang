package week2

import "fmt"

func Week2() {
	user1 := User{
		Name:  "Друг",
		Email: "drug@example.com",
		Age:   25,
	}

	err := user1.Validate()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Пользователь валиден:", user1.Name)
	}

	user2 := User{Name: "", Email: "bad", Age: -5}
	err = user2.Validate()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
