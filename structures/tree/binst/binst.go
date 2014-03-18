package binst

type TreeItem interface {
    Value() interface{}
    Compare(TreeItem) bool
}

type Tree struct {
    Item TreeItem
    Left *Tree
    Right *Tree
    Parent *Tree
}

func Insert(x TreeItem, t **Tree, parent *Tree) {
    if *t == nil {
        p := &Tree{}
        p.Item = x
        p.Parent = parent
        *t = p
        return
    }

    if x.Compare((*t).Item) {
        Insert(x, &((*t).Left), *t)
    } else {
        Insert(x, &((*t).Right), *t)
    }
}

func (t *Tree) Search(item TreeItem) *Tree {
    if item.Value() == t.Item.Value() {
        return t
    }

    if item.Compare(t.Item) {
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
