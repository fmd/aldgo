package binary

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

func (t *Tree) Traverse(fn func(*Tree)) {
    if t.Left != nil {
        t.Left.Traverse(fn)
    }

    fn(t)

    if t.Right != nil {
        t.Right.Traverse(fn)
    }
}
