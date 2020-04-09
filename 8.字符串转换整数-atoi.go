package main

import "fmt"

// @lc code=start
func myAtoi(str string) int {
	var begin = false
	var pos bool = false
	var nag bool = false
	var sum int64 = 0
	for _, s := range str {
		// fmt.Printf("%c, ", s)
		if s == ' ' {
			if begin {
				break
			} else {
				continue
			}
		}
		if s >= '0' && s <= '9' {
			begin = true
			sum = 10*sum + int64(s-'0')
			if nag {
				if sum > 2147483648 {
					return -2147483648
				}
			} else {
				if sum > 2147483647 {
					return 2147483647
				}
			}
		} else if s == '-' {
			if nag || pos || begin {
				break
			} else {
				begin = true
				nag = true
			}
		} else if s == '+' {
			if nag || pos || begin {
				break
			} else {
				begin = true
				pos = true
			}
		} else {
			break
		}
	}
	if nag {
		return int(0 - sum)
	} else {
		return int(sum)
	}
}

// @lc code=end

func main() {
	fmt.Println(myAtoi("+-2"))
}
