package popcount

import "sync"

// pc[i] is the population count of i.
var (
	pc     [256]byte
	pcOnce sync.Once
)

// PcInit ..
func PcInit() [256]byte {
	pcOnce.Do(func() {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
	return pc
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	pc = PcInit()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-
