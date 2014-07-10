package kmp

import "regexp/syntax"

// Search simply takes a pattern to find within a string,
// returning byte offsets for the beginning of each match.
// If you're going to be searching multiple strings, it would
// be better to compute the table yourself beforehand and use Find.
func Search(txt, pat string) (offsets []int) {
	return Find(txt, pat, FailT(pat))
}

// Contains looks for any matches, including substrings. Given a nil
// T, it will compute one for you. This behavior is similar for all funcs.
func Contains(txt, pat string, T []int) bool {
	return len(FindYourOwnWay(txt, pat, T, 1, noop)) > 0
}

// ContainsWord looks for any entire word matches.
func ContainsWord(txt, pat string, T []int) bool {
	return len(FindYourOwnWay(txt, pat, T, 1, isWord)) > 0
}

// Find is Search but would like to have a pre-computed failure table.
// If you don't want to give it one it will begrudgingly make one for you.
func Find(txt, pat string, T []int) (offsets []int) {
	return FindYourOwnWay(txt, pat, T, -1, noop)
}

// FindWords is similar to Find, but slightly modified to only
// find entire word matches.
func FindWords(txt, pat string, T []int) (offsets []int) {
	return FindYourOwnWay(txt, pat, T, -1, isWord)
}

// FailT computes the failure table. You might want one. They're nice.
// This is the O(m) part of KMP
func FailT(pat string) (T []int) {
	T = make([]int, len(pat))
	T[0] = -1

	for i := 1; i < len(pat); i++ {
		T[i] = T[i-1] + 1
		for T[i] > 0 && pat[i] != pat[T[i]-1] {
			T[i] = T[T[i]-1] + 1
		}
	}
	return T
}

// FindYourOwnWay will FindYourOwnWay 1 <= matches <= n with optional added conditional upon
// match, cond (see isWord). To FindYourOwnWay an infinite number of times, divide by zero, and
// see what the other side of a black hole looks like, let n = -1.
//
// Also the actual KMP algorithm, at least the N in O(M+N).
// FindYourOwnWay is not to be confused with an incorrect Journey lyric. Don't worry, it happens to the best of us.
func FindYourOwnWay(txt, pat string, T []int, n int, cond func(string, int, int) bool) (offsets []int) {
	if T == nil {
		T = FailT(pat)
	}
	M, N := len(pat), len(txt)
	for m, i := 0, 0; i < N; i++ {
		for m >= 0 && txt[i] != pat[m] {
			m = T[m]
		}
		m++
		if m >= M {
			if cond(txt, i-m+1, i) {
				offsets = append(offsets, i-m+1)
			}
			if n > 0 && len(offsets) >= n {
				return offsets
			}
			m = 0
		}
	}
	return offsets
}

// for sensible folk
func noop(txt string, x, y int) bool { return true }

// This function assumes it is given byte indices corresponding to a word,
// and i, n are the beginning and end of that word, respectively.
func isWord(txt string, i, n int) bool {
	if i == 0 {
		return !syntax.IsWordChar(rune(txt[n+1]))
	} else if n == len(txt)-1 {
		return !syntax.IsWordChar(rune(txt[i-1]))
	}
	return !syntax.IsWordChar(rune(txt[i-1])) && !syntax.IsWordChar(rune(txt[n+1]))
}
