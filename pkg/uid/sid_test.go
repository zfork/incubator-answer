package uid

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_ShortID(t *testing.T) {
	nums := []int64{
		10030000000000689, 10010000000000676, 10020000000000658, 10020000000000654,
		10030000009000689, 10010000009000676, 10020000009000658, 10020000009999999,
		10030000090000689, 10010000090000676, 10020000090000658, 10020000099999999,
		10030000900000689, 10010000900000676, 10020000900000658, 10020000999999999,
		10030009000000689, 10010009000000676, 10020009000000658, 10020009999999999,
		10030090000000689, 10010090000000676, 10020090000000658, 10020099999999999,
		10030900000000689, 10010900000000676, 10020900000000658, 10020999999999999,
		10039000000000689, 10019000000000676, 10029000000000658, 10029999999999999,
		10610000000000689, 10610000000000676, 10610000000000658, 10610000000000654,
	}
	for _, num := range nums {
		code := NumToShortID(num)
		denum := ShortIDToNum(code)
		fmt.Println(num, code, denum)
	}
}

func Test_EnDeShortID(t *testing.T) {
	nums := []string{"0", "1", "10", "100", "1000", "10000", "100000", "1234567", "10000000000000000", "10010000000001316", "10030000000001316", "99999999999999999", "999999999999999999", "1999999999999999999"}
	for _, num := range nums {
		code := EnShortID(num)
		denum := DeShortID(code)
		fmt.Println(num, code, denum)
	}
}

func Test_Demo(t *testing.T) {
	nums := []int64{0, 1, 10, 100, 1000, 10000, 100000, 1000000001316, 9000000001316, 10000000000000000, 10010000000001316, 10030000000001316, 99999999999999999, 999999999999999999, 1999999999999999999}
	for _, num := range nums {
		code := strconv.FormatInt(num, 36) //10 yo 16
		fmt.Println(num, code)
	}
}
