package telegram

import "fmt"

type Config struct {
}

type Telegram struct {
}

func ProvideTelegram() *Telegram {
	return &Telegram{}
}

func (t *Telegram) Run() {
	fmt.Println("Hello World!")
}
