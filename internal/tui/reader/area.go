package reader

import (
	"image"

	"github.com/arimatakao/comicread/internal/backend"
)

func imageAspect(img image.Image) float64 {
	if img == nil || img.Bounds().Dy() == 0 {
		return 1
	}
	return float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())
}

func imageArea(width, height int, imageAspect float64) backend.Area {
	return imageAreaWithCellAspect(width, height, imageAspect, 0.5)
}

func (m Model) imageArea(imageAspect float64) backend.Area {
	return imageAreaWithCellAspect(m.width, m.height, imageAspect, m.cellAspect())
}

func imageAreaWithCellAspect(width, height int, imageAspect, cellAspect float64) backend.Area {
	availableRows := height - 1 // one header row
	if width < 1 || availableRows < 1 {
		return backend.Area{}
	}
	cols, rows := fitToCells(width, availableRows, imageAspect, cellAspect)
	return backend.Area{
		X:    1 + (width-cols)/2,
		Y:    2 + (availableRows-rows)/2,
		Cols: cols,
		Rows: rows,
	}
}

func (m Model) bookImageArea(slot int, imageAspect float64) backend.Area {
	return bookImageAreaWithCellAspect(m.width, m.height, slot, imageAspect, m.cellAspect())
}

func bookImageAreaWithCellAspect(width, height, slot int, imageAspect, cellAspect float64) backend.Area {
	const gap = 1

	availableRows := height - 1
	if width < 2 || availableRows < 1 {
		return backend.Area{}
	}
	availableCols := width - gap
	leftCols := availableCols / 2
	panelCols, panelX := leftCols, 1
	if slot == 1 {
		panelCols = availableCols - leftCols
		panelX += leftCols + gap
	}
	if panelCols < 1 {
		return backend.Area{}
	}
	cols, rows := fitToCells(panelCols, availableRows, imageAspect, cellAspect)
	return backend.Area{
		X:    panelX + (panelCols-cols)/2,
		Y:    2 + (availableRows-rows)/2,
		Cols: cols,
		Rows: rows,
	}
}

// fitToCells sizes an image of the given aspect ratio to fill a panel of
// panelCols×availableRows terminal cells while keeping its proportions on
// screen; cellAspect is the width/height ratio of one cell.
func fitToCells(panelCols, availableRows int, imageAspect, cellAspect float64) (cols, rows int) {
	if imageAspect <= 0 {
		imageAspect = 1
	}
	if cellAspect <= 0 {
		cellAspect = 0.5
	}
	cols = panelCols
	rows = max(1, int(float64(cols)*cellAspect/imageAspect))
	if rows > availableRows {
		rows = availableRows
		cols = min(panelCols, max(1, int(float64(rows)*imageAspect/cellAspect)))
	}
	return cols, rows
}

func unionAreas(first, second backend.Area) backend.Area {
	if first.Cols < 1 || first.Rows < 1 {
		return second
	}
	if second.Cols < 1 || second.Rows < 1 {
		return first
	}
	left := min(first.X, second.X)
	top := min(first.Y, second.Y)
	right := max(first.X+first.Cols, second.X+second.Cols)
	bottom := max(first.Y+first.Rows, second.Y+second.Rows)
	return backend.Area{X: left, Y: top, Cols: right - left, Rows: bottom - top}
}
