package heap

// 基本行为
type Item struct {
	v int
	name string
}

func (i Item) Gt(j Item) bool {
	if i.v > j.v {
		return true
	}

	return false
}

func (i Item) Lt(j Item) bool {
	if i.v < j.v {
		return true
	}
	return false
}
