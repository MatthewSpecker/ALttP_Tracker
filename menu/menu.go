package menu

import (
	"fmt"
	"strconv"
	"net/url"
	"os"
  "math"

  "tracker/dungeon"
  "tracker/inventory"
  "tracker/save"
  "tracker/undo_redo"

	"fyne.io/fyne/v2"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
  "fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/spf13/viper"
)

func almostEqual (a, b float64) bool {
  tolerance := 0.001
  if diff := math.Abs(a - b); diff < tolerance {
    return true
  } else {
    return false
  }
}

func MakeMenu(myApp fyne.App, mainWindow fyne.Window, undoStack *undo_redo.UndoRedoStacks, preferencesConfig *viper.Viper, saveConfig *save.SaveFile, inventory *inventory.InventoryIcons, dungeon *dungeon.DungeonGrid) *fyne.MainMenu {
  categoryItem := fyne.NewMenuItem("Category", nil)
  defaultCatItem := fyne.NewMenuItem("Default", func() {
    
  })
  keysanityItem := fyne.NewMenuItem("Keysanity", nil)
  allDungeonsItem := fyne.NewMenuItem("All Dungeons", nil)
  fastGanonItem := fyne.NewMenuItem("Fast Ganon", nil)
  mapCompassBossShuffle := fyne.NewMenuItem("Map/Compass Boss Shuffle", nil)
  masterSwordItem := fyne.NewMenuItem("Master Sword Pedestal", nil)
  triforcePiecesItem := fyne.NewMenuItem("Triforce Pieces", nil)
  additionalOptionsItem := fyne.NewMenuItem("Additional Options", )
  categoryItem.ChildMenu = fyne.NewMenu("",
    defaultCatItem,
    keysanityItem,
    allDungeonsItem,
    fastGanonItem,
    mapCompassBossShuffle,
    masterSwordItem,
    triforcePiecesItem,
    additionalOptionsItem,
  )

  openPreferences := func() {
    prefWindow := myApp.NewWindow("App Preferences")
    prefWindow.Resize(mainWindow.Content().Size())

    progressive_BowsBool := preferencesConfig.GetBool("Progressive_Bows")
    progressive_BowsCheck := widget.NewCheck("Progressive Bows", func(value bool) {
      progressive_BowsBool = value
    })
    progressive_BowsCheck.Checked = preferencesConfig.GetBool("Progressive_Bows")
    progressive_BowsContainer := container.NewVBox(progressive_BowsCheck)

    bombBool := preferencesConfig.GetBool("Bombs")
    bombCheck := widget.NewCheck("Bombs", func(value bool) {
      bombBool = value
    })
    bombCheck.Checked = preferencesConfig.GetBool("Bombs")
    bombContainer := container.NewVBox(bombCheck)

    bottle_FullBool := preferencesConfig.GetBool("Bottle_Full")
    bottle_FullCheck := widget.NewCheck("Track Potions", func(value bool) {
      bottle_FullBool = value
    })
    bottle_FullCheck.Checked = preferencesConfig.GetBool("Bottle_Full")
    bottle_FullContainer := container.NewVBox(bottle_FullCheck)

    pseudo_BootsBool := preferencesConfig.GetBool("Pseudo_Boots")
    pseudo_BootsCheck := widget.NewCheck("Pseudo Boots", func(value bool) {
      pseudo_BootsBool = value
    })
    pseudo_BootsCheck.Checked = preferencesConfig.GetBool("Pseudo_Boots")
    pseudo_BootsContainer := container.NewVBox(pseudo_BootsCheck)

    swordBool := preferencesConfig.GetBool("Sword")
    swordCheck := widget.NewCheck("Sword", func(value bool) {
      swordBool = value
    })
    swordCheck.Checked = preferencesConfig.GetBool("Sword")
    swordContainer := container.NewVBox(swordCheck)

    shieldBool := preferencesConfig.GetBool("Shield")
    shieldCheck := widget.NewCheck("Shield", func(value bool) {
      shieldBool = value
    })
    shieldCheck.Checked = preferencesConfig.GetBool("Shield")
    shieldContainer := container.NewVBox(shieldCheck)

    mailBool := preferencesConfig.GetBool("Mail")
    mailCheck := widget.NewCheck("Mail", func(value bool) {
      mailBool = value
    })
    mailCheck.Checked = preferencesConfig.GetBool("Mail")
    mailContainer := container.NewVBox(mailCheck)

    halfMagicBool := preferencesConfig.GetBool("HalfMagic")
    halfMagicCheck := widget.NewCheck("Half-Magic", func(value bool) {
      halfMagicBool = value
    })
    halfMagicCheck.Checked = preferencesConfig.GetBool("HalfMagic")
    halfMagicContainer := container.NewVBox(halfMagicCheck)

    heart_PiecesBool := preferencesConfig.GetBool("Heart_Pieces")
    heart_PiecesCheck := widget.NewCheck("Heart Pieces", func(value bool) {
      heart_PiecesBool = value
    })
    heart_PiecesCheck.Checked = preferencesConfig.GetBool("Heart_Pieces")
    heart_PiecesContainer := container.NewVBox(heart_PiecesCheck)

    chest_CountBool := preferencesConfig.GetBool("Chest_Count")
    chest_CountCheck := widget.NewCheck("Track Dungeon Chests", func(value bool) {
      chest_CountBool = value
    })
    chest_CountCheck.Checked = preferencesConfig.GetBool("Chest_Count")
    chest_CountContainer := container.NewVBox(chest_CountCheck)

    mapsBool := preferencesConfig.GetBool("Maps")
    mapsCheck := widget.NewCheck("Maps", func(value bool) {
      mapsBool = value
    })
    mapsCheck.Checked = preferencesConfig.GetBool("Maps")
    mapsContainer := container.NewVBox(mapsCheck)

    compassesBool := preferencesConfig.GetBool("Compasses")
    compassesCheck := widget.NewCheck("Compasses", func(value bool) {
      compassesBool = value
    })
    compassesCheck.Checked = preferencesConfig.GetBool("Compasses")
    compassesContainer := container.NewVBox(compassesCheck)

    keysBool := preferencesConfig.GetBool("Keys")
    keysCheck := widget.NewCheck("Small Keys", func(value bool) {
      keysBool = value
    })
    keysCheck.Checked = preferencesConfig.GetBool("Keys")
    keysContainer := container.NewVBox(keysCheck)

    big_KeysBool := preferencesConfig.GetBool("Big_Keys")
    big_KeysCheck := widget.NewCheck("Big Keys", func(value bool) {
      big_KeysBool = value
    })
    big_KeysCheck.Checked = preferencesConfig.GetBool("Big_Keys")
    big_KeysContainer := container.NewVBox(big_KeysCheck)

    bossBool := preferencesConfig.GetBool("Bosses")
    bossCheck := widget.NewCheck("Bosses", func(value bool) {
      bossBool = value
    })
    bossCheck.Checked = preferencesConfig.GetBool("Bosses")
    bossContainer := container.NewVBox(bossCheck)
    mainWindow.Hide()

    applyButton := widget.NewButton("Apply Changes", func() {
      preferencesConfig.Set("Bombs", bombBool)
      preferencesConfig.Set("Progressive_Bows", progressive_BowsBool)
      preferencesConfig.Set("Bottle_Full", bottle_FullBool)
      preferencesConfig.Set("Pseudo_Boots", pseudo_BootsBool)
      preferencesConfig.Set("Sword", swordBool)
      preferencesConfig.Set("Shield", shieldBool)
      preferencesConfig.Set("Mail", mailBool)
      preferencesConfig.Set("HalfMagic", halfMagicBool)
      preferencesConfig.Set("Heart_Pieces", heart_PiecesBool)
      preferencesConfig.Set("Chest_Count", chest_CountBool)
      preferencesConfig.Set("Maps", mapsBool)
      preferencesConfig.Set("Compasses", compassesBool)
      preferencesConfig.Set("Keys", keysBool)
      preferencesConfig.Set("Big_Keys", big_KeysBool)
      preferencesConfig.Set("Bosses", bossBool)
      preferencesConfig.WriteConfig()
      inventory.PreferencesUpdate()
      dungeon.PreferencesUpdate()
      prefWindow.Close()
    })
    cancelButton := widget.NewButton("Cancel", func() {
      prefWindow.Close()
    })

    buttonContainer := container.NewHBox(applyButton, cancelButton)
    preferencesContainer := container.New(layout.NewGridLayout(1), progressive_BowsContainer, bombContainer, bottle_FullContainer,
      pseudo_BootsContainer, swordContainer, shieldContainer, mailContainer, halfMagicContainer, heart_PiecesContainer,  
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
    fyne.NewMenu("View", /*windowOnTopItem, */defaultZoomItem, zoomInItem, zoomOutItem, fyne.NewMenuItemSeparator(), fullScreenItem),
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
      if almostEqual(fyneScale, 0.1) == false && fyneScale > 0.1{
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