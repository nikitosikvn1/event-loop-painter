package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"golang.org/x/exp/shiny/screen"
) 

func TestLoop_Start(t *testing.T) {
	s := &mockScreen{}
	l := &Loop{
		Receiver: &testReceiver{},
	}
	l.Start(s)

	if l.next == nil || l.prev == nil {
		t.Error("unexpected nil texture")
	}

	l.StopAndWait()
}

func TestLoop_Post(t *testing.T){
	var (
		l Loop
		tr testReceiver
	)

	l.Receiver = &tr 

	l.Start(mockScreen{})
	l.Post(OperationFunc(WhiteFill))
	l.Post(OperationFunc(GreenFill))
	l.Post(UpdateOp)
	if tr.LastTexture != nil {
		t.Fatal("Reciever got the texture too early")
	}
	l.StopAndWait()

	tx, ok := tr.LastTexture.(*mockTexture)
	if !ok {
		t.Fatal("Reciever still has not texture")
	}
	if tx.FillCnt != 2 {
		t.Error("Unexpected number of fill calls:", tx.FillCnt)
	}
	
}

func TestMessageQueue_push_pull_empty(t *testing.T) {
	mq := &messageQueue{}
	op := &testOperation{}

	mq.push(op)

	if len(mq.ops) != 1 || mq.ops[0] != op {
		t.Error("failed to push operation into queue")
	}

	pulledOp := mq.pull()
	if pulledOp != op {
		t.Error("failed to pull operation from queue")
	}

	if !mq.empty() {
		t.Error("expected queue to be empty")
	}
}

func TestMessageQueue_push_blocked(t *testing.T) {
	mq := &messageQueue{}

	for i := 0; i < 10; i++ {
		mq.push(&testOperation{})
	}

	op := &testOperation{}

	// Push operation and ensure that it's blocked
	mq.push(op)

	if len(mq.ops) != 11 {
		t.Error("failed to push operation into queue")
	}

	if mq.blocked != nil {
		t.Error("expected message queue to be blocked")
	}

	// Remove operation from queue and ensure that it's unblocked
	mq.pull()

	if len(mq.ops) != 10 {
		t.Error("failed to pull operation from queue")
	}

	if mq.blocked != nil {
		t.Error("expected message queue to be unblocked")
	}

	if mq.empty() {
		t.Error("expected queue to not be empty")
	}
}

type testReceiver struct {
	LastTexture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.LastTexture = t
}

type testOperation struct {
	updated bool
}

func (op *testOperation) Do(t screen.Texture) bool {
	op.updated = true
	return true
}

type mockScreen struct {}

func (m mockScreen) NewBuffer(image.Point) (screen.Buffer, error){
	panic("implement me")
}

func (m mockScreen) NewTexture(image.Point) (screen.Texture, error){
	return new(mockTexture), nil
}

func (m mockScreen) NewWindow(*screen.NewWindowOptions) (screen.Window, error){
	panic("implement me")
}

type mockTexture struct {
	FillCnt int
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point { return size }

func (m *mockTexture) Bounds() image.Rectangle {
	 return image.Rectangle{Max: size} 
}

func (m *mockTexture) Upload( image.Point, screen.Buffer, image.Rectangle ) {
	panic("implement me")
}

func (m *mockTexture) Fill( image.Rectangle, color.Color, draw.Op ) {
	m.FillCnt++
}