package dungeon

import (
	"errors"
	"fmt"
	"image/color"

	"tracker/preferences"
	"tracker/save"
	"tracker/tappable_icons"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type dungeonIcons struct {
	nameText               *canvas.Text
	prizeTapIcon           *tappable_icons.TappablePrizeIcon
	completionTapIcon      *tappable_icons.TappableIcon
	chestTapIcon           *tappable_icons.TappableNumIconWithIcon
	mapTapIcon             *tappable_icons.TappableIcon
	compassTapIcon         *tappable_icons.TappableIcon
	keyTapIcon             *tappable_icons.TappableNumIconWithIcon
	bigKeyTapIcon          *tappable_icons.TappableIcon
	bossTapIcon            *tappable_icons.TappableBossIcon
	preferencesConfig      *preferences.PreferencesFile
	saveConfig             *save.SaveFile
	prizeBool              bool
	bossBool               bool
	bossInt                int
	mapBool                bool
	compassBool            bool
	bigKeyBool             bool
	keys                   int
	totalChecks            int
	nameTextContainer      *fyne.Container
	prizeTapContainer      *fyne.Container
	completionTapContainer *fyne.Container
	chestTapContainer      *fyne.Container
	mapTapContainer        *fyne.Container
	compassTapContainer    *fyne.Container
	keyTapContainer        *fyne.Container
	bigKeyTapContainer     *fyne.Container
	bossTapContainer       *fyne.Container
	DungeonRow             []*fyne.Container
}

func newDungeonIcons(undoStack *undo_redo.UndoRedoStacks, preferencesConfig *preferences.PreferencesFile, saveConfig *save.SaveFile, scaleConstant float32, name string, prizeBool bool, bossBool bool, bossInt int, mapBool bool, compassBool bool, bigKeyBool bool, keys int, totalChecks int) (*dungeonIcons, error) {
	if scaleConstant <= 0 {
		return nil, errors.New("'scaleConstant' must be greater than 0")
	}
	if name == "" {
		return nil, errors.New("'name' cannot be empty string")
	}
	if bossInt < -1 || bossInt > 9 {
		return nil, errors.New("'bossInt' must be int from -1 to 9")
	}
	if keys < 0 {
		return nil, errors.New("'keys' must be a non-negative integer")
	}
	if totalChecks < 0 {
		return nil, errors.New("'totalChecks' must be a non-negative integer")
	}

	var err error
	dungeon := &dungeonIcons{
		preferencesConfig: preferencesConfig,
		saveConfig:        saveConfig,
		prizeBool:         prizeBool,
		bossBool:          bossBool,
		bossInt:           bossInt,
		mapBool:           mapBool,
		compassBool:       compassBool,
		bigKeyBool:        bigKeyBool,
		keys:              keys,
		totalChecks:       totalChecks,
	}

	dungeon.nameText = canvas.NewText(name, color.White)
	dungeon.nameText.Alignment = fyne.TextAlignTrailing

	if prizeBool && bossBool {
		dungeon.prizeTapIcon, err = tappable_icons.NewTappablePrizeIcon(18*scaleConstant, undoStack, saveConfig, name)
		if err != nil {
			return nil, fmt.Errorf("Encountered error making prizeTapIcon: %w", err)
		}
	} else if prizeBool && bossBool == false {
		if name == "CT" {
			dungeon.completionTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceAgahnim1GrayPng, resourceAgahnim1Png}, 18*scaleConstant, undoStack, saveConfig, name+"_Agahnim1")
			if err != nil {
				return nil, fmt.Errorf("Encountered error making dungonCompletionTapIcon: %w", err)
			}
		} else {
			dungeon.completionTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceAgahnim2GrayPng, resourceAgahnim2Png}, 18*scaleConstant, undoStack, saveConfig, name+"_Agahnim2")
			if err != nil {
				return nil, fmt.Errorf("Encountered error making dungonCompletionTapIcon: %w", err)
			}
		}
	}

	chestCount := dungeon.calculateChestCount()
	dungeon.chestTapIcon, err = tappable_icons.NewTappableNumIconWithIcon([]fyne.Resource{resourceChestPng, resourceEmptyChestPng}, chestCount, false, 16*scaleConstant, undoStack, saveConfig, name+"_Chest")
	if err != nil {
		return nil, fmt.Errorf("Encountered error making chestTapIcon: %w", err)
	}

	if mapBool {
		dungeon.mapTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceMapGrayPng, resourceMapPng}, 16*scaleConstant, undoStack, saveConfig, name+"_Map")
		if err != nil {
			return nil, fmt.Errorf("Encountered error making mapTapIcon: %w", err)
		}
	}

	if compassBool {
		dungeon.compassTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceCompassGrayPng, resourceCompassPng}, 15*scaleConstant, undoStack, saveConfig, name+"_Compass")
		if err != nil {
			return nil, fmt.Errorf("Encountered error making compassTapIcon: %w", err)
		}
	}

	dungeon.keyTapIcon, err = tappable_icons.NewTappableNumIconWithIcon([]fyne.Resource{resourceKeyPng, resourceKeyPng}, keys, true, 16*scaleConstant, undoStack, saveConfig, name+"_Keys")
	if err != nil {
		return nil, fmt.Errorf("Encountered error making kepsTapIcon: %w", err)
	}

	if bigKeyBool {
		dungeon.bigKeyTapIcon, err = tappable_icons.NewTappableIcon([]fyne.Resource{resourceBigKeyGrayPng, resourceBigKeyPng}, 16*scaleConstant, undoStack, saveConfig, name+"_BigKey")
		if err != nil {
			return nil, fmt.Errorf("Encountered error making bigKeyTapIcon: %w", err)
		}
	}

	if bossBool {
		dungeon.bossTapIcon, err = tappable_icons.NewTappableBossIcon(bossInt, 20*scaleConstant, dungeon.prizeTapIcon, undoStack, saveConfig, name)
		if err != nil {
			return nil, fmt.Errorf("Encountered error making bossTapIcon: %w", err)
		}
	}

	return dungeon, nil
}

func (d *dungeonIcons) calculateChestCount() int {
	var mapInt, compassInt, keyInt, bigKeyInt int
	if d.preferencesConfig.GetPreferenceBool("Maps") || d.mapBool == false {
		mapInt = 0
	} else {
		mapInt = 1
	}
	if d.preferencesConfig.GetPreferenceBool("Compasses") || d.compassBool == false {
		compassInt = 0
	} else {
		compassInt = 1
	}
	if d.preferencesConfig.GetPreferenceBool("Keys") {
		keyInt = 0
	} else {
		keyInt = 1
	}
	if d.preferencesConfig.GetPreferenceBool("Big_Keys") || d.bigKeyBool == false {
		bigKeyInt = 0
	} else {
		bigKeyInt = 1
	}

	chestCount := d.totalChecks - d.keys*(keyInt) - bigKeyInt - mapInt - compassInt
	return chestCount
}

func (d *dungeonIcons) layout() []*fyne.Container {
	d.DungeonRow = []*fyne.Container{}

	d.nameTextContainer = container.New(layout.NewMaxLayout(), d.nameText)

	d.DungeonRow = append(d.DungeonRow, d.nameTextContainer)

	if d.prizeBool && d.bossBool {
		d.prizeTapContainer = d.prizeTapIcon.Layout()
	} else if d.prizeBool && d.bossBool == false {
		d.prizeTapContainer = d.completionTapIcon.Layout()
	} else {
		d.prizeTapContainer = createSpacer()
	}
	d.DungeonRow = append(d.DungeonRow, d.prizeTapContainer)

	d.chestTapContainer = d.chestTapIcon.Layout()
	d.DungeonRow = append(d.DungeonRow, d.chestTapContainer)

	if d.mapBool {
		d.mapTapContainer = d.mapTapIcon.Layout()
	} else {
		d.mapTapContainer = createSpacer()
	}
	d.DungeonRow = append(d.DungeonRow, d.mapTapContainer)

	if d.compassBool {
		d.compassTapContainer = d.compassTapIcon.Layout()
	} else {
		d.compassTapContainer = createSpacer()
	}
	d.DungeonRow = append(d.DungeonRow, d.compassTapContainer)

	d.keyTapContainer = d.keyTapIcon.Layout()
	d.DungeonRow = append(d.DungeonRow, d.keyTapContainer)

	if d.bigKeyBool {
		d.bigKeyTapContainer = d.bigKeyTapIcon.Layout()
	} else {
		d.bigKeyTapContainer = createSpacer()
	}
	d.DungeonRow = append(d.DungeonRow, d.bigKeyTapContainer)

	if d.bossBool {
		d.bossTapContainer = d.bossTapIcon.Layout()
	} else {
		d.bossTapContainer = createSpacer()
	}
	d.DungeonRow = append(d.DungeonRow, d.bossTapContainer)

	d.createSaveDefaults()
	d.saveUpdate()
	d.preferencesUpdate()

	return d.DungeonRow
}

func (d *dungeonIcons) saveUpdate() {
	if d.prizeBool && d.bossBool {
		d.prizeTapIcon.Update()
	}
	if d.prizeBool && d.bossBool == false {
		d.completionTapIcon.Update()
	}
	d.chestTapIcon.Update()
	if d.mapBool {
		d.mapTapIcon.Update()
	}
	if d.compassBool {
		d.compassTapIcon.Update()
	}
	d.keyTapIcon.Update()
	if d.bigKeyBool {
		d.bigKeyTapIcon.Update()
	}
	if d.bossBool {
		d.bossTapIcon.Update()
	}
}

func (d *dungeonIcons) preferencesUpdate() {
	colCounter := 8
	if d.prizeBool == false {
		colCounter--
	}
	if d.preferencesConfig.GetPreferenceBool("Chest_Count") {
		d.chestTapContainer.Show()
	} else {
		d.chestTapContainer.Hide()
		colCounter--
	}
	if d.preferencesConfig.GetPreferenceBool("Maps") {
		d.mapTapContainer.Show()
	} else {
		d.mapTapContainer.Hide()
		colCounter--
	}
	if d.preferencesConfig.GetPreferenceBool("Compasses") {
		d.compassTapContainer.Show()
	} else {
		d.compassTapContainer.Hide()
		colCounter--
	}
	if d.preferencesConfig.GetPreferenceBool("Keys") || d.preferencesConfig.GetPreferenceBool("Keys_Required") {
		d.keyTapContainer.Show()
	} else {
		d.keyTapContainer.Hide()
		colCounter--
	}
	if d.preferencesConfig.GetPreferenceBool("Big_Keys") || d.preferencesConfig.GetPreferenceBool("Big_Keys_Required") {
		d.bigKeyTapContainer.Show()
		if d.bigKeyBool == false {
			colCounter--
		}
	} else {
		d.bigKeyTapContainer.Hide()
		colCounter--
	}
	if d.preferencesConfig.GetPreferenceBool("Bosses") || d.preferencesConfig.GetPreferenceBool("Bosses_Required") {
		d.bossTapContainer.Show()
		if d.bossBool == false {
			colCounter--
		}
	} else {
		d.bossTapContainer.Hide()
		colCounter--
	}

	if colCounter == 1 {
		d.nameTextContainer.Hide()
		d.prizeTapContainer.Hide()
		d.chestTapContainer.Hide()
		d.mapTapContainer.Hide()
		d.compassTapContainer.Hide()
		d.keyTapContainer.Hide()
		d.bigKeyTapContainer.Hide()
		d.bossTapContainer.Hide()
	} else {
		d.nameTextContainer.Show()
		d.prizeTapContainer.Show()
	}

	for _, element := range d.DungeonRow {
		element.Refresh()
	}
}

func (d *dungeonIcons) createSaveDefaults() {
	if d.prizeBool && d.bossBool {
		d.prizeTapIcon.SetSaveDefaults()
	}
	if d.prizeBool && d.bossBool == false {
		d.completionTapIcon.SetSaveDefaults()
	}
	d.chestTapIcon.SetSaveDefaults()
	if d.mapBool {
		d.mapTapIcon.SetSaveDefaults()
	}
	if d.compassBool {
		d.compassTapIcon.SetSaveDefaults()
	}
	d.keyTapIcon.SetSaveDefaults()
	if d.bigKeyBool {
		d.bigKeyTapIcon.SetSaveDefaults()
	}
	if d.bossBool {
		d.bossTapIcon.SetSaveDefaults()
	}
}

func (d *dungeonIcons) restoreDefaults() {
	if d.prizeBool && d.bossBool {
		d.prizeTapIcon.GetSaveDefaults()
	}
	if d.prizeBool && d.bossBool == false {
		d.completionTapIcon.GetSaveDefaults()
	}
	d.chestTapIcon.GetSaveDefaults()
	if d.mapBool {
		d.mapTapIcon.GetSaveDefaults()
	}
	if d.compassBool {
		d.compassTapIcon.GetSaveDefaults()
	}
	d.keyTapIcon.GetSaveDefaults()
	if d.bigKeyBool {
		d.bigKeyTapIcon.GetSaveDefaults()
	}
	if d.bossBool {
		d.bossTapIcon.GetSaveDefaults()
	}
}

func createSpacer() *fyne.Container {
	return container.New(layout.NewCenterLayout(), layout.NewSpacer())
}
