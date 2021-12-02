package main

import "fmt"

type person interface {
	Talk()
	Count()
	Walk()
}

type Jimmy struct {
	name string
}

func (j *Jimmy) Talk() {
	fmt.Println("Hello this is me!")
}

func (j *Jimmy) Count() {
	fmt.Println("1, 2, 3 4, 5, 6, 7")
}

func (j *Jimmy) Walk() {
	fmt.Println("\\o/ I'm here and now I'm                                over here.\\o/ ")
}

func main() {
	jimmy := &Jimmy{}

	jimmy.Talk()
	jimmy.Count()
	jimmy.Walk()
}
