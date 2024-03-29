package main

import (
	"io/ioutil"
	"sort"
)

func readTypes() ([]string, error) {
	return readTypesDir(typeDir)
}

func readTypesDir(path string) ([]string, error) {
	types := []string{}
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fi := range l {
		switch {
		case !fi.IsDir(), fi.Name() == "_gen":
			continue
		}
		types = append(types, fi.Name())
	}

	sort.Strings(types)

	return types, nil
}
