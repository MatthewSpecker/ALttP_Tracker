package main

import (
	"fmt"

	"tracker/dungeon"
	"tracker/inventory"
	//"tracker/keyboard"
	"tracker/menu"
	"tracker/preferences"
	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func displayMainWindowContent(mainWindow fyne.Window, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) {
	mainGrid := container.NewHBox(inventory.Layout(), dungeon.Layout())

	mainWindow.SetContent(mainGrid)
}

func main() {
	var scale float32 = 2.0

	myApp := app.New()
	mainWindow := myApp.NewWindow("ALttP Randomizer Keyboard Tracker")
	mainWindow.SetMaster()
	saveConfig := save.NewSaveFile("./save/")
	preferencesConfig := preferences.NewPreferencesFile("./", mainWindow)
	undoStack := undo_redo.NewUndoRedoStacks()
	inventory, err := inventory.NewInventoryIcons(scale, undoStack, preferencesConfig, saveConfig)
	if err != nil {
		panic(fmt.Errorf("Failed to make inventory: %w", err))
	}

	dungeon, err := dungeon.NewDungeonGrid(scale, undoStack, preferencesConfig, saveConfig)
	if err != nil {
		panic(fmt.Errorf("Failed to make dungeonGrid: %w", err))
	}

	mainMenu := menu.MakeMenu(myApp, mainWindow, undoStack, preferencesConfig, saveConfig, inventory, dungeon)
	mainWindow.SetMainMenu(mainMenu)

	displayMainWindowContent(mainWindow, inventory, dungeon)

	mainWindow.SetFullScreen(preferencesConfig.GetPreferenceBool("Fullscreen"))
	mainWindow.Resize(preferencesConfig.GetWindowSize())

	//keyboard.KeyCheck(mainWindow, inventory)

	mainWindow.ShowAndRun()
	preferencesConfig.SetPreference("Fullscreen", mainWindow.FullScreen())
	preferencesConfig.SetWindowSize(mainWindow)
	preferencesConfig.SavePreferences()
	saveConfig.SaveState()
}
