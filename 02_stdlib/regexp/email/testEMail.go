package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gijsbers/go-pcre"
)

const reStr = "\\s*[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\\s+"

func main() {
	tBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	tString := string(tBytes)

	fmt.Printf("testing on %d bytes\n", len(tBytes))

	sTime := time.Now()
	re := regexp.MustCompile(reStr)
	i := 0
	res := regexp.MustCompile(reStr).FindAllIndex(tBytes, -1)
	i = len(res)

	fmt.Printf("stdlib regexp with compile: %fs, %d found\n", time.Now().Sub(sTime).Seconds(), i)

	sTime = time.Now()
	re2 := regexp2.MustCompile(reStr, 0)
	var m *regexp2.Match
	i = 0
	for m, err = re2.FindStringMatch(tString); m != nil && err == nil; m, err = re2.FindNextMatch(m) {
		i++
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("dlclark/regexp2 with compile: %fs, %d found\n", time.Now().Sub(sTime).Seconds(), i)

	sTime = time.Now()
	rep := pcre.MustCompile(reStr, 0)
	i = 0
	tBytesTmp := tBytes
	for loc := rep.FindIndex(tBytesTmp, 0); loc != nil && loc[1] < len(tBytesTmp); loc = rep.FindIndex(tBytesTmp, 0) {
		tBytesTmp = tBytesTmp[loc[1]:]
		i++
	}

	fmt.Printf("gijsbers/go-pcre with compile: %fs, %d found\n", time.Now().Sub(sTime).Seconds(), i)

	sTime = time.Now()
	res = re.FindAllIndex(tBytes, -1)
	i = len(res)

	fmt.Printf("stdlib regexp: %fs, %d found\n", time.Now().Sub(sTime).Seconds(), i)

	sTime = time.Now()
	i = 0
	for m, err = re2.FindStringMatch(tString); m != nil && err == nil; m, err = re2.FindNextMatch(m) {
		i++
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("dlclark/regexp2: %fs, %d found\n", time.Now().Sub(sTime).Seconds(), i)

	sTime = time.Now()
	i = 0
	tBytesTmp = tBytes
	for loc := rep.FindIndex(tBytesTmp, 0); loc != nil && loc[1] < len(tBytesTmp); loc = rep.FindIndex(tBytesTmp, 0) {
		tBytesTmp = tBytesTmp[loc[1]:]
		i++
	}

	fmt.Printf("gijsbers/go-pcre: %fs, %d found\n", time.Now().Sub(sTime).Seconds(), i)
}
