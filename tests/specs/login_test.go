package specs

import (
	"log"
	"testing"

	"example.com/v2/tests/base"
	"example.com/v2/tests/pages"
)

func TestLoginSuccess(t *testing.T) {
	// Arrange
	bt := base.NewBaseTest()
	bt.Setup(false)
	defer bt.Cleanup()

	loginPage := pages.NewLoginPage(bt.Ctx)
	inventoryPage := pages.NewInventoryPage(bt.Ctx)

	// Act
	if err := loginPage.Navigate(); err != nil {
		log.Fatal(err)
	}

	err := loginPage.SendUsername("standard_user")
	if err != nil {
		t.Fatal(err)
	}

	err = loginPage.SendPassword("secret_sauce")
	if err != nil {
		log.Fatal(err)
	}

	err = loginPage.ClickLogin()
	if err != nil {
		log.Fatal(err)
	}

	// Assert
	if err := inventoryPage.IsOnInventoryPage(); err != nil {
		log.Fatal(err)
	}
}
