package main

import "flag"
import "fmt"
import "os"
import "strconv"

type change struct {
    foo, bar, qix, baz int
}

func (self change) String() string {
    return fmt.Sprintf("{foo:%2d, bar:%2d, qix:%2d, baz:%2d} = %d",
        self.foo, self.bar, self.qix, self.baz,
        self.foo + self.bar * 7 + self.qix * 11 + self.baz * 21)
}
func (self *change) update(other *change) {
    if other.foo > 0 {
        self.foo = other.foo
    }
    if other.bar > 0 {
        self.bar = other.bar
    }
    if other.qix > 0 {
        self.qix = other.qix
    }
    if other.baz > 0 {
        self.baz = other.baz
    }
}
func newChange(coin, count int) *change {
    var c = &change{}
    if coin == 1 {
        c.foo = count
    } else if coin == 7 {
        c.bar = count
    } else if coin == 11 {
        c.qix = count
    } else if coin == 21 {
        c.baz = count
    }
    return c
}

func main() {
    flag.Parse()
    var arg = flag.Arg(0)
    var total, _ = strconv.Atoi(arg)

    var coins = []int{1, 7, 11, 21}
    nb_combinations(coins, total)
    var distinct_changes = changes(coins, total)
    for _, change0 := range distinct_changes {
        debug(change0)
    }
    debug(len(distinct_changes))
}

func changes(coins []int, total int) []*change {
    var len_coins = len(coins)
    var compositions = []*change{}
    var i, j int

    for i = len_coins - 1; i >= 0; i-- {
        var coin = coins[i]

        if coin > total {
            continue

        } else if coin == 1 {
            compositions = append(compositions, newChange(coin, total))
            break
        }

        // 2 = 17 / 7
        // 3 = 17 - (7 * 2)
        var nb_coins = total / coin
        for j = 1; j <= nb_coins; j++ {
            var modulus = total - (coin * j)

            if modulus == 0 {
                compositions = append(compositions, newChange(coin, j))
                break

            } else {
                var sub_changes = changes(coins[:i], modulus)
                for _, sub_change := range sub_changes {
                    var change0 = newChange(coin, j)
                    change0.update(sub_change)
                    compositions = append(compositions, change0)
                }
            }
        }
    }
    return compositions
}

func nb_combinations(coins []int, total int) int {
    var len_coins = len(coins)
    var lookupTable = make([][]int, total+1)
    var amount, iCoin int

    for ; amount <= total; amount++ {
        lookupTable[amount] = make([]int, len_coins+1)

        for iCoin = 0; iCoin <= len_coins; iCoin++ {
            if amount == 0 {
                lookupTable[amount][iCoin] = 1

            } else if iCoin == 0 {
                lookupTable[amount][iCoin] = 0

            } else if coins[iCoin-1] > amount {
                lookupTable[amount][iCoin] = lookupTable[amount][iCoin-1]

            } else {
                // [7 [0 1 2 2 2]]
                //       [6][1] + [7][0] = 1 + 0
                //         [0][2] + [7][1] = 1 + 1
                var a = lookupTable[amount-coins[iCoin-1]][iCoin]
                lookupTable[amount][iCoin] = a + lookupTable[amount][iCoin-1]
            }
        }
    }

    for i, row := range lookupTable {
        debug(i, row)
    }
    return lookupTable[total][len_coins]
}

func debug(args ...interface{}) {
    fmt.Fprintln(os.Stderr, args)
}
