package structures

type BinarySearchTree struct {
    Item interface{}
    parent *BinarySearchTree
    left *BinarySearchTree
    right *BinarySearchTree
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
