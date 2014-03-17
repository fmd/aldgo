package structures

// Doubly-Linked List
type Doubly struct {
    First *DoublyNode
    Last  *DoublyNode
}

// Doubly-Linked List Node
type DoublyNode struct {
    Next *DoublyNode
    Prev *DoublyNode
    Item interface{}
}

func (n *DoublyNode) Search(item interface{}) *DoublyNode {
    if n == nil {
        return nil
    }

    if n.Item == item {
        return n
    }

    return n.Next.Search(item)
}

func (d *Doubly) Search(item interface{}) *DoublyNode {
    if d.First == nil || d.Last == nil {
        return nil
    }

    return d.First.Search(item)
}

func (d *Doubly) InsertAfter(a *DoublyNode, item interface{}) {
    n := &DoublyNode{}
    n.Prev = a
    n.Next = a.Next
    n.Item = item

    if a.Next != nil {
        a.Next.Prev = n
    } else {
        d.Last = n
    }

    a.Next = n
}

func (d *Doubly) InsertBeginning(item interface{}) {
    if d.First == nil {
        node := &DoublyNode{}
        node.Item = item
        node.Prev = nil
        node.Next = nil

        d.First = node
        d.Last = node
    } else {
        d.InsertBefore(d.First, item)
    }
}

func (d *Doubly) InsertEnd(item interface{}) {
    if d.Last == nil {
        d.InsertBeginning(item)
    } else {
        d.InsertAfter(d.Last, item)
    }
}

func (d *Doubly) InsertBefore(b *DoublyNode, item interface{}) {
    n := &DoublyNode{}
    n.Prev = b.Prev
    n.Next = b
    n.Item = item

    if b.Prev != nil {
        b.Prev.Next = n
    } else {
        d.First = n
    }

    b.Prev = n
}

func (d *Doubly) Delete(del *DoublyNode) {

    if del.Prev == nil {
        d.First = del.Next
    } else {
        del.Prev.Next = del.Next
    }

    if del.Next == nil {
        d.Last = del.Prev
    } else {
        del.Next.Prev = del.Prev
    }
}
