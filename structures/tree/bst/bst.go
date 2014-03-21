package bst

type Tree interface {
    Item() interface{}
    Left() *Tree
    Right() *Tree
    Parent() *Tree
}
