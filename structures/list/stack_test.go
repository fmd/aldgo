package list

import "testing"

func createEmptySinglyStack() *Singly {
    return &Singly{}
}

func createEmptyDoublyStack() *Doubly {
   return &Doubly{}
}

func TestSinglyStackPush(t *testing.T) {
    s := createEmptySingly()
    s.Push(5)
    s.Push(6)
    s.Push(7)

    if s.First.Item != 7 {
        t.Error("Pushing new items onto a stack has failed.")
    }
}

func TestSinglyStackPop(t *testing.T) {
    s := createEmptySingly()
    s.Push(5)
    s.Push(6)
    s.Push(7)

    item := s.Pop()
    if item != 7 {
        t.Error("Popping items from a stack has failed.")
    }

    if s.First.Item != 6 {
        t.Error("Popped item was not properly removed from the stack.")
    }
}

func TestDoublyStackPush(t *testing.T) {
    d := createEmptyDoublyStack()
    d.Push(5)
    d.Push(6)
    d.Push(7)

    if d.First.Item != 7 {
        t.Error("Pushing new items onto a stack using a Doubly-Linked List has failed.")
    }
}

func TestDoublyStackPop(t *testing.T) {
    d := createEmptyDoublyStack()
    d.Push(5)
    d.Push(6)
    d.Push(7)

    item := d.Pop
}
