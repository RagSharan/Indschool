package even_odd

import "fmt"

func ScanFunc() {
	var length, even, odd int
	fmt.Scanln(&length)
	var numbers = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Println(numbers)
	for _, num := range numbers {
		var flag = isEven(num)
		if flag {
			even += num
		} else {
			odd += num
		}
	}
	if even > odd {
		fmt.Printf("Even")
	} else if even < odd {
		fmt.Printf("Odd")
	} else {
		fmt.Printf("Tied")
	}
}

func isEven(num int) (flag bool) {
	remain := num % 2
	flag = (remain == 0)
	return flag
}
