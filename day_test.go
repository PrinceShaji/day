package day

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLayout(t *testing.T) {

	res := GetLayout("DD-MM-YYYY")

	assert.Equal(t, "02-01-2006", res, "DD-MM-YYYY should be 02-01-2006")

}

func BenchmarkGetLayout(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = GetLayout("DD-MM-YYYY")

	}
}

func TestStrptime(t *testing.T) {

	res, err := Strptime("DD-MM-YYYY", "25-04-1996")

	assert.Equal(t, 1996, res.Year(), "year should be 1996")
	assert.Equal(t, "April", res.Month().String(), "month should be April")
	assert.Equal(t, 25, res.Day(), "date should be 25")

	assert.NoError(t, err, "error parsing date")

}

func BenchmarkStrptime(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, _ = Strptime("DD-MM-YYYY", "25-04-1996")

	}

}

func TestStrftime(t *testing.T) {

	date, _ := Strptime("DD-MM-YYYY", "25-04-1996")

	res, err := Strftime("DD-MM-YYYY", date)

	assert.Equal(t, "25-04-1996", res, "date should be 25-04-1996")

	assert.NoError(t, err, "error parsing date")

}

func BenchmarkStrftime(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, _ = Strftime("DD-MM-YYYY", time.Now())

	}

}
