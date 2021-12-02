package main

import (
	"fmt"
)

type person interface {
	Talk()
	Count()
	Walk()
}

type Human struct {
	Name string
	Age  int
}

func (h *Human) Talk() {
	fmt.Printf("Hello this is me! My name is %s and I am %d years old\n", h.Name, h.Age)
}

func (h *Human) Count() {
	fmt.Println("1, 2, 3 4, 5, 6, 7")
}

func (h *Human) Walk() {
	fmt.Println("\\o/ I'm here and now I'm                                over here.\\o/ ")
}

func main() {
	peter := Human{"Peter", 23}
	jimmy := Human{"Jimmy", 64}

	peter.Talk()
	peter.Count()
	peter.Walk()

	jimmy.Talk()
	jimmy.Count()
	jimmy.Walk()
}
