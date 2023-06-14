package data_structures

import "fmt"

type LRU[K comparable, V comparable] struct {
	capacity      uint
	length        uint
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
	verbose       bool
}

func NewLRU[K comparable, V comparable](capacity uint) *LRU[K, V] {
	lru := &LRU[K, V]{
		capacity:      capacity,
		length:        0,
		head:          nil,
		tail:          nil,
		lookup:        make(map[K]*Node[V]),
		reverseLookup: make(map[*Node[V]]K),
		verbose:       false,
	}
	return lru
}

type Node[V comparable] struct {
	next  *Node[V]
	prev  *Node[V]
	value V
}

func NewNode[V comparable](value V) *Node[V] {
	return &Node[V]{
		next:  nil,
		prev:  nil,
		value: value,
	}
}

func (lru *LRU[K, V]) SetVerbose(verbosity bool) {
	lru.verbose = verbosity
}

func (lru *LRU[K, V]) Update(key K, value V) {
	if node, exists := lru.lookup[key]; exists {
		node.value = value
		lru.detach(node)
		lru.prepend(node)
	} else {
		node := NewNode(value)
		lru.lookup[key] = node
		lru.reverseLookup[node] = key
		lru.prepend(node)
		lru.length++
		lru.trimCache()
	}
	if lru.verbose {
		lru.PrintLinkedList()
	}
}

func (lru *LRU[K, V]) Get(key K) (V, bool) {
	if node, exists := lru.lookup[key]; exists {
		lru.detach(node)
		lru.prepend(node)
		if lru.verbose {
			lru.PrintLinkedList()
		}
		return node.value, true
	}
	return *new(V), false
}

func (lru *LRU[K, V]) detach(node *Node[V]) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = node.next
	}

	if next != nil {
		next.prev = node.prev
	}

	if lru.head == node {
		lru.head = node.next
	}

	if lru.tail == node {
		lru.tail = node.prev
	}

	node.prev = nil
	node.next = nil
}

func (lru *LRU[K, V]) prepend(node *Node[V]) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}

	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}

func (lru *LRU[K, V]) trimCache() {
	if lru.length <= lru.capacity {
		return
	}

	tail := lru.tail
	key := lru.reverseLookup[tail]
	lru.detach(tail)
	delete(lru.lookup, key)
	delete(lru.reverseLookup, tail)
	lru.length--
}

func (lru *LRU[K, V]) PrintLinkedList() {
	fmt.Println("\nRecently Least Used")
	fmt.Println("-------------------")
	current := lru.head
	fmt.Printf("|--")
	for current != nil {
		key := lru.reverseLookup[current]
		fmt.Printf("(%v|%v)-->", key, current.value)
		current = current.next
	}
	fmt.Printf("END")
	fmt.Println()
}
