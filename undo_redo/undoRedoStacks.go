package undo_redo

/*To Do:
-Convert from Stacks to Doubly Linked Lists
-Add descriptions to functions
*/

import ()

const undoHistoryLength = 50

type UndoRedoStacks struct {
	undoStack      []func()
	undoStackOld   []func()
	redoStack      []func()
	redoStackStore []func()
}

func NewUndoRedoStacks() *UndoRedoStacks {
	stacks := &UndoRedoStacks{
		undoStack:      []func(){},
		undoStackOld:   []func(){},
		redoStack:      []func(){},
		redoStackStore: []func(){},
	}

	return stacks
}

func (u *UndoRedoStacks) StoreFunctions(undoFunc func(), redoFunc func()) {
	if len(u.undoStack) >= undoHistoryLength {
		u.undoStack[0] = nil
		u.undoStack = u.undoStack[1:]
		u.redoStackStore[0] = nil
		u.redoStackStore = u.redoStackStore[1:]
	}
	u.undoStack = append(u.undoStack, undoFunc)
	u.redoStackStore = append(u.redoStackStore, redoFunc)
	u.undoStackOld = nil
	u.redoStack = nil
}

func (u *UndoRedoStacks) Undo() {
	if len(u.undoStack) > 0 {
		u.undoStack[len(u.undoStack)-1]()
		u.undoStackOld = append(u.undoStackOld, u.undoStack[len(u.undoStack)-1])
		u.undoStack[len(u.undoStack)-1] = nil
		u.undoStack = u.undoStack[:len(u.undoStack)-1]

		u.redoStack = append(u.redoStack, u.redoStackStore[len(u.redoStackStore)-1])
		u.redoStackStore[len(u.redoStackStore)-1] = nil
		u.redoStackStore = u.redoStackStore[:len(u.redoStackStore)-1]
	}
}

func (u *UndoRedoStacks) Redo() {
	if len(u.redoStack) > 0 {
		redo := u.redoStack[len(u.redoStack)-1]
		redo()
		u.redoStack[len(u.redoStack)-1] = nil
		u.redoStack = u.redoStack[:len(u.redoStack)-1]

		undo := u.undoStackOld[len(u.undoStackOld)-1]
		u.undoStackOld[len(u.undoStackOld)-1] = nil
		u.undoStackOld = u.undoStackOld[:len(u.undoStackOld)-1]

		if len(u.undoStack) >= undoHistoryLength {
			u.undoStack[0] = nil
			u.undoStack = u.undoStack[1:]
			u.redoStackStore[0] = nil
			u.redoStackStore = u.redoStackStore[1:]
		}
		u.undoStack = append(u.undoStack, undo)
		u.redoStackStore = append(u.redoStackStore, redo)
	}
}
