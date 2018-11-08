package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fileName    = flag.String("file", "", "file name to be generated")
	nesting     = flag.Int("nesting", 0, "how deep the nesting should be")
	packageName = flag.String("package", "", "package name to be used")
)

func main() {
	flag.Parse()

	f, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, "package %s\n\n", *packageName)
	fmt.Fprintf(f, "import \"math/rand\"\n\n")

	for typeNum := 1; typeNum <= *nesting; typeNum++ {
		fmt.Fprintf(f, "type TV%3.3d struct {\n", typeNum)
		for fieldNum := 1; fieldNum <= typeNum; fieldNum++ {
			fmt.Fprintf(f, "I%3.3d int64\n", fieldNum)
			if typeNum > 1 {
				fmt.Fprintf(f, "S%3.3d TV%3.3d\n", fieldNum, typeNum-1)
			}
		}
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func NewTV%3.3d() TV%3.3d {\n", typeNum, typeNum)
		fmt.Fprintf(f, "return TV%3.3d {\n", typeNum)
		for fieldNum := 1; fieldNum <= typeNum; fieldNum++ {
			fmt.Fprintf(f, "I%3.3d: rand.Int63(),\n", fieldNum)
			if typeNum > 1 {
				fmt.Fprintf(f, "S%3.3d: NewTV%3.3d(),\n", fieldNum, typeNum-1)
			}
		}
		fmt.Fprintf(f, "}\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "type TP%3.3d struct {\n", typeNum)
		for fieldNum := 1; fieldNum <= typeNum; fieldNum++ {
			fmt.Fprintf(f, "I%3.3d int64\n", fieldNum)
			if typeNum > 1 {
				fmt.Fprintf(f, "S%3.3d *TP%3.3d\n", fieldNum, typeNum-1)
			}
		}
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func NewTP%3.3d() *TP%3.3d {\n", typeNum, typeNum)
		fmt.Fprintf(f, "return &TP%3.3d {\n", typeNum)
		for fieldNum := 1; fieldNum <= typeNum; fieldNum++ {
			fmt.Fprintf(f, "I%3.3d: rand.Int63(),\n", fieldNum)
			if typeNum > 1 {
				fmt.Fprintf(f, "S%3.3d: NewTP%3.3d(),\n", fieldNum, typeNum-1)
			}
		}
		fmt.Fprintf(f, "}\n")
		fmt.Fprintf(f, "}\n\n")

	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func numbersRange(start int, finish int) []int {
	res := make([]int, finish-start)

	for ri := start; ri < finish; ri++ {
		res[ri] = ri
	}

	return res
}
