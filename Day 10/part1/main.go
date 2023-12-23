package main

import (
    "fmt"
    "bufio"
    "os"
)

func constructMap(s []string) [][]int {
    m := make([][]int, len(s))

    for i := 0; i < len(s); i++ {
        m[i] = make([]int, len(s[i]))
        for j := 0; j < len(s[i]); j++ {
            if s[i][j] == 'S' {
                m[i][j] = 0
            } else if s[i][j] == '.' {
                m[i][j] = -2
            } else {
                m[i][j] = -1
            }
        }
    }

    return m
}

func getDistances(s []string, mp [][]int, d int) [][]int {
	step := d + 1
    nextStep := false
    m, n := len(s), len(s[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mp[i][j] == d {
                nextStep = true
                switch s[i][j] {
				case 'S':
					if j > 0 && mp[i][j-1] == -1 && (s[i][j-1] == '-' || s[i][j-1] == 'L' || s[i][j-1] == 'F') {
						mp[i][j-1] = step
					}
					if j < len(mp[i])-1 && mp[i][j+1] == -1 && (s[i][j+1] == '-' || s[i][j+1] == '7' || s[i][j+1] == 'J') {
						mp[i][j+1] = step
					}
					if i > 0 && mp[i-1][j] == -1 && (s[i-1][j] == '|' || s[i-1][j] == '7' || s[i-1][j] == 'F') {
						mp[i-1][j] = step
					}
					if i < len(mp)-1 && mp[i+1][j] == -1 && (s[i+1][j] == '|' || s[i+1][j] == 'J' || s[i+1][j] == 'L') {
						mp[i+1][j] = step
					}
				case '-':
					if j > 0 && mp[i][j-1] == -1 && (s[i][j-1] == '-' || s[i][j-1] == 'L' || s[i][j-1] == 'F') {
						mp[i][j-1] = step
					}
					if j < len(mp[i])-1 && mp[i][j+1] == -1 && (s[i][j+1] == '-' || s[i][j+1] == '7' || s[i][j+1] == 'J') {
						mp[i][j+1] = step
					}
				case '|':
					if i > 0 && mp[i-1][j] == -1 && (s[i-1][j] == '|' || s[i-1][j] == '7' || s[i-1][j] == 'F') {
						mp[i-1][j] = step
					}
					if i < len(mp)-1 && mp[i+1][j] == -1 && (s[i+1][j] == '|' || s[i+1][j] == 'J' || s[i+1][j] == 'L') {
						mp[i+1][j] = step
					}
				case '7':
					if j > 0 && mp[i][j-1] == -1 && (s[i][j-1] == '-' || s[i][j-1] == 'L' || s[i][j-1] == 'F') {
						mp[i][j-1] = step
					}
					if i < len(mp)-1 && mp[i+1][j] == -1 && (s[i+1][j] == '|' || s[i+1][j] == 'J' || s[i+1][j] == 'L') {
						mp[i+1][j] = step
					}
				case 'F':
					if j < len(mp[i])-1 && mp[i][j+1] == -1 && (s[i][j+1] == '-' || s[i][j+1] == '7' || s[i][j+1] == 'J') {
						mp[i][j+1] = step
					}
					if i < len(mp)-1 && mp[i+1][j] == -1 && (s[i+1][j] == '|' || s[i+1][j] == 'J' || s[i+1][j] == 'L') {
						mp[i+1][j] = step
					}
				case 'J':
					if j > 0 && mp[i][j-1] == -1 && (s[i][j-1] == '-' || s[i][j-1] == 'L' || s[i][j-1] == 'F') {
						mp[i][j-1] = step
					}
					if i > 0 && mp[i-1][j] == -1 && (s[i-1][j] == '|' || s[i-1][j] == '7' || s[i-1][j] == 'F') {
						mp[i-1][j] = step
					}
				case 'L':
					if j < len(mp[i])-1 && mp[i][j+1] == -1 && (s[i][j+1] == '-' || s[i][j+1] == '7' || s[i][j+1] == 'J') {
						mp[i][j+1] = step
					}
					if i > 0 && mp[i-1][j] == -1 && (s[i-1][j] == '|' || s[i-1][j] == '7' || s[i-1][j] == 'F') {
						mp[i-1][j] = step
					}
                }
            }
        }
    }

    if nextStep {
        return getDistances(s, mp, step)
    }

    return mp
}

func getFurthestDistance(s []string, mp [][]int) int {
	maxDist := 0
    mp = getDistances(s, mp, 0)

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if maxDist < mp[i][j] {
				maxDist = mp[i][j]
			}	
		}
	}
    return maxDist
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    s := []string{}

    for scanner.Scan() {
        str := scanner.Text()

        s = append(s, str)
    }

    mp := constructMap(s)

    fmt.Println(getFurthestDistance(s, mp))
}
