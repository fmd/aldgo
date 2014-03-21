package avl

import "github.com/fmd/aldgo/structures/tree/binst"

type TreeItem binst.TreeItem
type Tree binst.Tree

func (t *Tree) BalanceFactor() int {
    return t.Left.Height() - t.Right.Height()
}

func (t *Tree) RightRotate() {
    var l *Tree
    if t.Left != nil {
        l = t.Left
    }

    if l.Right != nil {
        t.Left = l.Right
        l.Right.Parent = t
    }

    l.Right = t
    t.Parent = l
}

func (t *Tree) LeftRotate() {
    var r *Tree
    if t.Right != nil {
        r = t.Right
    }

    if r.Left != nil {
        t.Right = r.Left
        r.Right.Parent = t
    }

    r.Left = t
    t.Parent = l
}

func (t *Tree) Balance() {
    if t.BalanceFactor() == 2 {
        var l *Tree
        if t.Left != nil {
            l = t.Left
            if l.BalanceFactor() == -1 {
                l.LeftRotate()
            }
        }
        t.RightRotate()
    } else {
        var r *Tree
        if t.Right != nil {
            r = t.Right
            if r.BalanceFactor() == -1 {
                r.RightRotate()
            }
        }
        t.LeftRotate()
    }
}

func Insert(x TreeItem, t **Tree, parent *Tree) {
    binst.Insert(x, t, parent)
    (*t).Balance()
}
