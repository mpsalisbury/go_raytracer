package raytracer

type Matrix struct {
	// Size of matrix.
	numRows, numCols int
	// Cell values in row-major order.
	val []float64
}

func MakeMatrixWithSize(rows, cols int) *Matrix {
	if rows < 1 || cols < 1 {
		panic("invalid matrix size")
	}
	m := &Matrix{rows, cols, make([]float64, rows*cols)}
	return m
}

func MakeMatrixIdentity(size int) *Matrix {
	m := MakeMatrixWithSize(size, size)
	for i := 0; i < size; i++ {
		m.set(i, i, 1.0)
	}
	return m
}

func MakeMatrix(cells [][]float64) *Matrix {
	numRows := len(cells)
	numCols := len(cells[0])

	m := &Matrix{numRows, numCols, make([]float64, numRows*numCols)}

	for r, rowVals := range cells {
		if len(rowVals) != numCols {
			panic("inconsistent number of columns")
		}
		for c, val := range rowVals {
			m.set(r, c, val)
		}
	}
	return m
}

func (m *Matrix) get(r, c int) float64 {
	i := m.index(r, c)
	return m.val[i]
}

func (m *Matrix) set(r, c int, val float64) {
	i := m.index(r, c)
	m.val[i] = val
}

func (m *Matrix) index(r, c int) int {
	if r < 0 || r >= m.numRows {
		panic("row out of bounds")
	}
	if c < 0 || c >= m.numCols {
		panic("col out of bounds")
	}
	return c*m.numRows + r
}

func (m1 *Matrix) times(m2 *Matrix) *Matrix {
	if m1.numCols != m2.numRows {
		panic("matrix sizes are incompatible for multiply")
	}
	size := m1.numCols
	numRows := m1.numRows
	numCols := m2.numCols
	m := MakeMatrixWithSize(numRows, numCols)
	for c := 0; c < numCols; c++ {
		for r := 0; r < numRows; r++ {
			v := 0.0
			for i := 0; i < size; i++ {
				v += m1.get(r, i) * m2.get(i, c)
			}
			m.set(r, c, v)
		}
	}

	return m
}

func (m *Matrix) transpose() *Matrix {
	t := MakeMatrixWithSize(m.numCols, m.numRows)
	for c := 0; c < m.numCols; c++ {
		for r := 0; r < m.numRows; r++ {
			t.set(c, r, m.get(r, c))
		}
	}

	return t
}
