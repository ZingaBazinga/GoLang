package week3

import "fmt"

func Week3() {
	email := EmailNotifier{To: "user@example.com"}
	slack := SlackNotifier{Channel: "#welcome"}

	NotifyUser(email, "Друг")
	NotifyUser(slack, "Друг")

	bad := EmailNotifier{To: ""}
	err := NotifyUser(bad, "Друг")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
