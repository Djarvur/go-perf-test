package language_test

import "fmt"

func switchIntSmall(p int) {
	switch p {
	case 0:
		func() { testInt64++ }()
	case 1:
		func() { testInt64++ }()
	case 2:
		func() { testInt64++ }()
	case 3:
		func() { testInt64++ }()
	case 4:
		func() { testInt64++ }()
	case 5:
		func() { testInt64++ }()
	case 6:
		func() { testInt64++ }()
	case 7:
		func() { testInt64++ }()
	case 8:
		func() { testInt64++ }()
	case 9:
		func() { testInt64++ }()
	default:
		panic(fmt.Errorf("Unexpected parameter: %v", p))
	}
}

var mapIntSmall = map[int]func(){
	0: func() { testInt64++ },
	1: func() { testInt64++ },
	2: func() { testInt64++ },
	3: func() { testInt64++ },
	4: func() { testInt64++ },
	5: func() { testInt64++ },
	6: func() { testInt64++ },
	7: func() { testInt64++ },
	8: func() { testInt64++ },
	9: func() { testInt64++ },
}

var sliceIntSmall = []func(){
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
	func() { testInt64++ },
}

var randIntSmall = []int{
	0,
	1,
	2,
	9,
	4,
	8,
	2,
	6,
	2,
	3,
}

func switchStringSmall(p string) {
	switch p {
	case "0":
		func() { testInt64++ }()
	case "1":
		func() { testInt64++ }()
	case "2":
		func() { testInt64++ }()
	case "3":
		func() { testInt64++ }()
	case "4":
		func() { testInt64++ }()
	case "5":
		func() { testInt64++ }()
	case "6":
		func() { testInt64++ }()
	case "7":
		func() { testInt64++ }()
	case "8":
		func() { testInt64++ }()
	case "9":
		func() { testInt64++ }()
	default:
		panic(fmt.Errorf("Unexpected parameter: %v", p))
	}
}

var mapStringSmall = map[string]func(){
	"0": func() { testInt64++ },
	"1": func() { testInt64++ },
	"2": func() { testInt64++ },
	"3": func() { testInt64++ },
	"4": func() { testInt64++ },
	"5": func() { testInt64++ },
	"6": func() { testInt64++ },
	"7": func() { testInt64++ },
	"8": func() { testInt64++ },
	"9": func() { testInt64++ },
}

var randStringSmall = []string{
	"6",
	"7",
	"0",
	"2",
	"0",
	"0",
	"3",
	"6",
	"6",
	"3",
}
