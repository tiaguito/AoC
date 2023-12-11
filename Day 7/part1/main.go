package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
)

type Hand struct {
    hand string
    score int
    kind int
}

var cardsValue = map[byte]int {'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7,
    '9': 8, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14 }

func getKind(str string) int {
    mp := make(map[byte]int, 0)

    for i := 0; i < len(str); i++ {
        mp[str[i]]++
    }
    
    if len(mp) == 1 {
        return 1
    } else if len(mp) == 2 {
        for _,v := range mp {
            if v == 4 {
                return 2
            }
        }
        return 3
    } else if len(mp) == 3 {
        for _,v := range mp {
            if v == 3 {
                return 4
            }
        }
        return 5
    } else if len(mp) == 4 {
        return 6
    } else {
        return 7
    }
}

type byKind []Hand

func (k byKind) Len() int {
    return len(k)
}

func (k byKind) Swap(i, j int) {
    k[i], k[j] = k[j], k[i]
}

func (k byKind) Less(i, j int) bool {
    if k[i].kind < k[j].kind {
        return false
    } else if k[i].kind == k[j].kind {
        for a := 0; a < len(k[i].hand); a++ {
            if cardsValue[k[i].hand[a]] < cardsValue[k[j].hand[a]] {
                return true
            } else if cardsValue[k[i].hand[a]] > cardsValue[k[j].hand[a]] {
                return false
            } else {
                continue
            }
        }
    } 

    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var hands []Hand

    for scanner.Scan() {
        str := scanner.Text()
        hand := Hand{}

        s := strings.Split(str, " ")

        hand.hand = s[0]


        if n,err := strconv.Atoi(strings.Trim(s[1], " ")); err == nil {
            hand.score = n
        }

        hand.kind = getKind(hand.hand)

        hands = append(hands, hand)
    }

    sort.Sort(byKind(hands))
    prod := 0

    for i := 0; i < len(hands); i++ {
       prod = prod + ((i + 1) * hands[i].score)
    }

    fmt.Println(prod)
}
