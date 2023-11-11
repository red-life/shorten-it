package base62

import (
	"github.com/red-life/shorten-it/internal/ports"
	"math"
	"strings"
)

var Seed = [62]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func NewConverter() ports.Converter {
	return &Base62{
		seed: Seed,
	}
}

type Base62 struct {
	seed [62]byte
}

func (b Base62) Encode(n int64) string {
	if n == 0 {
		return string(b.seed[0])
	}
	var result string
	base := int64(len(b.seed))
	for n > 0 {
		index := n % base
		result += string(b.seed[index])
		n = n / base
	}

	return result
}

func (b Base62) Decode(s string) int64 {
	result := int64(0)
	base := int64(len(b.seed))
	for i, char := range s {
		index := strings.Index(string(b.seed[:]), string(char))
		result += int64(index) * int64(math.Pow(float64(base), float64(i)))
	}
	return result
}
