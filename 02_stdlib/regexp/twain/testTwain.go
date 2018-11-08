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

var regexpStrings = []string{
	"Twain",
	"(?i)Twain",
	"[a-z]shing",
	"Huck[a-zA-Z]+|Saw[a-zA-Z]+",
	`\b\w+nn\b`,
	"[a-q][^u-z]{13}x",
	"Tom|Sawyer|Huckleberry|Finn",
	"(?i)Tom|Sawyer|Huckleberry|Finn",
	".{0,2}(Tom|Sawyer|Huckleberry|Finn)",
	".{2,4}(Tom|Sawyer|Huckleberry|Finn)",
	"Tom.{10,25}river|river.{10,25}Tom",
	"[a-zA-Z]+ing",
	`\s[a-zA-Z]{0,12}ing\s`,
	`([A-Za-z]awyer|[A-Za-z]inn)\s`,
	`["'][^"']{0,30}[?!\.][\"']`,
	`\p{Sm}`,
}

var regexps = func() []*regexp.Regexp {
	res := make([]*regexp.Regexp, len(regexpStrings))
	for ri := 0; ri < len(regexpStrings); ri++ {
		res[ri] = regexp.MustCompile(regexpStrings[ri])
	}
	return res
}()

var regexps2 = func() []*regexp2.Regexp {
	res := make([]*regexp2.Regexp, len(regexpStrings))
	for ri := 0; ri < len(regexpStrings); ri++ {
		res[ri] = regexp2.MustCompile(regexpStrings[ri], 0)
	}
	return res
}()

var regexpsp = func() []pcre.Regexp {
	res := make([]pcre.Regexp, len(regexpStrings))
	for ri := 0; ri < len(regexpStrings); ri++ {
		res[ri] = pcre.MustCompile(regexpStrings[ri], 0)
	}
	return res
}()

func main() {
	tBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	tString := string(tBytes)

	fmt.Printf("testing on %d bytes\n", len(tBytes))

	sTime := time.Now()
	for _, reStr := range regexpStrings {
		regexp.MustCompile(reStr).FindAllStringIndex(tString, -1)
	}

	fmt.Printf("stdlib regexp with compile: %v\n", time.Now().Sub(sTime))

	var m *regexp2.Match
	sTime = time.Now()
	for _, reStr := range regexpStrings {
		re := regexp2.MustCompile(reStr, 0)
		for m, err = re.FindStringMatch(tString); m != nil && err == nil; m, err = re.FindNextMatch(m) {
			// do nothing
		}
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("dlclark/regexp2 with compile: %v\n", time.Now().Sub(sTime))

	sTime = time.Now()
	for _, reStr := range regexpStrings {
		re := pcre.MustCompile(reStr, 0)
		tBytesTmp := tBytes
		for loc := re.FindIndex(tBytesTmp, 0); loc != nil && loc[1] < len(tBytesTmp); loc = re.FindIndex(tBytesTmp, 0) {
			tBytesTmp = tBytesTmp[loc[1]:]
		}
	}

	fmt.Printf("gijsbers/go-pcre with compile: %v\n", time.Now().Sub(sTime))

	sTime = time.Now()
	for _, re := range regexps {
		re.FindAllStringIndex(tString, -1)
	}

	fmt.Printf("stdlib regexp: %v\n", time.Now().Sub(sTime))

	sTime = time.Now()
	for _, re := range regexps2 {
		for m, err = re.FindStringMatch(tString); m != nil && err == nil; m, err = re.FindNextMatch(m) {
			// do nothing
		}
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("dlclark/regexp2: %v\n", time.Now().Sub(sTime))

	sTime = time.Now()
	for _, re := range regexpsp {
		tBytesTmp := tBytes
		for loc := re.FindIndex(tBytesTmp, 0); loc != nil && loc[1] < len(tBytesTmp); loc = re.FindIndex(tBytesTmp, 0) {
			tBytesTmp = tBytesTmp[loc[1]:]
		}
	}

	fmt.Printf("gijsbers/go-pcre: %v\n", time.Now().Sub(sTime))
}
