package list

// Singly-Linked List
type Singly struct {
    Next *Singly
    Item interface{}
}

func (s *Singly) Search(item interface{}) *Singly {
    if s == nil {
        return nil
    }

    if s.Item == item {
        return s
    }

    return s.Next.Search(item)
}

func (s *Singly) Insert(item interface{}) {
    sa := &s

    p :=  &Singly{}
    p.Item = item
    p.Next = s
    *sa = p
}

func (s *Singly) Delete(item interface{}) {
    sa := &s

    p := s.Search(item)
    if p != nil {
        prev := s.Predecessor(item)
        if prev == nil {
            *sa = p.Next
        } else {
            prev.Next = p.Next
        }
    }
}

func (s *Singly) Predecessor(item interface{}) *Singly {
    if s == nil || s.Next == nil {
        return nil
    }

    if s.Next.Item == item {
        return s
    }

    return s.Next.Predecessor(item)
}
