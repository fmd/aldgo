package tree

import "fmt"
import "testing"
import "github.com/fmd/aldgo/structures/tree/avl"

type AvlInt struct {
    value int
}

func (i AvlInt) Value() interface{} {
    return i.value
}

func (i AvlInt) Compare(b avl.TreeItem) bool {
    return i.value < b.Value().(int)
}

func TestBinaryTree(t *testing.T) {
    t := &avl.Tree{AvlInt{5}, nil, nil, nil}
    avl.Insert(AvlInt{4}, &t, nil)
    avl.Insert(AvlInt{3}, &t, nil)
    avl.Insert(AvlInt{12}, &t, nil)

    t.Traverse(func(t *avl.Tree) {
        fmt.Println(t.Item.Value())
    })
}
