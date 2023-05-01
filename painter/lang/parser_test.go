package lang

import (
	"strings"
	"testing"
	"errors"
	"image/color"

	"github.com/roman-mazur/architecture-lab-3/painter"
	"golang.org/x/exp/shiny/screen"
)

func TestParser_Parse(t *testing.T) {
	testCases := []struct {
		input     string
		expected []painter.Operation
		expectErr error
	}{
		{
			// Test case 1: Testing valid input with one figure and background color
			input: "white,figure 0.5 0.5,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.WhiteFill),
				(&painter.Figure{ 400, 400 }).DrawFigure(),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 2: Testing valid input with multiple figures and background color
			input: "green,figure 0.5 0.5,figure 0.7 0.7,update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.GreenFill),
				(&painter.Figure{ 400, 400 }).DrawFigure(),
				(&painter.Figure{ 560, 560 }).DrawFigure(),
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
				(&painter.Figure{ 400, 400 }).DrawFigure(),
				(&painter.Figure{ 560, 560 }).DrawFigure(),
				painter.BgRect(160, 160, 480, 480),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
		{
			// Test case 6: Testing invalid input with bgrect command with invalid number of arguments
			input: "bgrect 0.1 0.1 0.5,update",
			expected: []painter.Operation{
			},
			expectErr: errors.New("invalid number of arguments for bgrect"),
		},
		{
			// Test case 7: Testing invalid input with figure command with invalid arguments
			input: "figure 0.1,update",
			expected:  nil,
			expectErr: errors.New("invalid number of arguments for figure"),
		},
		{
			// Test case 8: Testing invalid input with move command with invalid arguments
			input: "move 0.1,update",
			expected:  nil,
			expectErr: errors.New("invalid number of arguments for move"),
		},
		{
			// Test case 9: Testing reset method
			input: "white,figure 0.5 0.5,reset,update",
			expected: []painter.Operation{
				painter.OperationFunc(func(t screen.Texture) {
					t.Fill(t.Bounds(), color.Black, screen.Src)
				}),
				painter.BgRect(0, 0, 0, 0),
				painter.UpdateOp,
			},
			expectErr: nil,
		},
	}

	for _, testCase := range testCases {
		parser := Parser{}
		t.Run(testCase.input, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			result, err := parser.Parse(reader)

			if err == nil && testCase.expectErr != nil {
				t.Errorf("Expected error: %v but got nil", testCase.expectErr)
			} else if err != nil && testCase.expectErr == nil {
				t.Errorf("Expected no error but got error: %v", err)
			} else if err != nil && testCase.expectErr != nil && !strings.Contains(err.Error(), testCase.expectErr.Error()) {
				t.Errorf("Expected error: %v but got error: %v", testCase.expectErr, err)
			}			

			if len(result) != len(testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, result)
			}

			for i, operation := range result {
				if operation == nil || testCase.expected[i] == nil {
					if operation != testCase.expected[i] {
						t.Errorf("Expected %v but got %v", testCase.expected, result)
					}
				}
			}
		})
	}
}
