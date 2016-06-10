package popcount

import 	(
	"fmt"
	"strconv"
)

var pc [256]byte

const debug = false

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		if debug {
			fmt.Printf("after:  i=%3v : pc[%3v]=%v : pc[%3v/2]=%v : byte(i&1)=%v %v\n", i,i , pc[i], i, pc[i/2], byte(i&1), strconv.FormatInt(int64(i),2))
		}
	}
}

func PopcountClear(x uint64) int {
	fmt.Printf("%v: %v\n", strconv.FormatInt(int64(x),2), x)
	fmt.Printf("%v: %v\n", strconv.FormatInt(int64(x&(x-1)),2), x&(x-1))

	return 0
}

func PopcountShift(x uint64) int {
	var sum int = 0
	for i := 0; i < 64; i++ {
		if debug {
			fmt.Printf("%v\n", strconv.FormatInt(int64(x),2))
		}
		if t := x&1; t == 1 {
			sum++
		}
		x = x>>1
	}
	return sum
}

func PopcountLoop(x uint64) int {
	var sum int = 0
	var i uint 
	for i = 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func Popcount(x uint64) int {
	if debug {
		fmt.Printf("x=%v : (x>>(1*8))=%v : byte(x>>(1*8))=%v\n",strconv.FormatInt(int64(x),2) ,strconv.FormatInt(int64(x>>(1*8)),2), byte(x>>(0*8)))
	}
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
