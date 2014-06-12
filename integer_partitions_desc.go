package main

import "fmt"
import "flag"
import "os"
import "strconv"

// http://www.geeksforgeeks.org/generate-unique-partitions-of-an-integer/
func partitions(n int) {
    var nums = make([]int, n)
    nums[0] = n
    var k = 0

    for {
        debug(nums[:k+1])

        var rem int
        for k >= 0 && nums[k] == 1 {
            rem += nums[k]
            k -= 1
        }
        // debug("  k:", k, "rem:", rem)

        if k < 0 {
            break
        }

        nums[k] -= 1
        rem += 1

        // uncomment and get unique permutations
        // for rem > nums[k] {
            // nums[k+1] = nums[k]
            // rem -= nums[k]
            // k += 1
        // }

        nums[k+1] = rem
        k += 1
        // debug("  k:", k, "rem:", rem)
    }
}

func main() {
    flag.Parse()
    var arg = flag.Arg(0)
    var n, _ = strconv.Atoi(arg)
    partitions(n)
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
