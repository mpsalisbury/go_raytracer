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

func (m *Matrix) timesPoint(p Point) Point {
	pm := toMatrix(p)
	tpm := m.times(pm)
	return tpm.toPoint()
}

func (m *Matrix) timesVector(v Vector) Vector {
	vm := toMatrix(v)
	tvm := m.times(vm)
	return tvm.toVector()
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

func (m *Matrix) determinant() float64 {
	if m.numRows != m.numCols {
		panic("cannot compute determinant of non-square matrix")
	}
	if m.numRows == 2 {
		return m.get(0, 0)*m.get(1, 1) - m.get(1, 0)*m.get(0, 1)
	} else {
		det := 0.0
		for c := 0; c < m.numCols; c++ {
			det += m.get(0, c) * m.cofactor(0, c)
		}
		return det
	}
}

func (m *Matrix) isInvertible() bool {
	return m.determinant() != 0.0
}

func (m *Matrix) inverse() *Matrix {
	if !m.isInvertible() {
		panic("cannot invert noninvertible matrix")
	}

	det := m.determinant()
	i := MakeMatrixWithSize(m.numRows, m.numCols)
	for r := 0; r < m.numRows; r++ {
		for c := 0; c < m.numCols; c++ {
			i.set(c, r, m.cofactor(r, c)/det)
		}
	}
	return i

}

func (m *Matrix) submatrix(rr, rc int) *Matrix {
	s := MakeMatrixWithSize(m.numRows-1, m.numCols-1)
	for c := 0; c < m.numCols-1; c++ {
		for r := 0; r < m.numRows-1; r++ {
			// Skip one row/col if at or beyond removed row/col.
			rowOffset := 0
			if r >= rr {
				rowOffset = 1
			}
			colOffset := 0
			if c >= rc {
				colOffset = 1
			}
			s.set(r, c, m.get(r+rowOffset, c+colOffset))
		}
	}
	return s
}

func (m *Matrix) minor(r, c int) float64 {
	return m.submatrix(r, c).determinant()
}

func (m *Matrix) cofactor(r, c int) float64 {
	cof := m.minor(r, c)
	if (r+c)%2 == 1 {
		cof = -cof
	}
	return cof
}
