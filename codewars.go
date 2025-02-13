package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(ToCamelCase("The_Stealth_Warrior"))

	// arr1 := nil
	var arr2 = []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
	fmt.Println(Comp(nil, arr2))
}

func ToCamelCase(s string) string {
	// result := ""

	arr := strings.Split(s, "")

	for i := 0; i < len(arr); i++ {
		if arr[i] == " " || arr[i] == "_" || arr[i] == "-" {
			fmt.Println("found " + arr[i])
			arr[i+1] = strings.ToUpper(string(arr[i+1]))
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return strings.Join(arr, "")
}

func Comp(array1 []int, array2 []int) bool {
	found := false
	if (array1 == nil || array2 == nil) || len(array1) != len(array2) {
		return false
	}
	for _, num1 := range array1 {
		for _, num2 := range array2 {
			if num2 == num1*num1 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
		found = false
	}
	return true
}
