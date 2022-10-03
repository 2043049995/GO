package codes

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if 4 != Add(1, 3) {
		t.Errorf("1+5!=6")
	}
}
func BenchmarkAdd(b *testing.B) {

}
