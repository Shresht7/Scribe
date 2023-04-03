package markdown

import "strings"

// * NodeTaskList * //

// NodeTask is a node that represents a task.
type NodeTask struct {
	// The task text.
	Text string
	// Whether the task is checked.
	Checked bool
}

// NewNodeTask creates a new task node.
func NewNodeTask(text string, checked bool) *NodeTask {
	return &NodeTask{
		Text:    text,
		Checked: checked,
	}
}

// String returns the task as a string.
func (n *NodeTask) String() string {
	if n.Checked {
		return "- [x] " + n.Text
	}
	return "- [ ] " + n.Text
}

// NodeTaskList is a node that represents a task list.
type NodeTaskList struct {
	tasks []*NodeTask
}

// NewNodeTaskList creates a new task list node.
func NewNodeTaskList() *NodeTaskList {
	return &NodeTaskList{}
}

// AddTask adds a task to the task list.
func (n *NodeTaskList) AddTask(text string, checked bool) {
	n.tasks = append(n.tasks, NewNodeTask(text, checked))
}

// String returns the task list as a string.
func (n *NodeTaskList) String() string {
	res := []string{}
	for _, task := range n.tasks {
		res = append(res, task.String())
	}
	return strings.Join(res, "\n")
}

// ToTaskList converts a map of tasks to a task list node.
func ToTaskList(tasks map[string]bool) *NodeTaskList {
	taskList := NewNodeTaskList()
	for text, checked := range tasks {
		taskList.AddTask(text, checked)
	}
	return taskList
}
