package lang

import (
	"bufio"
	"errors"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
	"golang.org/x/exp/shiny/screen"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	BgColor painter.OperationFunc
	Rect painter.OperationFunc
	Figures []*painter.Figure
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation
	
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), ",")
		if len(commands) == 0 {
			continue
		}
		
		for _, val := range commands {
			cmd := strings.Fields(val)
			
			switch cmd[0] {
			case "white":
				p.BgColor = painter.OperationFunc(painter.WhiteFill)
			
			case "green":
				p.BgColor = painter.OperationFunc(painter.GreenFill)

			case "update":
				if p.BgColor != nil {
					res = append(res, p.BgColor)
				}
				
				if p.Rect != nil {
					res = append(res, p.Rect)
				}

				for ind, figureInstance := range p.Figures {
					res = append(res, figureInstance.DrawFigure())
					fmt.Printf("FigureInstance %d: X: %d, Y: %d\n", ind, figureInstance.X, figureInstance.Y)
				}

				res = append(res, painter.UpdateOp)

			case "bgrect":
				if len(cmd) != 5 {
					return nil, errors.New("invalid number of arguments for bgrect")
				}

				x1, err1 := strconv.ParseFloat(cmd[1], 64)
				y1, err2 := strconv.ParseFloat(cmd[2], 64)
				x2, err3 := strconv.ParseFloat(cmd[3], 64)
				y2, err4 := strconv.ParseFloat(cmd[4], 64)

				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					return nil, errors.New("invalid arguments for bgrect")
				}

				x1Int := int(x1 * 800)
				y1Int := int(y1 * 800)
				x2Int := int(x2 * 800)
				y2Int := int(y2 * 800)

				p.Rect = painter.BgRect(x1Int, y1Int, x2Int, y2Int)

			case "figure":
				if len(cmd) != 3 {
					return nil, errors.New("invalid number of arguments for figure")
				}

				x, err1 := strconv.ParseFloat(cmd[1], 64)
				y, err2 := strconv.ParseFloat(cmd[2], 64)

				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid arguments for figure")
				}

				xInt := int(x * 800)
				yInt := int(y * 800)

				p.Figures = append(p.Figures, &painter.Figure{
					X: xInt,
					Y: yInt,
				})
			
			case "move":
				if len(cmd) != 3 {
					return nil, errors.New("invalid number of arguments for move")
				}
	
				dx, err1 := strconv.ParseFloat(cmd[1], 64)
				dy, err2 := strconv.ParseFloat(cmd[2], 64)

				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid arguments for move")
				}
				
				dxInt := int(dx * 800)
				dyInt := int(dy * 800)

				for _, figureInstance := range p.Figures {
					figureInstance.MoveFigure(dxInt, dyInt)
				}

			case "reset":
				p.BgColor = painter.OperationFunc(func(t screen.Texture) {
					t.Fill(t.Bounds(), color.Black, screen.Src)
				})
				p.Rect = painter.BgRect(0, 0, 0, 0)
				p.Figures = nil
			}
		}
	}

	// TODO: Реалізувати парсинг команд.
	// res = append(res, painter.OperationFunc(painter.WhiteFill))
	// res = append(res, painter.UpdateOp)

	return res, nil
}
