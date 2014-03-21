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

func (d *Tree) Delete() {
    if d.Left != nil && d.Right != nil {
        m := d.Right.Minimum()
        d.Item = m.Item
        m.Delete()
    } else if d.Left != nil {
        d.Replace(d.Left)
    } else if d.Right != nil {
        d.Replace(d.Right)
    } else {
        d.Replace(nil)
    }
}

func (t *Tree) Height() int {
    if t == nil {
        return 0
    }

    l := 0
    if t.Left != nil {
        l := t.Left.Height()
    }

    r := 0
    if t.Right != nil {
        r := t.Right.Height()
    }

    if l > r {
        return l + 1
    }

    return r + 1
}

func (t *Tree) Replace(r *Tree) {
    if t.Parent != nil {
        if t == t.Parent.Left {
            t.Parent.Left = r
        } else if t == t.Parent.Right {
            t.Parent.Right = r
        }
    }

    if r != nil {
        r.Parent = t.Parent
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
