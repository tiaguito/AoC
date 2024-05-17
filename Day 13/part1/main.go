package main

import(
    "fmt"
    "bufio"
    "os"
)

func checkReflectionAcrossVerticalAxis(mp []string, left int, right int) bool {
    for ; left >= 0 && right < len(mp[0]); left, right = left - 1, right + 1 {
        if mp[0][left] == mp[0][right] {
            for c := 0; c < len(mp); c++ {
                if mp[c][left] != mp[c][right] {
                    return false
                }
            }
        } else {
            return false
        }
    }

    return true
}

func hasVerticalReflection(mp []string) (bool, int) {
    found := true

    for r := 0; r < len(mp[0]) - 1; r++ {
        if mp[0][r] == mp[0][r+1] {
            for c := 1; c < len(mp); c++ {
                if mp[c][r] != mp[c][r+1] {
                    found = false
                    break
                }
            }
            if found {
                if checkReflectionAcrossVerticalAxis(mp, r - 1, r + 2) {
                    return true, r + 1
                } 
            } else {
                // reset found variable
                found = true
            }
        }
    }

    return false, 0
}

func checkReflectionAcrossHorizontalAxis(mp []string, up int, down int) bool {
    for ; up >= 0 && down < len(mp); up, down = up - 1, down + 1 {
        if mp[up][0] == mp[down][0] {
            for r := 0; r < len(mp[0]); r++ {
                if mp[up][r] != mp[down][r] {
                    return false
                }
            }
        } else {
            return false
        }
    }

    return true
}

func hasHorizontalReflection(mp []string) (bool, int) {
    found := true

    for c := 0; c < len(mp) - 1; c++ {
        if mp[c][0] == mp[c+1][0] {
            for r := 1; r < len(mp[0]); r++ {
                if mp[c][r] != mp[c+1][r] {
                    found = false
                    break
                }
            }
            if found {
                if checkReflectionAcrossHorizontalAxis(mp, c - 1, c + 2) {
                    return true, c + 1
                } 
            } else {
                // reset found variable
                found = true
            }
        }
    }

    return false, 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    var mp []string

    for i := 1; scanner.Scan(); i++ {
        str := scanner.Text()

        if len(str) == 0 {
            found := false
            x, y := 0, 0

            if found, x = hasVerticalReflection(mp); found {
                fmt.Printf("Found vertical reflection at points (%d, %d)\n", x, x + 1)                
            } 

            if found, y = hasHorizontalReflection(mp); found {
                fmt.Printf("Found horizontal reflection at points (%d, %d)\n", y, y + 1)                
            }

            sum = sum + int(max(float64(x), 100.0 * float64(y)))

            mp = nil
        } else {
            mp = append(mp, str)
        }
    }

    found := false
    x, y := 0, 0

    if found, x = hasVerticalReflection(mp); found {
        fmt.Printf("Found vertical reflection at points (%d, %d)\n", x, x + 1)                
    } 

    if found, y = hasHorizontalReflection(mp); found {
        fmt.Printf("Found horizontal reflection at points (%d, %d)\n", y, y + 1)                
    }

    sum = sum + int(max(float64(x), 100.0 * float64(y)))

    mp = nil


    fmt.Println(sum)
}
