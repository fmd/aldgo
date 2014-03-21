package avl

import "github.com/fmd/aldgo/structures/tree/binst"

type TreeItem binst.TreeItem
type Tree binst.Tree

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

func (t *Tree) BalanceFactor() int {
    return t.Left.Height() - t.Right.Height()
}
