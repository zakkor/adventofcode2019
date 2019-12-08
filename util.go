package main

import "io/ioutil"

func readInput(file string) string {
	inputData, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(inputData)
}
