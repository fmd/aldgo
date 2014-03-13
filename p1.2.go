package main

import "fmt"
import "sort"
import "strconv"

type Interval struct {
    Label string
    Start int
    End int
}

type IntervalSet []Interval

func (slice IntervalSet) Len() int {
    return len(slice)
}

func (slice IntervalSet) Less (i, j int) bool {
    return slice[i].End < slice[j].End
}

func (slice IntervalSet) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func DeleteIntersects(I IntervalSet, r Interval) IntervalSet {
    tmp := IntervalSet{}

    for i := range I {
        el := I[i]
        if !(el.Start < r.Start &&
             el.End > r.Start   ||
             el.End > r.End     &&
             el.Start < r.End   ||
             el == r) {

            tmp = append(tmp, el)
        }
    }

    return tmp
}

func rightJobs (I IntervalSet) IntervalSet {
    sort.Sort(I)
    subset := IntervalSet{}

    for len(I) > 0 {
        el := I[0]
        subset = append(subset, el)
        I = DeleteIntersects(I, el)
    }

    return subset
}

func Problem1_2() {
    fmt.Println(" -- Problem 1.2 -- ")

    problemSet := IntervalSet{
        Interval{"The President's Algorist", 0, 5},
        Interval{"Discrete Mathematics", 1, 3},
        Interval{"Tarjan of the Jungle", 2, 7},
        Interval{"Halting State", 4, 8},
        Interval{"Steiner's Tree", 6, 10},
        Interval{"The Four Volume Problem", 9, 15},
        Interval{"Programming Challenges", 11, 13},
        Interval{"Process Terminated", 12, 16},
        Interval{"Calculated Bets", 14, 17},
    }

    subset := rightJobs(problemSet)
    for i := range subset {
        s := subset[i]
        fmt.Println(s.Label, ": ", strconv.Itoa(s.Start), " -> ", strconv.Itoa(s.End))
    }
}
