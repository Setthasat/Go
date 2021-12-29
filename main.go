package main

import "fmt"

func divide() {
	var x int
	fmt.Println("[first] / second = ?")
	fmt.Print("frist = ")
	fmt.Scanln(&x)
	fmt.Print(x, " / second = ?")
	var j int
	fmt.Print("second = ")
	fmt.Scanln(&j)
	var ans = x / j
	fmt.Print(x, " / ", j, " = ", ans)
}

func multiplay() {
	var x int
	fmt.Println("[first] * second = ?")
	fmt.Print("frist = ")
	fmt.Scanln(&x)
	var j int
	fmt.Println(x, " * [second] = ?")
	fmt.Print("second = ")
	fmt.Scanln(&j)
	var ans = x * j
	fmt.Print(x, " * ", j, " = ", ans)
}

func subtract() {
	var x int
	fmt.Println("[first] - second = ?")
	fmt.Print("frist = ")
	fmt.Scanln(&x)
	var j int
	fmt.Println(x, " - [second] = ?")
	fmt.Print("second = ")
	fmt.Scanln(&j)
	var ans = x - j
	fmt.Print(x, " - ", j, " = ", ans)
}

func plus() {
	var x int
	fmt.Println("[first] + second = ?")
	fmt.Print("frist = ")
	fmt.Scanln(&x)
	var j int
	fmt.Println(x, " + [second] = ?")
	fmt.Print("second = ")
	fmt.Scanln(&j)
	var ans = x + j
	fmt.Print(x, " + ", j, " = ", ans)
}

func main() {
	var choice int
	fmt.Println("pass 1 to plus")
	fmt.Println("pass 2 to subtract")
	fmt.Println("pass 3 to multiplay")
	fmt.Println("pass 4 to divide")
	fmt.Print("place your choice here : ")
	fmt.Scanln(&choice)
	if choice == 1 {
		plus()
	} else if choice == 2 {
		subtract()
	} else if choice == 3 {
		multiplay()
	} else if choice == 4 {
		divide()
	} else {
		fmt.Println("place only 1-4!!")
	}

}
