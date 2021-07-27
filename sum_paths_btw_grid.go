package main

//the proble is stated thus:
//we have a grid specified by its width and its height
//we wish to know how many paths we can take from the point (0,0) to (n,m)
//to do this we first think about the smallest case
//we leace it to the caller to keep the contract that neither of the two values of x and y would be 0
func sumPathsBtwGrid(x, y int) int {
	//base case
	//when n,m = 1
	// we realize that here, there can only be one path when either of the two dimensions is one
	//and that path will be straight
	if x == 1 || y == 1 {
		return 1
	} else {
		//done with the base case, so what else?
		//what if neither of the dimensions is 0?
		//looking at exmples, we find that the number of ways to sum paths (n,m)
		//is equal to sum(n-1, m) + sum(n, m-1)
		// this is quite easy to see as all that is needed in both cases is to shift the
		//smaller-by-one cases to the right:  (append?) or shift down, depending on
		//the side we're considering
		return sumPathsBtwGrid(x-1, y) + sumPathsBtwGrid(x, y-1)
	}
	//we have faith in recursion to fix the solution up for us.
	//since our pattern works out fine heuristically, we need not worry
	//bring as many test cases as you want, we're at home for you

}
