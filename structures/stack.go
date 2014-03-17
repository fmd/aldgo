package structures 

type Stack interface {
    Push(interface{})
    Pop() interface{}
}

func (s *Singly) Push(item interface{}) {
    s.Insert(item)
}

func (s *Singly) Pop() interface{} {
    item := s.First.Item
    s.Delete(s.First)
    return item
}

func (d *Doubly) Push(item interface{}) {
    d.InsertBeginning(item)
}

func (d *Doubly) Pop() interface{} {
    item := d.First.Item
    d.Delete(d.First)
    return item
}
