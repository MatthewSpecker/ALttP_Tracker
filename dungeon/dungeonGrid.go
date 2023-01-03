package dungeon

import (
	"fmt"

	"tracker/preferences"
	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type DungeonGrid struct {
	easternPalace    *dungeonIcons
	desertPalace     *dungeonIcons
	towerOfHera      *dungeonIcons
	hyruleCastle     *dungeonIcons
	castleTower      *dungeonIcons
	palaceOfDarkness *dungeonIcons
	swampPalace      *dungeonIcons
	skullWoods       *dungeonIcons
	thievesTown      *dungeonIcons
	icePalace        *dungeonIcons
	miseryMire       *dungeonIcons
	turtleRock       *dungeonIcons
	ganonsTower      *dungeonIcons
	preferencesFile  *preferences.PreferencesFile
	saveFile         *save.SaveFile
	dungeonGrid      *fyne.Container
}

func NewDungeonGrid(undoStack *undo_redo.UndoRedoStacks, preferencesConfig *preferences.PreferencesFile, saveConfig *save.SaveFile) (*DungeonGrid, error) {
	const scaleConstant = 3.0 / 5.0

	var err error
	grid := &DungeonGrid{
		preferencesFile: preferencesConfig,
		saveFile:        saveConfig,
	}
	grid.easternPalace, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "EP", true, true, 0, true, true, true, 0, 6)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making easternPalace: %w", err)
	}
	grid.desertPalace, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "DP", true, true, 1, true, true, true, 1, 6)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making desertPalace: %w", err)
	}
	grid.towerOfHera, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "TH", true, true, 2, true, true, true, 1, 6)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making towerOfHera: %w", err)
	}
	grid.hyruleCastle, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "HC", false, false, -1, true, false, false, 1, 8)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making hyruleCastle: %w", err)
	}
	grid.castleTower, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "CT", true, false, -1, false, false, false, 2, 2)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making castleTower: %w", err)
	}
	grid.palaceOfDarkness, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "PD", true, true, 3, true, true, true, 6, 14)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making palaceOfDarkness: %w", err)
	}
	grid.swampPalace, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "SP", true, true, 4, true, true, true, 1, 10)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making swampPalace: %w", err)
	}
	grid.skullWoods, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "SW", true, true, 5, true, true, true, 3, 8)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making skullWoods: %w", err)
	}
	grid.thievesTown, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "TT", true, true, 6, true, true, true, 1, 8)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making thievesTown: %w", err)
	}
	grid.icePalace, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "IP", true, true, 7, true, true, true, 2, 8)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making icePalace: %w", err)
	}
	grid.miseryMire, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "MM", true, true, 8, true, true, true, 3, 8)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making miseryMire: %w", err)
	}
	grid.turtleRock, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "TR", true, true, 9, true, true, true, 4, 12)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making turtleRock: %w", err)
	}
	grid.ganonsTower, err = newDungeonIcons(undoStack, preferencesConfig, saveConfig, scaleConstant, "GT", true, false, -1, true, true, true, 4, 27)
	if err != nil {
		return nil, fmt.Errorf("Encountered error making ganonsTower: %w", err)
	}

	return grid, nil
}

func (d *DungeonGrid) Layout() *fyne.Container {
	gridCol := 2

	if d.preferencesFile.GetPreferenceBool("Chest_Count") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Maps") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Compasses") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Keys") || d.preferencesFile.GetPreferenceBool("Keys_Required") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Big_Keys") || d.preferencesFile.GetPreferenceBool("Big_Keys_Required") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Bosses") || d.preferencesFile.GetPreferenceBool("Bosses_Required") {
		gridCol++
	}

	d.dungeonGrid = container.New(layout.NewGridLayout(gridCol))

	for _, element := range d.easternPalace.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.desertPalace.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.towerOfHera.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.hyruleCastle.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.castleTower.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.palaceOfDarkness.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.swampPalace.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.skullWoods.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.thievesTown.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.icePalace.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.miseryMire.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.turtleRock.layout() {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.ganonsTower.layout() {
		d.dungeonGrid.Add(element)
	}

	d.saveUpdate()

	return d.dungeonGrid
}

func (d *DungeonGrid) saveUpdate() {
	d.easternPalace.saveUpdate()
	d.desertPalace.saveUpdate()
	d.towerOfHera.saveUpdate()
	d.hyruleCastle.saveUpdate()
	d.castleTower.saveUpdate()
	d.palaceOfDarkness.saveUpdate()
	d.swampPalace.saveUpdate()
	d.skullWoods.saveUpdate()
	d.thievesTown.saveUpdate()
	d.icePalace.saveUpdate()
	d.miseryMire.saveUpdate()
	d.turtleRock.saveUpdate()
	d.ganonsTower.saveUpdate()
}

func (d *DungeonGrid) PreferencesUpdate() {
	d.easternPalace.preferencesUpdate()
	d.desertPalace.preferencesUpdate()
	d.towerOfHera.preferencesUpdate()
	d.hyruleCastle.preferencesUpdate()
	d.castleTower.preferencesUpdate()
	d.palaceOfDarkness.preferencesUpdate()
	d.swampPalace.preferencesUpdate()
	d.skullWoods.preferencesUpdate()
	d.thievesTown.preferencesUpdate()
	d.icePalace.preferencesUpdate()
	d.miseryMire.preferencesUpdate()
	d.turtleRock.preferencesUpdate()
	d.ganonsTower.preferencesUpdate()

	gridCol := 2

	if d.preferencesFile.GetPreferenceBool("Chest_Count") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Maps") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Compasses") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Keys") || d.preferencesFile.GetPreferenceBool("Keys_Required") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Big_Keys") || d.preferencesFile.GetPreferenceBool("Big_Keys_Required") {
		gridCol++
	}
	if d.preferencesFile.GetPreferenceBool("Bosses") || d.preferencesFile.GetPreferenceBool("Bosses_Required") {
		gridCol++
	}

	d.dungeonGrid.RemoveAll()
	d.dungeonGrid.Layout = layout.NewGridLayout(gridCol)

	for _, element := range d.easternPalace.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.desertPalace.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.towerOfHera.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.hyruleCastle.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.castleTower.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.palaceOfDarkness.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.swampPalace.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.skullWoods.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.thievesTown.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.icePalace.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.miseryMire.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.turtleRock.DungeonRow {
		d.dungeonGrid.Add(element)
	}
	for _, element := range d.ganonsTower.DungeonRow {
		d.dungeonGrid.Add(element)
	}
}

func (d *DungeonGrid) RestoreDefaults() {
	d.easternPalace.restoreDefaults()
	d.desertPalace.restoreDefaults()
	d.towerOfHera.restoreDefaults()
	d.hyruleCastle.restoreDefaults()
	d.castleTower.restoreDefaults()
	d.palaceOfDarkness.restoreDefaults()
	d.swampPalace.restoreDefaults()
	d.skullWoods.restoreDefaults()
	d.thievesTown.restoreDefaults()
	d.icePalace.restoreDefaults()
	d.miseryMire.restoreDefaults()
	d.turtleRock.restoreDefaults()
	d.ganonsTower.restoreDefaults()
}
