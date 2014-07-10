package kmp_test

import (
	"testing"

	"github.com/rdallman/kmp"
)

func TestAFewThings(t *testing.T) {
	pat := "hi"
	T := kmp.FailT(pat)
	txt := `bacon ipsum dolor`

	one := kmp.Contains(txt, pat, T)
	two := kmp.Contains(txt, pat, nil)
	if one != two {
		t.Fatal("your whole life is a lie")
	} else if one || two {
		t.Fatal("time to get a pet unicorn")
	}

	txt = "hihi hi hehihe hexehi"

	if !kmp.Contains(txt, pat, T) {
		t.Fatal("dinosaurs are a myth")
	}

	three := kmp.Find(txt, pat, T)
	if len(three) != 5 {
		t.Fatal("your dad isn't really proud of you")
	}

	four := kmp.FindWords(txt, pat, T)
	if len(four) != 1 {
		t.Fatal("nobody loves you", four)
	}

	five := kmp.Find(txt, pat, nil)
	if five[0] != 0 || five[1] != 2 || five[2] != 5 ||
		five[3] != 10 || five[4] != 19 {
		t.Fatal(`I'm glad somebody found my work and all, but you really
        should probably go outside`, five)
	}

}
