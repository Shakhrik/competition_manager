package competition

import (
	"bufio"
	"fmt"
	"small_projects/socker_table/models"
	"small_projects/socker_table/storage"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Competition struct {
	db               *sqlx.DB
	storage          storage.StorageI
	reader           *bufio.Reader
	competitionModel models.Competition
}

func NewCompetition(reader *bufio.Reader) *Competition {
	return &Competition{
		// db:      db,
		// storage: stg,
		reader: reader,
	}
}

// this function is for terminal
func (c Competition) Create() error {
	// get competetion title
	for {
		err := c.setandGetCompotetionTitle()
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// get all teams count in the competition
	for {
		err := c.setandGetTeamsCount()
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// get all tables count in the competition
	for {
		err := c.setandGetTablesCount()
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// _, err := c.storage.Competition().Create(c.competitionModel)
	// if err != nil {
	// 	fmt.Println("Internal server error. Please try again!")
	// 	return err
	// }

	fmt.Println("Competition is created successfully!")
	return nil
}

func (c *Competition) setandGetCompotetionTitle() error {
	fmt.Printf("\nEnter Competition name: ")
	title, err := c.reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read competition name")
	}
	c.competitionModel.Title = title
	return nil
}

func (c *Competition) setandGetTeamsCount() error {
	fmt.Printf("\nEnter Competition all teams count: ")
	s, err := c.reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read competition teams count")
	}

	teamsCount, err := strconv.ParseInt(strings.TrimSpace(s), 10, 10)
	if err != nil {
		fmt.Println("error:", err)
		return fmt.Errorf("this is not correct number, Please enter correct number")
	}

	if teamsCount <= 0 {
		return fmt.Errorf("teams count can not be less than or equal to zero")
	}

	c.competitionModel.TeamsCount = int(teamsCount)
	return nil
}

func (c *Competition) setandGetTablesCount() error {
	fmt.Printf("\nEnter Competition all tables count: ")
	s, err := c.reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read competition tables count")
	}

	tablesCount, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("this is not correct number, Please enter correct number")
	}

	if tablesCount <= 0 {
		return fmt.Errorf("tables count can not be less than or equal to zero")
	}

	c.competitionModel.TablesCount = tablesCount
	return nil
}
