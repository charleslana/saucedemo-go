package base

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

type BaseTest struct {
	Ctx      context.Context
	Cancel   context.CancelFunc
	AllocCtx context.Context
	Alloc    context.CancelFunc
}

func NewBaseTest() *BaseTest {
	return &BaseTest{}
}

func (b *BaseTest) Setup(headless bool) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		chromedp.WindowSize(1280, 720),
	)

	b.AllocCtx, b.Alloc = chromedp.NewExecAllocator(context.Background(), opts...)

	b.Ctx, b.Cancel = context.WithTimeout(b.AllocCtx, 15*time.Second)

	b.Ctx, _ = chromedp.NewContext(b.Ctx)
}

func (b *BaseTest) Cleanup() {
	if b.Cancel != nil {
		b.Cancel()
	}
	if b.Alloc != nil {
		b.Alloc()
	}
}
