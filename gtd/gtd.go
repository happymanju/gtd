package gtd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	INBOX   string = "INBOX"
	ACTION  string = "ACTION"
	LATER   string = "LATER"
	WAITING string = "WAITING"
	DONE    string = "DONE"
)

type TaskManager struct {
	Tasks    []Task
	Projects []string
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:    []Task{},
		Projects: []string{},
	}
}

func (tm *TaskManager) pruneProjects() {
	toKeep := []string{}
	for _, project := range tm.Projects {
		foundProject := false
		for _, task := range tm.Tasks {
			if task.project == project {
				foundProject = true
				break
			}
		}
		if foundProject {
			toKeep = append(toKeep, project)
		}
	}
	tm.Projects = toKeep

}

func (tm *TaskManager) Add(newTask Task) {
	for _, task := range tm.Tasks {
		if newTask.content == task.content {
			return
		}
	}
	tm.Tasks = append(tm.Tasks, newTask)
	if newTask.project != "" {
		if !(slices.Contains(tm.Projects, newTask.project)) {
			tm.Projects = append(tm.Projects, newTask.project)
			slices.Sort(tm.Projects)
		}
	}
}

func (tm *TaskManager) Remove(idx int) (Task, error) {
	if idx > len(tm.Tasks) || idx < 0 {
		return Task{}, errors.New("No task found: task index out of bounds")
	}
	oldTask := tm.Tasks[idx]
	tm.Tasks = slices.Delete(tm.Tasks, idx, idx)
	tm.pruneProjects()
	return oldTask, nil
}

func (tm *TaskManager) GetTasksByStatus(status string) map[int]Task {
	foundTasks := make(map[int]Task)
	for idx, task := range tm.Tasks {
		if task.status == status {
			foundTasks[idx] = task
		}
	}
	return foundTasks
}

// todo
func (tm *TaskManager) GetTasksByProject(project string) {}

func (tm *TaskManager) Load() error {
	f, err := os.Open("gtd.txt")
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		tm.Add(NewTaskFromString(sc.Text()))
	}

	return nil
}

func (tm *TaskManager) Save() error {
	f, err := os.Create("gtd.txt")
	if err != nil {
		return err
	}
	bw := bufio.NewWriter(f)

	for _, task := range tm.Tasks {
		_, err = bw.Write([]byte(task.ToString()))
		return err
	}
	err = bw.Flush()
	if err != nil {
		return err
	}

	return nil
}

type Task struct {
	status  string
	content string
	project string
	url     string
	note    string
}

func NewTaskFromString(taskString string) Task {
	t := Task{}
	parts := strings.Split(taskString, ";")
	t.status = parts[0]
	t.content = parts[1]
	t.project = parts[2]
	t.url = parts[3]
	t.note = parts[4]
	return t
}

func (t Task) ToString() string {
	s := fmt.Sprintf("%s;%s;%s;%s;%s\n", t.status, t.content, t.project, t.url, t.note)
	return s
}
