package main

import(
    "fmt"
    "bufio"
    "os"
)

func checkReflectionAcrossVerticalAxis(mp []string, left int, right int, diff int) bool {
    for ; left >= 0 && right < len(mp[0]); left, right = left - 1, right + 1 {
        for c := 0; c < len(mp); c++ {
            if mp[c][left] != mp[c][right] {
                diff++
            }
        }
    }

    if diff == 1 {
        return true
    } else {
        return false
    }
}

func hasVerticalReflection(mp []string) (bool, int) {
    found := true
    diff := 0

    for r := 0; r < len(mp[0]) - 1; r++ {
        for c := 0; c < len(mp); c++ {
            if mp[c][r] != mp[c][r+1] {
                if diff > 1 {
                    found = false
                    break
                } else {
                    diff++
                }
            } 
        }
        if found {
            if checkReflectionAcrossVerticalAxis(mp, r - 1, r + 2, diff) {
                return true, r + 1
            } 
        } else {
            // reset found variable
            found = true
        }
        diff = 0
    }

    return false, 0
}

func checkReflectionAcrossHorizontalAxis(mp []string, up int, down int, diff int) bool {
    for ; up >= 0 && down < len(mp); up, down = up - 1, down + 1 {
        for r := 0; r < len(mp[0]); r++ {
            if mp[up][r] != mp[down][r] {
                diff++
            }
        }
    }

    if diff == 1 {
        return true
    } else {
        return false
    }
}

func hasHorizontalReflection(mp []string) (bool, int) {
    found := true
    diff := 0

    for c := 0; c < len(mp) - 1; c++ {
        for r := 0; r < len(mp[0]); r++ {
            if mp[c][r] != mp[c+1][r] {
                if diff > 1 {
                    found = false
                    break
                } else {
                    diff++
                }
            }
        }
        if found {
            if checkReflectionAcrossHorizontalAxis(mp, c - 1, c + 2, diff) {
                return true, c + 1
            } 
        } else {
            // reset found variable
            found = true
        }
        diff = 0
    }

    return false, 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    var mp []string

    bp := 1
    for i := 1; scanner.Scan(); i++ {
        str := scanner.Text()

        if len(str) == 0 {
            found := false
            x, y := 0, 0

            if found, x = hasVerticalReflection(mp); found {
                fmt.Printf("On map %d, Found vertical reflection at points (%d, %d)\n", bp, x, x + 1)                
            } 

            if found, y = hasHorizontalReflection(mp); found {
                fmt.Printf("On map %d, Found horizontal reflection at points (%d, %d)\n", bp, y, y + 1)                
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
        fmt.Printf("On map %d, Found vertical reflection at points (%d, %d)\n", bp, x, x + 1)                
    } 

    if found, y = hasHorizontalReflection(mp); found {
        fmt.Printf("On map %d, Found horizontal reflection at points (%d, %d)\n", bp, y, y + 1)                
    }

    sum = sum + int(max(float64(x), 100.0 * float64(y)))

    mp = nil

    fmt.Println(sum)
}
