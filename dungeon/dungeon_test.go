package dungeon_test

import (
	"testing"
	"image/color"

	"tracker/dungeon"
	"tracker/preferences"
	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
)

func TestNewDungeonGrid(t *testing.T) {
	t.Parallel()

	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	var scale float32 = 1.0
	dungeon, err := dungeon.NewDungeonGrid(scale, undo_redo.NewUndoRedoStacks(), preferences.NewPreferencesFile(t.TempDir(), testWindow), save.NewSaveFile(t.TempDir()))
	if err != nil {
		t.Fatalf("Failed to make dungeonGrid: %v", err)
	}

	if dungeon == nil {
		t.Error("got nil from NewDungeonGrid, but expected a DungeonGrid")
	}
}

func TestLayout(t *testing.T) {
	t.Parallel()

	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	var scale float32 = 1.0

	dungeon, err := dungeon.NewDungeonGrid(scale, undo_redo.NewUndoRedoStacks(), preferences.NewPreferencesFile(t.TempDir(), testWindow), save.NewSaveFile(t.TempDir()))
	if err != nil {
		t.Fatalf("Failed to make dungeonGrid: %v", err)
	}

	dungeonGrid := dungeon.Layout()

	if dungeonGrid == nil {
		t.Error("got nil from Layout, but expected *fyne.Container")
	}
}

func TestScreenUpdate(t *testing.T) {
	t.Parallel()

	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)
	preferences := preferences.NewPreferencesFile(t.TempDir(), testWindow)
	saves := save.NewSaveFile(t.TempDir())
	var scale float32 = 1.0

	dungeon, err := dungeon.NewDungeonGrid(scale, undo_redo.NewUndoRedoStacks(), preferences, saves)
	if err != nil {
		t.Fatalf("Failed to make dungeonGrid: %v", err)
	}

	dungeon.Layout()

	preferences.SetPreference("Chest_Count", true)

	dungeon.ScreenUpdate()
	chest := dungeon.GetEasternPalaceRow()

	if chest.Visible() != true {
		t.Error("found Chest_Count to be hidden, but expected Chest_Count to be visible")
	}
}
