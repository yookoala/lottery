package reader

// Collection returns Row of nth row
type Collection interface {
	Len() int
	Row(rowNum int) (Row, error)
}

// Row reads the nth cell from the given row
type Row interface {
	Len() int
	ReadString(cellNum int) string
}

// ReadMulti reads multiple cell string values in a collection
func ReadMulti(c Collection, cellNums ...int) func(rowNum int) (strs []string, err error) {
	return func(rowNum int) (strs []string, err error) {
		r, err := c.Row(rowNum)
		if err != nil {
			return
		}
		strs = make([]string, 0, len(cellNums))
		for _, cellNum := range cellNums {
			strs = append(strs, r.ReadString(cellNum))
		}
		return
	}
}
