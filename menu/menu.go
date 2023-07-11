package menu

import (
	"fmt"
	"image/color"
	"math"
	"net/url"
	"os"
	"strconv"

	"tracker/dungeon"
	"tracker/inventory"
	"tracker/preferences"
	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func almostEqual(a, b float64) bool {
	tolerance := 0.001
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	} else {
		return false
	}
}

const minFyneScale = 0.5
const defaultFyneScale = 1.0
const maxFyneScale = 2.5

func fyneScaleInit(zoomInItem *fyne.MenuItem, zoomOutItem *fyne.MenuItem, defaultZoomItem *fyne.MenuItem) {
	fyneScaleString := os.Getenv("FYNE_SCALE")
	if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
		if almostEqual(fyneScale, maxFyneScale) || fyneScale > maxFyneScale {
			zoomInItem.Disabled = true
			zoomOutItem.Disabled = false
		} else if almostEqual(fyneScale, minFyneScale) || fyneScale < minFyneScale {
			zoomInItem.Disabled = false
			zoomOutItem.Disabled = true
		} else {
			zoomInItem.Disabled = false
			zoomOutItem.Disabled = false
		}
		if almostEqual(fyneScale, defaultFyneScale) {
			defaultZoomItem.Disabled = true
		} else {
			defaultZoomItem.Disabled = false
		}
	}
}

func defaultZoom(defaultZoomItem *fyne.MenuItem, main *fyne.MainMenu) {
	fyneScaleString := os.Getenv("FYNE_SCALE")
	if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
		if almostEqual(fyneScale, defaultFyneScale) == false {
			defaultZoomItem.Disabled = true
			fyneScale = float64(defaultFyneScale)
			os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))
		}
	}
	main.Refresh()
}

func zoomIn(zoomInItem *fyne.MenuItem, zoomOutItem *fyne.MenuItem, defaultZoomItem *fyne.MenuItem, main *fyne.MainMenu) {
	fyneScaleString := os.Getenv("FYNE_SCALE")
	if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
		fyneScale = fyneScale + float64(minFyneScale)
		if almostEqual(fyneScale, maxFyneScale) == false && fyneScale < maxFyneScale {
			zoomInItem.Disabled = false
		} else if almostEqual(fyneScale, maxFyneScale) || fyneScale > maxFyneScale {
			fyneScale = float64(maxFyneScale)
			zoomInItem.Disabled = true
		}
		if almostEqual(fyneScale, minFyneScale) == false && fyneScale > minFyneScale {
			zoomOutItem.Disabled = false
		}
		if almostEqual(fyneScale, defaultFyneScale) {
			defaultZoomItem.Disabled = true
		} else {
			defaultZoomItem.Disabled = false
		}
		os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))
	}
	main.Refresh()
}

func zoomOut(zoomOutItem *fyne.MenuItem, zoomInItem *fyne.MenuItem, defaultZoomItem *fyne.MenuItem, main *fyne.MainMenu) {
	fyneScaleString := os.Getenv("FYNE_SCALE")
	if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
		fyneScale = fyneScale - float64(minFyneScale)
		if almostEqual(fyneScale, minFyneScale) == false && fyneScale > minFyneScale {
			zoomOutItem.Disabled = false
		} else if almostEqual(fyneScale, minFyneScale) || fyneScale < minFyneScale {
			fyneScale = float64(minFyneScale)
			zoomOutItem.Disabled = true
		}
		if almostEqual(fyneScale, maxFyneScale) == false && fyneScale < maxFyneScale {
			zoomInItem.Disabled = false
		}
		if almostEqual(fyneScale, defaultFyneScale) {
			defaultZoomItem.Disabled = true
		} else {
			defaultZoomItem.Disabled = false
		}
		os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))
	}
	main.Refresh()
}

func fullScreen(mainWindow fyne.Window, fullScreenItem *fyne.MenuItem, main *fyne.MainMenu) {
	if mainWindow.FullScreen() {
		mainWindow.SetFullScreen(false)
		fullScreenItem.Checked = false
	} else {
		mainWindow.SetFullScreen(true)
		fullScreenItem.Checked = true
	}
	main.Refresh()
}

func defaultWindowSize(mainWindow fyne.Window, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) {
	mainGrid := container.NewHBox(inventory.Layout(), dungeon.Layout())
	mainWindow.Resize(mainGrid.Size())
}

func openMode(myApp fyne.App, mainWindow fyne.Window, preferencesConfig *preferences.PreferencesFile, saveConfig *save.SaveFile, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) {
	modeWindow := myApp.NewWindow("Mode Options")
	modeWindow.Resize(mainWindow.Content().Size())

	ganonGoalText := canvas.NewText("Ganon Crystal Requirement", color.White)
	ganonGoalInt := saveConfig.GetSaveInt("Ganon Goal_Current")
	ganonGoalSelect := widget.NewSelect([]string{"Unknown", "0", "1", "2", "3", "4", "5", "6", "7", "All Dungeons"}, func(value string) {
		if value == "Unknown" {
			ganonGoalInt = -1
		} else if value == "All Dungeons" {
			ganonGoalInt = 8
		} else {
			ganonGoalInt, _ = strconv.Atoi(value)
		}
	})
	ganonGoalContainer := container.NewVBox(ganonGoalText, ganonGoalSelect)

	modeGoalText := canvas.NewText("Mode Goal Type", color.White)
	modeGoalInt := saveConfig.GetSaveInt("Goal")
	modeGoalSelect := widget.NewSelect([]string{"Ganon Goal", "Master Sword Pedestal", "Triforce Pieces"}, func(value string) {
		if value == "Ganon Goal" {
			modeGoalInt = 0
		} else if value == "Master Sword Pedestal" {
			modeGoalInt = 1
		} else if value == "Triforce Pieces" {
			modeGoalInt = 2
		}
	})
	modeGoalContainer := container.NewVBox(modeGoalText, modeGoalSelect)

	modeContainer := container.NewVBox(ganonGoalContainer, modeGoalContainer)

	modeChecks := [][]string{{"Progressive_Bows", "Progressive Bows"}, {"Pseudo_Boots", "Pseudo Boots"}, {"Maps_Required", "Shuffled Maps"}, {"Compasses_Required", "Shuffled Compasses"}, 
		{"Keys_Required", "Shuffled Small Keys"}, {"Big_Keys_Required", "Shuffled Big Keys"}, {"Bosses_Required", "Shuffled Bosses"}}

	modeBool := []bool{}
	modeCheck := []*widget.Check{}
	for index, element := range modeChecks {
		modeBool = append(modeBool, saveConfig.GetSaveBool(element[0]))
		currIndex := index
		description := element[1]
		modeCheck = append(modeCheck, widget.NewCheck(description, func(value bool) {
			modeBool[currIndex] = value
		}))
		modeCheck[index].Checked = saveConfig.GetSaveBool(element[0])
		modeContainer.Add(modeCheck[index])
	}

	mainWindow.Hide()

	applyButton := widget.NewButton("Apply Changes", func() {
		saveConfig.SetSave("Ganon Goal_Current", ganonGoalInt)
		saveConfig.SetSave("Goal", modeGoalInt)
		for index, element := range modeChecks {
			saveConfig.SetSave(element[0], modeBool[index])
		}
		inventory.SaveUpdate()
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
		modeWindow.Close()
	})
	cancelButton := widget.NewButton("Cancel", func() {
		modeWindow.Close()
	})

	buttonContainer := container.NewHBox(applyButton, cancelButton)
	mainContainer := container.NewVBox(modeContainer, buttonContainer)
	scrollContainer := container.NewVScroll(mainContainer)

	modeWindow.SetContent(scrollContainer)
	modeWindow.Show()
	modeWindow.SetOnClosed(mainWindow.Show)
}

func openPreferences(myApp fyne.App, mainWindow fyne.Window, preferencesConfig *preferences.PreferencesFile, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) {
	prefWindow := myApp.NewWindow("App Preferences")
	prefWindow.Resize(mainWindow.Content().Size())

	preferencesContainer := container.NewVBox()

	prefChecks := [][]string{{"Bombs", "Bombs"}, {"Bottle_Full", "Track Potions"}, {"Sword", "Sword"},
		{"Shield", "Shield"}, {"Mail", "Mail"}, {"HalfMagic", "Half-Magic"}, {"Heart_Pieces", "Heart Pieces"},
		{"Chest_Count", "Track Dungeon Chests"}, {"Maps", "Maps"}, {"Compasses", "Compasses"},
		{"Keys", "Small Keys"}, {"Big_Keys", "Big Keys"}, {"Bosses", "Bosses"}}

	prefBool := []bool{}
	prefCheck := []*widget.Check{}
	for index, element := range prefChecks {
		prefBool = append(prefBool, preferencesConfig.GetPreferenceBool(element[0]))
		currIndex := index
		description := element[1]
		prefCheck = append(prefCheck, widget.NewCheck(description, func(value bool) {
			prefBool[currIndex] = value
		}))
		prefCheck[index].Checked = preferencesConfig.GetPreferenceBool(element[0])
		preferencesContainer.Add(prefCheck[index])
	}

	mainWindow.Hide()

	applyButton := widget.NewButton("Apply Changes", func() {
		for index, element := range prefChecks {
			preferencesConfig.SetPreference(element[0], prefBool[index])
		}
		preferencesConfig.SavePreferences()
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
		prefWindow.Close()
	})
	cancelButton := widget.NewButton("Cancel", func() {
		prefWindow.Close()
	})

	buttonContainer := container.NewHBox(applyButton, cancelButton)
	mainContainer := container.NewVBox(preferencesContainer, buttonContainer)
	scrollContainer := container.NewVScroll(mainContainer)

	prefWindow.SetContent(scrollContainer)
	prefWindow.Show()
	prefWindow.SetOnClosed(mainWindow.Show)
}

func MakeMenu(myApp fyne.App, mainWindow fyne.Window, undoStack *undo_redo.UndoRedoStacks, preferencesConfig *preferences.PreferencesFile, saveConfig *save.SaveFile, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) *fyne.MainMenu {
	modeItem := fyne.NewMenuItem("Mode", nil)
	defaultModeItem := fyne.NewMenuItem("Default", func() {
		//change ganon tower goal number to 7
		saveConfig.SetSave("Ganon Goal_Current", 7)
		saveConfig.SetSave("Goal", 0)
		saveConfig.SetSave("Maps_Required", false)
		saveConfig.SetSave("Compasses_Required", false)
		saveConfig.SetSave("Keys_Required", false)
		saveConfig.SetSave("Big_Keys_Required", false)
		saveConfig.SetSave("Bosses_Required", false)
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
	})
	keysanityItem := fyne.NewMenuItem("Keysanity", func() {
		//change ganon tower goal number to 7
		saveConfig.SetSave("Ganon Goal_Current", 7)
		saveConfig.SetSave("Goal", 0)
		saveConfig.SetSave("Maps_Required", true)
		saveConfig.SetSave("Compasses_Required", true)
		saveConfig.SetSave("Keys_Required", true)
		saveConfig.SetSave("Big_Keys_Required", true)
		saveConfig.SetSave("Bosses_Required", false)
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
	})
	allDungeonsItem := fyne.NewMenuItem("AD Keys", func() {
		//change ganon tower goal number to AD
		saveConfig.SetSave("Ganon Goal_Current", 8)
		saveConfig.SetSave("Goal", 0)
		saveConfig.SetSave("Maps_Required", true)
		saveConfig.SetSave("Compasses_Required", true)
		saveConfig.SetSave("Keys_Required", true)
		saveConfig.SetSave("Big_Keys_Required", true)
		saveConfig.SetSave("Bosses_Required", false)
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
	})
	mapCompassBossShuffle := fyne.NewMenuItem("Map/Compass Boss Shuffle", func() {
		//change ganon tower goal number to 7
		saveConfig.SetSave("Ganon Goal_Current", 7)
		saveConfig.SetSave("Goal", 0)
		saveConfig.SetSave("Maps_Required", true)
		saveConfig.SetSave("Compasses_Required", true)
		saveConfig.SetSave("Keys_Required", false)
		saveConfig.SetSave("Big_Keys_Required", false)
		saveConfig.SetSave("Bosses_Required", true)
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
	})
	masterSwordItem := fyne.NewMenuItem("Master Sword Pedestal", func() {
		saveConfig.SetSave("Goal", 1)
		saveConfig.SetSave("Maps_Required", false)
		saveConfig.SetSave("Compasses_Required", false)
		saveConfig.SetSave("Keys_Required", false)
		saveConfig.SetSave("Big_Keys_Required", false)
		saveConfig.SetSave("Bosses_Required", false)
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
	})
	triforcePiecesItem := fyne.NewMenuItem("Triforce Pieces", func() {
		saveConfig.SetSave("Goal", 2)
		saveConfig.SetSave("Maps_Required", false)
		saveConfig.SetSave("Compasses_Required", false)
		saveConfig.SetSave("Keys_Required", false)
		saveConfig.SetSave("Big_Keys_Required", false)
		saveConfig.SetSave("Bosses_Required", false)
		inventory.ScreenUpdate()
		dungeon.ScreenUpdate()
	})
	additionalOptionsItem := fyne.NewMenuItem("Additional Options", nil)
	additionalOptionsItem.Action = func() {
		openMode(myApp, mainWindow, preferencesConfig, saveConfig, inventory, dungeon)
	}
	additionalOptionsItem.Icon = theme.SettingsIcon()
	modeItem.ChildMenu = fyne.NewMenu("",
		defaultModeItem,
		keysanityItem,
		allDungeonsItem,
		mapCompassBossShuffle,
		masterSwordItem,
		triforcePiecesItem,
		additionalOptionsItem,
	)

	preferencesItem := fyne.NewMenuItem("Preferences", nil)
	preferencesItem.Action = func() {
		openPreferences(myApp, mainWindow, preferencesConfig, inventory, dungeon)
	}
	preferencesItem.Icon = theme.SettingsIcon()
	preferencesShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	preferencesItem.Shortcut = preferencesShortcut
	mainWindow.Canvas().AddShortcut(preferencesShortcut, func(shortcut fyne.Shortcut) {
		openPreferences(myApp, mainWindow, preferencesConfig, inventory, dungeon)
	})

	undoItem := fyne.NewMenuItem("Undo", func() {
		undoStack.Undo()
	})
	undoItem.Icon = theme.ContentUndoIcon()
	redoItem := fyne.NewMenuItem("Redo", func() {
		undoStack.Redo()
	})
	redoItem.Icon = theme.ContentRedoIcon()
	refreshItem := fyne.NewMenuItem("Refresh", func() {
		inventory.RestoreDefaults()
		dungeon.RestoreDefaults()
	})
	refreshItem.Icon = theme.ViewRefreshIcon()

	defaultWindowSizeItem := fyne.NewMenuItem("Default Size", nil)
	//defaultWindowSizeItem.Icon = theme.ZoomFitIcon()
	defaultZoomItem := fyne.NewMenuItem("Default Zoom", nil)
	defaultZoomItem.Icon = theme.ZoomFitIcon()
	zoomInItem := fyne.NewMenuItem("Zoom In", nil)
	zoomInItem.Icon = theme.ZoomInIcon()
	zoomOutItem := fyne.NewMenuItem("Zoom Out", nil)
	zoomOutItem.Icon = theme.ZoomOutIcon()

	fullScreenItem := fyne.NewMenuItem("Fullscreen", nil)
	fullScreenItem.Icon = theme.ViewFullScreenIcon()

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = myApp.OpenURL(u)
		}))
	//helpMenu.Icon = theme.HelpIcon()

	// a quit item will be appended to our first (File) menu
	options := fyne.NewMenu("Options", modeItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		options.Items = append(options.Items, fyne.NewMenuItemSeparator(), preferencesItem)
	}
	main := fyne.NewMainMenu(
		options,
		fyne.NewMenu("Edit", undoItem, redoItem, fyne.NewMenuItemSeparator(), refreshItem),
		fyne.NewMenu("View" /*windowOnTopItem, */, defaultWindowSizeItem, fyne.NewMenuItemSeparator(), defaultZoomItem, zoomInItem, zoomOutItem, fyne.NewMenuItemSeparator(), fullScreenItem),
		helpMenu,
	)

	if mainWindow.FullScreen() {
		fullScreenItem.Checked = true
	} else {
		fullScreenItem.Checked = false
	}

	fyneScaleInit(zoomInItem, zoomOutItem, defaultZoomItem)

	defaultWindowSizeItem.Action = func() {
		defaultWindowSize(mainWindow, inventory, dungeon)
	}

	defaultZoomItem.Action = func() {
		defaultZoom(defaultZoomItem, main)
	}

	zoomInItem.Action = func() {
		zoomIn(zoomInItem, zoomOutItem, defaultZoomItem, main)
	}

	zoomOutItem.Action = func() {
		zoomOut(zoomOutItem, zoomInItem, defaultZoomItem, main)
	}

	fullScreenItem.Action = func() {
		fullScreen(mainWindow, fullScreenItem, main)
	}

	return main
}
