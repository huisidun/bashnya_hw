package main

import "fmt"

func main() {

	var num int

	fmt.Println("введите целое число меньшее 12307")
	fmt.Scanf("%d", &num)

	if num > 12307 {
		fmt.Println("число должно быть меньше 12307")
	} else {
		for (num < 12307) {
			if num < 0 {
				num *= -1
			} else if num%7==0 {
				num *= 39
			} else if num%9==0 {
				num *= 13
				num += 1
				continue
			} else {
				num += 2 
				num *= 3
			}
			if num%13==0 &&num%9==0 {
				fmt.Println("service error")
				break
			} else {
				num +=1
			}
		}
		fmt.Printf("новое значение: %v, тип переменной: %T", num, num)
	}
}