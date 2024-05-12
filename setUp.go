package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Node struct {
	x, y, val int
	// race      int // 0, 1
	// blood     int // 0 to 100
}

func main() {

	node1 := Node{1, 1, 0}
	node2 := Node{4, 4, 0}
	var output = floorDistAB(node1, node2, 4.0)
	fmt.Println(output)

	matrix := generateAndPopulateMatrix(5, 5, 4)
	adjacencyList := matrizToListAdja(matrix, 5.0)
	printMatrix(matrix)

	if true {
		fmt.Println("Lista de Adjacência:")
		for vertex, neighbors := range adjacencyList {
			fmt.Printf("%d -> %v\n", vertex, neighbors)
		}
	}

	i, j := indicesFromListIndex(10, len(matrix[0])) //exemplo de como acessar node a partir da lista
	fmt.Println(matrix[i][j].val)

}

func indicesFromListIndex(index, cols int) (int, int) {
	i := index / cols
	j := index % cols
	return i, j
}

func generateAndPopulateMatrix(rows, cols, densidade int) [][]Node {
	rand.Seed(time.Now().UnixNano())

	var populated int

	// Create the matrix
	matrix := make([][]Node, rows)
	for i := range matrix {
		matrix[i] = make([]Node, cols)
		for j := range matrix[i] {
			if rand.Intn(densidade) == 0 { //func rand[0,n) fechado aberto
				populated = 1
			} else {
				populated = 0
			}

			matrix[i][j] = Node{x: i, y: j, val: populated}
		}
	}

	return matrix
}

func matrizToListAdja(matrix [][]Node, raioConexao float64) map[int][]int {
	adjacencyList := make(map[int][]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for m := i; m <= i+int(raioConexao); m++ {
				for n := j; n <= j+int(raioConexao); n++ {
					// Verifica se o ponto (m, n) é válido e se não é o ponto (i, j) original
					valid := isValidPoint(matrix, m, n)
					if valid && (m != i || n != j) && matrix[m][n].val == 1 && matrix[i][j].val == 1 {
						adjacencyList[i*len(matrix[0])+j] = append(adjacencyList[i*len(matrix[0])+j], m*len(matrix[0])+n)
					}
				}
			}
		}
	}

	return adjacencyList
}

func printMatrix(matrix [][]Node) {
	for _, row := range matrix {
		for _, node := range row {
			fmt.Printf("%d  ", node.val)
		}
		fmt.Println()
	}
}

func isValidPoint(matrix [][]Node, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0])
}

func floorDistAB(pA Node, pB Node, x float64) float64 {
	ax, bx, ay, by := float64(pA.x), float64(pB.x), float64(pA.y), float64(pB.y)
	dist := math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
	return dist
}

func printCommaMatrix(matrix [][]int) {
	for _, row := range matrix {
		rowStrings := make([]string, len(row))
		for i, val := range row {
			rowStrings[i] = fmt.Sprintf("%d", val)
		}
		fmt.Println(strings.Join(rowStrings, ","))
	}
}
