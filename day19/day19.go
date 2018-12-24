package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

const regCount = 6

type registers [regCount]int
type opFunc = func(a, b, c int, in registers) registers
type emitFunc = func(a, b, c int) string

type instruction struct {
	in    string
	opStr string
	op    opFunc
	emit  emitFunc
	a     int
	b     int
	c     int
}

func (i *instruction) run(r registers) registers {
	return i.op(i.a, i.b, i.c, r)
}

func (i *instruction) emitOp() string {
	return i.emit(i.a, i.b, i.c)
}

func addr(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] + in[b]
	return out
}

func addrEmit(a, b, c int) string {
	return regEmit(a, b, c, "+")
}

func regEmit(a, b, c int, op string) string {
	return fmt.Sprintf("r%d = r%d %s r%d", c, a, op, b)
}

func immEmit(a, b, c int, op string) string {
	return fmt.Sprintf("r%d = r%d %s %d", c, a, op, b)
}

func addi(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] + b
	return out
}

func addiEmit(a, b, c int) string {
	return immEmit(a, b, c, "+")
}

func mulr(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] * in[b]
	return out
}

func mulrEmit(a, b, c int) string {
	return regEmit(a, b, c, "*")
}

func muli(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] * b
	return out
}

func muliEmit(a, b, c int) string {
	return immEmit(a, b, c, "*")
}

func banr(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] & in[b]
	return out
}

func banrEmit(a, b, c int) string {
	return regEmit(a, b, c, "&")
}

func bani(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] & b
	return out
}

func baniEmit(a, b, c int) string {
	return immEmit(a, b, c, "&")
}

func borr(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] | in[b]
	return out
}

func borrEmit(a, b, c int) string {
	return regEmit(a, b, c, "|")
}

func bori(a, b, c int, in registers) registers {
	out := in
	out[c] = in[a] | b
	return out
}

func boriEmit(a, b, c int) string {
	return immEmit(a, b, c, "|")
}

func setr(a, _, c int, in registers) registers {
	out := in
	out[c] = in[a]
	return out
}

func setrEmit(a, _, c int) string {
	return fmt.Sprintf("r%d = r%d", c, a)
}

func seti(a, _, c int, in registers) registers {
	out := in
	out[c] = a
	return out
}

func setiEmit(a, _, c int) string {
	return fmt.Sprintf("r%d = %d", c, a)
}

func gtir(a, b, c int, in registers) registers {
	out := in
	out[c] = 0
	if a > in[b] {
		out[c] = 1
	}
	return out
}

func gtirEmit(a, b, c int) string {
	return fmt.Sprintf("out := int64(0)\nif %d > r%d {\n\tout = 1\n} \n r%d = out", a, b, c)
}

func gtri(a, b, c int, in registers) registers {
	out := in
	out[c] = 0
	if in[a] > b {
		out[c] = 1
	}
	return out
}

func gtriEmit(a, b, c int) string {
	return fmt.Sprintf("out := int64(0)\nif r%d > %d {\n\tout = 1\n} r%d = out", a, b, c)
}

func gtrr(a, b, c int, in registers) registers {
	out := in
	out[c] = 0
	if in[a] > in[b] {
		out[c] = 1
	}
	return out
}

func gtrrEmit(a, b, c int) string {
	return fmt.Sprintf("out := int64(0)\nif r%d > r%d {\n\tout = 1\n} \n  r%d = out", a, b, c)
}

func eqir(a, b, c int, in registers) registers {
	out := in
	out[c] = 0
	if a == in[b] {
		out[c] = 1
	}
	return out
}

func eqirEmit(a, b, c int) string {
	return fmt.Sprintf("out := int64(0)\nif %d == r%d {\n\tout = 1\n} \n r%d = out", a, b, c)
}

func eqri(a, b, c int, in registers) registers {
	out := in
	out[c] = 0
	if in[a] == b {
		out[c] = 1
	}
	return out
}

func eqriEmit(a, b, c int) string {
	return fmt.Sprintf("out := int64(0)\nif r%d == %d {\n\tout = 1\n} \n r%d = out", a, b, c)
}

func eqrr(a, b, c int, in registers) registers {
	out := in
	out[c] = 0
	if in[a] == in[b] {
		out[c] = 1
	}
	return out
}

func eqrrEmit(a, b, c int) string {
	return fmt.Sprintf("out := int64(0)\nif r%d == r%d {\n\tout = 1\n} \n r%d = out", a, b, c)
}

var opMap = map[string]opFunc{
	"addr": addr, "addi": addi,
	"mulr": mulr, "muli": muli,
	"banr": banr, "bani": bani,
	"borr": borr, "bori": bori,
	"gtir": gtir, "gtri": gtri, "gtrr": gtrr,
	"eqir": eqir, "eqri": eqri, "eqrr": eqrr,
	"seti": seti, "setr": setr}

var emitMap = map[string]emitFunc{
	"addr": addrEmit, "addi": addiEmit,
	"mulr": mulrEmit, "muli": muliEmit,
	"banr": banrEmit, "bani": baniEmit,
	"borr": borrEmit, "bori": boriEmit,
	"gtir": gtirEmit, "gtri": gtriEmit, "gtrr": gtrrEmit,
	"eqir": eqirEmit, "eqri": eqriEmit, "eqrr": eqrrEmit,
	"seti": setiEmit, "setr": setrEmit}

func findIP(s string) int {
	var ip int
	_, err := fmt.Sscanf(s, "#ip %d", &ip)
	if err != nil {
		panic(err)
	}

	return ip
}

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	bs, err := ioutil.ReadFile("day19.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")
	ip := findIP(lines[0])

	var ins []instruction

	for _, line := range lines[1:] {
		in := parseInstruction(line)
		ins = append(ins, in)
	}

	reg := registers{}
	for reg[ip] >= 0 && reg[ip] < len(ins) {
		curIP := reg[ip]
		reg = ins[curIP].run(reg)
		reg[ip]++
	}

	log.Printf("Part1: %d", reg[0])

	fmt.Println("package main")
	for i := 0; i < regCount; i++ {
		fmt.Printf("var r%d int64\n", i)

	}

	fmt.Println("\nfunc main() {\n")
	fmt.Println("profile()")
	fmt.Printf("for r%d >= 0 && r%d < %d {\n", ip, ip, len(ins))

	fmt.Printf("switch r%d {\n", ip)

	for i := 0; i < len(ins); i++ {
		fmt.Printf("// %s\n", ins[i].in)
		fmt.Printf("case %d:\n", i)
		fmt.Printf("invoke%d()\n", i)
	}

	fmt.Printf("\n}\n}\nfmt.Printf(\"Output: %%d\", r0)\n}\n")

	for i := 0; i < len(ins); i++ {
		fmt.Printf("func invoke%d() {\n", i)
		for j := i; j < len(ins); j++ {
			fmt.Printf("// %s\n", ins[j].in)
			fmt.Printf("%s\n", ins[j].emitOp())
			if ins[j].c == ip {
				if ins[j].opStr == "seti" {
					j = ins[j].a
				} else if ins[j].opStr == "addi" && ins[j].a == ip {
					j += ins[i].b
				} else {
					fmt.Printf("r%d++ // %d\n", ip, j+1)
					break
				}
			}
			fmt.Printf("r%d++ // %d\n", ip, j+1)
		}
		fmt.Println("}\n")
	}
	/*		if ins[j].c == ip {
			if ins[j].opStr == "seti" {
				j = ins[j].a
			} else if ins[j].opStr == "addi" && ins[j].a == ip {
				j += ins[j].b
			} else {
				break
			}
		}
	}

	for j := i; j < len(ins); j++ {
		fmt.Printf("")
		fmt.Printf("%s\n", ins[j].emitOp())
		fmt.Printf("r%d++ // %d \n", ip, j+1)
		if ins[j].c == ip {
			if ins[j].opStr == "seti" {
				j = ins[j].a
			} else if ins[j].opStr == "addi" && ins[j].a == ip {
				j += ins[j].b
			} else {
				break
			}
		}
	}*/
}

func parseInstruction(s string) instruction {
	var opName string
	var a int
	var b int
	var c int

	_, err := fmt.Sscanf(s, "%s %d %d %d", &opName, &a, &b, &c)
	if err != nil {
		panic(err)
	}

	return instruction{s, opName, opMap[opName], emitMap[opName], a, b, c}
}
