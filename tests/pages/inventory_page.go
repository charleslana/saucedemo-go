package pages

import (
	"context"
	"fmt"

	"example.com/v2/tests/base"
	"github.com/chromedp/chromedp"
)

const (
	InventoryURL    = "https://www.saucedemo.com/inventory.html"
	InventoryTitle  = "Swag Labs"
	InventoryButton = `.inventory_item`
)

type InventoryPage struct {
	ctx context.Context
}

func NewInventoryPage(ctx context.Context) *InventoryPage {
	return &InventoryPage{ctx}
}

func (p *InventoryPage) IsOnInventoryPage() error {
	var title string

	err := chromedp.Run(p.ctx, chromedp.Title(&title))
	if err != nil {
		return err
	}

	if title != InventoryTitle {
		return fmt.Errorf("esperado título '%s', mas obtido '%s'", InventoryTitle, title)
	}

	err = chromedp.Run(p.ctx, chromedp.WaitVisible(InventoryButton))
	if err != nil {
		base.TakeScreenshot(p.ctx, "inventory_page_visible")
		return fmt.Errorf("o item de inventário não está visível: %v", err)
	}

	return nil
}
