package structures 

type Queue interface {
    Enqueue(interface{})
    Dequeue() interface{}
}

func (s *Singly) Enqueue(item interface{}) {
    s.Insert(item)
}

func (s *Singly) Dequeue() interface{} {
    l := s.LastNode()
    item := l.Item
    s.Delete(l)
    return item
}
func (d *Doubly) Enqueue(item interface{}) {
    d.InsertBeginning(item)
}

func (d *Doubly) Dequeue() interface{} {
    item := d.Last.Item
    d.Delete(d.Last)
    return item
}
