package main

import "fmt"

func main() {
	shifts := kmpFailure("hi")
	fmt.Println(kmp("hihihithi", "hi", shifts))
}

func kmpFailure(pat string) (T []int) {
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

func kmp(txt, pat string, T []int) (offsets []int) {
	M, N := len(pat), len(txt)
	for m, i := 0, 0; i < N; i++ {
		for m > 0 && txt[i] != pat[m] {
			m = T[m]
		}
		m++
		if m >= M {
			offsets = append(offsets, i-m+1)
			m = 0
		}
	}
	return offsets
}
