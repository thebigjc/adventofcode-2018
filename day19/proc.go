package main

import "fmt"

var r0 int64
var r1 int64
var r2 int64
var r3 int64
var r4 int64
var r5 int64

func main() {

	profile()
	for r1 >= 0 && r1 < 36 {
		switch r1 {
		// addi 1 16 1
		case 0:
			invoke0()
		// seti 1 2 5
		case 1:
			invoke1()
		// seti 1 2 2
		case 2:
			invoke2()
		// mulr 5 2 3
		case 3:
			invoke3()
		// eqrr 3 4 3
		case 4:
			invoke4()
		// addr 3 1 1
		case 5:
			invoke5()
		// addi 1 1 1
		case 6:
			invoke6()
		// addr 5 0 0
		case 7:
			invoke7()
		// addi 2 1 2
		case 8:
			invoke8()
		// gtrr 2 4 3
		case 9:
			invoke9()
		// addr 1 3 1
		case 10:
			invoke10()
		// seti 2 8 1
		case 11:
			invoke11()
		// addi 5 1 5
		case 12:
			invoke12()
		// gtrr 5 4 3
		case 13:
			invoke13()
		// addr 3 1 1
		case 14:
			invoke14()
		// seti 1 1 1
		case 15:
			invoke15()
		// mulr 1 1 1
		case 16:
			invoke16()
		// addi 4 2 4
		case 17:
			invoke17()
		// mulr 4 4 4
		case 18:
			invoke18()
		// mulr 1 4 4
		case 19:
			invoke19()
		// muli 4 11 4
		case 20:
			invoke20()
		// addi 3 3 3
		case 21:
			invoke21()
		// mulr 3 1 3
		case 22:
			invoke22()
		// addi 3 4 3
		case 23:
			invoke23()
		// addr 4 3 4
		case 24:
			invoke24()
		// addr 1 0 1
		case 25:
			invoke25()
		// seti 0 0 1
		case 26:
			invoke26()
		// setr 1 5 3
		case 27:
			invoke27()
		// mulr 3 1 3
		case 28:
			invoke28()
		// addr 1 3 3
		case 29:
			invoke29()
		// mulr 1 3 3
		case 30:
			invoke30()
		// muli 3 14 3
		case 31:
			invoke31()
		// mulr 3 1 3
		case 32:
			invoke32()
		// addr 4 3 4
		case 33:
			invoke33()
		// seti 0 0 0
		case 34:
			invoke34()
		// seti 0 1 1
		case 35:
			invoke35()

		}
	}
	fmt.Printf("Output: %d", r0)
}
func invoke0() {
	// addi 1 16 1
	r1 = r1 + 16
	r1++ // 17
	// addi 4 2 4
	r4 = r4 + 2
	r1++ // 18
	// mulr 4 4 4
	r4 = r4 * r4
	r1++ // 19
	// mulr 1 4 4
	r4 = r1 * r4
	r1++ // 20
	// muli 4 11 4
	r4 = r4 * 11
	r1++ // 21
	// addi 3 3 3
	r3 = r3 + 3
	r1++ // 22
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke1() {
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke2() {
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke3() {
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke4() {
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke5() {
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke6() {
	r2++

	if r2 > r4 {
		r3 = 1
		r1 = 12
	} else {
		r3 = 0
		r1 = 11
	}
}

func invoke7() {
	// addr 5 0 0
	r0 += r5
	r2++
	if r2 > r4 {
		r3 = 1
		r1 = 12
	} else {
		r3 = 0
		r1 = 11
	}
}

func invoke8() {
	// addi 2 1 2
	r2 = r2 + 1
	r1++ // 9
	// gtrr 2 4 3
	out := int64(0)
	if r2 > r4 {
		out = 1
	}
	r3 = out
	r1++ // 10
	// addr 1 3 1
	r1 = r1 + r3
	r1++ // 11
}

func invoke9() {
	// gtrr 2 4 3
	out := int64(0)
	if r2 > r4 {
		out = 1
	}
	r3 = out
	r1++ // 10
	// addr 1 3 1
	r1 = r1 + r3
	r1++ // 11
}

func invoke10() {
	// addr 1 3 1
	r1 = r1 + r3
	r1++ // 11
}

func invoke11() {
	if r4 != r5 && r4%r5 == 0 {
		r0 += r5
		r5++
		if r5 > r4 {
			r1 = 257
			return
		} else {
			r2 = 1
			if r5*r2 == r4 {
				r3 = 1
				r1 = 7
				return
			} else {
				r3 = 0
				r1 = 6
				return
			}
		}
	} else {
		r5++
		if r5 > r4 {
			r1 = 257
			return
		} else {
			r2 = 1
			if r5*r2 == r4 {
				r3 = 1
				r1 = 7
			} else {
				r3 = 0
				r1 = 6
			}
		}
	}
}

func invoke12() {
	r5++
	if r5 > r4 {
		r1 = 257
		return
	} else {
		r2 = 1
		if r5*r2 == r4 {
			r3 = 1
			r1 = 7
		} else {
			r3 = 0
			r1 = 6
		}
	}
}

func invoke13() {
	// gtrr 5 4 3
	out := int64(0)
	if r5 > r4 {
		out = 1
	}
	r3 = out
	r1++ // 14
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 15
}

func invoke14() {
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 15
}

func invoke15() {
	r2 = 1
	if r5*r2 == r4 {
		r3 = 1
		r1 = 7
	} else {
		r3 = 0
		r1 = 6
	}
}

func invoke16() {
	// mulr 1 1 1
	r1 = r1 * r1
	r1++ // 17
}

func invoke17() {
	// addi 4 2 4
	r4 = r4 + 2
	r1++ // 18
	// mulr 4 4 4
	r4 = r4 * r4
	r1++ // 19
	// mulr 1 4 4
	r4 = r1 * r4
	r1++ // 20
	// muli 4 11 4
	r4 = r4 * 11
	r1++ // 21
	// addi 3 3 3
	r3 = r3 + 3
	r1++ // 22
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke18() {
	// mulr 4 4 4
	r4 = r4 * r4
	r1++ // 19
	// mulr 1 4 4
	r4 = r1 * r4
	r1++ // 20
	// muli 4 11 4
	r4 = r4 * 11
	r1++ // 21
	// addi 3 3 3
	r3 = r3 + 3
	r1++ // 22
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke19() {
	// mulr 1 4 4
	r4 = r1 * r4
	r1++ // 20
	// muli 4 11 4
	r4 = r4 * 11
	r1++ // 21
	// addi 3 3 3
	r3 = r3 + 3
	r1++ // 22
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke20() {
	// muli 4 11 4
	r4 = r4 * 11
	r1++ // 21
	// addi 3 3 3
	r3 = r3 + 3
	r1++ // 22
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke21() {
	// addi 3 3 3
	r3 = r3 + 3
	r1++ // 22
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke22() {
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 23
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke23() {
	// addi 3 4 3
	r3 = r3 + 4
	r1++ // 24
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke24() {
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 25
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke25() {
	// addr 1 0 1
	r1 = r1 + r0
	r1++ // 26
}

func invoke26() {
	// seti 0 0 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke27() {
	// setr 1 5 3
	r3 = r1
	r1++ // 28
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 29
	// addr 1 3 3
	r3 = r1 + r3
	r1++ // 30
	// mulr 1 3 3
	r3 = r1 * r3
	r1++ // 31
	// muli 3 14 3
	r3 = r3 * 14
	r1++ // 32
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 33
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke28() {
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 29
	// addr 1 3 3
	r3 = r1 + r3
	r1++ // 30
	// mulr 1 3 3
	r3 = r1 * r3
	r1++ // 31
	// muli 3 14 3
	r3 = r3 * 14
	r1++ // 32
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 33
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke29() {
	// addr 1 3 3
	r3 = r1 + r3
	r1++ // 30
	// mulr 1 3 3
	r3 = r1 * r3
	r1++ // 31
	// muli 3 14 3
	r3 = r3 * 14
	r1++ // 32
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 33
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke30() {
	// mulr 1 3 3
	r3 = r1 * r3
	r1++ // 31
	// muli 3 14 3
	r3 = r3 * 14
	r1++ // 32
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 33
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke31() {
	// muli 3 14 3
	r3 = r3 * 14
	r1++ // 32
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 33
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke32() {
	// mulr 3 1 3
	r3 = r3 * r1
	r1++ // 33
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke33() {
	// addr 4 3 4
	r4 = r4 + r3
	r1++ // 34
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke34() {
	// seti 0 0 0
	r0 = 0
	r1++ // 35
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}

func invoke35() {
	// seti 0 1 1
	r1 = 0
	r1++ // 1
	// seti 1 2 5
	r5 = 1
	r1++ // 2
	// seti 1 2 2
	r2 = 1
	r1++ // 3
	// mulr 5 2 3
	r3 = r5 * r2
	r1++ // 4
	// eqrr 3 4 3
	out := int64(0)
	if r3 == r4 {
		out = 1
	}
	r3 = out
	r1++ // 5
	// addr 3 1 1
	r1 = r3 + r1
	r1++ // 6
}
