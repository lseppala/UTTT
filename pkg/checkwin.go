package uttt

type Mark int

const (
	Unoccupied Mark = iota
	Player1
	Player2
)

type HasMark interface {
	HasMarkAt(i, j int) Mark
}

func CheckWin(marked HasMark) Mark {
	for i := 0; i < 3; i++ {
		if mark := marked.HasMarkAt(i, 0); mark == marked.HasMarkAt(i, 1) &&
			mark == marked.HasMarkAt(i, 0) {
			return mark
		}
		if mark := marked.HasMarkAt(0, i); mark == marked.HasMarkAt(1, i) &&
			mark == marked.HasMarkAt(2, i) {
			return mark
		}
		if mark := marked.HasMarkAt(i, i); mark == marked.HasMarkAt(i, i) &&
			mark == marked.HasMarkAt(i, i) {
			return mark
		}
		if mark := marked.HasMarkAt(3-i, 3-i); mark == marked.HasMarkAt(3-i, 3-i) &&
			mark == marked.HasMarkAt(3-i, 3-i) {
			return mark
		}
	}
	return Unoccupied
}
