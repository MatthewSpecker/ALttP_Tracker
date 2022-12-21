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
	"fyne.io/fyne/v2/layout"
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

func MakeMenu(myApp fyne.App, mainWindow fyne.Window, undoStack *undo_redo.UndoRedoStacks, preferencesConfig *preferences.PreferencesFile, saveConfig *save.SaveFile, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) *fyne.MainMenu {
	openCategory := func() {
		categoryWindow := myApp.NewWindow("Category Options")
		categoryWindow.Resize(mainWindow.Content().Size())

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

		categoryGoalText := canvas.NewText("Category Goal Type", color.White)
		categoryGoalInt := preferencesConfig.GetPreferenceInt("Goal")
		categoryGoalSelect := widget.NewSelect([]string{"Ganon Goal", "Master Sword Pedestal", "Triforce Pieces"}, func(value string) {
			if value == "Ganon Goal" {
				categoryGoalInt = 0
			} else if value == "Master Sword Pedestal" {
				categoryGoalInt = 1
			} else if value == "Triforce Pieces" {
				categoryGoalInt = 2
			}
		})
		categoryGoalContainer := container.NewVBox(categoryGoalText, categoryGoalSelect)

		progressive_BowsBool := preferencesConfig.GetPreferenceBool("Progressive_Bows")
		progressive_BowsCheck := widget.NewCheck("Progressive Bows", func(value bool) {
			progressive_BowsBool = value
		})
		progressive_BowsCheck.Checked = preferencesConfig.GetPreferenceBool("Progressive_Bows")
		progressive_BowsContainer := container.NewVBox(progressive_BowsCheck)

		pseudo_BootsBool := preferencesConfig.GetPreferenceBool("Pseudo_Boots")
		pseudo_BootsCheck := widget.NewCheck("Pseudo Boots", func(value bool) {
			pseudo_BootsBool = value
		})
		pseudo_BootsCheck.Checked = preferencesConfig.GetPreferenceBool("Pseudo_Boots")
		pseudo_BootsContainer := container.NewVBox(pseudo_BootsCheck)

		keysBool := preferencesConfig.GetPreferenceBool("Keys_Required")
		keysCheck := widget.NewCheck("Shuffled Small Keys", func(value bool) {
			keysBool = value
		})
		keysCheck.Checked = preferencesConfig.GetPreferenceBool("Keys_Required")
		keysContainer := container.NewVBox(keysCheck)

		big_KeysBool := preferencesConfig.GetPreferenceBool("Big_Keys_Required")
		big_KeysCheck := widget.NewCheck("Shuffled Big Keys", func(value bool) {
			big_KeysBool = value
		})
		big_KeysCheck.Checked = preferencesConfig.GetPreferenceBool("Big_Keys_Required")
		big_KeysContainer := container.NewVBox(big_KeysCheck)

		bossBool := preferencesConfig.GetPreferenceBool("Bosses_Required")
		bossCheck := widget.NewCheck("Shuffled Bosses", func(value bool) {
			bossBool = value
		})
		bossCheck.Checked = preferencesConfig.GetPreferenceBool("Bosses_Required")
		bossContainer := container.NewVBox(bossCheck)
		mainWindow.Hide()

		applyButton := widget.NewButton("Apply Changes", func() {
			inventory.UpdateGanonGoal(ganonGoalInt)
			preferencesConfig.SetPreference("Goal", categoryGoalInt)
			preferencesConfig.SetPreference("Progressive_Bows", progressive_BowsBool)
			preferencesConfig.SetPreference("Pseudo_Boots", pseudo_BootsBool)
			preferencesConfig.SetPreference("Key_Requireds", keysBool)
			preferencesConfig.SetPreference("Big_Keys_Required", big_KeysBool)
			preferencesConfig.SetPreference("Bosses_Required", bossBool)
			preferencesConfig.SavePreferences()
			inventory.PreferencesUpdate()
			dungeon.PreferencesUpdate()
			categoryWindow.Close()
		})
		cancelButton := widget.NewButton("Cancel", func() {
			categoryWindow.Close()
		})

		buttonContainer := container.NewHBox(applyButton, cancelButton)
		categoryContainer := container.New(layout.NewGridLayout(1), ganonGoalContainer, categoryGoalContainer,
			progressive_BowsContainer, pseudo_BootsContainer, keysContainer, big_KeysContainer, bossContainer)
		mainContainer := container.NewVBox(categoryContainer, buttonContainer)
		scrollContainer := container.NewVScroll(mainContainer)

		categoryWindow.SetContent(scrollContainer)
		categoryWindow.Show()
		categoryWindow.SetOnClosed(mainWindow.Show)
	}

	categoryItem := fyne.NewMenuItem("Category", nil)
	defaultCatItem := fyne.NewMenuItem("Default", func() {
		//change ganon tower goal number to 7
		inventory.UpdateGanonGoal(7)
		preferencesConfig.SetPreference("Goal", 0)
		preferencesConfig.SetPreference("Keys_Required", false)
		preferencesConfig.SetPreference("Big_Keys_Required", false)
		preferencesConfig.SetPreference("Bosses_Required", false)
		preferencesConfig.SavePreferences()
		inventory.PreferencesUpdate()
		dungeon.PreferencesUpdate()
	})
	keysanityItem := fyne.NewMenuItem("Keysanity", func() {
		//change ganon tower goal number to 7
		inventory.UpdateGanonGoal(7)
		preferencesConfig.SetPreference("Goal", 0)
		preferencesConfig.SetPreference("Keys_Required", true)
		preferencesConfig.SetPreference("Big_Keys_Required", true)
		preferencesConfig.SetPreference("Bosses_Required", false)
		preferencesConfig.SavePreferences()
		inventory.PreferencesUpdate()
		dungeon.PreferencesUpdate()
	})
	allDungeonsItem := fyne.NewMenuItem("All Dungeons", func() {
		//change ganon tower goal number to 7
		inventory.UpdateGanonGoal(8)
		preferencesConfig.SetPreference("Goal", 0)
		preferencesConfig.SetPreference("Keys_Required", false)
		preferencesConfig.SetPreference("Big_Keys_Required", false)
		preferencesConfig.SetPreference("Bosses_Required", false)
		preferencesConfig.SavePreferences()
		inventory.PreferencesUpdate()
		dungeon.PreferencesUpdate()
	})
	mapCompassBossShuffle := fyne.NewMenuItem("Map/Compass Boss Shuffle", func() {
		//change ganon tower goal number to 7
		inventory.UpdateGanonGoal(7)
		preferencesConfig.SetPreference("Goal", 0)
		preferencesConfig.SetPreference("Keys_Required", false)
		preferencesConfig.SetPreference("Big_Keys_Required", false)
		preferencesConfig.SetPreference("Bosses_Required", true)
		preferencesConfig.SavePreferences()
		inventory.PreferencesUpdate()
		dungeon.PreferencesUpdate()
	})
	masterSwordItem := fyne.NewMenuItem("Master Sword Pedestal", func() {
		//change ganon tower goal number to 7
		preferencesConfig.SetPreference("Goal", 1)
		preferencesConfig.SetPreference("Keys_Required", false)
		preferencesConfig.SetPreference("Big_Keys_Required", false)
		preferencesConfig.SetPreference("Bosses_Required", false)
		preferencesConfig.SavePreferences()
		inventory.PreferencesUpdate()
		dungeon.PreferencesUpdate()
	})
	triforcePiecesItem := fyne.NewMenuItem("Triforce Pieces", func() {
		//change ganon tower goal number to 7
		preferencesConfig.SetPreference("Goal", 2)
		preferencesConfig.SetPreference("Keys_Required", false)
		preferencesConfig.SetPreference("Big_Keys_Required", false)
		preferencesConfig.SetPreference("Bosses_Required", false)
		preferencesConfig.SavePreferences()
		inventory.PreferencesUpdate()
		dungeon.PreferencesUpdate()
	})
	additionalOptionsItem := fyne.NewMenuItem("Additional Options", openCategory)
	additionalOptionsItem.Icon = theme.SettingsIcon()
	categoryItem.ChildMenu = fyne.NewMenu("",
		defaultCatItem,
		keysanityItem,
		allDungeonsItem,
		mapCompassBossShuffle,
		masterSwordItem,
		triforcePiecesItem,
		additionalOptionsItem,
	)

	openPreferences := func() {
		prefWindow := myApp.NewWindow("App Preferences")
		prefWindow.Resize(mainWindow.Content().Size())

		bombBool := preferencesConfig.GetPreferenceBool("Bombs")
		bombCheck := widget.NewCheck("Bombs", func(value bool) {
			bombBool = value
		})
		bombCheck.Checked = preferencesConfig.GetPreferenceBool("Bombs")
		bombContainer := container.NewVBox(bombCheck)

		bottle_FullBool := preferencesConfig.GetPreferenceBool("Bottle_Full")
		bottle_FullCheck := widget.NewCheck("Track Potions", func(value bool) {
			bottle_FullBool = value
		})
		bottle_FullCheck.Checked = preferencesConfig.GetPreferenceBool("Bottle_Full")
		bottle_FullContainer := container.NewVBox(bottle_FullCheck)

		swordBool := preferencesConfig.GetPreferenceBool("Sword")
		swordCheck := widget.NewCheck("Sword", func(value bool) {
			swordBool = value
		})
		swordCheck.Checked = preferencesConfig.GetPreferenceBool("Sword")
		swordContainer := container.NewVBox(swordCheck)

		shieldBool := preferencesConfig.GetPreferenceBool("Shield")
		shieldCheck := widget.NewCheck("Shield", func(value bool) {
			shieldBool = value
		})
		shieldCheck.Checked = preferencesConfig.GetPreferenceBool("Shield")
		shieldContainer := container.NewVBox(shieldCheck)

		mailBool := preferencesConfig.GetPreferenceBool("Mail")
		mailCheck := widget.NewCheck("Mail", func(value bool) {
			mailBool = value
		})
		mailCheck.Checked = preferencesConfig.GetPreferenceBool("Mail")
		mailContainer := container.NewVBox(mailCheck)

		halfMagicBool := preferencesConfig.GetPreferenceBool("HalfMagic")
		halfMagicCheck := widget.NewCheck("Half-Magic", func(value bool) {
			halfMagicBool = value
		})
		halfMagicCheck.Checked = preferencesConfig.GetPreferenceBool("HalfMagic")
		halfMagicContainer := container.NewVBox(halfMagicCheck)

		heart_PiecesBool := preferencesConfig.GetPreferenceBool("Heart_Pieces")
		heart_PiecesCheck := widget.NewCheck("Heart Pieces", func(value bool) {
			heart_PiecesBool = value
		})
		heart_PiecesCheck.Checked = preferencesConfig.GetPreferenceBool("Heart_Pieces")
		heart_PiecesContainer := container.NewVBox(heart_PiecesCheck)

		chest_CountBool := preferencesConfig.GetPreferenceBool("Chest_Count")
		chest_CountCheck := widget.NewCheck("Track Dungeon Chests", func(value bool) {
			chest_CountBool = value
		})
		chest_CountCheck.Checked = preferencesConfig.GetPreferenceBool("Chest_Count")
		chest_CountContainer := container.NewVBox(chest_CountCheck)

		mapsBool := preferencesConfig.GetPreferenceBool("Maps")
		mapsCheck := widget.NewCheck("Maps", func(value bool) {
			mapsBool = value
		})
		mapsCheck.Checked = preferencesConfig.GetPreferenceBool("Maps")
		mapsContainer := container.NewVBox(mapsCheck)

		compassesBool := preferencesConfig.GetPreferenceBool("Compasses")
		compassesCheck := widget.NewCheck("Compasses", func(value bool) {
			compassesBool = value
		})
		compassesCheck.Checked = preferencesConfig.GetPreferenceBool("Compasses")
		compassesContainer := container.NewVBox(compassesCheck)

		keysBool := preferencesConfig.GetPreferenceBool("Keys")
		keysCheck := widget.NewCheck("Small Keys", func(value bool) {
			keysBool = value
		})
		keysCheck.Checked = preferencesConfig.GetPreferenceBool("Keys")
		keysContainer := container.NewVBox(keysCheck)

		big_KeysBool := preferencesConfig.GetPreferenceBool("Big_Keys")
		big_KeysCheck := widget.NewCheck("Big Keys", func(value bool) {
			big_KeysBool = value
		})
		big_KeysCheck.Checked = preferencesConfig.GetPreferenceBool("Big_Keys")
		big_KeysContainer := container.NewVBox(big_KeysCheck)

		bossBool := preferencesConfig.GetPreferenceBool("Bosses")
		bossCheck := widget.NewCheck("Bosses", func(value bool) {
			bossBool = value
		})
		bossCheck.Checked = preferencesConfig.GetPreferenceBool("Bosses")
		bossContainer := container.NewVBox(bossCheck)
		mainWindow.Hide()

		applyButton := widget.NewButton("Apply Changes", func() {
			preferencesConfig.SetPreference("Bombs", bombBool)
			preferencesConfig.SetPreference("Bottle_Full", bottle_FullBool)
			preferencesConfig.SetPreference("Sword", swordBool)
			preferencesConfig.SetPreference("Shield", shieldBool)
			preferencesConfig.SetPreference("Mail", mailBool)
			preferencesConfig.SetPreference("HalfMagic", halfMagicBool)
			preferencesConfig.SetPreference("Heart_Pieces", heart_PiecesBool)
			preferencesConfig.SetPreference("Chest_Count", chest_CountBool)
			preferencesConfig.SetPreference("Maps", mapsBool)
			preferencesConfig.SetPreference("Compasses", compassesBool)
			preferencesConfig.SetPreference("Keys", keysBool)
			preferencesConfig.SetPreference("Big_Keys", big_KeysBool)
			preferencesConfig.SetPreference("Bosses", bossBool)
			preferencesConfig.SavePreferences()
			inventory.PreferencesUpdate()
			dungeon.PreferencesUpdate()
			prefWindow.Close()
		})
		cancelButton := widget.NewButton("Cancel", func() {
			prefWindow.Close()
		})

		buttonContainer := container.NewHBox(applyButton, cancelButton)
		preferencesContainer := container.New(layout.NewGridLayout(1), bombContainer, bottle_FullContainer,
			swordContainer, shieldContainer, mailContainer, halfMagicContainer, heart_PiecesContainer,
			chest_CountContainer, mapsContainer, compassesContainer, keysContainer, big_KeysContainer, bossContainer)
		mainContainer := container.NewVBox(preferencesContainer, buttonContainer)
		scrollContainer := container.NewVScroll(mainContainer)

		prefWindow.SetContent(scrollContainer)
		prefWindow.Show()
		prefWindow.SetOnClosed(mainWindow.Show)
	}

	preferencesItem := fyne.NewMenuItem("Preferences", openPreferences)
	preferencesItem.Icon = theme.SettingsIcon()
	preferencesShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	preferencesItem.Shortcut = preferencesShortcut
	mainWindow.Canvas().AddShortcut(preferencesShortcut, func(shortcut fyne.Shortcut) {
		openPreferences()
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

	defaultZoomItem := fyne.NewMenuItem("Default Zoom", nil)
	defaultZoomItem.Icon = theme.ZoomFitIcon()
	zoomInItem := fyne.NewMenuItem("Zoom In", nil)
	zoomInItem.Icon = theme.ZoomInIcon()
	zoomOutItem := fyne.NewMenuItem("Zoom Out", nil)
	zoomOutItem.Icon = theme.ZoomOutIcon()

	fullScreenItem := fyne.NewMenuItem("Fullscreen", nil)
	fullScreenItem.Icon = theme.ViewFullScreenIcon()

	/*cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	  cutItem := fyne.NewMenuItem("Cut", func() {
	    shortcutFocused(cutShortcut, w)
	  })
	  cutItem.Shortcut = cutShortcut
	  copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	  copyItem := fyne.NewMenuItem("Copy", func() {
	    shortcutFocused(copyShortcut, w)
	  })
	  copyItem.Shortcut = copyShortcut
	  pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	  pasteItem := fyne.NewMenuItem("Paste", func() {
	    shortcutFocused(pasteShortcut, w)
	  })
	  pasteItem.Shortcut = pasteShortcut*/

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = myApp.OpenURL(u)
		}))
	//helpMenu.Icon = theme.HelpIcon()

	// a quit item will be appended to our first (File) menu
	options := fyne.NewMenu("Options", categoryItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		options.Items = append(options.Items, fyne.NewMenuItemSeparator(), preferencesItem)
	}
	main := fyne.NewMainMenu(
		options,
		fyne.NewMenu("Edit", undoItem, redoItem, fyne.NewMenuItemSeparator(), refreshItem),
		fyne.NewMenu("View" /*windowOnTopItem, */, defaultZoomItem, zoomInItem, zoomOutItem, fyne.NewMenuItemSeparator(), fullScreenItem),
		helpMenu,
	)

	fyneScaleString := os.Getenv("FYNE_SCALE")
	if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
		if almostEqual(fyneScale, 2.0) || fyneScale > 2.0 {
			zoomInItem.Disabled = true
			zoomOutItem.Disabled = false
		} else if almostEqual(fyneScale, 0.1) || fyneScale < 0.1 {
			zoomInItem.Disabled = false
			zoomOutItem.Disabled = true
		} else {
			zoomInItem.Disabled = false
			zoomOutItem.Disabled = false
		}
		if almostEqual(fyneScale, 1.0) {
			defaultZoomItem.Disabled = true
		} else {
			defaultZoomItem.Disabled = false
		}
	}

	if mainWindow.FullScreen() {
		fullScreenItem.Checked = true
	} else {
		fullScreenItem.Checked = false
	}

	defaultZoomItem.Action = func() {
		fyneScaleString := os.Getenv("FYNE_SCALE")
		if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
			if almostEqual(fyneScale, 1.0) == false {
				defaultZoomItem.Disabled = true
				fyneScale = float64(1.0)
				os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))
			}
		}
		main.Refresh()
	}

	zoomInItem.Action = func() {
		fyneScaleString := os.Getenv("FYNE_SCALE")
		if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
			fyneScale = fyneScale + float64(0.1)
			if almostEqual(fyneScale, 2.0) == false && fyneScale < 2.0 {
				zoomInItem.Disabled = false
			} else if almostEqual(fyneScale, 2.0) || fyneScale > 2.0 {
				fyneScale = float64(2.0)
				zoomInItem.Disabled = true
			}
			if almostEqual(fyneScale, 0.1) == false && fyneScale > 0.1 {
				zoomOutItem.Disabled = false
			}
			if almostEqual(fyneScale, 1.0) {
				defaultZoomItem.Disabled = true
			} else {
				defaultZoomItem.Disabled = false
			}
			os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))
		}
		main.Refresh()
	}

	zoomOutItem.Action = func() {
		fyneScaleString := os.Getenv("FYNE_SCALE")
		if fyneScale, err := strconv.ParseFloat(fyneScaleString, 32); err == nil {
			fyneScale = fyneScale - float64(0.1)
			if almostEqual(fyneScale, 0.1) == false && fyneScale > 0.1 {
				zoomOutItem.Disabled = false
			} else if almostEqual(fyneScale, 0.1) || fyneScale < 0.1 {
				fyneScale = float64(0.1)
				zoomOutItem.Disabled = true
			}
			if almostEqual(fyneScale, 2.0) == false && fyneScale < 2.0 {
				zoomInItem.Disabled = false
			}
			if almostEqual(fyneScale, 1.0) {
				defaultZoomItem.Disabled = true
			} else {
				defaultZoomItem.Disabled = false
			}
			os.Setenv("FYNE_SCALE", fmt.Sprintf("%f", fyneScale))
		}
		main.Refresh()
	}

	fullScreenItem.Action = func() {
		if mainWindow.FullScreen() {
			mainWindow.SetFullScreen(false)
			fullScreenItem.Checked = false
		} else {
			mainWindow.SetFullScreen(true)
			fullScreenItem.Checked = true
		}
		main.Refresh()
	}

	return main
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
