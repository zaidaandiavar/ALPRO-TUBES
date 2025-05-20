package main

import "fmt"

var count int

type as struct {
	ID [100]int
}

func addID() {
	var d as
	var IDnew int
	fmt.Printf("Masukkan ID: ")
	fmt.Printf("\n")
	fmt.Scan("%d", &IDnew)
	d.ID[count] = IDnew
	count++
	fmt.Printf("ID %d berhasil ditambahkan\n", IDnew)
	addasessment()
}
