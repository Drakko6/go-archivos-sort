package main

import (
	"fmt"
	"os"
	"sort"
)

func main (){

	var n uint
	fmt.Println("¿Cuántos strings quiere?")
	fmt.Scan(&n)

	var sliceStrings []string

	var i uint =1
	var s string
	for i<= n{
		fmt.Println("String", i, ":")
		fmt.Scan(&s)
		sliceStrings = append(sliceStrings, s)
		i ++
	}

	sort.Strings(sliceStrings)
	ascendente := sliceStrings

	
	archivo, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer archivo.Close()


	archivo2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivo2.Close()

	for _, c:= range ascendente {
		archivo.WriteString(c + "\n")
	}

	sort.Sort(sort.Reverse(sort.StringSlice(sliceStrings)))
	descendente := sliceStrings

	for _, c:= range descendente {
		archivo2.WriteString(c + "\n")
	}

	
}