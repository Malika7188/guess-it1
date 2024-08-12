package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	mathskils "mathskils/functions"
)

func main() {
	win := 5
	var values []float64 // initializing variable to store the final data read from filedat

	scanner := bufio.NewScanner(os.Stdin) // scan the filedata line by line
	for scanner.Scan() {
		filecontent := scanner.Text()
		if filecontent == "" { // checks if the file data is empty
			return
		}
		num, err := strconv.ParseFloat(filecontent, 64) // converts the filedata contents to float
		if err != nil {
			fmt.Println(err)
			return
		}
		values = append(values, num) // append the data read from filedata in value variable
		// calling all the math skills functions
		if len(values) >= win {
			values = values[len(values)-win:]
		}
		
		lowerLimit, upperLimit := mathskils.ChecLimit(values)
	
		fmt.Printf("%v  %v\n", int(lowerLimit), int(upperLimit))
	}

}