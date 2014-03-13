package main

import "flag"

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
}
