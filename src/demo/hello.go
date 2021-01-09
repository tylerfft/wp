package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{2, 8, 4, 6}
	B := []int{7, 9, 11, 5}
	fmt.Println(Shuffle(A, B))

}

// give two arrays of random integers with same length A, B
// Shuffle A to maximize the count of A[i] > b[i] (i => 0 .. N - 1)

//
//A =  [2,  8,   4,  6],
//B =  [7,  9,  11 , 5]
//
//A' = [8 ,  2,  4,  6]
//A' = [8 ,  4,  2,  6]

//vector<int> Shuffle(vector<int> A, vector<int> B){
//
//
//}

func Shuffle(A []int, B []int) []int {
	if len(A) != len(B) {
		return nil
	}
	n := len(A)
	if n == 0 {
		return nil
	}
	Aidx := make([]int, n)
	Bidx := make([]int, n)
	for i := 0; i < len(A); i++ {
		Aidx[i] = i
		Bidx[i] = i
	}
	rst := make([]int, n)
	sort.Slice(Aidx, func(i, j int) bool {
		return A[Aidx[i]] < A[Aidx[j]]
	})
	sort.Slice(Bidx, func(i, j int) bool {
		return B[Bidx[i]] < B[Bidx[j]]
	})

	idxa_left, idxa_right, idxb := 0, n-1, n-1
	for idxa_right >= idxa_left {
		if A[Aidx[idxa_right]] <= B[Bidx[idxb]] {
			rst[Bidx[idxb]] = A[Aidx[idxa_left]]
			idxa_left++
		} else {
			rst[Bidx[idxb]] = A[Aidx[idxa_right]]
			idxa_right--
		}
		idxb--
	}

	return rst
}
