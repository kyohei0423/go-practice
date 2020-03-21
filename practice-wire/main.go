package main

import "github.com/kyohei0423/go-practice/practice-wire/animal"

func main() {
	// wireを使う前に想定されるコード群
	// 必要な構造体を順に作成し、依存を注入していく
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)

	// event.Start()

	e := InitializeEvent()
	e.Start()

	a := animal.InitializeDog()
	a.Cry()
}
