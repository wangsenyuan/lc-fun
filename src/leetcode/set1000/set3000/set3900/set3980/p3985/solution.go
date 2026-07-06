package p3985

func getSum(nums []int) int64 {

	var best int64

	L := fastLongestPalindromes(nums)

	n := len(nums)
	sum := make([]int64, n+1)

	for i, v := range nums {
		sum[i+1] = sum[i] + int64(v)
	}

	for i := range L {
		s := i/2 - L[i]/2
		e := s + L[i]
		best = max(best, sum[e]-sum[s])
	}

	return best

}

/*
*
*

	"""
	   Given a sequence seq, returns a list l such that l[2 * i + 1]
	   holds the length of the longest palindrome centered at seq[i]
	   (which must be odd), l[2 * i] holds the length of the longest
	   palindrome centered between seq[i - 1] and seq[i] (which must be
	   even), and l[2 * len(seq)] holds the length of the longest
	   palindrome centered past the last element of seq (which must be 0,
	   as is l[0]).

	   The actual palindrome for l[i] is seq[s:(s + l[i])] where s is i
	   // 2 - l[i] // 2. (// is integer division.)
*/
func fastLongestPalindromes(seq []int) []int {
	n := len(seq)
	var l []int
	var palLen int
	// Loop invariant: seq[(i - palLen):i] is a palindrome.
	// Loop invariant: len(l) >= 2 * i - palLen. The code path that
	// increments palLen skips the l-filling inner-loop.
	// Loop invariant: len(l) < 2 * i + 1. Any code path that
	// increments i past seqLen - 1 exits the loop early and so skips
	// the l-filling inner loop.
	for i := 0; i < n; {
		// First, see if we can extend the current palindrome.  Note
		// that the center of the palindrome remains fixed.
		if i > palLen && seq[i-palLen-1] == seq[i] {
			palLen += 2
			i++
			continue
		}

		l = append(l, palLen)
		// Now to make further progress, we look for a smaller
		// palindrome sharing the right edge with the current
		// palindrome.  If we find one, we can try to expand it and see
		// where that takes us.  At the same time, we can fill the
		// values for l that we neglected during the loop above. We
		// make use of our knowledge of the length of the previous
		// palindrome (palLen) and the fact that the values of l for
		// positions on the right half of the palindrome are closely
		// related to the values of the corresponding positions on the
		// left half of the palindrome.

		// Traverse backwards starting from the second-to-last index up
		// to the edge of the last palindrome.
		s := len(l) - 2
		e := s - palLen
		found := false
		for j := s; j > e; j-- {
			// d is the value l[j] must have in order for the
			// palindrome centered there to share the left edge with
			// the last palindrome.  (Drawing it out is helpful to
			// understanding why the - 1 is there.)
			d := j - e - 1
			if l[j] == d {
				found = true
				palLen = d
				// We actually want to go to the beginning of the outer
				// loop,
				break
			}
			// Otherwise, we just copy the value over to the right
			// side.  We have to bound l[i] because palindromes on the
			// left side could extend past the left edge of the last
			// palindrome, whereas their counterparts won't extend past
			// the right edge.
			l = append(l, min(d, l[j]))
		}
		if !found {
			palLen = 1
			i++
		}
	}
	// We know from the loop invariant that len(l) < 2 * seqLen + 1, so
	// we must fill in the remaining values of l.

	// Obviously, the last palindrome we're looking at can't grow any
	// more.
	l = append(l, palLen)
	s := len(l) - 2
	e := s - (2*n + 1 - len(l))
	for i := s; i > e; i-- {
		// The d here uses the same formula as the d in the inner loop
		// above.  (Computes distance to left edge of the last
		// palindrome.)
		d := i - e - 1
		l = append(l, min(d, l[i]))
	}
	return l
}
