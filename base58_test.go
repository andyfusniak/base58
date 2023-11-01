package base58

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

const numStringsToGenerate = 50

var regExpBase58 = regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]+$`)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestBase58GenerateStrings(t *testing.T) {
	for i := 0; i < 1000; i++ {
		l := randNumberOverRange(1, numStringsToGenerate)
		str, err := RandString(l)
		if err != nil {
			t.Fatalf("unexpected rand string failed")
		}

		if len(str) != l {
			t.Fatalf("expected string of length %d, got %d", l, len(str))
		}

		if !regExpBase58.Match([]byte(str)) {
			t.Fatalf("generated base58 string %q failed regex test", str)
		}

		t.Logf("%s", str)
	}
}

func randNumberOverRange(min, max int) int {
	return r.Intn(max-min+1) + min
}
