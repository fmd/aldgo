package bst

type TreeItem interface {
    Value() interface{}
    Compare(TreeItem) bool
}

type Tree struct {
    Item TreeItem
    Left *Tree
    Right *Tree
    Parent *Tree
}
