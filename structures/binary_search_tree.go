package structures

type BinarySearchTree struct {
    Item interface{}
    parent *BinarySearchTree
    left *BinarySearchTree
    right *BinarySearchTree
}

func (b *BinarySearchTree) Insert(parent *BinarySearchTree, x interface{}) {
    var bp **BinarySearchTree = &b
    if *bp == nil {
        n := &BinarySearchTree{}
        n.Item = x
        n.parent = parent
        *bp = n
        return
    }

    if x < b.Item {
        b.left.Insert(b, x)
    } else {
        b.right.Insert(b, x)
    }
}

func (b *BinarySearchTree) Search(x interface{}) *BinarySearchTree {
    if b.Item == x {
        return b
    }

    if b.Item < x {
        return b.left.Search(x)
    } else {
        return b.right.Search(x)
    }
}

func (b *BinarySearchTree) Minimum() *BinarySearchTree {
    min := b
    for min.left {
        min = min.left
    }

    return min
}

func (b *BinarySearchTree) Maximum() *BinarySearchTree {
    max := b
    for max.right {
        max = max.right
    }

    return max
}

func (b *BinarySearchTree) Traverse(fn func(*BinarySearchTree)) {
    b.left.Traverse(fn)
    fn(b)
    b.right.Traverse(fn)
}
