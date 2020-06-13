/*
 *Read input from STDIN. Print your output to STDOUT
 *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
 */

package main

import "fmt"

func main() {
	createLinkedList(1)
}

type node struct {
	data int
	next *node
}

func createLinkedList(data int) {
	newNode := &node{data, nil}
	nodeList := newNode
	nodeList = insertNode(newNode, nodeList)
	printList(nodeList)

}

func printList(nodeList *node) {
	for p := nodeList; p != nil; p = p.next {
		fmt.Println(p)
	}
}
func insertNode(newNode, nodeList *node) *node {
	if nodeList == nil {
		return nodeList
	}
	for p := nodeList; p != nil; p = p.next {
		if p.next == nil {
			p.next = newNode
			return nodeList
		}
	}
	return nodeList
}
