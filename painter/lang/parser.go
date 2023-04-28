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
			fmt.Println(cmd)
			
			switch cmd[0] {
			case "white":
				res = append(res, painter.OperationFunc(painter.WhiteFill))
			
			case "green":
				res = append(res, painter.OperationFunc(painter.GreenFill))

			case "update":
				res = append(res, painter.UpdateOp)

			case "bgrect":
				if len(cmd) != 5 {
					return nil, errors.New("invalid number of arguments for bgrect")
				}

				x1, err1 := strconv.ParseFloat(cmd[1], 64)
				y1, err2 := strconv.ParseFloat(cmd[2], 64)
				x2, err3 := strconv.ParseFloat(cmd[3], 64)
				y2, err4 := strconv.ParseFloat(cmd[4], 64)
				fmt.Println(x1, y1, x2, y2)

				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					return nil, errors.New("invalid arguments for bgrect")
				}

				// Danya pishi tut svoi metod blya

			case "figure":
				if len(cmd) != 3 {
					return nil, errors.New("invalid number of arguments for figure")
				}

				x, err1 := strconv.ParseFloat(cmd[1], 64)
				y, err2 := strconv.ParseFloat(cmd[2], 64)
				fmt.Println(x, y)

				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid arguments for figure")
				}

				// i tut tozhe blya
			
			case "move":
				if len(cmd) != 3 {
					return nil, errors.New("invalid number of arguments for move")
				}
	
				dx, err1 := strconv.ParseFloat(cmd[1], 64)
				dy, err2 := strconv.ParseFloat(cmd[2], 64)
				fmt.Println(dx, dy)
	
				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid arguments for move")
				}

				// tut bi tozhe ne pomeshalo, no hz kak ono rabotat dolzhno blya

			case "reset":
				res = []painter.Operation{}
				res = append(res, painter.OperationFunc(func(t screen.Texture) {
					t.Fill(t.Bounds(), color.Black, screen.Src)
				}))
			}
		}
	}

	// TODO: Реалізувати парсинг команд.
	// res = append(res, painter.OperationFunc(painter.WhiteFill))
	// res = append(res, painter.UpdateOp)

	return res, nil
}
