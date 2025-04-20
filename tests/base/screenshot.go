package base

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chromedp/chromedp"
)

func TakeScreenshot(ctx context.Context, filename string) error {
	screenshotDir := filepath.Join("..", "..", "screenshots")

	fmt.Println("Criando diretório de screenshots:", screenshotDir)
	if err := os.MkdirAll(screenshotDir, os.ModePerm); err != nil {
		return fmt.Errorf("erro ao criar diretório para screenshots: %v", err)
	}

	filePath := filepath.Join(screenshotDir, filename)
	fmt.Println("Caminho do arquivo de screenshot:", filePath)

	var buf []byte
	fmt.Println("Capturando screenshot...")
	if err := chromedp.Run(ctx, chromedp.FullScreenshot(&buf, 90)); err != nil {
		return fmt.Errorf("erro ao capturar screenshot: %v", err)
	}

	fmt.Println("Salvando screenshot no arquivo...")
	if err := os.WriteFile(filePath, buf, 0644); err != nil {
		return fmt.Errorf("erro ao salvar o screenshot no arquivo: %v", err)
	}

	fmt.Printf("Screenshot salvo em: %s\n", filePath)
	return nil
}
