package entry

import (
	"fmt"
	"sort"
	"time"
)

func (els Els) Limit(n int) Els {
	if len(els) <= n {
		return els
	}
	return els[:n]
}

func (els Els) Args(lang string) Args {
	args := Args{}
	for _, e := range els {
		args = append(args, &Arg{
			El:   e,
			Lang: lang,
		})
	}
	return args
}

func (els Els) Year(year int) Els {
	nl := Els{}

	for _, e := range els {
		if Date(e).Year() == year {
			nl = append(nl, e)
			continue
		}
		break
	}
	return nl
}

func (els Els) Offset(start, end int) Els {
	l := len(els)
	if l < start {
		return Els{}
	}
	if end > l || end <= 0 {
		return els[start:]
	}
	return els[start:end]
}

/*
func (els Els) LazyLoad() Els {
	imgs := 0
	nl := Els{}
	for _, e := range els {
		i, ok := e.(*Image)
		if ok {
			imgs++
			if imgs >= 12 {
				ni := i
				ni.LazyLoad = true
				nl = append(nl, ni)
				continue
			}
		}
		nl = append(nl, e)
	}
	return nl
}
*/

type Asc Els

func (a Asc) Len() int {
	return len(a)
}

func (a Asc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Asc) Less(i, j int) bool {
	return Date(a[i]).Unix() < Date(a[j]).Unix()
}

type Desc Els

func (a Desc) Len() int {
	return len(a)
}

func (a Desc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Desc) Less(i, j int) bool {
	return Date(a[i]).Unix() > Date(a[j]).Unix()
}

/*
type Desc Els

func (a Desc) Len() int {
	return len(a)
}

func (a Desc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Desc) Less(i, j int) bool {
	return Id(a[i]) > Id(a[j])
}
*/

func (sets Sets) Lookup(id string) (*Set, error) {
	for _, s := range sets {
		if s.File.Id != id {
			continue
		}
		return s, nil
	}
	return nil, NotFoundErr(id)
}

func (els Els) LookupPosition(id string) (int, error) {
	for i, e := range els {
		if Id(e) != id {
			continue
		}
		return i, nil
	}
	return -1, NotFoundErr(id)
}

func (els Els) Lookup(id string) (interface{}, error) {
	for _, e := range els {
		if Id(e) != id {
			continue
		}
		return e, nil
	}
	return nil, NotFoundErr(id)
}

/*
func (els Els) FindImage(filename string) (*Image, error) {
	println(filename)
	for _, e := range els {
		switch e.(type) {
		case *Image:
			if e.(*Image).File.Base() == filename {
				return e.(*Image), nil
			}
		}
	}
	return nil, fmt.Errorf("Couldnt find image %v in given Els.", filename)
}
*/

func NotFoundErr(id string) error {
	return fmt.Errorf("Lookup: el not found: %v", id)
}

type Day struct {
	Date time.Time
	Els  Els
}

func (els Els) Desc() Els {
	sort.Sort(Desc(els))
	return els
}

func (els Els) Months() []Els {
	months := []Els{}
	cm := Els{}
	for i, e := range els {
		if i > 0 && isNewMonth(e, els[i-1]) {
			months = append(months, cm)
			cm = Els{e}
			continue
		}
		cm = append(cm, e)
	}
	months = append(months, cm)
	return months
}

func isNewMonth(current, before interface{}) bool {
	cd, err := DateSafe(current)
	if err != nil {
		return true
	}
	bd, err := DateSafe(before)
	if err != nil {
		return true
	}
	return cd.Month() != bd.Month()
}

func (els Els) Days() []Els {
	days := []Els{}
	cd := Els{}
	for i, e := range els {
		if i == 0 || !isNewDay(e, els[i-1]) {
			cd = append(cd, e)
			continue
		}
		days = append(days, cd.Reverse())
		cd = Els{e}
	}
	if len(cd) > 0 {
		days = append(days, cd.Reverse())
	}
	return days
}

func (els Els) DayOrder() Els {
	nl := Els{}
	for _, d := range els.Days() {
		for _, e := range d {
			nl = append(nl, e)
		}
	}
	return nl
}

func (els Els) Exclude() Els {
	l := Els{}
	for _, e := range els {
		i := InfoSafe(e)
		if i["hidden"] == "true" || i["exclude"] == "true" {
			continue
		}
		l = append(l, e)
	}
	return l
}

func (els Els) Reverse() Els {
	n := Els{}
	for i := len(els) - 1; i >= 0; i-- {
		n = append(n, els[i])
	}
	return n
}

func isNewDay(current, before interface{}) bool {
	cdate, err := DateSafe(current)
	if err != nil {
		return true
	}
	bdate, err := DateSafe(before)
	if err != nil {
		return true
	}
	if cdate.Hour() <= 5 {
		cdate = cdate.Add(-(time.Hour * 5))
	}
	if bdate.Hour() <= 5 {
		bdate = bdate.Add(-(time.Hour * 5))
	}
	if bdate.Day() != cdate.Day() {
		return true
	}
	return false
}

func (els Els) MediaGroups() []Els {
	packs := []Els{}

	cp := Els{}

	for i, e := range els {
		if i > 0 && isNewGroup(els[i-1], e) {
			packs = append(packs, cp)
			cp = Els{e}
			continue
		}
		cp = append(cp, e)
	}
	packs = append(packs, cp)

	return packs
}

func isNewGroup(a, b interface{}) bool {
	before := Type(a)
	now := Type(b)
	if before == now {
		return false
	}
	if amongImages(before) != amongImages(now) {
		return true
	}
	return false
}

func amongImages(typ string) bool {
	if typ == "text" {
		return false
	}
	return true
}