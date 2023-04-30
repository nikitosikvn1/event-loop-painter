package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{R: 85, G: 217, B: 104}, screen.Src)
}

// BgRect малює прямокутник по координатам лівого верхнього та правого нижнього кута.
func BgRect(x1, y1, x2, y2 int) OperationFunc {
	return func(t screen.Texture) {
		t.Fill(image.Rect(x1, y1, x2, y2), color.Black, screen.Src)
	}
}

// Структура, яка представляє фігуру варіанту
type Figure struct {
	X int
	Y int
}

// DrawFigure повертає Operation, яка малює фігуру варіанту по координатам центру
func (f *Figure) DrawFigure() OperationFunc {
	return func(t screen.Texture) {
		t.Fill(image.Rect(f.X - 150, f.Y - 100, f.X + 150, f.Y), color.RGBA{255, 255, 0, 1}, screen.Src)
    	t.Fill(image.Rect(f.X - 50, f.Y, f.X + 50, f.Y + 100), color.RGBA{255, 255, 0, 1}, screen.Src)
	}
}

// MoveFigure змінює координати центру фігури
func (f *Figure) MoveFigure(x, y int) {
    f.X += x
    f.Y += y
}