package day

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLayout(t *testing.T) {

	res := generateLayout("DD-MM-YYYY")

	assert.Equal(t, "02-01-2006", res, "DD-MM-YYYY should be 02-01-2006")

}

func BenchmarkGenerateLayout(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = generateLayout("DD-MM-YYYY")

	}
}
