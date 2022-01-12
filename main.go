package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

//insert node as head
func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}

	if L.head != nil {
		L.head.prev = list
	}
	L.head = list
	L.length += 1
	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

//insert new tail
func (L *List) Push(key interface{}) {
	list := &Node{
		prev: L.tail,
		key:  key,
	}

	if L.head != nil {
		L.tail.next = list

	}
	L.tail = list
	L.length += 1
	l := L.tail
	for l.prev != nil {
		l = l.prev
	}
	L.head = l
}

//take tail out from list
func (L *List) Pop() interface{} {
	list := L.tail
	if L.tail.prev != nil {
		L.tail = L.tail.prev
		L.tail.next = nil
	} else {
		L.head = nil
	}
	L.length -= 1
	return list.key
}

//print list value
func (L *List) Display() {
	list := L.head
	for list != nil {
		fmt.Printf(" %+v ->", list.key)
		list = list.next
	}
	fmt.Println()
	fmt.Println("Length List : ", L.length)
}

//take head out from the list
func (L *List) Shifting() interface{} {
	list := L.head
	if list.next != nil {
		L.head = L.head.next
		L.head.prev = nil
	} else {
		L.head = nil
	}
	L.length -= 1
	return list.key
}

// search value at specific index
func (L *List) Get(idx int) *Node {
	if idx < 0 || idx > L.length {
		fmt.Printf("Invalid index")
		fmt.Println()
		return nil
	}
	l := L.head
	count := 1
	for l.next != nil {
		if count == idx {
			fmt.Printf("Index %+v are %+v", idx, l.key)
			fmt.Println()
			return l
		} else {
			count += 1
			l = l.next
		}
	}
	return nil
}

func (L *List) Set(key interface{}, idx int) bool {
	val := L.Get(idx)
	if val == nil {
		return false
	} else {
		L.Get(idx).key = key
		return true
	}
}

func main() {
	link := List{}
	//link.Insert(5)
	//link.Insert(9)
	//link.Insert(13)
	//link.Insert(22)
	//link.Insert(28)
	//link.Insert("asd")
	//link.Display()
	//fmt.Println(link.Shifting())
	link.Push(32)
	link.Get(1)
	link.Display()
	link.Set(12, 1)
	link.Display()

}
