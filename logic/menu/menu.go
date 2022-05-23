package menu

import "fmt"

type Menu struct {
	context []string
}

func NewMenu(context []string) *Menu {
	return &Menu{
		context: context,
	}
}

func (m Menu) getMenuContext() []string {
	return m.context
}

func (m Menu) GetTemplate() int {
	var menu int
	fmt.Printf("\nChoose any menu you want:\n")
	for i, line := range m.getMenuContext() {
		fmt.Printf("%d. %s\n", i+1, line)
	}

	_, err := fmt.Scanf("%d", &menu)
	if err != nil {
		fmt.Printf("Unable to read the entered number: %v", err)
		return 0
	}
	return menu
}
