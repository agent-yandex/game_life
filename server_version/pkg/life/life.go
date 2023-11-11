package life

import "math/rand"

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) Seed() {
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
	// w.Cells[0][1] = true
	// w.Cells[1][2] = true
	// w.Cells[2][0] = true
	// w.Cells[2][1] = true
	// w.Cells[2][2] = true

	// w.Cells[5][6] = true
	// w.Cells[6][6] = true
	// w.Cells[7][6] = true
}

func (w *World) Neighbours(x, y int) int {
	n_count := 0
	positions := [][]int{
		{-1, -1}, {-1, 0},
		{-1, 1}, {0, -1},
		{0, 1}, {1, -1},
		{1, 0}, {1, 1},
	}

	for _, pos := range positions {
		new_x, new_y := x+pos[0], y+pos[1]

		if new_x >= 0 && new_x < w.Height && new_y >= 0 && new_y < w.Width {
			if w.Cells[new_x][new_y] {
				n_count++
			}
		}
	}

	return n_count
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbours(x, y)
	alive := w.Cells[y][x]

	if n < 4 && n > 1 && alive {
		return true
	}

	if n == 3 && !alive {
		return true
	}

	return false
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}
