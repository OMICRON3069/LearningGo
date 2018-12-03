package main

import "fmt"

func HeadSort(piece []int, position int, num int) {
	current := piece[position]
	for i := 2*(position+1) - 1; i < num; i = 2*(i+1) - 1 {
		// start from fist left child
		// use for loop to travel all child node and select the biggest value
		// if section to find biggest value of 3 node
		if i+1 < num { //make sure will not cross border
			if piece[i] < piece[i+1] {
				i++ // so piece[i] is the biggest value among child node
			}
		}
		if current >= piece[i] {
			break // stop loop because current parent is biggest
		}
		piece[position] = piece[i] // set parent to the biggest value
		position = i               // continue loop from swapped node, but currently node's value haven't been changed
	}
	piece[position] = current
}

func HeapSort(jibai []int, num int) {

	for i := (num / 2) - 1; i >= 0; i-- {
		// start loop from last one which have child node
		// I choose to create max-heap here
		HeadSort(jibai, i, num)
	}

	for i := num - 1; i >= 0; i-- {
		// swap first node with last
		jibai[0], jibai[i] = jibai[i], jibai[0]
		// reformat max-heap
		HeadSort(jibai, 0, i)
	}
}

func main() {
	// this is an experimental code,
	// in traditional heap structure,
	// array start in 1, so I need some trick to deal with this
	jibai := [9]int{0, 50, 60, 40, 70, 80, 20, 30, 10}
	HeapSort(jibai[0:9], 9)
	fmt.Println(jibai[0:9])
}
