package tree 
import "fmt"
import "testing"
import "github.com/fmd/aldgo/structures/tree/binary"

type BstInt struct {
    value int
}

func (i BstInt) Value() interface{} {
    return i.value * i.value
}

func (i BstInt) LessThan(b binary.TreeItem) bool {
    return i.value > b.Value().(int)
}

func TestCreateBinaryTree(t *testing.T) {
    b := &binary.Tree{BstInt{5},nil,nil}
    binary.Insert(BstInt{4}, &b)
    binary.Insert(BstInt{6}, &b)
    binary.Insert(BstInt{3}, &b)

    b.Traverse(func(b *binary.Tree) {
        fmt.Println(b.Item.Value())
    })
}
