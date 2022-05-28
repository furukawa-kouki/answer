package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ### 問 1
	// あなたは濃度 N%の水溶液を作ろうとしています。
	// 手元には濃度 A_1%,A_2%,A_3%...A_n%の水溶液がそれぞれ十分な量あるとします。
	// これらを用いて N%の水溶液を作成することができるか判定してください。

	// #### 制約
	// - 水溶液の数:1 <= n <= 100
	// - 目標の濃度:1 <= N <= 100
	// - 水溶液の濃度:1 <= A1 < ... < An <= 99
	// - N、n、A1、...、An は整数とする

	// #### 入力形式
	// n N
	// A1 A2 ... An

	// #### 出力形式
	// "Yes" or "No"を出力してください。

	// go run .\problem1.go
	var n int
	fmt.Print("水溶液の数:")
	fmt.Scan(&n)

	var N int
	fmt.Print("目標の濃度:")
	fmt.Scan(&N)

	var A [100]int

	var ans = "No"
	for index := 0; index < n; index++ {
		fmt.Print("水溶液" + strconv.Itoa(index+1) + "の濃度:")
		// fmt.Scan(&A[index])
		A[index] = index + 3
		fmt.Println(A[index])
		if A[index] >= N {
			ans = "Yes"
		}
	}
	fmt.Println(ans)

}
