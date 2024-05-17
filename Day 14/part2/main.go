package main

import (
    "fmt"
    "bufio"
    "os"
)

func tiltNorth(satelite *[]string) {
    m, n := len((*satelite)), len((*satelite)[0])
    for c := 0; c < n; c++ {
        for r := 0; r < m; r++ {
            if (*satelite)[r][c] == 'O' {
                y := r - 1
                if y == -1 {
                    (*satelite)[0] = (*satelite)[0][:c] + "O" + (*satelite)[0][c+1:]
                    continue
                }

                updated := false
                for ; y >= 0; y-- {
                    if (*satelite)[y][c] == '#' || (*satelite)[y][c] == 'O' {
                        (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                        (*satelite)[y+1] = (*satelite)[y+1][:c] + "O" + (*satelite)[y+1][c+1:]
                        updated = true
                        break
                    }
                }

                if !updated {
                    (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                    (*satelite)[0] = (*satelite)[0][:c] + "O" + (*satelite)[0][c+1:]
                }
            }
        }
    }
}

func tiltSouth(satelite *[]string) {
    m, n := len((*satelite)), len((*satelite)[0])
    for c := 0; c < n; c++ {
        for r := m-1; r >= 0; r-- {
            if (*satelite)[r][c] == 'O' {
                y := r + 1
                if y == m {
                    (*satelite)[n-1] = (*satelite)[n-1][:c] + "O" + (*satelite)[n-1][c+1:]
                    continue
                }

                updated := false
                for ; y < m; y++ {
                    if (*satelite)[y][c] == '#' || (*satelite)[y][c] == 'O' {
                        (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                        (*satelite)[y-1] = (*satelite)[y-1][:c] + "O" + (*satelite)[y-1][c+1:]
                        updated = true
                        break
                    }
                }

                if !updated {
                    (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                    (*satelite)[n-1] = (*satelite)[n-1][:c] + "O" + (*satelite)[n-1][c+1:]
                }
            }
        }
    }
}

func tiltWest(satelite *[]string) {
    m, n := len((*satelite)), len((*satelite)[0])
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if (*satelite)[r][c] == 'O' {
                x := c - 1
                if x == -1 {
                    //(*satelite)[0] = "O" + (*satelite)[r][1:]
                    continue
                }

                updated := false
                for ; x >= 0; x-- {
                    if (*satelite)[r][x] == '#' || (*satelite)[r][x] == 'O' {
                        (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                        (*satelite)[r] = (*satelite)[r][:x+1] + "O" + (*satelite)[r][x+2:]
                        updated = true
                        break
                    }
                }

                if !updated {
                    (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                    (*satelite)[r] = "O" + (*satelite)[r][1:]
                }
            }
        }
    }
}

func tiltEast(satelite *[]string) {
    m, n := len((*satelite)), len((*satelite)[0])
    for r := 0; r < m; r++ {
        for c := n-1; c >= 0; c-- {
            if (*satelite)[r][c] == 'O' {
                x := c + 1
                if x == n {
                    //(*satelite)[0] = "O" + (*satelite)[r][1:]
                    continue
                }

                updated := false
                for ; x < n; x++ {
                    if (*satelite)[r][x] == '#' || (*satelite)[r][x] == 'O' {
                        (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                        (*satelite)[r] = (*satelite)[r][:x-1] + "O" + (*satelite)[r][x:]
                        updated = true
                        break
                    }
                }

                if !updated {
                    (*satelite)[r] = (*satelite)[r][:c] + "." + (*satelite)[r][c+1:]
                    (*satelite)[r] = (*satelite)[r][:m-1] + "O"
                }
            }
        }
    }
}


func load_north(satelite []string) int {
    load := 0

    m, n := len(satelite), len(satelite[0])

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if satelite[r][c] == 'O' {
                load += m - r
            }
        }
    }

    return load
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var satelite []string

    m := make(map[string]int)

    for scanner.Scan() {
        str := scanner.Text()
        satelite = append(satelite, str)
    }

    cycle := 1
    var first int
    for ;cycle <= 1000; cycle++ {
        tiltNorth(&satelite)
        tiltWest(&satelite)
        tiltSouth(&satelite)
        tiltEast(&satelite)

        var str string
        for r := 0; r < len(satelite); r++ {
            for c := 0; c < len(satelite[0]); c++ {
                str = fmt.Sprintf(str, satelite[r][c])
            }
        }

        if(m[str] != 0) {
            first = m[str]
            fmt.Printf("Cycle repeating from %d\n", cycle)
            fmt.Printf("Initial found location %d\n", m[str])
            break
        } else {
            m[str] = cycle
        }

        fmt.Printf("Cycle %d has load %d\n\n", cycle, load_north(satelite))
    }

    idx := (1000000000 - first) % (cycle - first) + first
    fmt.Println(idx)

    //fmt.Printf("%d\n", load_north(satelite))
}
