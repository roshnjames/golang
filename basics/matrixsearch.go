package main

import "fmt"

func main() {
	var n, m, ele, loc, loc1 int
	var arr [20][20]int
	var flag bool
	fmt.Println("Enter the no of rows for the matrix")
	fmt.Scanln(&n)
	fmt.Println("Enter the no of columns for the matrix")
	fmt.Scanln(&m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Println("\nENter the element :")
			fmt.Scanln(&arr[i][j])
		}
	}

	fmt.Println("ENtered Matrix is :")
	for i := 0; i < n; i++ {
		fmt.Println()
		for j := 0; j < m; j++ {
			fmt.Printf("%d \t", arr[i][j])
		}
	}

	fmt.Println("\nENter the element you want to search for:")
	fmt.Scanln(&ele)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == ele {
				flag = !flag
				loc, loc1 = i+1, j+1
				break
			}
		}
	}

	if flag {
		fmt.Printf("\n\nElement found at location :ROw:%d COLUMN:%d", loc, loc1)
	} else {
		fmt.Println("\nElement is absent in the matrix")
	}

}
