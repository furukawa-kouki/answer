package main

import (
	"fmt"
	"strconv"
)

type N struct {
	A int
	B int
	C int
}

func isOK(n N) bool {
	if n.A == n.B {
		return false
	}
	if n.A == n.C {
		return false
	}
	if n.B == n.C {
		return false
	}

	if n.A < n.B && n.B < n.C {
		return false
	}

	if n.A > n.B && n.B > n.C {
		return false
	}

	return true
}

func isMinus(n N) bool {
	if n.A > 0 && n.B > 0 && n.C > 0 {
		return false
	}
	return true
}

func main() {

	// ### 問 2
	// ３つの自然数 A,B,C が与えられたときに、それぞれの数字が異なり２番目に大きい整数が A または C である状態を作りたい。
	// 以下の操作を任意回行えるとしたとき、上述の条件を満たすのに最小の操作回数を出力せよ。
	// なお、条件を満たすことができない場合は -1 を出力せよ。

	// - A = A - 1
	// - B = B - 1
	// - C = C - 1

	// この問題では１つの入力に対して複数のクエリが与えられます。

	// #### 制約
	// - 1 <= T <= 10000
	// - 1 <= Ai,Bi,Ci <= 10^9

	// #### 入力形式
	// T
	// A1,B1,C1
	// ...
	// AT,BT,CT

	// go run .\problem2.go
	var T int
	fmt.Print("行数:")
	fmt.Scan(&T)

	var n [10000]N

	var ans = -1
	for index := 0; index < T; index++ {
		fmt.Print("自然数A" + strconv.Itoa(index+1) + ":")
		fmt.Scan(&n[index].A)
		fmt.Print("自然数B" + strconv.Itoa(index+1) + ":")
		fmt.Scan(&n[index].B)
		fmt.Print("自然数C" + strconv.Itoa(index+1) + ":")
		fmt.Scan(&n[index].C)
	}

	for index := 0; index < T; index++ {
		if n[index].A == n[index].B && n[index].A == n[index].C && n[index].B == n[index].C {
			fmt.Println(3)
			continue
		}

		if n[index].A != n[index].C && n[index].A < n[index].B && n[index].C < n[index].B {
			fmt.Println(0)
			continue
		}

		if n[index].A != n[index].C && n[index].A > n[index].B && n[index].C > n[index].B {
			fmt.Println(0)
			continue
		}

		var spcount = 0
		if n[index].A == n[index].C && n[index].A > 1 {
			n[index].A = n[index].A - 1
			spcount = 1
			if isOK(n[index]) {
				fmt.Println(1)
				continue
			}
		}

		var countA = 0
		var countB = 0
		var countC = 0

		var nA = n[index]
		for {
			nA.A = nA.A - 1
			countA = countA + 1
			if isMinus(nA) {
				countA = -1
				break
			}

			if isOK(nA) {
				break
			}
		}

		var nB = n[index]
		for {
			nB.B = nB.B - 1
			countB = countB + 1
			if isMinus(nB) {
				countB = -1
				break
			}
			if isOK(nB) {
				break
			}
		}

		var nC = n[index]
		for {
			nC.C = nC.C - 1
			countC = countC + 1
			if isMinus(nC) {
				countC = -1
				break
			}
			if isOK(nC) {
				break
			}
		}

		if countA != -1 {
			ans = countA
		}
		if ans == -1 || (countB != -1 && countB < ans) {
			ans = countB
		}
		if ans == -1 || (countC != -1 && countC < ans) {
			ans = countC
		}

		fmt.Println(ans + spcount)
		ans = -1

	}

}
