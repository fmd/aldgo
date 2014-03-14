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

func Insert(s **Singly, item interface{}) {
    p :=  &Singly{}
    p.Item = item
    p.Next = *s
    *s = p
}

func (s *Singly) Predecessor(s *Singly, item interface{}) *Singly {
    if s == nil || s.Next == nil {
        return nil
    }

    if s.Next.Item == item {
        return s
    }

    return s.Next.Predecessor(s.Next, item)
 }
