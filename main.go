package main

import (
	"fmt"
	"strconv"
)

func main() {

	firstPrintedPage := 7
	lastPrintedPage := 613

	sheetsInMiniBook := 10 // 12-47 11-43 10-39 9-35 8-31 7-27 6-23 5-19 4-15 3-11 2-7 1-3
	//sheetsInLastMiniBook := 9

	lastPageInMiniBook := sheetsInMiniBook*4 + firstPrintedPage
	pagesInMiniBook := sheetsInMiniBook * 4

	for firstPageInMiniBook := firstPrintedPage; firstPageInMiniBook < lastPrintedPage; firstPageInMiniBook += pagesInMiniBook {
		//if firstPageInMiniBook > 573 {
		//	lastPageInMiniBook = sheetsInLastMiniBook * 4
		//	pagesInMiniBook = sheetsInLastMiniBook * 4
		//}

		var str1 []byte
		var str2 []byte
		str1 = str1[:0]
		str2 = str2[:0]

		countMinus := 4

		for pageNum := firstPageInMiniBook; pageNum < lastPageInMiniBook-pagesInMiniBook/2; pageNum += 2 {

			str1 = append(str1, []byte(strconv.Itoa(pageNum+1))...)
			str1 = append(str1, []byte(",")...)
			str1 = append(str1, []byte(strconv.Itoa(lastPageInMiniBook-countMinus+2))...)
			str1 = append(str1, []byte(",")...)
			str2 = append(str2, []byte(strconv.Itoa(lastPageInMiniBook-countMinus+3))...)
			str2 = append(str2, []byte(",")...)
			str2 = append(str2, []byte(strconv.Itoa(pageNum))...)
			str2 = append(str2, []byte(",")...)
			countMinus += 2

		}

		fmt.Println(string(str1))
		fmt.Println(string(str2))
		lastPageInMiniBook += pagesInMiniBook
	}

	//err := os.WriteFile("pages.txt", data, 0666)
	//if err != nil {
	//	panic(err)
	//}
}
