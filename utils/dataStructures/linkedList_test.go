package dataStructures

import (
	"testing"
)

func TestLinkedList_PushFront(t *testing.T) {
	ll := NewLinkedList()
	err := ll.PushFront(12)
	if err != nil {
		t.Error(err)
	}

	err = ll.PushFront("string")
	if err == nil {
		t.Error("expected linked list to keep only slements of one type")
	}
}

func TestLinkedList_PushBack(t *testing.T) {
	ll := NewLinkedList()
	err := ll.PushFront(12)
	if err != nil {
		t.Error(err)
	}

	err = ll.PushFront("string")
	if err == nil {
		t.Error("expected linked list to keep only slements of one type")
	}
}

func TestLinkedList_Pop(t *testing.T) {
	ll := NewLinkedList()
	value := int(12)
	err := ll.PushFront(value)
	if err != nil {
		t.Error(err)
	}

	inst := ll.Pop()
	if err != nil {
		t.Error(err)
	}

	v, ok := inst.(int)
	if ok != true {
		t.Error("can't convert")
	}

	if v != value {
		t.Error("Put and got values are not the same")
	}
}
