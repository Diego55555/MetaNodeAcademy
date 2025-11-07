package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// 只出现一次的数字
	numberArray := []int{1, 2, 3, 4, 3, 2, 1}
	number, result := getNumberAppearsOnlyOnce(numberArray)
	if result {
		fmt.Println("The Number Appears Only Once:", number)
	} else {
		fmt.Println("No Number Appears Only Once:", number)
	}

	// 判断是否是回文
	number = 1236321
	result = isPalindrome(number)
	if result {
		fmt.Println(number, " is Palindrome")
	} else {
		fmt.Println(number, " is't Palindrome")
	}
	number = 123
	result = isPalindrome(number)
	if result {
		fmt.Println(number, " is Palindrome")
	} else {
		fmt.Println(number, " is't Palindrome")
	}

	// 括号是否闭合
	text := "{12312(22)}"
	result = isParenthesesClosed(text)
	if result {
		fmt.Println(text, " is Closed")
	} else {
		fmt.Println(text, " is't Closed")
	}
	text = "{12312(2}2)}"
	result = isParenthesesClosed(text)
	if result {
		fmt.Println(text, " is Closed")
	} else {
		fmt.Println(text, " is't Closed")
	}

	// 最长公共前缀
	texts := [...]string{"flower", "flow", "flight"}
	prefix := getMaxPrefix(texts[:])
	fmt.Println("Max prefix :", prefix)
	texts = [...]string{"dog", "racecar", "car"}
	prefix = getMaxPrefix(texts[:])
	fmt.Println("Max prefix :", prefix)

	// 数字数组加一
	digitArray := []byte{1, 2, 3}
	resultDigitArray := digitArrayAddOne2(digitArray)
	fmt.Println("resultDigitArray :", resultDigitArray)
	digitArray = []byte{4, 3, 2, 1}
	resultDigitArray = digitArrayAddOne2(digitArray)
	fmt.Println("resultDigitArray :", resultDigitArray)
	digitArray = []byte{9}
	resultDigitArray = digitArrayAddOne2(digitArray)
	fmt.Println("resultDigitArray :", resultDigitArray)

	// 删除有序数组中的重复项
	digitArray = []byte{1, 1, 2}
	length := removeDuplicates(digitArray)
	fmt.Println("DigitArray length :", length, ", DigitArray:", digitArray)
	digitArray = []byte{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length = removeDuplicates(digitArray)
	fmt.Println("DigitArray length :", length, ", DigitArray:", digitArray)

	// 区间合并
	intervals := [][2]uint64{{0, 10}, {10, 20}, {25, 35}}
	mergedIntervals := mergeIntervals(intervals)
	fmt.Println("MergedIntervals :", mergedIntervals)
	intervals = [][2]uint64{{10, 15}, {10, 15}, {17, 25}, {22, 30}}
	mergedIntervals = mergeIntervals(intervals)
	fmt.Println("MergedIntervals :", mergedIntervals)
	intervals = [][2]uint64{{23, 56}, {24, 25}, {10, 15}, {17, 25}, {22, 30}, {25, 30}}
	mergedIntervals = mergeIntervals(intervals)
	fmt.Println("MergedIntervals :", mergedIntervals)

	// 两数之和
	numberArray = []int{2, 7, 11, 15}
	indexArray := twoSum(numberArray, 18)
	fmt.Println("indexArray :", indexArray)
	numberArray = []int{5, 3, 18, 2, 17, 5}
	indexArray = twoSum(numberArray, 10)
	fmt.Println("indexArray :", indexArray)
}

// 只出现一次的数字
func getNumberAppearsOnlyOnce(numbers []int) (int, bool) {
	numberMap := make(map[int]int, 10)
	for _, value := range numbers {
		numberMap[value] += 1
	}

	for key, value := range numberMap {
		if value == 1 {
			return key, true
		}
	}

	return 0, false
}

// 判断是否是回文
func isPalindrome(number int) bool {
	numberString := strconv.Itoa(number)
	numberLength := len(numberString)
	count := numberLength / 2
	for i := 0; i < count; i++ {
		if numberString[i] != numberString[numberLength-1-i] {
			return false
		}
	}

	return true
}

// 括号是否闭合
func isParenthesesClosed(text string) bool {
	parenthesesStack := []rune{}

	runes := []rune(text)
	for _, value := range runes {
		switch value {
		case '(', '{', '[':
			parenthesesStack = append(parenthesesStack, value)
		case ')', '}', ']':
			if len(parenthesesStack) <= 0 {
				return false
			}

			if (value == ')' && parenthesesStack[len(parenthesesStack)-1] != '(') ||
				(value == '}' && parenthesesStack[len(parenthesesStack)-1] != '{') ||
				(value == ']' && parenthesesStack[len(parenthesesStack)-1] != '[') {
				return false
			}

			parenthesesStack = parenthesesStack[:len(parenthesesStack)-1]
		}
	}

	return len(parenthesesStack) == 0
}

// 最长公共前缀
func getMaxPrefix(texts []string) string {
	prefix := []rune{}

	for index, value := range texts {
		if index == 0 {
			prefix = []rune(value)
		} else {
			valueRune := []rune(value)
			for subIndex := range valueRune {
				if subIndex >= len(prefix) || valueRune[subIndex] != prefix[subIndex] {
					if subIndex == 0 {
						return ""
					} else {
						prefix = prefix[:subIndex]
						break
					}
				}
			}
		}
	}

	return string(prefix)
}

// 数字数组加一
func digitArrayAddOne1(digitArray []byte) []byte {
	if len(digitArray) == 0 {
		return []byte{0}
	}

	number := uint64(0)
	for index, value := range digitArray {
		if index == 0 && value == 0 {
			return []byte{0}
		}

		if value > 9 {
			return []byte{0}
		}

		number = number*10 + uint64(value)
	}

	number++
	numberByte := []byte(strconv.Itoa(int(number)))
	for i := 0; i < len(numberByte); i++ {
		numberInt, _ := strconv.Atoi(string(numberByte[i]))
		numberByte[i] = byte(numberInt)
	}

	return numberByte
}

func digitArrayAddOne2(digitArray []byte) []byte {
	count := len(digitArray)
	if count == 0 {
		return []byte{0}
	}

	for i := count - 1; i > 0; i-- {
		if digitArray[i] < 9 {
			digitArray[i]++
			return digitArray
		}

		digitArray[i] = 0
	}

	resultDigitArray := make([]byte, count+1)
	resultDigitArray[0] = 1

	return resultDigitArray
}

// 删除有序数组中的重复项
func removeDuplicates(digitArray []byte) int {
	count := len(digitArray)
	digitMap := map[byte]byte{}
	preIndex := 0
	for i := 0; i < count; i++ {
		_, exist := digitMap[digitArray[i]]
		if !exist {
			digitMap[digitArray[i]] = 0
			if i > preIndex {
				digitArray[preIndex] = digitArray[i]
			}
			preIndex++
		}
	}

	return preIndex
}

// 区间合并
func mergeIntervals(intervals [][2]uint64) [][2]uint64 {
	count := len(intervals)
	if count == 0 {
		return [][2]uint64{}
	}

	sortIntervals2(intervals)

	mergedIntervals := make([][2]uint64, 0, count)
	mergedIntervals = append(mergedIntervals, intervals[0])
	for i := 1; i < count; i++ {
		if intervals[i][0] >= mergedIntervals[len(mergedIntervals)-1][1] {
			mergedIntervals = append(mergedIntervals, intervals[i])
			continue
		}

		if intervals[i][1] > mergedIntervals[len(mergedIntervals)-1][1] {
			mergedIntervals[len(mergedIntervals)-1][1] = intervals[i][1]
		}
	}

	return mergedIntervals
}

// 区间排序
func sortIntervals1(intervals [][2]uint64) {
	count := len(intervals)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if (intervals[j][0] < intervals[i][0]) ||
				(intervals[j][0] == intervals[i][0] && intervals[j][1] < intervals[i][1]) {
				startIndex, endIndex := intervals[i][0], intervals[i][1]
				intervals[i][0], intervals[i][1] = intervals[j][0], intervals[j][1]
				intervals[j][0], intervals[j][1] = startIndex, endIndex
			}
		}
	}
}

func sortIntervals2(intervals [][2]uint64) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})
}

// 两数之和未目标值
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, exist := numMap[complement]; exist {
			return []int{index, i}
		}

		numMap[num] = i
	}

	return []int{}
}
