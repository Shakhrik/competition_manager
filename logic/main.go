package logic

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"small_projects/socker_table/logic/competition"
	"small_projects/socker_table/logic/menu"
)

func Start() {
	// storage.NewStoragePG()
	reader := bufio.NewReader(os.Stdin)
	menu := menu.NewMenu(MenuContext)
	competition := competition.NewCompetition(reader)
	for {
		num := menu.GetTemplate()
		CallClear()

		switch num {
		case 1:
			competition.Create()
		}
	}
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
