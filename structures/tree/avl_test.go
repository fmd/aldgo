package tree

import "fmt"
import "testing"
import "github.com/fmd/aldgo/structures/tree/binst"

type AvlInt struct {
    value int
}

func (i AvlInt) Value() interface{} {
    return i.value
}

func (i AvlInt) Compare(b binst.TreeItem) bool {
    return i.value < b.Value().(int)
}

func TestAvlTree(t *testing.T) {
    b := &binst.Tree{AvlInt{5}, nil, nil, nil}
    binst.Insert(AvlInt{4}, &b, nil)
    binst.Insert(AvlInt{3}, &b, nil)
    binst.Insert(AvlInt{12}, &b, nil)

    b.Traverse(func(t *binst.Tree) {
        fmt.Println(t.Item.Value())
    })
}
