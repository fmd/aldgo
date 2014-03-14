package list

import "testing"

func TestCreate(t *testing.T) {
    var s *Singly
    s = &Singly{}

    if s == nil {
        t.Error("Could not create Singly-Linked List.")
    }
}

func TestSetItem(t *testing.T) {
    var s *Singly
    s = &Singly{}
    s.Item = "Test String"

    if s.Item != "Test String" {
        t.Error("Could not set item for Singly-Linked List.")
    }
}

func TestAttachNodes(t *testing.T) {
    var p *Singly
    for i := 0; i < 10; i++ {
        s := &Singly{}

        if p != nil {
            s.Next = p
        }

        s.Item = i
        p = s
    }

    iCount := 0

    for p != nil {
        iCount += p.Item
        p = p.Next
    }

    if iCount != 45 {
        t.Error("Nodes were not properly attached to list.")
    }
}

func TestInsert(t *testing.T) {
    p := &Singly{}
    p.Item = 5

    Insert(&p, 6)

    iCount := 0
    for p != nil {
        iCount += p.Item
    }

    if iCount != 11 {
        t.Error("Could not insert new node into list")
    }
}
