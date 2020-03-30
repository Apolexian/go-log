package test

import (
	"../golog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	assert := assert.New(t)
	testLogger := golog.Initialise("testLogger", "test.log")
	assert.Equal(testLogger.GetName(), "testLogger")
	assert.Equal(testLogger.GetPath(), "test.log")

	testLogger = testLogger.SetName("changedName")
	assert.Equal(testLogger.GetName(), "changedName")
	testLogger = testLogger.SetPath("newTest.log")
	assert.Equal(testLogger.GetPath(), "newTest.log")
}
