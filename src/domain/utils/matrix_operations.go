package domain_utils

type Pos struct {
	R int
	C int
}

type Transform struct {
	Name   string
	Grid   [][]rune
	PosMap [][]Pos
}

// This function builds a new transformed Grid based on the provided mapping function
func BuildTransform(m [][]rune, Rnew, Cnew int, mapper func(r, c int) (int, int), Name string) Transform {
	g := make([][]rune, Rnew)
	PosMap := make([][]Pos, Rnew)
	for r := 0; r < Rnew; r++ {
		g[r] = make([]rune, Cnew)
		PosMap[r] = make([]Pos, Cnew)
		for c := 0; c < Cnew; c++ {
			origR, origC := mapper(r, c)
			g[r][c] = m[origR][origC]
			PosMap[r][c] = Pos{R: origR, C: origC}
		}
	}
	return Transform{Name: Name, Grid: g, PosMap: PosMap}
}

func MakeTransform(Name string, Grid [][]rune, diag bool) []Transform {
	nR := len(Grid)
	nC := len(Grid[0])

	transform := []Transform{}

	R := nR
	C := nC

	// Identity
	transform = append(transform, BuildTransform(Grid, R, C,
		func(r, c int) (int, int) {
			return r, c
		},
		Name+"Identity"),
	)

	// Rotate 90 degrees
	// transform = append(transform, BuildTransform(Grid, R, C,
	// 	func(r, c int) (int, int) {
	// 		return nR - 1 - r, c
	// 	},
	// 	"Rotate90"),
	// )

	// // Rotate 180 degrees
	// transform = append(transform, BuildTransform(Grid, R, C,
	// 	func(r, c int) (int, int) {
	// 		return nR - 1 - r, nC - 1 - c
	// 	},
	// 	"Rotate180"),
	// )

	// // Rotate 270 degrees

	if !diag {
		transform = append(transform, BuildTransform(Grid, R, C,
			func(r, c int) (int, int) {
				return c, nC - 1 - r
			},
			Name+"Rotate270"),
		)
	}

	// Reflect Horizontal
	transform = append(transform, BuildTransform(Grid, R, C,
		func(r, c int) (int, int) {
			return r, nC - 1 - c
		},
		Name+"ReflectH"),
	)

	// Reflect Vertical
	// transform = append(transform, BuildTransform(Grid, R, C,
	// 	func(r, c int) (int, int) {
	// 		return nR - 1 - r, c
	// 	},
	// 	"ReflectV"),
	// )

	// Reflect Diagonal (top-left to bottom-right)
	// transform = append(transform, BuildTransform(Grid, C, R,
	// 	func(r, c int) (int, int) {
	// 		return c, r
	// 	},
	// 	"ReflectD1"),
	// )

	// Reflect Diagonal (top-right to bottom-left)
	if !diag {
		transform = append(transform, BuildTransform(Grid, C, R,
			func(r, c int) (int, int) {
				return nR - 1 - c, nC - 1 - r
			},
			Name+"ReflectD2"),
		)
	}

	return transform
}

// Construye matriz de diagonales ↘ (top-left a bottom-right)
func BuildDiagonalMatrixSE(Grid [][]rune) [][]rune {
	n := len(Grid)
	diagonals := [][]rune{}

	// Primera mitad: diagonales desde fila 0
	for startC := 0; startC < n; startC++ {
		var diag []rune
		r, c := 0, startC
		for r < n && c < n {
			diag = append(diag, Grid[r][c])
			r++
			c++
		}
		diagonals = append(diagonals, diag)
	}

	// Segunda mitad: diagonales desde col 0 (excepto la primera)
	for startR := 1; startR < n; startR++ {
		var diag []rune
		r, c := startR, 0
		for r < n && c < n {
			diag = append(diag, Grid[r][c])
			r++
			c++
		}
		diagonals = append(diagonals, diag)
	}

	combinedDiagonals := combineShortDiagonals(diagonals, n)

	return combinedDiagonals
}

// Construye matriz de diagonales ↙ (top-right a bottom-left)
func BuildDiagonalMatrixSW(Grid [][]rune) [][]rune {
	n := len(Grid)
	diagonals := [][]rune{}

	// Primera mitad: diagonales desde fila 0
	for startC := n - 1; startC >= 0; startC-- {
		var diag []rune
		r, c := 0, startC
		for r < n && c >= 0 {
			diag = append(diag, Grid[r][c])
			r++
			c--
		}
		diagonals = append(diagonals, diag)
	}

	// Segunda mitad: diagonales desde última col (excepto la primera)
	for startR := 1; startR < n; startR++ {
		var diag []rune
		r, c := startR, n-1
		for r < n && c >= 0 {
			diag = append(diag, Grid[r][c])
			r++
			c--
		}
		diagonals = append(diagonals, diag)
	}
	combinedDiagonals := combineShortDiagonals(diagonals, n)

	return combinedDiagonals
}

func combineShortDiagonals(diags [][]rune, targetLen int) [][]rune {
	var combined [][]rune
	i := 0
	for i < targetLen {
		current := diags[i]
		if len(current) < targetLen {
			current = append(current, diags[len(diags)-i]...)
		}
		combined = append(combined, current)
		i++
	}
	return combined
}
