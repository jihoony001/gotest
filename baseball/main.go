package main

import (
	"fmt"
	"time"
	"math/rand"
)

 func main() {
	// rand seed 값 적용
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

	 for {
		 fmt.Println("겹치지 않는 숫자 3개를 입력해주세요")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			//겹치는지 확인
			duplicated := false
			for j :=0; j < idx; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false 
				break
			}

			if idx >= 3{
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if success && idx < 3{
			fmt.Println("3개의 숫자를 입력하세요")
			success = false
		}

		if !success{
			continue
		}
		break
	 }
	 rst[0], rst[2] = rst[2], rst[0]
	 fmt.Println(rst)
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
