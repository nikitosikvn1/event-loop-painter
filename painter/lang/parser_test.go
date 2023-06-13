package lang

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

func TestParser_Parse(t *testing.T) {
	testCases := []struct {
		input     string
		expected  []painter.Operation
		expectErr error
		expectedFigures []painter.Figure
	}{
		{
			// Test case 1: Testing valid input with one figure and background color
			input: "white,figure 0.5 0.5,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.WhiteFill),
				(&painter.Figure{X: 400, Y: 400}).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 2: Testing valid input with multiple figures and background color
			input: "green,figure 0.5 0.5,figure 0.7 0.7,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.GreenFill),
				(&painter.Figure{X: 400, Y: 400}).DrawFigure(),
				(&painter.Figure{X: 560, Y: 560}).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 3: Testing valid input with single bgrect and background color
			input: "green,bgrect 0.1 0.1 0.5 0.5,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.GreenFill),
				painter.BgRect(80, 80, 400, 400),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 4: Testing valid input with multiple bgrect and background color
			input: "white,bgrect 0.1 0.1 0.5 0.5,bgrect 0.2 0.2 0.6 0.6,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.WhiteFill),
				painter.BgRect(160, 160, 480, 480),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 5: Testing valid input with multiple figures, multiple bgrect and background color
			input: "white,figure 0.5 0.5,figure 0.7 0.7,bgrect 0.1 0.1 0.5 0.5,bgrect 0.2 0.2 0.6 0.6,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.WhiteFill),
				painter.BgRect(160, 160, 480, 480),
				(&painter.Figure{X: 400, Y: 400}).DrawFigure(),
				(&painter.Figure{X: 560, Y: 560}).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 6: Testing invalid input with bgrect command with invalid number of arguments
			input:     "bgrect 0.1 0.1 0.5,update",
			expected:  []painter.Operation{},
			expectErr: errors.New("invalid number of arguments for bgrect"),
		},
		{
			// Test case 7: Testing invalid input with figure command with invalid arguments
			input:     "figure 0.1,update",
			expected:  nil,
			expectErr: errors.New("invalid number of arguments for figure"),
		},
		{
			// Test case 8: Testing invalid input with move command with invalid arguments
			input:     "move 0.1,update",
			expected:  nil,
			expectErr: errors.New("invalid number of arguments for move"),
		},
		{
			// Test case 9: Testing reset method
			input: "white,figure 0.5 0.5,reset,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.BlackFill),
				painter.BgRect(0, 0, 0, 0),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 10: Testing move method with one figure
			input: "green,figure 0.4 0.4,move 0.2 0.1,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.GreenFill),
				(&painter.Figure{X: 320, Y: 320}).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
			expectedFigures: []painter.Figure{
				{X: 480, Y: 400},
			},
		},
		{
			// Test case 11: Testing move method with several figures
			input: "green,figure 0.3 0.5,figure 0.7 0.4,move 0.2 -0.1,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.GreenFill),
				(&painter.Figure{X: 240, Y: 400}).DrawFigure(),
				(&painter.Figure{X: 560, Y: 320}).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
			expectedFigures: []painter.Figure{
				{X: 400, Y: 320},
				{X: 720, Y: 240},
			},
		},
		{
			// Test case 12: Check if the coordinates of a shape drawn after the “move” command change
			input: "green,figure 0.3 0.5,figure 0.7 0.4,move 0.2 -0.1,figure 0.1 0.15,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.GreenFill),
				(&painter.Figure{X: 240, Y: 400}).DrawFigure(),
				(&painter.Figure{X: 560, Y: 320}).DrawFigure(),
				(&painter.Figure{X: 80, Y: 120}).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
			expectedFigures: []painter.Figure{
				{X: 400, Y: 320},
				{X: 720, Y: 240},
				{X: 80, Y: 120},
			},
		},
	}

	for _, testCase := range testCases {
		state := NewCanvasState()
		parser := NewParserWithState(state)

		t.Run(testCase.input, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			result, err := parser.Parse(reader)

			// Error checking
			if err == nil && testCase.expectErr != nil {
				t.Errorf("Expected error: %v but got nil", testCase.expectErr)

			} else if err != nil && testCase.expectErr == nil {
				t.Errorf("Expected no error but got error: %v", err)

			} else if err != nil && testCase.expectErr != nil && !strings.Contains(err.Error(), testCase.expectErr.Error()) {
				t.Errorf("Expected error: %v but got error: %v", testCase.expectErr, err)
			}

			// Comparisons of pointers in slices
			if !equalOperations(testCase.expected, result) {
				t.Errorf("Expected slice: %v but got: %v", testCase.expected, result)
			}

			// Coordinate comparisons (used to check for move)
			if testCase.expectedFigures != nil {
				for i, figureInstance := range testCase.expectedFigures {
					if figureInstance.X != parser.State.Figures[i].X || figureInstance.Y != parser.State.Figures[i].Y {
						t.Errorf("Expected coords: (%d, %d) but got: (%d, %d)",figureInstance.X, figureInstance.Y, parser.State.Figures[i].X, parser.State.Figures[i].Y)
					}
				}
			}
		})
	}
}

func equalOperations(op1, op2 []painter.Operation) bool {
	if len(op1) != len(op2) {
		return false
	}

	for i, op := range op1 {
		if fmt.Sprintf("%v", op) != fmt.Sprintf("%v", op2[i]) {
			return false
		}
	}
	return true
}
