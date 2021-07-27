package main

//if given a number n, and another number m >= 0, how many ways can we partition n with pieces
//never greater than m?
//this does not look like a case where recursion would help, but it does help!
func countPartitions(n, m int) int {
	//first things first
	//we think of the smallest possible r allowable input we can get
	// if we would not allow some inputs, we can either make this function private(or inaccessible) to other code
	// and create a wrapper function that encloses it while keeping the contract
	//or we could do some error handling. I prefer enclosing te function in another
	//back to base cases!

	//what is possible?
	//if n=0, we have 1, regardles of the number of m. Tell me, how many ways can   we partition zero?
	//if m=0, we have no possible way, so we call it zero.
	if n == 0 {
		return 1
	} else if m == 0 || n < 0 {
		return 0
	} else {
		//but what if we have greater than zreo for either of them
		// we can see that the number of partitions in (5,4)
		//is equal to the sum of two already gotten solutions.
		//they are: countPartitions(n, m-1) which tells us the number of ways to partition n
		//when we have only to use a max of (m-1). Look ath that again, we just lowered the bar by
		//one number, what that means is that we have gotten all the partitions wheere we're limited to
		//one less than our given partition
		//apparently, these partitions will also feature in the bigger one (here, I mean m) because
		//when we move to m, the n has not changed, what has changed is the allowable max partition we're given,
		//so we gain more partitions not loose
		//the second part of the solution is where we determine exactly how partitions that include
		//
		//but what if (n-m) goes less than zero, which it would, we have to fix this case:
		// so we go back to fix the bse case
		//we don't need to worry about when m is negative, because we're subtracting  by 1 from it
		return countPartitions(n-m, m) + countPartitions(n, m-1)
	}
}
