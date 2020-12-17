package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type node struct {
	value interface{}
	prev  *node
	next  *node
}
type LinkedQueue struct {
	head *node
	tail *node
	size int
}

func (queue *LinkedQueue) Size() int {
	return queue.size
}
func (queue *LinkedQueue) Peek() interface{} {
	if queue.head == nil {
		panic("Empty queue.")
	}
	return queue.head.value
	// value, err := queue.head.value
	// if err != nil {
	// 	return nil, err
	// }
	// return value, nil
}
func (queue *LinkedQueue) Add(value interface{}) {
	new_node := &node{value, queue.tail, nil}
	if queue.tail == nil {
		queue.head = new_node
		queue.tail = new_node
	} else {
		queue.tail.next = new_node
		queue.tail = new_node
	}
	queue.size++
	new_node = nil
}
func (queue *LinkedQueue) Remove() {
	if queue.head == nil {
		panic("Empty queue.")
	}
	first_node := queue.head
	queue.head = first_node.next
	first_node.next = nil
	first_node.value = nil
	queue.size--
	first_node = nil
}

// func (string) ParseQueryStringFromFuture(future string, ) {

// }

type Selector struct {
	Members []SelectorMember `json:"$or"`
}

type SelectorMember struct {
	ID string `json:"future"`
}

func GetStringForSelctorMemberListFromString(future string) string {
	var memberList []SelectorMember
	strList := strings.Split(future, "|")
	if len(strList) == 0 {
		return ""
	}
	for i := 0; i < len(strList); i++ {
		// fmt.Println("i=", i, strList[i])
		if strList[i] != "" {
			var member SelectorMember
			member.ID = strList[i]
			memberList = append(memberList, member)
		}
	}
	selector := Selector{
		Members: memberList,
	}
	q, err := json.Marshal(selector)
	if err != nil {
	}
	finalQstring := `{"selector":` + string(q) + `}`

	return finalQstring
}

func main() {
	// var queue LinkedQueue
	// queue.Add(1)
	// queue.Add(2)

	// queue.Remove()
	// println(queue.Peek().(int))
	// queue.Remove()
	// queue.Peek()
	test := "abcd|efgh|hijk|elmn|"
	list := GetStringForSelctorMemberListFromString(test)
	fmt.Println(list)

	// fmt.Println(a, len(a))
	// fmt.Println(test)
	// fmt.Println(a[0])
	// s := Selector{
	// 	Members: list,
	// }
	// j, err := json.Marshal(s)
	// if err != nil {
	// }
	// fmt.Println(string(j))

}
