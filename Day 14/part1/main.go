package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var satelite []string

    for scanner.Scan() {
        str := scanner.Text()
        satelite = append(satelite, str)
    }

    m, n := len(satelite), len(satelite[0])
    sum := 0
    for c := 0; c < n; c++ {
        for r := 0; r < m; r++ {
            if satelite[r][c] == 'O' {
                y := r - 1
                if y == -1 {
                    sum += m
                    satelite[0] = satelite[0][:c] + "O" + satelite[0][c+1:]
                    continue
                }

                updated := false
                for ; y >= 0; y-- {
                    if satelite[y][c] == '#' || satelite[y][c] == 'O' {
                        sum += m - (y + 1)
                        satelite[r] = satelite[r][:c] + "." + satelite[r][c+1:]
                        satelite[y+1] = satelite[y+1][:c] + "O" + satelite[y+1][c+1:]
                        updated = true
                        break
                    }
                }

                if !updated {
                    sum += m
                    satelite[r] = satelite[r][:c] + "." + satelite[r][c+1:]
                    satelite[0] = satelite[0][:c] + "O" + satelite[0][c+1:]
                }
            }
        }
    }

    for r := 0; r < m; r++ {
        fmt.Println(satelite[r])
    }

    fmt.Println(sum)
}
