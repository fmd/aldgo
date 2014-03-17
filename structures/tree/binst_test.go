package tree

import "testing"
import "github.com/fmd/aldgo/structures/tree/binst"

type BstInt struct {
    value int
}

func (i BstInt) Value() interface{} {
    return i.value
}

func (i BstInt) LessThan(b binst.TreeItem) bool {
    return i.value < b.Value().(int)
}

func TestBinaryTree(t *testing.T) {
    b := &binst.Tree{BstInt{5},nil,nil}
    binst.Insert(BstInt{4}, &b)
    binst.Insert(BstInt{6}, &b)
    binst.Insert(BstInt{3}, &b)
    binst.Insert(BstInt{5}, &b)
    binst.Insert(BstInt{19}, &b)

    if b.Minimum().Item.Value() != 3 {
        t.Error("Minimum function did not work on Binary Search Tree")
    }

    if b.Maximum().Item.Value() != 19 {
        t.Error("Maximum function did not work on Binary Search Tree")
    }

    if b.Search(BstInt{6}) == nil {
        t.Error("Search function did not work on Binary Search Tree")
    }
}
