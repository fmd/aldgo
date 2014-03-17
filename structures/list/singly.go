package structures

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

func (s* Singly) Delete(n *SinglyNode) {
    prev := s.First.Predecessor(n)

    if prev == nil {
        s.First = n.Next
    } else {
        prev.Next = n.Next
    }
}

func (n *SinglyNode) Predecessor(p *SinglyNode) *SinglyNode {
    if n == nil || n.Next == nil {
        return nil
    }

    if n.Next == p {
        return n
    }

    return n.Next.Predecessor(p)
}
