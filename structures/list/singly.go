package list

// Singly-Linked List
type Singly struct {
    First *SinglyNode
}

type SinglyNode struct {
    Next *SinglyNode
    Item interface{}
}

func (s *Singly) Search(item interface{}) *SinglyNode {
    return s.First.Search(item)
}

func (n *SinglyNode) Search(item interface{}) *SinglyNode {
    if n == nil {
        return nil
    }

    if n.Item == item {
        return n
    }

    return n.Next.Search(item)
}

func (n *SinglyNode) LastNode() *SinglyNode {
    if n.Next != nil {
        return n.Next.LastNode()
    }

    return n
}

func (s *Singly) LastNode() *SinglyNode {
    return s.First.LastNode()
}

func (s *Singly) Insert(item interface{}) {
    p := &SinglyNode{}
    p.Item = item
    p.Next = s.First
    s.First = p
}

func (s* Singly) Delete(item interface{}) {
    p := s.Search(item)

    if p == nil {
        return
    }

    prev := s.First.Predecessor(item)

    if prev == nil {
        s.First = p.Next
    } else {
        prev.Next = p.Next
    }
}

func (n *SinglyNode) Predecessor(item interface{}) *SinglyNode {
    if n == nil || n.Next == nil {
        return nil
    }

    if n.Next.Item == item {
        return n
    }

    return n.Next.Predecessor(item)
}
