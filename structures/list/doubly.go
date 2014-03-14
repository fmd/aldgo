package list

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

func (d *Doubly) InsertAfter(after interface{}, item interface{}) {
    afterNode := d.Search(after)

    if afterNode == nil {
        return
    }

    n := &DoublyNode{}
    n.Prev = afterNode
    n.Next = afterNode.Next
    n.Item = item

    if afterNode.Next != nil {
        afterNode.Next.Prev = n
    } else {
        d.Last = n
    }

    afterNode.Next = n
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
        d.InsertBefore(d.First.Item, item)
    }
}

func (d *Doubly) InsertEnd(item interface{}) {
    if d.Last == nil {
        d.InsertBeginning(item)
    } else {
        d.InsertAfter(d.Last.Item, item)
    }
}

func (d *Doubly) InsertBefore(before interface{}, item interface{}) {
    beforeNode := d.Search(before)

    if beforeNode == nil {
        return
    }

    n := &DoublyNode{}
    n.Prev = beforeNode.Prev
    n.Next = beforeNode
    n.Item = item

    if beforeNode.Prev != nil {
        beforeNode.Prev.Next = n
    } else {
        d.First = n
    }

    beforeNode.Prev = n
}

func (d *Doubly) Delete(item interface{}) {
    node := d.Search(item)

    if node.Prev == nil {
        d.First = node.Next
    } else {
        node.Prev.Next = node.Next
    }

    if node.Next == nil {
        d.Last = node.Prev
    } else {
        node.Next.Prev = node.Prev
    }
}
