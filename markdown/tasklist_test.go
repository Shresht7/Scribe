package markdown

import (
	"testing"

	helpers "github.com/Shresht7/goutils/strings"
)

func TestTask(t *testing.T) {

	// Create test cases
	testCases := []struct {
		name     string
		task     string
		expected string
	}{
		{
			name:     "Empty",
			task:     NewNodeTask("", false).String(),
			expected: "- [ ] ",
		},
		{
			name:     "Checked",
			task:     NewNodeTask("Task 1", true).String(),
			expected: "- [x] Task 1",
		},
		{
			name:     "Unchecked",
			task:     NewNodeTask("Task 2", false).String(),
			expected: "- [ ] Task 2",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.task != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.task)
		}
	}

}

func TestTaskList(t *testing.T) {

	// Create test cases
	testCases := []struct {
		name     string
		taskList string
		expected string
	}{
		{
			name:     "Empty",
			taskList: NewNodeTaskList().String(),
			expected: "",
		},
		{
			name: "Single",
			taskList: func() string {
				tl := NewNodeTaskList()
				tl.AddTask("Task 1", true)
				return tl.String()
			}(),
			expected: "- [x] Task 1",
		},
		{
			name: "Multiple",
			taskList: func() string {
				tl := NewNodeTaskList()
				tl.AddTask("Task 1", true)
				tl.AddTask("Task 2", false)
				return tl.String()
			}(),
			expected: helpers.HereDoc(`
				- [x] Task 1
				- [ ] Task 2
			`),
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.taskList != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.taskList)
		}
	}

}

func TestToTaskList(t *testing.T) {

	// Task List
	tasks := map[string]bool{
		"Task 1": true,
		"Task 2": false,
	}

	// Convert to task list
	taskList := ToTaskList(tasks)

	// Expected task list
	expected := helpers.HereDoc(`
		- [x] Task 1
		- [ ] Task 2
	`)

	// Check
	if taskList.String() != expected {
		t.Errorf("Expected %s, got %s", expected, taskList.String())
	}

}
