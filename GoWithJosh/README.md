# GamesEngineering

* **Author:** Ruslan Gavrilov & Joshua Boyce Hyland
* **Created:** 17/10/2024

### Licence Details
This program is free software: you can redistribute it and/or modify  
it under the terms of the GNU General Public License as published by  
the Free Software Foundation, either version 3 of the License, or (at  
your option) any later version.
 
This program is distributed in the hope that it will be useful, but  
WITHOUT ANY WARRANTY; without even the implied warranty of  
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU  
General Public License for more details.  
You should have received a copy of the GNU General Public License  
along with GNU Emacs.  If not, see the [GNU licence web page](http://www.gnu.org/licenses/).  

##  Description:
Problem 1 : Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

### Example 1:
**Input**: nums = [0,0,1,1,1,2,2,3,3,4]
**Output**: 5, nums = [0,1,2,3,4]
**Explanation**: function should return new array without duplicated variables, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
  
 **Constraints:**  
1. 1 <= nums.length <= 3 * 104
2. -100 <= nums[i] <= 100
3. nums is sorted in non-decreasing order.

##  Description:
Problem 2 : Given a string, the program will convert the string to a decimal number. This is under the cicrumstance a valid roman numeral

### Example 1:
**Input**: romanNumeral = "MMCCCXXXIII" 
**Output**: 2333

### Example 2:
**Input**: romanNumeral = "MCMXCIX" 
**Output**: 1999

### Example 3:
**Input**: romanNumeral = "MMMCDXLIV" 
**Output**: 3444

 **Constraints:**  
1.  <= s.length <= 15
2. s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
3. It is guaranteed that s is a valid roman numeral in the range [1, 3999].

## Project Files
1. **samples.go** source code for solution
2. **Readme.md** this Readme file
3. **removeDupes_test.go** code containing Unit Tests for remove duplicates code
4. **romanNumerals_tests** code containing Unit Tests for roman numerals code

## Running and Installation Code
To run the code do:
> go run samples.go

To run the tests do:
> go test 

or to see more output from the tests:
> go test -v

To install first create an executable with:
> go build samples.go


