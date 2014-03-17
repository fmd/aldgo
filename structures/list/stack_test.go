package structures

import "testing"

func GenericStackPushPop(s Stack) string {
    s.Push(5)
    s.Push(6)
    s.Push(7)

    item := s.Pop()
    if item != 7 {
        return "Popping items from the stack has failed."
    }

    item = s.Pop()
    if item != 6 {
        return "Popped item was not properly removed from the stack"
    }

    return ""
}

func TestSinglyStackPushPop(t *testing.T) {
    s := &Singly{}
    e := GenericStackPushPop(s)
    if len(e) > 0 {
        t.Error(e)
    }
}

func TestDoublyStackPushPop(t *testing.T) {
    d := &Doubly{}
    e := GenericStackPushPop(d)
    if len(e) > 0 {
        t.Error(e)
    }
}
