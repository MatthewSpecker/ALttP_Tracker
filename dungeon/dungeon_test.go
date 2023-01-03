package dungeon_test

import (
	"testing"

	"tracker/dungeon"
	"tracker/preferences"
	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2/test"
)

func TestNewDungeonGrid(t *testing.T) {
	t.Parallel()

	dungeon, err := dungeon.NewDungeonGrid(undo_redo.NewUndoRedoStacks(), preferences.NewPreferencesFile(t.TempDir()), save.NewSaveFile(t.TempDir()))
	if err != nil {
		t.Fatalf("Failed to make dungeonGrid: %v", err)
	}

	if dungeon == nil  {
		t.Error("got nil from NewDungeonGrid, but expected a DungeonGrid")
	}
}


func TestLayout(t *testing.T) {
	t.Parallel()

	test.NewApp()

	dungeon, err := dungeon.NewDungeonGrid(undo_redo.NewUndoRedoStacks(), preferences.NewPreferencesFile(t.TempDir()), save.NewSaveFile(t.TempDir()))
	if err != nil {
		t.Fatalf("Failed to make dungeonGrid: %v", err)
	}

	dungeonGrid := dungeon.Layout()

	if dungeonGrid == nil  {
		t.Error("got nil from Layout, but expected *fyne.Container")
	}
}