package main

import "fmt" 


//Lists 
// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
}

// Displays prints the elements of  the linked list
func (ll *LinkedList) DisplayList(){
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

// Stacks : A stack is a linear data structure that follows the Last-In, First-Out (LIFO) principle
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (int,error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is emptty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

// QUEUES

// Queue represents a queue implemented using a slice
type Queue struct {
	items []int
}

// Enqueue adds an element ot the rear of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element from the front of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	front := q.items[0]
	q.items = q.items[1:] //Create a new slince exclusing the first element
	return front, nil
}

// Peek returns the element at the front of the queue without removing it
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Maps Or Hash Tables



func main() {
	var a [5]int // Declares an array 'a' of 5 integers
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50

	fmt.Println(a)

	b := [3]string{"apple","banana","cherry"}
	fmt.Println(b)

	//Now we work with slices. dynamically-sized way to work with sequences of elements
	var s []int
	fmt.Println(s, len(s), cap(s) ) // output : [] 0 0

	// len(), the number of elements currently in the slice
	// cap(), the maximum number of elements the underlying array can hold starting from the slice first element.

	numbers := []int{1,2,3,4,5}
	fmt.Println(numbers, len(numbers), cap(numbers))

	// Creating a slice from an array
	arr := [5]int{10,20,30,40,50}
	slice := arr[1:4]
	fmt.Println(slice, len(slice), cap(slice))

	//We can use the append()
	number2 := []int{8,9,10}
	number2 = append(number2,11)
	fmt.Println(number2, len(number2), cap(number2))

	number2 = append(number2,5,6)
	fmt.Println(number2, len(number2),cap(number2))


	// Data into the linkedlist
	myList := LinkedList{}
	myList.InsertAtBeginning(3)
	myList.InsertAtBeginning(2)
	myList.InsertAtBeginning(1)
	myList.DisplayList()

	// Data into the Stacks

	var myStack Stack
	myStack.Push(11)
	myStack.Push(22)
	myStack.Push(33)

	fmt.Println("Stack:", myStack.items)
	top,err := myStack.Peek()
	if err != nil {
		fmt.Println("Error peeking:", err)
	} else {
		fmt.Println("Top:", top)
	}
	
	item, err := myStack.Pop()
	if err != nil {
		fmt.Println("Poppod:",item)
		fmt.Println("Stack after pop:", myStack.items)
	}

	fmt.Println("Is empty?", myStack.IsEmpty())
	fmt.Println("Size:", myStack.Size())


	// Data into queues
	var myQueue Queue
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)

	fmt.Println("Queue:", myQueue.items)   // Output: Queue: [1 2 3]
	ite,er := myQueue.Peek()    // Output: Front: 1
	if er != nil {
		fmt.Println("Error peeking:", er)
	} else {
		fmt.Println("Front:", ite)
	}

	item, error := myQueue.Dequeue()
	if error == nil {
		fmt.Println("Dequeued:", item)
		fmt.Println("Queue after dequeue:", myQueue.items)
	} else {
		fmt.Println("Error dequeuing:", error)
	}

	fmt.Println("Is empty?", myQueue.IsEmpty()) // Output: Is empty? false
	fmt.Println("Size:", myQueue.Size())      // Output: Size: 2


	// Hash Tables
	var ages map[string]int
	fmt.Println(ages == nil ) // Output: true (maps must be initialized before used)

	// Init a map using make
	ages = make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	fmt.Println(ages)

	//Declare and initialize a map with values

	scores := map[string]float64 {
		"Math":    95.5,
		"Science": 88.0,
		"English": 92.3,
	}
	fmt.Println(scores)

	// Accessing values using keys
	fmt.Println("Alice's age:", ages["Alice"])

	// Checking if a key exists
	age, ok := ages["David"]
	if ok {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David's age not found")
	}


	// Adding ne key-value pairs
	ages["Eve"] = 28
	fmt.Println(ages)

	//Deleting a key-value pair
	delete(ages,"Bob")
	fmt.Println(ages)

	// Getting the number of key-value pairs
	fmt.Println("Number of entries is ages:", len(ages))
}