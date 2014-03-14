package main

import "fmt"
import "flag"
import "github.com/fmd/aldgo/structures/list"

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

    var p *list.Singly

    for i := 0; i < 10; i++ {
        s := &list.Singly{}

        if p != nil {
            s.Next = p
        }

        s.Item = i
        p = s
    }

    fmt.Println(p.Search(5).Item)

    list.Insert(&p, 141)
    list.Insert(&p, 140)


    fmt.Println(p.Item)

    list.Delete(&p, 5)

    for p != nil {
        fmt.Println(p.Item)
        p = p.Next
    }
}
