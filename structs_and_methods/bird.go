package main

import "fmt"

type Bird struct {
	name string
}

func (b Bird) bird() {
	fmt.Printf("I am a bird, my name is %s\n", b.name)
}
func (b Bird) fly() {
	fmt.Println("I am flying")
}
func (b Bird) changeName(newName string) {
	b.name = newName;
}

func main() {
	bird := Bird { "parrot" }
	bird.bird()
	bird.fly()

	bird.changeName("eagle")
	bird.bird()
}