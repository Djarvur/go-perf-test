package passingparams_test

import "math/rand"

type TV001 struct {
	I001 int64
}

func NewTV001() TV001 {
	return TV001{
		I001: rand.Int63(),
	}
}

type TP001 struct {
	I001 int64
}

func NewTP001() *TP001 {
	return &TP001{
		I001: rand.Int63(),
	}
}

type TV002 struct {
	I001 int64
	S001 TV001
	I002 int64
	S002 TV001
}

func NewTV002() TV002 {
	return TV002{
		I001: rand.Int63(),
		S001: NewTV001(),
		I002: rand.Int63(),
		S002: NewTV001(),
	}
}

type TP002 struct {
	I001 int64
	S001 *TP001
	I002 int64
	S002 *TP001
}

func NewTP002() *TP002 {
	return &TP002{
		I001: rand.Int63(),
		S001: NewTP001(),
		I002: rand.Int63(),
		S002: NewTP001(),
	}
}

type TV003 struct {
	I001 int64
	S001 TV002
	I002 int64
	S002 TV002
	I003 int64
	S003 TV002
}

func NewTV003() TV003 {
	return TV003{
		I001: rand.Int63(),
		S001: NewTV002(),
		I002: rand.Int63(),
		S002: NewTV002(),
		I003: rand.Int63(),
		S003: NewTV002(),
	}
}

type TP003 struct {
	I001 int64
	S001 *TP002
	I002 int64
	S002 *TP002
	I003 int64
	S003 *TP002
}

func NewTP003() *TP003 {
	return &TP003{
		I001: rand.Int63(),
		S001: NewTP002(),
		I002: rand.Int63(),
		S002: NewTP002(),
		I003: rand.Int63(),
		S003: NewTP002(),
	}
}

type TV004 struct {
	I001 int64
	S001 TV003
	I002 int64
	S002 TV003
	I003 int64
	S003 TV003
	I004 int64
	S004 TV003
}

func NewTV004() TV004 {
	return TV004{
		I001: rand.Int63(),
		S001: NewTV003(),
		I002: rand.Int63(),
		S002: NewTV003(),
		I003: rand.Int63(),
		S003: NewTV003(),
		I004: rand.Int63(),
		S004: NewTV003(),
	}
}

type TP004 struct {
	I001 int64
	S001 *TP003
	I002 int64
	S002 *TP003
	I003 int64
	S003 *TP003
	I004 int64
	S004 *TP003
}

func NewTP004() *TP004 {
	return &TP004{
		I001: rand.Int63(),
		S001: NewTP003(),
		I002: rand.Int63(),
		S002: NewTP003(),
		I003: rand.Int63(),
		S003: NewTP003(),
		I004: rand.Int63(),
		S004: NewTP003(),
	}
}

type TV005 struct {
	I001 int64
	S001 TV004
	I002 int64
	S002 TV004
	I003 int64
	S003 TV004
	I004 int64
	S004 TV004
	I005 int64
	S005 TV004
}

func NewTV005() TV005 {
	return TV005{
		I001: rand.Int63(),
		S001: NewTV004(),
		I002: rand.Int63(),
		S002: NewTV004(),
		I003: rand.Int63(),
		S003: NewTV004(),
		I004: rand.Int63(),
		S004: NewTV004(),
		I005: rand.Int63(),
		S005: NewTV004(),
	}
}

type TP005 struct {
	I001 int64
	S001 *TP004
	I002 int64
	S002 *TP004
	I003 int64
	S003 *TP004
	I004 int64
	S004 *TP004
	I005 int64
	S005 *TP004
}

func NewTP005() *TP005 {
	return &TP005{
		I001: rand.Int63(),
		S001: NewTP004(),
		I002: rand.Int63(),
		S002: NewTP004(),
		I003: rand.Int63(),
		S003: NewTP004(),
		I004: rand.Int63(),
		S004: NewTP004(),
		I005: rand.Int63(),
		S005: NewTP004(),
	}
}

type TV006 struct {
	I001 int64
	S001 TV005
	I002 int64
	S002 TV005
	I003 int64
	S003 TV005
	I004 int64
	S004 TV005
	I005 int64
	S005 TV005
	I006 int64
	S006 TV005
}

func NewTV006() TV006 {
	return TV006{
		I001: rand.Int63(),
		S001: NewTV005(),
		I002: rand.Int63(),
		S002: NewTV005(),
		I003: rand.Int63(),
		S003: NewTV005(),
		I004: rand.Int63(),
		S004: NewTV005(),
		I005: rand.Int63(),
		S005: NewTV005(),
		I006: rand.Int63(),
		S006: NewTV005(),
	}
}

type TP006 struct {
	I001 int64
	S001 *TP005
	I002 int64
	S002 *TP005
	I003 int64
	S003 *TP005
	I004 int64
	S004 *TP005
	I005 int64
	S005 *TP005
	I006 int64
	S006 *TP005
}

func NewTP006() *TP006 {
	return &TP006{
		I001: rand.Int63(),
		S001: NewTP005(),
		I002: rand.Int63(),
		S002: NewTP005(),
		I003: rand.Int63(),
		S003: NewTP005(),
		I004: rand.Int63(),
		S004: NewTP005(),
		I005: rand.Int63(),
		S005: NewTP005(),
		I006: rand.Int63(),
		S006: NewTP005(),
	}
}

type TV007 struct {
	I001 int64
	S001 TV006
	I002 int64
	S002 TV006
	I003 int64
	S003 TV006
	I004 int64
	S004 TV006
	I005 int64
	S005 TV006
	I006 int64
	S006 TV006
	I007 int64
	S007 TV006
}

func NewTV007() TV007 {
	return TV007{
		I001: rand.Int63(),
		S001: NewTV006(),
		I002: rand.Int63(),
		S002: NewTV006(),
		I003: rand.Int63(),
		S003: NewTV006(),
		I004: rand.Int63(),
		S004: NewTV006(),
		I005: rand.Int63(),
		S005: NewTV006(),
		I006: rand.Int63(),
		S006: NewTV006(),
		I007: rand.Int63(),
		S007: NewTV006(),
	}
}

type TP007 struct {
	I001 int64
	S001 *TP006
	I002 int64
	S002 *TP006
	I003 int64
	S003 *TP006
	I004 int64
	S004 *TP006
	I005 int64
	S005 *TP006
	I006 int64
	S006 *TP006
	I007 int64
	S007 *TP006
}

func NewTP007() *TP007 {
	return &TP007{
		I001: rand.Int63(),
		S001: NewTP006(),
		I002: rand.Int63(),
		S002: NewTP006(),
		I003: rand.Int63(),
		S003: NewTP006(),
		I004: rand.Int63(),
		S004: NewTP006(),
		I005: rand.Int63(),
		S005: NewTP006(),
		I006: rand.Int63(),
		S006: NewTP006(),
		I007: rand.Int63(),
		S007: NewTP006(),
	}
}

type TV008 struct {
	I001 int64
	S001 TV007
	I002 int64
	S002 TV007
	I003 int64
	S003 TV007
	I004 int64
	S004 TV007
	I005 int64
	S005 TV007
	I006 int64
	S006 TV007
	I007 int64
	S007 TV007
	I008 int64
	S008 TV007
}

func NewTV008() TV008 {
	return TV008{
		I001: rand.Int63(),
		S001: NewTV007(),
		I002: rand.Int63(),
		S002: NewTV007(),
		I003: rand.Int63(),
		S003: NewTV007(),
		I004: rand.Int63(),
		S004: NewTV007(),
		I005: rand.Int63(),
		S005: NewTV007(),
		I006: rand.Int63(),
		S006: NewTV007(),
		I007: rand.Int63(),
		S007: NewTV007(),
		I008: rand.Int63(),
		S008: NewTV007(),
	}
}

type TP008 struct {
	I001 int64
	S001 *TP007
	I002 int64
	S002 *TP007
	I003 int64
	S003 *TP007
	I004 int64
	S004 *TP007
	I005 int64
	S005 *TP007
	I006 int64
	S006 *TP007
	I007 int64
	S007 *TP007
	I008 int64
	S008 *TP007
}

func NewTP008() *TP008 {
	return &TP008{
		I001: rand.Int63(),
		S001: NewTP007(),
		I002: rand.Int63(),
		S002: NewTP007(),
		I003: rand.Int63(),
		S003: NewTP007(),
		I004: rand.Int63(),
		S004: NewTP007(),
		I005: rand.Int63(),
		S005: NewTP007(),
		I006: rand.Int63(),
		S006: NewTP007(),
		I007: rand.Int63(),
		S007: NewTP007(),
		I008: rand.Int63(),
		S008: NewTP007(),
	}
}

type TV009 struct {
	I001 int64
	S001 TV008
	I002 int64
	S002 TV008
	I003 int64
	S003 TV008
	I004 int64
	S004 TV008
	I005 int64
	S005 TV008
	I006 int64
	S006 TV008
	I007 int64
	S007 TV008
	I008 int64
	S008 TV008
	I009 int64
	S009 TV008
}

func NewTV009() TV009 {
	return TV009{
		I001: rand.Int63(),
		S001: NewTV008(),
		I002: rand.Int63(),
		S002: NewTV008(),
		I003: rand.Int63(),
		S003: NewTV008(),
		I004: rand.Int63(),
		S004: NewTV008(),
		I005: rand.Int63(),
		S005: NewTV008(),
		I006: rand.Int63(),
		S006: NewTV008(),
		I007: rand.Int63(),
		S007: NewTV008(),
		I008: rand.Int63(),
		S008: NewTV008(),
		I009: rand.Int63(),
		S009: NewTV008(),
	}
}

type TP009 struct {
	I001 int64
	S001 *TP008
	I002 int64
	S002 *TP008
	I003 int64
	S003 *TP008
	I004 int64
	S004 *TP008
	I005 int64
	S005 *TP008
	I006 int64
	S006 *TP008
	I007 int64
	S007 *TP008
	I008 int64
	S008 *TP008
	I009 int64
	S009 *TP008
}

func NewTP009() *TP009 {
	return &TP009{
		I001: rand.Int63(),
		S001: NewTP008(),
		I002: rand.Int63(),
		S002: NewTP008(),
		I003: rand.Int63(),
		S003: NewTP008(),
		I004: rand.Int63(),
		S004: NewTP008(),
		I005: rand.Int63(),
		S005: NewTP008(),
		I006: rand.Int63(),
		S006: NewTP008(),
		I007: rand.Int63(),
		S007: NewTP008(),
		I008: rand.Int63(),
		S008: NewTP008(),
		I009: rand.Int63(),
		S009: NewTP008(),
	}
}

type TV010 struct {
	I001 int64
	S001 TV009
	I002 int64
	S002 TV009
	I003 int64
	S003 TV009
	I004 int64
	S004 TV009
	I005 int64
	S005 TV009
	I006 int64
	S006 TV009
	I007 int64
	S007 TV009
	I008 int64
	S008 TV009
	I009 int64
	S009 TV009
	I010 int64
	S010 TV009
}

func NewTV010() TV010 {
	return TV010{
		I001: rand.Int63(),
		S001: NewTV009(),
		I002: rand.Int63(),
		S002: NewTV009(),
		I003: rand.Int63(),
		S003: NewTV009(),
		I004: rand.Int63(),
		S004: NewTV009(),
		I005: rand.Int63(),
		S005: NewTV009(),
		I006: rand.Int63(),
		S006: NewTV009(),
		I007: rand.Int63(),
		S007: NewTV009(),
		I008: rand.Int63(),
		S008: NewTV009(),
		I009: rand.Int63(),
		S009: NewTV009(),
		I010: rand.Int63(),
		S010: NewTV009(),
	}
}

type TP010 struct {
	I001 int64
	S001 *TP009
	I002 int64
	S002 *TP009
	I003 int64
	S003 *TP009
	I004 int64
	S004 *TP009
	I005 int64
	S005 *TP009
	I006 int64
	S006 *TP009
	I007 int64
	S007 *TP009
	I008 int64
	S008 *TP009
	I009 int64
	S009 *TP009
	I010 int64
	S010 *TP009
}

func NewTP010() *TP010 {
	return &TP010{
		I001: rand.Int63(),
		S001: NewTP009(),
		I002: rand.Int63(),
		S002: NewTP009(),
		I003: rand.Int63(),
		S003: NewTP009(),
		I004: rand.Int63(),
		S004: NewTP009(),
		I005: rand.Int63(),
		S005: NewTP009(),
		I006: rand.Int63(),
		S006: NewTP009(),
		I007: rand.Int63(),
		S007: NewTP009(),
		I008: rand.Int63(),
		S008: NewTP009(),
		I009: rand.Int63(),
		S009: NewTP009(),
		I010: rand.Int63(),
		S010: NewTP009(),
	}
}
