package display

import (
	"tracker/dungeon"
	"tracker/inventory"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func DisplayMainWindowContent(mainWindow fyne.Window, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) {
	mainGrid := container.NewHBox(inventory.Layout(), dungeon.Layout())

	mainWindow.SetContent(mainGrid)
}