package main

import (
	"fmt"
)

type Node struct {
	x           int
	left, right *Node
}

func size(u *Node) int {
	if u == nil {
		return 0
	}
	return size(u.left) + size(u.right) + 1
}

func min(u *Node) int {
	for u.left != nil {
		u = u.left
	}
	return u.x
}

func printInOrder(u *Node) {
	if u != nil {
		printInOrder(u.left)
		fmt.Println(u.x)
		printInOrder(u.right)
	}
}

func printPreOrder(level string, u *Node) {
	if u != nil {
		fmt.Println(level, u.x)
		printPreOrder(level+">", u.left)
		printPreOrder(level+">", u.right)

	}

}

func printPostOrder(u *Node) {
	if u != nil {
		printPostOrder(u.left)
		printPostOrder(u.right)
		fmt.Println(u.x)
	}
}

func find(u *Node, x int) *Node {
	fmt.Println("looking for", x)
	for u != nil {
		fmt.Printf("Inspecting location %p. Finding %v\n", u, *u)
		if x < u.x {
			fmt.Printf("%v < %v, so going left\n", x, u.x)
			u = u.left
		} else if x > u.x {
			fmt.Printf("%v > %v, so going right\n", x, u.x)
			u = u.right
		} else { //found
			fmt.Println("found")
			return u
		}
	}
	fmt.Println("not found")
	return nil
}

func main() {
	z := &Node{x: 7}
	z.left = &Node{x: 3}
	z.right = &Node{x: 11}
	z.left.left = &Node{x: 1}
	z.left.right = &Node{x: 5}

	z.right.left = &Node{x: 9}
	z.right.right = &Node{x: 13}
	z.left.right.left = &Node{x: 4}
	z.left.right.right = &Node{x: 6}
	z.right.left.left = &Node{x: 8}
	
	fmt.Println(find(z, 6))
	fmt.Println(find(z, 10)) // no 10

}
