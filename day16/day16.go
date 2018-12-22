package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"runtime"
	"strings"
)

func addr(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] + in[b]
	return out
}

func addi(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] + b
	return out
}

func mulr(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] * in[b]
	return out
}

func muli(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] * b
	return out
}

func banr(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] & in[b]
	return out
}

func bani(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] & b
	return out
}

func borr(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] | in[b]
	return out
}

func bori(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a] | b
	return out
}

func setr(a, _, c int, in [4]int) [4]int {
	out := in
	out[c] = in[a]
	return out
}

func seti(a, _, c int, in [4]int) [4]int {
	out := in
	out[c] = a
	return out
}

func gtir(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = 0
	if a > in[b] {
		out[c] = 1
	}
	return out
}

func gtri(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = 0
	if in[a] > b {
		out[c] = 1
	}
	return out
}

func gtrr(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = 0
	if in[a] > in[b] {
		out[c] = 1
	}
	return out
}

func eqir(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = 0
	if a == in[b] {
		out[c] = 1
	}
	return out
}

func eqri(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = 0
	if in[a] == b {
		out[c] = 1
	}
	return out
}

func eqrr(a, b, c int, in [4]int) [4]int {
	out := in
	out[c] = 0
	if in[a] == in[b] {
		out[c] = 1
	}
	return out
}

func parseReg(s string) [4]int {
	var ret [4]int

	ret[0] = int(s[9] - '0')
	ret[1] = int(s[12] - '0')
	ret[2] = int(s[15] - '0')
	ret[3] = int(s[18] - '0')

	for i := 0; i < 4; i++ {
		if ret[i] < 0 || ret[i] > 9 {
			log.Fatalf("Failed to parse line: %s, %d, %d", s, i, ret[i])
		}
	}

	return ret
}

func parseOp(s string) [4]int {
	var ret [4]int
	_, err := fmt.Sscanf(s, "%d %d %d %d", &ret[0], &ret[1], &ret[2], &ret[3])
	if err != nil {
		panic(err)
	}

	return ret
}

type opFunc = func(a, b, c int, in [4]int) [4]int

func cmp(a [4]int, b [4]int) bool {
	ret := true
	ret = ret && a[0] == b[0]
	ret = ret && a[1] == b[1]
	ret = ret && a[2] == b[2]
	ret = ret && a[3] == b[3]

	return ret
}

var opFuncs = []opFunc{addr, addi,
	mulr, muli,
	banr, bani,
	borr, bori,
	gtir, gtri, gtrr,
	eqir, eqri, eqrr,
	seti, setr}

var opMap [16][]opFunc

type input struct {
	before [4]int
	op     [4]int
	after  [4]int
}

func main() {
	bs, err := ioutil.ReadFile("day16-1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	inputs := make([]input, 0, len(lines)/4)
	for i := 0; i < len(lines); i += 4 {
		beforeS := strings.TrimSpace(lines[i])
		opS := strings.TrimSpace(lines[i+1])
		afterS := strings.TrimSpace(lines[i+2])
		// Blank
		before := parseReg(beforeS)
		after := parseReg(afterS)
		op := parseOp(opS)

		inputs = append(inputs, input{before, op, after})
	}

	for i := 0; i < len(opMap); i++ {
		opMap[i] = opFuncs
	}

	output := runInputs(inputs)
	log.Printf("Part 1:%d", output)

	knownOps := []opFunc{}

	for len(knownOps) < len(opFuncs) {
		knownOps = []opFunc{}
		for _, opFs := range opMap {
			if len(opFs) == 1 {
				knownOps = append(knownOps, opFs[0])
			}
		}
		log.Printf("%s are known", funcNames(knownOps))
		// Remove the known ones and loop again
		for n, opFs := range opMap {
			if len(opFs) > 1 {
				opMap[n] = removeKnown(opFs, knownOps)
			}
		}
	}

	bs, err = ioutil.ReadFile("day16-2.txt")
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(bs), "\n")
	reg := [4]int{0, 0, 0, 0}
	for _, line := range lines {
		ops := parseOp(line)
		reg = opMap[ops[0]][0](ops[1], ops[2], ops[3], reg)
	}
	log.Printf("Part 2: %d", reg[0])
}

func runInputs(inputs []input) int {
	output := 0
	for _, inp := range inputs {
		opN := inp.op[0]

		matches := runAllOps(inp.before, inp.op, inp.after)
		cnt := len(matches)

		log.Printf("%v -> %v -> %v: %d", inp.before, inp.op, inp.after, cnt)

		opMap[opN] = intersect(matches, opMap[opN])

		matchStr := funcNames(opMap[opN])
		log.Printf("Matches: %s", matchStr)

		if len(matches) == 0 {
			log.Fatalf("Op %d impossible", opN)
		}

		if cnt >= 3 {
			output++
		}
	}

	return output
}

func funcName(f opFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func funcNames(matches []opFunc) string {
	matchNames := []string{}
	for _, match := range matches {
		matchNames = append(matchNames, funcName(match))
	}
	matchStr := strings.Join(matchNames, ",")
	return matchStr
}

func equalFunc(a, b opFunc) bool {
	return reflect.ValueOf(a) == reflect.ValueOf(b)
}

func intersect(funcsA, funcsB []opFunc) []opFunc {
	outFuncs := []opFunc{}

	for _, a := range funcsA {
		for _, b := range funcsB {
			if equalFunc(a, b) {
				outFuncs = append(outFuncs, a)
			}
		}
	}

	return outFuncs
}

func removeKnown(fs []opFunc, knownFs []opFunc) []opFunc {
	out := []opFunc{}
	for _, f := range fs {
		known := false
		for _, b := range knownFs {
			if equalFunc(f, b) {
				known = true
			}
		}
		if !known {
			out = append(out, f)
		}
	}

	return out
}

func runAllOps(before, op, after [4]int) []opFunc {
	matches := []opFunc{}

	for _, opF := range opFuncs {
		reg := opF(op[1], op[2], op[3], before)

		if cmp(reg, after) {
			matches = append(matches, opF)
		}
	}

	return matches
}
