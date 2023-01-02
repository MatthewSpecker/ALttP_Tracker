package save_test

import (
	"testing"

	"tracker/save"
)

func TestNewSaveFile(t *testing.T) {
	t.Parallel()

	if save.NewSaveFile(t.TempDir()) == nil {
		t.Error("got nil from NewSaveFile, but expected a SaveFile")
	}
}

func TestSaveState(t *testing.T) {
	t.Parallel()

	saveFile := save.NewSaveFile(t.TempDir())

	saveFile.SaveState()
}

func TestSetSave(t *testing.T) {
	t.Parallel()

	saveFileDirectory := t.TempDir()

	saveFile := save.NewSaveFile(saveFileDirectory)
	saveFile.SetSave("is_dark", true)
	err := saveFile.SaveState()
	if err != nil {
		t.Fatalf("error saving save file state: %v", err)
	}

	saveFile2 := save.NewSaveFile(saveFileDirectory)
	retrievedBool := saveFile2.GetSaveBool("is_dark")

	if retrievedBool != true {
		t.Errorf("expected retrievedBool to be true, but got false")
	}
}
