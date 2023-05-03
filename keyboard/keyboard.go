package keyboard

import (
	//"fmt"
	"log"

	//"tracker/dungeon"
	"tracker/inventory"
	//"tracker/menu"
	//"tracker/preferences"
	//"tracker/save"
	//"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"golang.design/x/hotkey"
)

func inputCompare(input []string, shortcut []string) bool {
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
	keyShortcut = append(keyShortcut, []string{"B", "O"})
	keyShortcut = append(keyShortcut, []string{"M", "U"})
	keyShortcut = append(keyShortcut, []string{"M", "P"})
	keyShortcut = append(keyShortcut, []string{"F", "R"})
	keyShortcut = append(keyShortcut, []string{"I", "R"})
	keyShortcut = append(keyShortcut, []string{"B", "M"})
	keyShortcut = append(keyShortcut, []string{"E", "M"})
	keyShortcut = append(keyShortcut, []string{"Q", "M"})
	keyShortcut = append(keyShortcut, []string{"L", "M"})
	keyShortcut = append(keyShortcut, []string{"H", "A"})
	keyShortcut = append(keyShortcut, []string{"F", "L"})
	keyShortcut = append(keyShortcut, []string{"S", "V"})
	keyShortcut = append(keyShortcut, []string{"B", "N"})
	keyShortcut = append(keyShortcut, []string{"B", "K"})
	keyShortcut = append(keyShortcut, []string{"B", "T"})
	keyShortcut = append(keyShortcut, []string{"B", "1"})
	keyShortcut = append(keyShortcut, []string{"B", "2"})
	keyShortcut = append(keyShortcut, []string{"B", "3"})
	keyShortcut = append(keyShortcut, []string{"B", "4"})
	keyShortcut = append(keyShortcut, []string{"C", "S"})
	keyShortcut = append(keyShortcut, []string{"C", "B"})
	keyShortcut = append(keyShortcut, []string{"M", "C"})
	keyShortcut = append(keyShortcut, []string{"M", "M"})
	keyShortcut = append(keyShortcut, []string{"P", "B"})
	keyShortcut = append(keyShortcut, []string{"G", "L"})
	keyShortcut = append(keyShortcut, []string{"F", "P"})
	keyShortcut = append(keyShortcut, []string{"M", "L"})
	keyShortcut = append(keyShortcut, []string{"S", "W"})
	keyShortcut = append(keyShortcut, []string{"S", "H"})
	keyShortcut = append(keyShortcut, []string{"M", "A"})
	keyShortcut = append(keyShortcut, []string{"H", "M"})
	keyShortcut = append(keyShortcut, []string{"H", "P"})
	keyShortcut = append(keyShortcut, []string{"G", "G"})
	keyShortcut = append(keyShortcut, []string{"P", "G"})
	keyShortcut = append(keyShortcut, []string{"T", "G"})
	keyShortcut = append(keyShortcut, []string{"G", "T"})

	return keyShortcut
}

var keyShortcut = keyShortcutConstructor()

func findResult(inputSave []string, inventory *inventory.InventoryIcons) {
	log.Print(inputSave)
	result := inputCheck(inputSave, keyShortcut)
	if result >= 0 {
		inventory.Keys(result)
		inputSave = make([]string, 0)
	}
}

var inputSave = make([]string, 0)

func KeyCheck(mainWindow fyne.Window, inventory *inventory.InventoryIcons) {
	go func() {
		// Register a desired hotkey.
		//hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyZ)
		undoHK := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyZ)
		if err := undoHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		// Start listen hotkey event whenever it is ready.
		for range undoHK.Keydown() {
			inventory.Keys(40)
		}
	}()

	go func() {
		// Register a desired hotkey.
		//hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyZ)
		redoHK := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyY)
		if err := redoHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		// Start listen hotkey event whenever it is ready.
		for range redoHK.Keydown() {
			inventory.Keys(41)
		}
	}()

	go func() {
		aHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyA)
		if err := aHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range aHK.Keydown() {
			inputSave = append(inputSave, aHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		bHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyB)
		if err := bHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range bHK.Keydown() {
			inputSave = append(inputSave, bHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		cHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyC)
		if err := cHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range cHK.Keydown() {
			inputSave = append(inputSave, cHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		dHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyD)
		if err := dHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range dHK.Keydown() {
			inputSave = append(inputSave, dHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		eHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyE)
		if err := eHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range eHK.Keydown() {
			inputSave = append(inputSave, eHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		fHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyF)
		if err := fHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range fHK.Keydown() {
			inputSave = append(inputSave, fHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		gHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyG)
		if err := gHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range gHK.Keydown() {
			inputSave = append(inputSave, gHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		hHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyH)
		if err := hHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range hHK.Keydown() {
			inputSave = append(inputSave, hHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		iHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyI)
		if err := iHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range iHK.Keydown() {
			inputSave = append(inputSave, iHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		jHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyJ)
		if err := jHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range jHK.Keydown() {
			inputSave = append(inputSave, jHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		kHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyK)
		if err := kHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range kHK.Keydown() {
			inputSave = append(inputSave, kHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		lHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyL)
		if err := lHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range lHK.Keydown() {
			inputSave = append(inputSave, lHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		mHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyM)
		if err := mHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range mHK.Keydown() {
			inputSave = append(inputSave, mHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		nHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyN)
		if err := nHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range nHK.Keydown() {
			inputSave = append(inputSave, nHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		oHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyO)
		if err := oHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range oHK.Keydown() {
			inputSave = append(inputSave, oHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		pHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyP)
		if err := pHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range pHK.Keydown() {
			inputSave = append(inputSave, pHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		qHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyQ)
		if err := qHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range qHK.Keydown() {
			inputSave = append(inputSave, qHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		rHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyR)
		if err := rHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range rHK.Keydown() {
			inputSave = append(inputSave, rHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		sHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyS)
		if err := sHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range sHK.Keydown() {
			inputSave = append(inputSave, sHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		tHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyT)
		if err := tHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range tHK.Keydown() {
			inputSave = append(inputSave, tHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		uHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyU)
		if err := uHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range uHK.Keydown() {
			inputSave = append(inputSave, uHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		vHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyV)
		if err := vHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range vHK.Keydown() {
			inputSave = append(inputSave, vHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		wHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyW)
		if err := wHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range wHK.Keydown() {
			inputSave = append(inputSave, wHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		xHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyX)
		if err := xHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range xHK.Keydown() {
			inputSave = append(inputSave, xHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		yHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyY)
		if err := yHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range yHK.Keydown() {
			inputSave = append(inputSave, yHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		zHK := hotkey.New([]hotkey.Modifier{}, hotkey.KeyZ)
		if err := zHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range zHK.Keydown() {
			inputSave = append(inputSave, zHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		zeroHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key0)
		if err := zeroHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range zeroHK.Keydown() {
			inputSave = append(inputSave, zeroHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		oneHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key1)
		if err := oneHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range oneHK.Keydown() {
			inputSave = append(inputSave, oneHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		twoHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key2)
		if err := twoHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range twoHK.Keydown() {
			inputSave = append(inputSave, twoHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		threeHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key3)
		if err := threeHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range threeHK.Keydown() {
			inputSave = append(inputSave, threeHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		fourHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key4)
		if err := fourHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range fourHK.Keydown() {
			inputSave = append(inputSave, fourHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		fiveHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key5)
		if err := fiveHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range fiveHK.Keydown() {
			inputSave = append(inputSave, fiveHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		sixHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key6)
		if err := sixHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range sixHK.Keydown() {
			inputSave = append(inputSave, sixHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		sevenHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key7)
		if err := sevenHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range sevenHK.Keydown() {
			inputSave = append(inputSave, sevenHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		eightHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key8)
		if err := eightHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range eightHK.Keydown() {
			inputSave = append(inputSave, eightHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		nineHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key9)
		if err := nineHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range nineHK.Keydown() {
			inputSave = append(inputSave, nineHK.String())
			findResult(inputSave, inventory)
		}
	}()

	go func() {
		delHK := hotkey.New([]hotkey.Modifier{}, hotkey.Key(0x2E))
		if err := delHK.Register(); err != nil {
			panic("hotkey registration failed")
		}
		for range delHK.Keydown() {
			inputSave = make([]string, 0)
		}
	}()
}
