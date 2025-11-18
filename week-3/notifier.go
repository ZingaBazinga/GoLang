package week3

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	To string
}

func (e EmailNotifier) Send(message string) error {
	if e.To == "" {
		return errors.New("адрес email не указан")
	}
	fmt.Printf("[Email] Отправлено на %s: %s\n", e.To, message)
	return nil
}

type SlackNotifier struct {
	Channel string
}

func (s SlackNotifier) Send(message string) error {
	if s.Channel == "" {
		return errors.New("канал Slack не указан")
	}
	fmt.Printf("[Slack] Отправлено в %s: %s\n", s.Channel, message)
	return nil
}

func NotifyUser(n Notifier, name string) error {
	message := fmt.Sprintf("Привет, %s! Добро пожаловать.", name)
	return n.Send(message)
}
