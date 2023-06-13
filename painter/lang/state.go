package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
)

type CanvasState struct {
	BgColor painter.OperationFunc
	Rect    painter.OperationFunc
	Figures []*painter.Figure
}

func NewCanvasState() *CanvasState {
	return &CanvasState{}
}

func (cs *CanvasState) SetBgColor(op painter.OperationFunc) {
	cs.BgColor = op
}

func (cs *CanvasState) SetRect(op painter.OperationFunc) {
	cs.Rect = op
}

func (cs *CanvasState) AddFigure(f *painter.Figure) {
	cs.Figures = append(cs.Figures, f)
}

func (cs *CanvasState) Reset() {
	cs.BgColor = painter.OperationFunc(painter.BlackFill)
	cs.Rect = painter.BgRect(0, 0, 0, 0)
	cs.Figures = nil
}

func (cs *CanvasState) Update() []painter.Operation {
	var res []painter.Operation

	if cs.BgColor != nil {
		res = append(res, cs.BgColor)
	}

	if cs.Rect != nil {
		res = append(res, cs.Rect)
	}

	for _, figureInstance := range cs.Figures {
		res = append(res, figureInstance.DrawFigure())
	}

	res = append(res, painter.UpdateOp)

	return res
}

func (cs *CanvasState) MoveFigures(dx, dy int) {
	for _, figureInstance := range cs.Figures {
		figureInstance.MoveFigure(dx, dy)
	}
}
