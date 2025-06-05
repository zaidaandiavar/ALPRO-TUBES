package main

import "fmt"

func tomboll() {
	fmt.Printf("=========================\n")
	fmt.Printf("|   Menu Assessment     |\n")
	fmt.Printf("=========================\n")
	fmt.Printf("|   1 --> Add Asessment |\n")
	fmt.Printf("|   2 --> History       |\n")
	fmt.Printf("|   3 --> Statistic     | \n")
	fmt.Printf("|   4 --> Exit          |\n")
	fmt.Printf("=========================\n")
	var tombol int
	fmt.Scan(&tombol)
	switch tombol {
	case 1:
		for {
			addID()
			break
		}
	case 2:
		fmt.Println("History")
	case 3:
		fmt.Println("Statistic")
	case 4:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid option")
	}
}
