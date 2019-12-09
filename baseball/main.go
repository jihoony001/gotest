package main

import (
	"fmt"
	"time"
	"math/rand"
)

 func main() {
	 rand.Seed(time.Now().UnixNano())
	 // 무작위 숫자 3개를 만든다
	 numbers := MakeNumbers()

	cnt := 0

	for {
		cnt++
		// 사용자의 입력을 받는다
		inputNumbers := InputNumbers()

		// 결과를 비교한다
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S인가?
		if IsGameEnd(result){
			// 게임끝
			break 
		}
	}

	// 게임끝. 몇번만에 출력
	fmt.Printf("%d 번만에 맞쳤습니다.\n", cnt)
 }

 func MakeNumbers() [3]int {
	// 0~9사이의 겹치지 않는 무작위 숫자 3개를 반환한다.
	var rst [3]int

	for i :=0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j :=0; j < i; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if !duplicated { // 겹치지 않으면 숫자를 배열에 추가
				rst[i] = n
				break
			}
		}
	}

	fmt.Println(rst)
	return rst
 }

 func InputNumbers() [3]int {
	 // 키보드로부터 0~9사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	 var rst [3]int

	 return rst
 }

 func CompareNumbers(numbers, inputNumbers [3]int) bool {
	 // 두개의 숫자 3개를 비교해서 결과를 반환한다.
	 return true
 }

func PrintResult(result bool) {
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	// 비교 결과가 3스트라이크인지 확인?
	return result
}
