package lang

import (
	"bufio"
	"strings"
	"strconv"
	"errors"
	"io"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

type Parser struct {
	State *CanvasState
}

func NewParserWithState(state *CanvasState) *Parser {
    return &Parser{State: state}
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
				p.State.SetBgColor(painter.OperationFunc(painter.WhiteFill))
			case "green":
				p.State.SetBgColor(painter.OperationFunc(painter.GreenFill))
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

				p.State.Rect = painter.BgRect(x1Int, y1Int, x2Int, y2Int)
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

				p.State.AddFigure(&painter.Figure{
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

				p.State.MoveFigures(dxInt, dyInt)
			case "update":
				res = append(res, p.State.Update()...)
			case "reset":
				p.State.Reset()
			default:
				return nil, errors.New("invalid command")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
