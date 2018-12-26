package main

import (
	"fmt"
)

var r0 int64
var r1 int64
var r2 int64
var r3 int64
var r4 int64
var r5 int64
var n int64
var r3s map[int64]int64 = make(map[int64]int64)

func loop(maxn int64) int64 {

	for r1 >= 0 && r1 < 31 && n < maxn {
		switch r1 {
		// seti 123 0 3
		case 0:
			invoke0()
		// bani 3 456 3
		case 1:
			invoke1()
		// eqri 3 72 3
		case 2:
			invoke2()
		// addr 3 1 1
		case 3:
			invoke3()
		// seti 0 0 1
		case 4:
			invoke4()
		// seti 0 9 3
		case 5:
			invoke5()
		// bori 3 65536 5
		case 6:
			invoke6()
		// seti 15028787 4 3
		case 7:
			invoke7()
		// bani 5 255 2
		case 8:
			invoke8()
		// addr 3 2 3
		case 9:
			invoke9()
		// bani 3 16777215 3
		case 10:
			invoke10()
		// muli 3 65899 3
		case 11:
			invoke11()
		// bani 3 16777215 3
		case 12:
			invoke12()
		// gtir 256 5 2
		case 13:
			invoke13()
		// addr 2 1 1
		case 14:
			invoke14()
		// addi 1 1 1
		case 15:
			invoke15()
		// seti 27 3 1
		case 16:
			invoke16()
		// seti 0 9 2
		case 17:
			invoke17()
		// addi 2 1 4
		case 18:
			invoke18()
		// muli 4 256 4
		case 19:
			invoke19()
		// gtrr 4 5 4
		case 20:
			invoke20()
		// addr 4 1 1
		case 21:
			invoke21()
		// addi 1 1 1
		case 22:
			invoke22()
		// seti 25 1 1
		case 23:
			invoke23()
		// addi 2 1 2
		case 24:
			invoke24()
		// seti 17 8 1
		case 25:
			invoke25()
		// setr 2 4 5
		case 26:
			invoke26()
		// seti 7 3 1
		case 27:
			invoke27()
		// eqrr 3 0 2
		case 28:
			invoke28()
		// addr 2 1 1
		case 29:
			invoke29()
		// seti 5 3 1
		case 30:
			invoke30()

		}
	}
	fmt.Printf("Output: %d, n: %d\n", r0, n)

	return n
}
func invoke0() {
	// seti 123 0 3
	r3 = 123
	r1++ // 1
	n++
	// bani 3 456 3
	r3 &= 456
	r1++ // 2
	n++
	// eqri 3 72 3
	if r3 == 72 {
		r3 = 1
	} else {
		r3 = 0
	}
	r1++ // 3
	n++
	// addr 3 1 1
	r1 += r3
	r1++ // 4
	n++
}

func invoke1() {
	// bani 3 456 3
	r3 &= 456
	r1++ // 2
	n++
	// eqri 3 72 3
	if r3 == 72 {
		r3 = 1
	} else {
		r3 = 0
	}
	r1++ // 3
	n++
	// addr 3 1 1
	r1 += r3
	r1++ // 4
	n++
}

func invoke2() {
	// eqri 3 72 3
	if r3 == 72 {
		r3 = 1
	} else {
		r3 = 0
	}
	r1++ // 3
	n++
	// addr 3 1 1
	r1 += r3
	r1++ // 4
	n++
}

func invoke3() {
	// addr 3 1 1
	r1 += r3
	r1++ // 4
	n++
}

func invoke4() {
	// seti 0 0 1
	r1 = 0
	r1++ // 5
	n++
	// bani 3 456 3
	r3 &= 456
	r1++ // 2
	n++
	// eqri 3 72 3
	if r3 == 72 {
		r3 = 1
	} else {
		r3 = 0
	}
	r1++ // 3
	n++
	// addr 3 1 1
	r1 += r3
	r1++ // 4
	n++
}

func invoke5() {
	// seti 0 9 3
	r3 = 0
	r1++ // 6
	n++
	// bori 3 65536 5
	r5 = r3 | 65536
	r1++ // 7
	n++
	// seti 15028787 4 3
	r3 = 15028787
	r1++ // 8
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke6() {
	// bori 3 65536 5
	r5 = r3 | 65536
	r1++ // 7
	n++
	// seti 15028787 4 3
	r3 = 15028787
	r1++ // 8
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke7() {
	// seti 15028787 4 3
	r3 = 15028787
	r1++ // 8
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke8() {
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke9() {
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke10() {
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke11() {
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke12() {
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke13() {
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke14() {
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke15() {
	// addi 1 1 1
	r1 += 1
	r1++ // 16
	n++
	// seti 0 9 2
	r2 = 0
	r1++ // 18
	n++
	// addi 2 1 4
	r4 = r2 + 1
	r1++ // 19
	n++
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke16() {
	// seti 27 3 1
	r1 = 27
	r1++ // 17
	n++
	// eqrr 3 0 2
	if _, ok := r3s[r3]; !ok {
		r3s[r3] = n
		if len(r3s) == 13004 {
			fmt.Printf("%d\n", r3)
		}
	}

	if r3 == r0 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 29
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 30
	n++
}

func invoke17() {
	// seti 0 9 2
	r2 = 0
	r1++ // 18
	n++
	// addi 2 1 4
	r4 = r2 + 1
	r1++ // 19
	n++
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke18() {
	// addi 2 1 4
	r4 = r2 + 1
	r1++ // 19
	n++
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke19() {
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke20() {
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke21() {
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke22() {
	// addi 1 1 1
	r1 += 1
	r1++ // 23
	n++
	// addi 2 1 2
	r2 += 1
	r1++ // 25
	n++
	// seti 17 8 1
	r1 = 17
	r1++ // 26
	n++
	// addi 2 1 4
	r4 = r2 + 1
	r1++ // 19
	n++
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke23() {
	// seti 25 1 1
	r1 = 25
	r1++ // 24
	n++
	// setr 2 4 5
	r5 = r2
	r1++ // 27
	n++
	// seti 7 3 1
	r1 = 7
	r1++ // 28
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke24() {
	// addi 2 1 2
	r2 += 1
	r1++ // 25
	n++
	// seti 17 8 1
	r1 = 17
	r1++ // 26
	n++
	// addi 2 1 4
	r4 = r2 + 1
	r1++ // 19
	n++
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke25() {
	// seti 17 8 1
	r1 = 17
	r1++ // 26
	n++
	// addi 2 1 4
	r4 = r2 + 1
	r1++ // 19
	n++
	// muli 4 256 4
	r4 *= 256
	r1++ // 20
	n++
	// gtrr 4 5 4
	if r4 > r5 {
		r4 = 1
	} else {
		r4 = 0
	}
	r1++ // 21
	n++
	// addr 4 1 1
	r1 += r4
	r1++ // 22
	n++
}

func invoke26() {
	// setr 2 4 5
	r5 = r2
	r1++ // 27
	n++
	// seti 7 3 1
	r1 = 7
	r1++ // 28
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke27() {
	// seti 7 3 1
	r1 = 7
	r1++ // 28
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}

func invoke28() {
	// eqrr 3 0 2
	if _, ok := r3s[r3]; !ok {
		r3s[r3] = n
		if len(r3s) == 13004 {
			fmt.Printf("%d\n", r3)
		}
	}

	if r3 == r0 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 29
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 30
	n++
}

func invoke29() {
	// addr 2 1 1
	r1 += r2
	r1++ // 30
	n++
}

func invoke30() {
	// seti 5 3 1
	r1 = 5
	r1++ // 31
	n++
	// bori 3 65536 5
	r5 = r3 | 65536
	r1++ // 7
	n++
	// seti 15028787 4 3
	r3 = 15028787
	r1++ // 8
	n++
	// bani 5 255 2
	r2 = r5 & 255
	r1++ // 9
	n++
	// addr 3 2 3
	r3 += r2
	r1++ // 10
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 11
	n++
	// muli 3 65899 3
	r3 *= 65899
	r1++ // 12
	n++
	// bani 3 16777215 3
	r3 &= 16777215
	r1++ // 13
	n++
	// gtir 256 5 2
	if 256 > r5 {
		r2 = 1
	} else {
		r2 = 0
	}
	r1++ // 14
	n++
	// addr 2 1 1
	r1 += r2
	r1++ // 15
	n++
}
