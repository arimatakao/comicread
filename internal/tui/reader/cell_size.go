package reader

import "github.com/arimatakao/comicread/internal/backend"

func (m *Model) updateCellSize() {
	renderer, ok := m.backend.(backend.CellSizeRenderer)
	if !ok {
		return
	}
	width, height := terminalCellSize(m.width, m.height)
	renderer.SetCellSize(width, height)
}

func (m Model) cellAspect() float64 {
	if renderer, ok := m.backend.(backend.CellSizeRenderer); ok {
		return renderer.CellAspect()
	}
	return 0.5
}
