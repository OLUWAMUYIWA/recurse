package main

//to recursively solve this problem
//first we have to understan the relationship between one solution for n-1 and n
//and that is quite simple
//turns out that sum(n) = sum(n-1) + n
func sum_numbers(n int) int {
	return n + sum_numbers(n-1)
}
