package main

import (
	"fmt"
	"os"
)

type node struct {
	data int
	next *node
}

var head, ptr *node

// insertion at beginning
func insertbeg(val int) {
	newnode := node{data: val}
	if head == nil {
		newnode.next = nil
		head = &newnode

		fmt.Println("Inserted at beginning")
	} else {
		j := &newnode
		j.next = head
		head = j
		fmt.Println("inserted at beginning")
	}
}

// insertion at a position
func insertpos(val, pos int) {
	newnode := node{data: val}
	if head == nil {
		newnode.next = nil
		head = &newnode
		fmt.Println("Inserted at beginning")
	} else {
		j := head
		for i := 0; i < pos-2; i++ {
			j = j.next

		}

		temp := j.next
		m := &newnode
		m.next = temp
		j.next = m
		fmt.Println("Element inserted")

	}
}

func insertend(val int) {
	newnode := node{data: val}
	if head == nil {
		newnode.next = nil
		head = &newnode
		fmt.Println("Inserted at beginning")
	} else {
		var i, j *node
		for i = head; i.next != nil; i = i.next {
			continue
		}
		j = &newnode
		j.next = nil
		i.next = j
		fmt.Println("inserted at end")
	}
}

func deletebeg() {
	if head == nil {
		fmt.Println("List empty")
	} else {
		var j *node
		ptr = head
		j = ptr.next
		head = j
		fmt.Println("Succesfully deleted")

	}
}

func deletepos(pos int) {
	if head == nil {
		fmt.Println("List empty")
	} else {
		j := head
		for i := 0; i < pos-2; i++ {
			j = j.next

		}
		fmt.Println(j.data)
		j.next = j.next.next
		fmt.Println("Element deleted")
	}
}

func deleteend() {
	if head == nil {
		fmt.Println("List empty")
	} else {
		var j *node
		for i := head; i.next != nil; i = i.next {
			j = i
		}
		j.next = nil
		fmt.Println("deleted")
	}
}

func display() {
	if head == nil {
		fmt.Println("List empty")
	} else {
		var i *node
		for i = head; i != nil; i = i.next {
			fmt.Print("->", i.data)
		}
	}
}

func main() {
	var op, val, pos int
	for i := 0; i != 8; {
		fmt.Print("\n\nOptions-----------------------------\n1.Insertion at beginning\n2.Insertion at a position\n3.Insertion at end\n4.Deletion at beginning\n5.Deletion at position\n6.Deletion at end\n7.Display\n8.EXIT\n-------------------------------------\nENter your Choice :")
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Println("Enter the new value")
			fmt.Scanln(&val)
			insertbeg(val)
		case 2:
			fmt.Println("Enter the new value")
			fmt.Scanln(&val)
			fmt.Println("Enter the location")
			fmt.Scanln(&pos)
			insertpos(val, pos)
		case 3:
			fmt.Println("Enter the new value")
			fmt.Scanln(&val)
			insertend(val)
		case 4:
			deletebeg()
		case 5:
			fmt.Println("Enter the location")
			fmt.Scanln(&pos)
			deletepos(pos)
		case 6:
			deleteend()
		case 7:
			display()
		case 8:
			os.Exit(0)
		default:
			fmt.Println("Invalid Entry")
		}
		i = op
	}
}
