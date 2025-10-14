package enums

type Genre int32

const (
	UNSPECIFIED Genre = iota
	ACTION
	COMEDY
	THRILLER
	HORROR
	DRAMA
	BIOGRAPHY
)

var (
	Genre_value = map[string]int32 {
		"UNSPECIFIED": 0,
		"ACTION": 1,
		"COMEDY": 2,
		"THRILLER": 3,
		"HORROR": 4,
		"DRAMA": 5,
		"BIOGRAPHY": 6,
	}
)