package main

import "fmt"

type addMetdata interface {
	insertNewRow()
}

type updateMetadata interface {
	softDeleteCurrentRow()
	insertNewRow()
}

type offering struct {
	offerPrice int
}

type secondryMarket struct {
	buyPrice int
}

func (o offering) insertNewRow() {
	fmt.Printf("New row added")
}

func (o offering) softDeleteCurrentRow() {
	fmt.Printf("Current row soft deleted")
}

func launchPlayerTokenInOffering(add addMetdata) {
	add.insertNewRow()
}

func updatePlayerTokenInOffering(update updateMetadata) {
	update.softDeleteCurrentRow()
	update.insertNewRow()
}

func main() {
	fmt.Printf("hello..")
	x := 0
	for x < 3 {
		launchPlayerTokenInOffering()
		x = x + 1
	}
}
