package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var (
	templateName = flag.String("template", "", "template file to be used")
	fileName     = flag.String("file", "", "file name to be generated")
	packageName  = flag.String("package", "", "package name to be used")
)

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	tTxt, err := ioutil.ReadFile(*templateName)
	if err != nil {
		panic(err)
	}

	t := template.Must(
		template.New("switch_data").Funcs(
			template.FuncMap{
				"numbersRange": numbersRange,
				"randIntn":     rand.Intn,
			},
		).Parse(string(tTxt)),
	)

	f, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, packageName)
	if err != nil {
		panic(err)
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
