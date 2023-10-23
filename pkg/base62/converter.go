package base62

type Converter interface {
	Encode(n int64) string
	Decode(s string) int64
}
