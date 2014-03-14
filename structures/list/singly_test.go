package list

import "testing"

func createEmptySingly() *Singly {
    return &Singly{}
}

func createTenNodeSingly() *Singly {
    s := &Singly{}

    var p *SinglyNode
    for i := 0; i < 10; i++ {
        n := &SinglyNode{}

        if p != nil {
            n.Next = p
        }

        n.Item = i
        p = n
    }
    s.First = p
    return s
}

func TestCreateEmptySingly(t *testing.T) {
    s := createEmptySingly()

    if s == nil {
        t.Error("Could not create Singly-Linked List.")
    }
}

func TestSetItemSingly(t *testing.T) {
    s := createEmptySingly()
    s.First = &SinglyNode{}

    s.First.Item = "Test String"

    if s.First.Item != "Test String" {
        t.Error("Could not set item for Singly-Linked List.")
    }
}

func TestCreateTenNodeSingly(t *testing.T) {
    p := createTenNodeSingly()
    n := p.First
    iCount := 0

    for n != nil {
        iCount += n.Item.(int)
        n = n.Next
    }

    if iCount != 45 {
        t.Error("Nodes were not properly attached to Singly-Linked List.")
    }
}

func TestInsertSingly(t *testing.T) {
    p := createEmptySingly()

    p.Insert(6)

    n := p.First
    if n.Item != 6 {
        t.Error("Could not insert new node into Singly-Linked List.")
    }
}

func TestDeleteSingly(t *testing.T) {
    p := createTenNodeSingly()

    iCount := 0
    p.Delete(5)

    n := p.First

    for n != nil {
        iCount += n.Item.(int)
        n = n.Next
    }

    if iCount != 40 {
        t.Error("Deletion from Singly-Linked List failed.")
    }
}
