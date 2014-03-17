package binst

type TreeItem interface {
    Value() interface{}
    LessThan(TreeItem) bool
}

type Tree struct {
    Item TreeItem
    Left *Tree
    Right *Tree
}

func Insert(x TreeItem, t **Tree) {
    if *t == nil {
        p := &Tree{}
        p.Item = x
        *t = p
        return
    }

    if x.LessThan((*t).Item) {
        Insert(x, &((*t).Left))
    } else {
        Insert(x, &((*t).Right))
    }
}

func (t *Tree) Search(item TreeItem) *Tree {
    if item.Value() == t.Item.Value() {
        return t
    }

    if item.LessThan(t.Item) {
        return t.Left.Search(item)
    }

    return t.Right.Search(item)
}

func (t *Tree) Minimum() *Tree {
    min := t
    for min.Left != nil {
        min = min.Left
    }

    return min
}

func (t *Tree) Maximum() *Tree {
    max := t
    for max.Right != nil {
        max = max.Right
    }

    return max
}

func (t *Tree) Traverse(fn func(*Tree)) {
    if t.Left != nil {
        t.Left.Traverse(fn)
    }

    fn(t)

    if t.Right != nil {
        t.Right.Traverse(fn)
    }
}
