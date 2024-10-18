// Roman Numeral & Array Duplication Question
// Authors: Ruslan Gavrilov, Joshua Boyce Hyland
// Date: 17-10-24
package main

import "fmt"

func allCharactersAreRomanNumeral(input string) bool {

	numOfRomanNumerals := (0)
	romanNumeralStr := [7]rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}

	for i := int(0); i < len(input); i++ {

		for k := int(0); k < len(romanNumeralStr); k++ {

			if romanNumeralStr[k] == rune(input[i]) {
				numOfRomanNumerals++
			}
		}

	}

	return numOfRomanNumerals == len(input)
}

func romanNumeral(input string) int {

	year := int(0)

	if len(input) > 0 && len(input) < 15 {
		MAX_ROMAN_NUMERALS := int(7)
		romanNumeralStr := [7]rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
		romanNumeralValues := [7]int{1, 5, 10, 50, 100, 500, 1000}

		var previousRomanNumeral rune

		for i := len(input) - 1; i >= 0; i-- {

			for j := int(0); j < MAX_ROMAN_NUMERALS; j++ {

				if romanNumeralStr[j] == rune(input[i]) {

					if romanNumeralValues[j] < year && romanNumeralStr[j] != previousRomanNumeral {
						year -= romanNumeralValues[j] //
					} else {
						year += romanNumeralValues[j]
					}
					previousRomanNumeral = romanNumeralStr[j]

				}
			}

		}

		if year > 3999 || !allCharactersAreRomanNumeral(input) {
			year = 0
		}
	}

	return year
}

// valid in range 1 <= nums.length <= 3 * 104
func arrayIsValid(array []int) bool {
	return len(array) > 1 && len(array) < 3*10^4
} //end of arrayIsValid func

// individual members are -100 <= nums[i] <= 100
func arrayValuesAreInRange(array []int) bool {
	for value := range array {
		if value < -100 || value > 100 {
			return false
		}
	}
	return true
} //end of arrayValuesAreInRange func

// nums is sorted in non-decreasing order
func arraySortedInNonDecreaseOrder(array []int) bool {
	for value := 0; value < len(array)-1; value++ {
		if array[value] > array[value+1] {
			return false
		}
	}
	return true
} // end of arraySortedInNonDecreaseOrder fun

//Checking above conditional functions
func arrayConditions(array []int) bool {
	return arrayIsValid(array) &&
		arrayValuesAreInRange(array) &&
		arraySortedInNonDecreaseOrder(array)
} //end of arrayConditions func

//Function to remove duplicates from array
func removeDuplicates(array []int) (int, []int, error) {
	var tempArray []int
	var dupeArray []int

	if !arrayConditions(array) {
		return 0, nil, fmt.Errorf("array conditions aren't met")
	}

	for i := range len(array) {
		current := array[i]
		dupeTaken := false
		for k := 0; k < len(array); k++ {
			//if lhs value == rhs value, the position of the values are not the same ( don't get yourself )
			if current == array[k] && i != k {
				for dupes := range len(dupeArray) {
					if dupeArray[dupes] == array[k] {
						dupeTaken = true
					}
				}
				if !dupeTaken {
					tempArray = append(tempArray, array[i])
					dupeArray = append(dupeArray, array[i])
					dupeTaken = true
				}
			}
		}
		if !dupeTaken {
			tempArray = append(tempArray, array[i])
		}
	}
	//fmt.Printf("len=%d cap=%d %v\n", len(tempArray), cap(tempArray), tempArray)
	return len(tempArray), tempArray, nil
} //end of remove dupes func

func main() {
	input := []int{3, 3, 4, 4, 5}

	newArraySize, newArray, err := removeDuplicates(input)

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("len=%d %v\n", newArraySize, newArray)

	myArray := make([]string, 5)
	myArray[0] = "IV"
	myArray[1] = "pen"
	myArray[2] = "V"
	myArray[3] = "MMII"
	myArray[4] = "poI"

	for i := 0; i < len(myArray); i++ {

		fmt.Printf(" %s = %d\n", myArray[i], romanNumeral(myArray[i]))
	}
}
