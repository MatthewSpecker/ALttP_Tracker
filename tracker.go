package main

import (
	"fmt"

	"tracker/dungeon"
	"tracker/inventory"
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
	myApp := app.New()
	mainWindow := myApp.NewWindow("ALttP Randomizer Keyboard Tracker")
	mainWindow.SetMaster()
	saveConfig := save.NewSaveFile()
	preferencesConfig := preferences.NewPreferencesFile()
	undoStack := undo_redo.NewUndoRedoStacks()
	inventory, err := inventory.NewInventoryIcons(undoStack, preferencesConfig, saveConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to make inventory: %w", err))
	}

	dungeon, err := dungeon.NewDungeonGrid(undoStack, preferencesConfig, saveConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to make dungeonGrid: %w", err))
	}

	mainMenu := menu.MakeMenu(myApp, mainWindow, undoStack, preferencesConfig, saveConfig, inventory, dungeon)
	mainWindow.SetMainMenu(mainMenu)

	displayMainWindowContent(mainWindow, inventory, dungeon)

	mainWindow.SetFullScreen(preferencesConfig.GetPreferenceBool("Fullscreen"))

	mainWindow.ShowAndRun()
	preferencesConfig.SetPreference("Fullscreen", mainWindow.FullScreen())
	preferencesConfig.SavePreferences()
	saveConfig.SaveState()
}
