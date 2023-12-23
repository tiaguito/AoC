package main

import (
    "fmt"
    "bufio"
    "os"
	"regexp"
	"strings"
)

func constructMap(s []string) ([][]int, int, int) {
    m := make([][]int, len(s))
    var sx, sy int

    for i := 0; i < len(s); i++ {
        m[i] = make([]int, len(s[i]))
        for j := 0; j < len(s[i]); j++ {
            if s[i][j] == 'S' {
                sx = i
                sy = j
                m[i][j] = 0
            } else if s[i][j] == '.' {
                m[i][j] = -2
            } else {
                m[i][j] = -1
            }
        }
    }

    return m, sx, sy
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

func getStartSymbol(s []string, sy, sx int) uint8 {
	var (
		up    uint8 = '.'
		down  uint8 = '.'
		left  uint8 = '.'
		right uint8 = '.'
	)
	if sy > 0 {
		up = s[sy-1][sx]
	}
	if sy < len(s)-1 {
		down = s[sy+1][sx]
	}
	if sx > 0 {
		left = s[sy][sx-1]
	}
	if sx < len(s[0])-1 {
		right = s[sy][sx+1]
	}
	upCond := up == '7' || up == 'F' || up == '|'
	downCond := down == 'J' || down == 'L' || down == '|'
	leftCond := left == 'F' || left == 'L' || left == '-'
	rightCond := right == 'J' || right == '7' || right == '-'

	if upCond && downCond {
		return '|'
	}
	if leftCond && rightCond {
		return '-'
	}
	if leftCond && downCond {
		return '7'
	}
	if rightCond && downCond {
		return 'F'
	}
	if rightCond && upCond {
		return 'L'
	}
	if leftCond && upCond {
		return 'J'
	}
	return 'S'
}

func getEnclosedTiles(s []string) int {
    mp,sx,sy := constructMap(s)

    mp = getDistances(s, mp, 0)

    s[sy] = strings.Replace(s[sy], "S", string(getStartSymbol(s, sy, sx)), -1)

	for y := 0; y < len(s); y++ {
		newS := make([]byte, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			if mp[y][x] < 0 {
				newS[x] = '.'
			} else {
				newS[x] = s[y][x]
			}
		}
		s[y] = string(newS)
	}
	noWall := regexp.MustCompile(`F-*7|L-*J`)
	wall := regexp.MustCompile(`F-*J|L-*7`)
	for i, ss := range s {
		s1 := noWall.ReplaceAllString(ss, " ")
		s2 := wall.ReplaceAllString(s1, "|")
		s[i] = s2
	}
	var (
		parity int
		count  int = 0
	)
	for y := 0; y < len(s); y++ {
		parity = 0
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == '|' {
				parity++
			}
			if s[y][x] == '.' && parity%2 == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    s := []string{}

    for scanner.Scan() {
        str := scanner.Text()

        s = append(s, str)
    }

    fmt.Println(getEnclosedTiles(s))
}
