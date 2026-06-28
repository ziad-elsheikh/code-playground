package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	quota = 25
	total_usage = 22 // read only
	max_capacity = 10
	containers = make([]int, 3)
)

func main() {
	containers[0] = 10
	containers[1] = 10
	containers[2] = 2

	mainloop()
}


func mainloop() {
LoopOut:
	for {
		choice := menu()
		switch choice {
		case 1: dashboard()
		case 2: add_container()
		case 3: add_value()
		case 4:  
			clearScreen()
			break LoopOut
		default: 
			fmt.Println("Invalid choice, try again.")
		}
	}

	fmt.Println("eb'a t3ala tany:) ta7yaty")
}


func menu() int {

	var choice int
	fmt.Println("\n==============================")
	fmt.Println("QUOTA ALLOCATOR SYSTEM")
	fmt.Println("==============================")
	fmt.Println("1. view dashboard")
	fmt.Println("2. create a new container")
	fmt.Println("3. add items to an existing container")
	fmt.Println("4. exit")
	fmt.Print("> ")

	fmt.Scan(&choice)
	return choice
}


func add_value() {
	var val int
	var cntnr int

	clearScreen()
	fmt.Printf("Enter number of the container (0:%v):\n", len(containers) - 1)
	fmt.Print("> ")
	fmt.Scan(&cntnr)

	if cntnr < 0 || cntnr >= len(containers) {
		fmt.Println("invalid container number")
		return
	}

	available := calcAvailableItems(cntnr)
	inContainer := containers[cntnr]

	clearScreen()
	fmt.Printf("Enter a value to add/remove - (max add %v, max remove %v)\n", available, inContainer)
	fmt.Print("to remove enter negative value\n")
	fmt.Print("> ")
	fmt.Scan(&val)

	if val <= available && val >= -inContainer {
		containers[cntnr] += val
		fmt.Printf("Number of used slots in container %v is %v\n", cntnr, containers[cntnr])
		fmt.Printf("available slots in container %v is %v\n", cntnr, available - val)
	}else {
		fmt.Println("You exceeded the max available slots")
	}
}


func calcAvailableItems(conIndex int) int {
	
	dynamic_total := 0
	for _, v := range containers {
		dynamic_total += v
	}
	
	remaining_quota := quota - dynamic_total
	remaining_capacity := max_capacity - containers[conIndex]

	return min(remaining_quota, remaining_capacity)
}

func add_container() {
	clearScreen()
	containers = append(containers, 0)
	fmt.Println("New container added")
	fmt.Println("==============================")	
	fmt.Println("the containers in system")
	fmt.Println("==============================")
	for k, v := range containers{
		fmt.Printf("- container %v has %v busy slots\n", k, v)
	}
}


func dashboard(){
	clearScreen()
	fmt.Println("==============================")	
	fmt.Println("the containers in system")
	fmt.Println("==============================")
	for k, v := range containers{
		fmt.Printf("- container %v has %v busy slots\n", k, v)
	}
}
func clearScreen(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// func calcAvailableItems(conIndex int) int {

// 	// remaining = min(Quota - total, max - current)
// 	other_usage := total_usage - containers[conIndex]
// 	total := other_usage + containers[conIndex]
// 	return min(quota - total, max_capacity - containers[conIndex])
// }