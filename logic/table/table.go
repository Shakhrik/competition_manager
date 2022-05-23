package table

type Tables struct {
	Table      map[string]Table
	TableCount int
	TeamsCount int
}
type Table struct {
	// Title      string
	Teams      []string
	TeamsCount int
}

func NewTable() *Tables {
	return &Tables{
		Table: make(map[string]Table),
	}
}

func (t *Tables) AddTableCount(count int) {
	t.TableCount = count
}

func (t *Tables) AddTeamCount(teamsCount int) {
	t.TeamsCount = teamsCount
}

// func (t *Tables) DivideTeamsIntoTeamsNum() {
// 	reminder := t.TeamsCount % t.TableCount
// 	if reminder == 0 {
// 		// t.Table.TeamsCount = t.TeamsCount / t.TableCount
// 	}
// }

func (t *Tables) CreateTable(title string, tableCount int, teams []string) {
	// get table names from console and write to db or txt file
}

// func (t *Tables) CreateMatches(tableTitle string) error {
// 	table, ok := t.Table[tableTitle]
// 	teams := table.Teams
// 	if !ok {
// 		return errors.New("This table not exist")
// 	}

// 	file, err := filesystem.GetOrCreateFile("Matches_" + tableTitle + ".txt")
// 	if err != nil {
// 		return err
// 	}

// 	for i := 0; i < table.TeamsCount; i++ {
// 		for j := i + 1; ; j++ {
// 			s := []byte{
// 				byte(teams[0]),
// 			}

// 			file.WriteString(table.Teams[i])
// 		}
// 	}
// }
