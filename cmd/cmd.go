package cmd

import "github.com/happymanju/gtd/gtd"

type AppState struct {
	menuStack    []Menu
	ModifiedFlag bool
	tm           *gtd.TaskManager
}

// add task
// delete task
// edit task
// complete task
// view by status
// view by project
// return to main menu
// go back one menu

func (a *AppState) parseOptions(option string) {
	switch option {
	case "n":

	}

}

type Menu struct {
	DisplayText string
	Options     []string
}

func Run(args []string) int {
	return 0
}
