package main

import (
	"fmt"
	"strconv"
)

/**
Given a time in -hour AM/PM format, convert it to military (-hour) time.

Note: Midnight is  on a -hour clock, and  on a -hour clock. Noon is  on a -hour clock, and  on a -hour clock.

Input Format

A single string containing a time in -hour clock format (i.e.:  or ), where  and .

Output Format

Convert and print the given time in -hour format, where .

Sample Input

07:05:45PM
Sample Output

19:05:45
 */
func main(){
	const PM  = "PM"
	const hourTwelve = "12"
	var hourAmPm, HH string
	fmt.Scanln(&hourAmPm)
	h := hourAmPm[0:2]
	ampm := hourAmPm[8:]
	minseg := hourAmPm[2:8]

	if ampm == PM {
		if h == hourTwelve {
			HH = "12"
		} else {
			hint,_ := strconv.Atoi(h)
			HH = strconv.Itoa(hint + 12)
		}
	} else {
		if h == hourTwelve {
			HH = "00"
		} else {
			HH = h
		}
	}
	fmt.Println(HH+minseg)
}
