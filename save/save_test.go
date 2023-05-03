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
	err := saveFile.SetSave("is_test", true)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = saveFile.SaveState()
	if err != nil {
		t.Fatalf("error saving save file state: %v", err)
	}

	saveFile2 := save.NewSaveFile(saveFileDirectory)
	retrievedBool := saveFile2.GetSaveBool("is_test")

	if retrievedBool != true {
		t.Error("expected retrievedBool to be true, but got false")
	}
}

func TestGetSaveInt(t *testing.T) {
	t.Parallel()

	saveFileDirectory := t.TempDir()

	saveFile := save.NewSaveFile(saveFileDirectory)
	err := saveFile.SetSave("is_testInt", 1)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = saveFile.SaveState()
	if err != nil {
		t.Fatalf("error saving save file state: %v", err)
	}

	saveFile2 := save.NewSaveFile(saveFileDirectory)
	retrievedInt := saveFile2.GetSaveInt("is_testInt")

	if retrievedInt != 1 {
		t.Errorf("expected retrievedInt to be 1, but got %d", retrievedInt)
	}
}

func TestGetSaveBool(t *testing.T) {
	t.Parallel()

	saveFileDirectory := t.TempDir()

	saveFile := save.NewSaveFile(saveFileDirectory)
	err := saveFile.SetSave("is_testBool", true)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = saveFile.SaveState()
	if err != nil {
		t.Fatalf("error saving save file state: %v", err)
	}

	saveFile2 := save.NewSaveFile(saveFileDirectory)
	retrievedBool := saveFile2.GetSaveBool("is_testBool")

	if retrievedBool != true {
		t.Error("expected retrievedBool to be true, but got false")
	}
}

func TestSetDefault(t *testing.T) {
	t.Parallel()

	saveFileDirectory := t.TempDir()

	saveFile := save.NewSaveFile(saveFileDirectory)
	err := saveFile.SetDefault("is_test", 2)
	if err != nil {
		t.Fatalf("error saving SetDefault: %v", err)
	}

	retrievedInt := saveFile.GetSaveInt("is_test")

	if retrievedInt != 2 {
		t.Errorf("expected retrievedInt to be 2, but got %d", retrievedInt)
	}
}
