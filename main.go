package main

import "fmt"
import "flag"
import "strconv"

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
        problem1_2()
    default:
        runAllProblems()
    }
}

func runAllProblems() {
    problem1_2()
}

/** 
 *  Problem 1.2 - Selecting the Right Jobs
 */
type Interval struct {
    Label string
    Start int
    End int
}

func rightJobs(I []Interval) []Interval {
    subset := []Interval{}
    return subset
}

func problem1_2() {

    // Set up an example problem set and test the method.
    problemSet := []Interval{}
    problemSet = append(problemSet, Interval{"The President's Algorist", 0, 5})
    problemSet = append(problemSet, Interval{"Discrete Mathematics", 1, 3})
    problemSet = append(problemSet, Interval{"Tarjan of the Jungle", 2, 7})
    problemSet = append(problemSet, Interval{"Halting State", 4, 8})
    problemSet = append(problemSet, Interval{"Steiner's Tree", 6, 10})
    problemSet = append(problemSet, Interval{"The Four Volume Problem", 9, 15})
    problemSet = append(problemSet, Interval{"Programming Challenges", 11, 13})
    problemSet = append(problemSet, Interval{"Process Terminated", 12, 16})
    problemSet = append(problemSet, Interval{"Calculated Bets", 14, 17})

    subset := rightJobs(problemSet)
    for idx := range subset {
        s := subset[idx]
        fmt.Println(s.Label, ": ", strconv.Itoa(s.Start), " -> ", strconv.Itoa(s.End))
    }
}
