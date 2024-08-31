package main

import "fmt"

func init() {
	fmt.Println("Hello from init 2")
}

func main() {

	fmt.Println(true)

	// iteration
	var colors = [...]string{"merah", "kuning", "hijau"}
	fmt.Println(len(colors))
	fmt.Println(colors)

	var i = 0
	for {
		if i%2 != 0 {
			i++
			continue
		} else if i == 10 {
			break
		}

		fmt.Println("Nilai : ", i)
		i++
	}

}

func init() {
	fmt.Println("Hello from init 1")
}
