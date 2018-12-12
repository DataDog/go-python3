package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreadInitialization(t *testing.T) {
	Py_Initialize()
	PyEval_InitThreads()

	assert.True(t, PyEval_ThreadsInitialized())

	PyEval_ReInitThreads()
}

func TestGIL(t *testing.T) {
	Py_Initialize()
	PyEval_InitThreads()

	gil := PyGILState_Ensure()

	assert.True(t, PyGILState_Check())

	PyGILState_Release(gil)
}

func TestThreadState(t *testing.T) {
	Py_Initialize()
	PyEval_InitThreads()

	threadState := PyGILState_GetThisThreadState()

	threadState2 := PyThreadState_Get()

	assert.Equal(t, threadState, threadState2)

	threadState3 := PyThreadState_Swap(threadState)

	assert.Equal(t, threadState, threadState3)
}

func TestThreadSaveRestore(t *testing.T) {
	Py_Initialize()
	PyEval_InitThreads()

	threadState := PyEval_SaveThread()

	assert.False(t, PyGILState_Check())

	PyEval_RestoreThread(threadState)

}
