package list

import "testing"

func createEmptyDoubly() *Doubly {
    return &Doubly{}
}

func createTenNodeDoubly() *Doubly {
    d := &Doubly{}

    for i := 0; i < 10; i++ {
        d.InsertBeginning(i)
    }

    return d
}

func TestCreateEmptyDoubly(t *testing.T) {
    d := createEmptyDoubly()

    if d == nil {
        t.Error("Could not create Doubly-Linked List.")
    }
}

func TestCreateDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    n := d.First
    iCount := 0
    for n != nil {
        iCount += n.Item.(int)
        n = n.Next
    }

    if iCount != 45 {
        t.Error("Could not create and populate a ten node Doubly-Linked List.")
    }
}

func TestSearchDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    node := d.Search(5)

    if node.Item.(int) != 5 {
        t.Error("Could not search for node halfway through Doubly-Linked List.")
    }
}

func TestInsertEmptyBeginningDoubly(t *testing.T) {
    d := createEmptyDoubly()

    item := 5
    d.InsertBeginning(item)

    if d.First.Item.(int) != 5 {
        t.Error("Could not insert value into beginning of empty Doubly-Linked List.")
    }
}

func TestInsertEmptyEndDoubly(t *testing.T) {
    d := createEmptyDoubly()

    item := 5
    d.InsertEnd(item)

    if d.Last.Item.(int) != 5 {
        t.Error("Could not insert value into end of empty Doubly-Linked List.")
    }
}

func TestInsertBeginningDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    item := -1
    d.InsertBeginning(item)

    if d.First.Item.(int) != -1 {
        t.Error("Could not insert value into beginning of ten node Doubly-Linked List.")
    }
}

func TestInsertEndDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    item := 10
    d.InsertEnd(item)

    if d.Last.Item.(int) != 10 {
        t.Error("Could not insert value into end of ten node Doubly-Linked List.")
    }
}

func TestInsertBeforeDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    item := 10
    d.InsertBefore(5, item)

    n := d.Search(5)
    if n.Prev.Item.(int) != 10 {
        t.Error("Could not insert value before a node in a ten node Doubly-Linked List.")
    }
}

func TestInsertAfterDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    item := 10
    d.InsertAfter(5, item)

    n := d.Search(5)
    if n.Next.Item.(int) != 10 {
        t.Error("Could not insert value after a node in a ten node Doubly-Linked List.")
    }
}

func TestDeleteDoubly(t *testing.T) {
    d := createTenNodeDoubly()

    d.Delete(5)

    iCount := 0
    p := d.First
    for p != nil {
        iCount += p.Item.(int)
        p = p.Next
    }

    if iCount != 40 {
        t.Error("Could not delete a node from a ten node Doubly-Linked List.")
    }
}
