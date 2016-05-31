package popcount

import 	(
	"fmt"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		//fmt.Printf("before: i=%v : pc[%v]=%v : pc[%v/2]=%v : byte(i&1)=%v\n", i, i, pc[i], i, pc[i/2], byte(i&1)) 
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("after:  i=%v : pc[%v]=%v : pc[%v/2]=%v : byte(i&1)=%v %v\n", i,i , pc[i], i, pc[i/2], byte(i&1), strconv.FormatInt(int64(i),2))
		
	}
}

func Popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
