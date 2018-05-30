package image

// Format は画像の形式を表す
//go:generate enumer -type=Format -transform=kebab format.go
type Format int

const (
	Png Format = iota
	Jpeg
	Gif
)
