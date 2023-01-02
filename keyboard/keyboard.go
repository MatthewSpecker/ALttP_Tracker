package keyboard

import (
//"fmt"

//"tracker/dungeon"
//"tracker/inventory"
//"tracker/menu"
//"tracker/preferences"
//"tracker/save"
//"tracker/undo_redo"

// "fyne.io/fyne/v2"
// "golang.design/x/hotkey"
)

/*func inputCompare(input []string, shortcut []string) bool {
	if len(input) > 0 && len(input) == len(shortcut) {
		for i, e := range input {
			if e != shortcut[i] {
				return false
			}
		}
		return true
	}
	return false
}

func inputCheck(input []string, keyShortcut [][]string) int {
	for i, s := range keyShortcut {
		if inputCompare(input, s) == true {
			return i
		}
	}
	return -1
}

func keyShortcutConstructor() [][]string {
	keyShortcut := make([][]string, 0)
	keyShortcut = append(keyShortcut, []string{"B", "W"})
	keyShortcut = append(keyShortcut, []string{"B", "B"})
	keyShortcut = append(keyShortcut, []string{"B", "R"})
	keyShortcut = append(keyShortcut, []string{"H", "K"})
	keyShortcut = append(keyShortcut, []string{"M", "P"})
	keyShortcut = append(keyShortcut, []string{"M", "U"})
	keyShortcut = append(keyShortcut, []string{"F", "R"})
	keyShortcut = append(keyShortcut, []string{"I", "R"})
	keyShortcut = append(keyShortcut, []string{"B", "M"})
	keyShortcut = append(keyShortcut, []string{"E", "M"})
	keyShortcut = append(keyShortcut, []string{"Q", "M"})
	keyShortcut = append(keyShortcut, []string{"S", "V"})
	keyShortcut = append(keyShortcut, []string{"L", "M"})
	keyShortcut = append(keyShortcut, []string{"H", "M"})
	keyShortcut = append(keyShortcut, []string{"F", "L"})
	keyShortcut = append(keyShortcut, []string{"B", "N"})
	keyShortcut = append(keyShortcut, []string{"B", "K"})
	keyShortcut = append(keyShortcut, []string{"P", "R"})
	keyShortcut = append(keyShortcut, []string{"B", "T"})
	keyShortcut = append(keyShortcut, []string{"C", "S"})
	keyShortcut = append(keyShortcut, []string{"C", "B"})
	keyShortcut = append(keyShortcut, []string{"M", "C"})
	keyShortcut = append(keyShortcut, []string{"M", "M"})
	keyShortcut = append(keyShortcut, []string{"G", "L"})
	keyShortcut = append(keyShortcut, []string{"P", "B"})
	keyShortcut = append(keyShortcut, []string{"F", "P"})
	keyShortcut = append(keyShortcut, []string{"S", "W"})
	keyShortcut = append(keyShortcut, []string{"S", "H"})
	keyShortcut = append(keyShortcut, []string{"M", "L"})
	keyShortcut = append(keyShortcut, []string{"Z"})
	keyShortcut = append(keyShortcut, []string{"Y"})

	return keyShortcut
}

keyShortcut := keyShortcutConstructor()
inputSave := make([]string, 0)
result := -1

mainWindow.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {

	inputSave = append(inputSave, string(k.Name))
	result = inputCheck(inputSave, keyShortcut)
	if result >= 0 || k.Name == "Delete" {
		inputSave = make([]string, 0)
	}

	if result == 0 {
		//bowTapIcon.Keyed()
	} else if result == 1 {
		//blueBoomerangTapIcon.Keyed()
	} else if result == 2 {
		//redBoomerangTapIcon.Keyed()
	} else if result == 3 {
		//hookshotTapIcon.Keyed()
	} else if result == 4 {
		//magicPowderTapIcon.Keyed()
	} else if result == 5 {
		//mushroomTapIcon.Keyed()
	} else if result == 6 {
		//fireRodTapIcon.Keyed()
	} else if result == 7 {
		//iceRodTapIcon.Keyed()
	} else if result == 8 {
		//bombosTapIcon.Keyed()
	} else if result == 9 {
		//etherTapIcon.Keyed()
	} else if result == 10 {
		//quakeTapIcon.Keyed()
	} else if result == 11 {
		//shovelTapIcon.Keyed()
	} else if result == 12 {
		//lampTapIcon.Keyed()
	} else if result == 13 {
		//hammerTapIcon.Keyed()
	} else if result == 14 {
		//fluteTapIcon.Keyed()
	} else if result == 15 {
		//bugNetTapIcon.Keyed()
	} else if result == 16 {
		//bookOfMudoraTapIcon.Keyed()
	} else if result == 17 {
		//moonPearlTapIcon.Keyed()
	} else if result == 18 {
		//bottleTotalTapIcon.Keyed()
	} else if result == 19 {
		//caneOfSomariaTapIcon.Keyed()
	} else if result == 20 {
		//caneOfByrnaTapIcon.Keyed()
	} else if result == 21 {
		//magicCapeTapIcon.Keyed()
	} else if result == 22 {
		//magicMirrorTapIcon.Keyed()
	} else if result == 23 {
		//glovesTapIcon.Keyed()
	} else if result == 24 {
		//pegasusBootsTapIcon.Keyed()
	} else if result == 25 {
		//flippersTapIcon.Keyed()
	} else if result == 26 {
		//swordTapIcon.Keyed()
	} else if result == 27 {
		//shieldTapIcon.Keyed()
	} else if result == 28 {
		//mailTapIcon.Keyed()
	} else if result == 29 {
		undoStack.Undo()
	} else if result == 30 {
		undoStack.Redo()
	}

})

go func() {
	// Register a desired hotkey.
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	if err := hk.Register(); err != nil {
		panic("hotkey registration failed")
	}
	// Start listen hotkey event whenever it is ready.
	for range hk.Keydown() {
		//bowTapIcon.Keyed()
	}
}()*/
