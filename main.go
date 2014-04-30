package main

import "fmt"
import "flag"
import list "github.com/fmd/aldgo/structures/list"

/** 
 *  Set up the program to run through all problems unless a particular problem is supplied.
 */
var problemFlag = flag.String("p", "all", "The problem function to run")

func main() {
    flag.Parse()
    runProblems(*problemFlag)
}

func runProblems(p string) {
    switch p {
    case "1.2":
        Problem1_2()
    case "12.1":
        Problem12_1()
    default:
        runAllProblems()
    }
}

func runAllProblems() {
    Problem1_2()
    Problem12_1()

    // Not actual tests. Remove when finished.
    Tests()
}

func Tests() {
    fmt.Println(" -- Tests -- ")

    sn := &list.Singly{}
    var p *list.SinglyNode

    for i := 0; i < 10; i++ {
        s := &list.SinglyNode{}
        sn.First = s

        if p != nil {
            s.Next = p
        }

        s.Item = i
        p = s
    }

    fmt.Println(p.Search(5).Item)

    sn.Insert(141)
    sn.Insert(140)

    fmt.Println(p.Item)

    sn.Delete(sn.Search(5))

    for p != nil {
        fmt.Println(p.Item)
        p = p.Next
    }
}
