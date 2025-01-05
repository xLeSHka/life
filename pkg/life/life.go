package life

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	// Создаём тип World с количеством слайсов hight (количество строк)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // Создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}
func (w *World) Neighbors(x, y int) int {
	offset := make([][]int, 3)
	for i := range offset {
		offset[i] = make([]int, 3)
	}
	res := 0
	for i := range offset {
		for j := range offset[i] {
			if j == 1 && i == 1 {
				continue
			}
			if w.Cells[(w.Height+y+i-1)%w.Height][(w.Width+x+j-1)%w.Width] == true {
				res += 1
			}
		}
	}
	return res
}
func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)       // Получим количество живых соседей
	alive := w.Cells[y][x]       // Текущее состояние клетки
	if n < 4 && n > 1 && alive { // Если соседей двое или трое, а клетка жива,
		return true // то следующее её состояние — жива
	}
	if n == 3 && !alive { // Если клетка мертва, но у неё трое соседей,
		return true // клетка оживает
	}

	return false // В любых других случаях — клетка мертва
}
func NextState(oldWorld, newWorld *World) {
	// Переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// Для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}
func (w *World) Seed() {
	// Снова переберём все клетки
	for _, row := range w.Cells {
		for i := range row {
			//rand.Intn(10) возвращает случайное число из диапазона	от 0 до 9
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}
func (w *World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	res := ""
	for row := range w.Cells {
		for _, b := range w.Cells[row] {
			if !b {
				res += brownSquare
			} else {
				res += greenSquare
			}
		}
		if row != w.Height-1 {
			res += "\n"

		}

	}
	return res
}
func (w *World) SaveState(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	res := ""
	for row := range w.Cells {
		for _, b := range w.Cells[row] {
			if !b {
				res += "0"
			} else {
				res += "1"
			}
		}
		if row != w.Height-1 {
			res += "\n"

		}

	}
	_, err = f.WriteString(res)
	if err != nil {
		return err
	}
	return nil
}
func (w *World) LoadState(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	height := 0
	width := 0

	file := make([]string, 0)
	for scanner.Scan() {
		height += 1
		file = append(file, scanner.Text())
		if width == 0 {
			width = len(file[len(file)-1])
		}
		if width != len(file[len(file)-1]) {
			return fmt.Errorf("width not match")
		}
	}
	w.Height = height
	w.Width = width
	cells := make([][]bool, height)
	for row := range cells {
		cells[row] = make([]bool, width)
	}
	w.Cells = cells
	for y, row := range w.Cells {
		for x := range row {
			if string(file[y][x]) == "1" {
				row[x] = true
			}
		}
	}
	return nil
}
func (w *World)RandInit(max int) {
	for _, row := range w.Cells {
		for i := range row {
			//rand.Intn(10) возвращает случайное число из диапазона	от 0 до 9
			if rand.Intn(max) == 1 {
				row[i] = true
			}
		}
	}
}