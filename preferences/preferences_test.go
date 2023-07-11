package preferences_test

import (
	"testing"
	"image/color"

	"tracker/preferences"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
)

func TestNewPreferencesFile(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	if preferences.NewPreferencesFile(t.TempDir(), testWindow) == nil {
		t.Error("got nil from NewPreferencesFile, but expected a PreferencesFile")
	}
}

func TestSavePreferences(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFile := preferences.NewPreferencesFile(t.TempDir(), testWindow)

	preferencesFile.SavePreferences()
}

func TestSetPreference(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFileDirectory := t.TempDir()

	preferencesFile := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	err := preferencesFile.SetPreference("is_test", true)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = preferencesFile.SavePreferences()
	if err != nil {
		t.Fatalf("error saving preferences file: %v", err)
	}

	preferencesFile2 := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	retrievedBool := preferencesFile2.GetPreferenceBool("is_test")

	if retrievedBool != true {
		t.Error("expected retrievedBool to be true, but got false")
	}
}

func TestGetPreferenceInt(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFileDirectory := t.TempDir()

	preferencesFile := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	err := preferencesFile.SetPreference("is_testInt", 1)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = preferencesFile.SavePreferences()
	if err != nil {
		t.Fatalf("error saving preferences file: %v", err)
	}

	preferencesFile2 := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	retrievedInt := preferencesFile2.GetPreferenceInt("is_testInt")

	if retrievedInt != 1 {
		t.Errorf("expected retrievedInt to be 1, but got %d", retrievedInt)
	}
}

func TestGetPreferenceFloat(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFileDirectory := t.TempDir()

	preferencesFile := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	err := preferencesFile.SetPreference("is_testFloat", 1.123456)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = preferencesFile.SavePreferences()
	if err != nil {
		t.Fatalf("error saving preferences file: %v", err)
	}

	preferencesFile2 := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	retrievedFloat := preferencesFile2.GetPreferenceFloat("is_testFloat")

	if retrievedFloat != 1.123456 {
		t.Errorf("expected retrievedInt to be 1.123456, but got %f", retrievedFloat)
	}
}

func TestGetPreferenceBool(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFileDirectory := t.TempDir()

	preferencesFile := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	err := preferencesFile.SetPreference("is_testBool", true)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = preferencesFile.SavePreferences()
	if err != nil {
		t.Fatalf("error saving preferences file: %v", err)
	}

	preferencesFile2 := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	retrievedBool := preferencesFile2.GetPreferenceBool("is_testBool")

	if retrievedBool != true {
		t.Error("expected retrievedBool to be true, but got false")
	}
}

func TestCreateDefaults(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFileDirectory := t.TempDir()

	preferencesFile := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	preferencesFile.CreateDefaults(testWindow)

	retrievedInt := preferencesFile.GetPreferenceInt("")
	retrievedFloat := preferencesFile.GetPreferenceFloat("Fyne_Scale")
	retrievedBool := preferencesFile.GetPreferenceBool("Big_Keys")

	if retrievedInt != 0 {
		t.Errorf("expected retrievedInt to be 0, but got %d", retrievedInt)
	}
	if retrievedFloat != 1.000000 {
		t.Errorf("expected retrievedFloat to be 1.000000, but got %f", retrievedFloat)
	}
	if retrievedBool != false {
		t.Error("expected retrievedBool to be false, but got true")
	}
}

func TestRestoreDefaults(t *testing.T) {
	text := canvas.NewText("test", color.White)
	testWindow := test.NewWindow(text)

	preferencesFileDirectory := t.TempDir()

	preferencesFile := preferences.NewPreferencesFile(preferencesFileDirectory, testWindow)
	preferencesFile.CreateDefaults(testWindow)

	err := preferencesFile.SetPreference("Fyne_Scale", 2.000000)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}
	err = preferencesFile.SetPreference("Big_Keys", true)
	if err != nil {
		t.Fatalf("error saving value: %v", err)
	}

	preferencesFile.RestoreDefaults(testWindow)

	retrievedFloat := preferencesFile.GetPreferenceFloat("Fyne_Scale")
	retrievedBool := preferencesFile.GetPreferenceBool("Big_Keys")

	if retrievedFloat != 1.000000 {
		t.Errorf("expected retrievedFloat to be 1.000000, but got %f", retrievedFloat)
	}
	if retrievedBool != false {
		t.Error("expected retrievedBool to be false, but got true")
	}
}
