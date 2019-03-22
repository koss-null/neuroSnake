package dataStructures

import (
	"reflect"
	"errors"
)

var typeMismatchError = errors.New("list type mismatch")

type (
	node struct {
		value interface{}
		next  *node
	}

	linkedList struct {
		head     *node
		tail     *node
		dataType reflect.Type
	}

	LinkedList interface {
		PushFront(interface{}) error
		PushBack(interface{}) error
		Pop() interface{}
		Head() interface{}
		Slice() []interface{}
	}
)

func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (ll *linkedList) PushFront(item interface{}) error {
	if ll.head == nil {
		n := node{item, nil}
		ll.head, ll.tail = &n, &n
		ll.dataType = reflect.TypeOf(item)
		return nil
	}

	if ll.dataType != reflect.TypeOf(item) {
		return typeMismatchError
	}

	n := node{item, ll.head}
	ll.head = &n
	return nil
}

func (ll *linkedList) PushBack(item interface{}) error {
	if ll.head == nil {
		n := node{item, nil}
		ll.head, ll.tail = &n, &n
		ll.dataType = reflect.TypeOf(item)
		return nil
	}

	if ll.dataType != reflect.TypeOf(item) {
		return typeMismatchError
	}

	n := node{item, nil}
	ll.tail.next = &n
	ll.tail = &n
	return nil
}

func (ll *linkedList) Pop() interface{} {
	if ll.head == nil {
		return nil
	}

	this := ll.head
	for this.next != ll.tail {
		this = this.next
	}
	defer func() { ll.tail = this }()
	return ll.tail.value
}

func (ll *linkedList) Head() interface{} {
	return ll.head.value
}

func (ll *linkedList) Slice() []interface{} {
	this := ll.head
	sl := make([]interface{}, 0)
	for this != nil {
		sl = append(sl, this.value)
		this = this.next
	}

	return sl
}
