package tree

import "fmt"
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

    b.Traverse(func(b *binst.Tree) {
        fmt.Println(b.Item.Value())
    })
}
