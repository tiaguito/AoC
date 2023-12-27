package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func expandUniverse(universe []string, galaxies [][]int, scalingFactor int) [][]int {
    columnsWithGalaxies := make(map[int]bool)
    rowsWithGalaxies := make(map[int]bool)
    
    for _, galaxy := range galaxies {
        rowsWithGalaxies[galaxy[0]] = true
        columnsWithGalaxies[galaxy[1]] = true
    }

    minRow, maxRow := getKeyRange(rowsWithGalaxies)
    minCol, maxCol := getKeyRange(columnsWithGalaxies)

    var rowsWithoutGalaxies []int
    for row := minRow; row <= maxRow; row++ {
        if !rowsWithGalaxies[row] {
            rowsWithoutGalaxies = append(rowsWithoutGalaxies, row)
        }
    }

    var columnsWithoutGalaxies []int
    for col := minCol; col <= maxCol; col++ {
        if !columnsWithGalaxies[col] {
            columnsWithoutGalaxies = append(columnsWithoutGalaxies, col)
        }
    }

    var expanded [][]int

    for _, galaxy := range galaxies {
        var e []int
        e = append(e, galaxy...)
        for _, row := range rowsWithoutGalaxies {
            if row > galaxy[0] {
                break
            }
            e[0] += scalingFactor - 1
        }

        for _, col := range columnsWithoutGalaxies {
            if col > galaxy[1] {
                break
            }
            e[1] += scalingFactor - 1
        }
        expanded = append(expanded, e)
    }
    
    return expanded
}

func getKeyRange(m map[int]bool) (int,int) {
    min, max := math.MaxInt, math.MinInt

    for k := range m {
        if k < min {
            min = k
        }
        if k > max {
            max = k
        }
    }
    return min, max
}

func sumOfDistances(galaxies [][]int) int {
    var sum int

    for i := range galaxies {
        for j := i + 1; j < len(galaxies); j++ {
            sum += distance(galaxies[i], galaxies[j])
        }
    }

    return sum
}

func distance(a, b []int) int {
    return abs(a[0] - b[0]) + abs(a[1] - b[1])
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    universe := []string{}

    for scanner.Scan() {
        str := scanner.Text()
        universe = append(universe, str)
    }

    galaxies := [][]int{}

    for i := 0; i < len(universe); i++ {
        for j := 0; j < len(universe[i]); j++ {
            if universe[i][j] == '#' {
                galaxies = append(galaxies, []int{i,j})
            }
        }
    }

    expandedGalaxies := expandUniverse(universe, galaxies, 1000000)

    sum := sumOfDistances(expandedGalaxies)

    fmt.Println(sum)
}
