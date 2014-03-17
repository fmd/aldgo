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
}
