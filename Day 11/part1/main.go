package main

import (
    "fmt"
	"bufio"
	"os"
)

var DX = [4]int{-1, 0, 1, 0}
var DY = [4]int{0, -1, 0, 1}

func transposeUniverse(universe []string) []string{
    var resultUniverse []string
    for i := 0; i < len(universe[0]); i++ {
        s := ""
        for j := 0; j < len(universe); j++ {
            s += string(universe[j][i])
        }
        resultUniverse = append(resultUniverse, s)
    }

    return resultUniverse
}

func expandUniverse(universe []string) []string {
    var expandedUniverse []string

    for i := 0; i < len(universe); i++ {
        isEmpty := true
        for j := 0; j < len(universe[i]); j++ {
            if universe[i][j] == '#' {
                isEmpty = false
                break
            }
        }

        if isEmpty {
            expandedUniverse = append(expandedUniverse, universe[i])
        }
        expandedUniverse = append(expandedUniverse, universe[i])
    }

    return expandedUniverse
}

func isValid(x, y, rowsUniverse, colsUniverse int) bool {
    if (x >= 0 && x < rowsUniverse) && (y >= 0 && y < colsUniverse){
        return true
    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    universe := []string{}

    for scanner.Scan() {
        str := scanner.Text()
        universe = append(universe, str)
    }

    universe = expandUniverse(universe)
    universe = transposeUniverse(universe)
    universe = expandUniverse(universe)
    universe = transposeUniverse(universe)

    rowsUniverse := len(universe)
    colsUniverse := len(universe[0])

    galaxies := [][]int{}

    for i := 0; i < len(universe); i++ {
        for j := 0; j < len(universe[i]); j++ {
            if universe[i][j] == '#' {
                galaxies = append(galaxies, []int{i,j})
            }
        }
    }

    n := len(galaxies)

    sum := 0
    for i := 1; i <= n; i++ {
        visited := make([][]bool, rowsUniverse)
        distances := make([][]int, rowsUniverse)
        for x := 0; x < rowsUniverse; x++ {
            distances[x] = make([]int, colsUniverse)
            visited[x] = make([]bool, colsUniverse)
        }

        queue := [][]int{}
        queue = append(queue, galaxies[i-1])
        visited[galaxies[i-1][0]][galaxies[i-1][1]] = true

        for len(queue) > 0 {
            node := queue[0]
            x,y := node[0], node[1]

            for j := 0; j < 4; j++ {
                if isValid(x + DX[j], y + DY[j], rowsUniverse, colsUniverse) && !visited[x + DX[j]][y + DY[j]]  {
                    queue = append(queue, []int{x + DX[j], y + DY[j]})
                    visited[x + DX[j]][y + DY[j]] = true
                    distances[x + DX[j]][y + DY[j]] = distances[x][y] + 1
                }
            }
            queue = queue[1:]
        }

        for j := i + 1; j <= n; j++ {
            coordinates := galaxies[j-1] 
            sum += distances[coordinates[0]][coordinates[1]]
        }
    }

    fmt.Println(sum)
}
