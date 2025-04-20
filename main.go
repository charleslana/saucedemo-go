package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.WindowSize(1280, 720),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := context.WithTimeout(allocCtx, 15*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var pageTitle string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://example.com`),
		chromedp.Title(&pageTitle),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Título da página:", pageTitle)
}
