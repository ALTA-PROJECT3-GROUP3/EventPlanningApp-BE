package main

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/database"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	database.Migrate(db)
}
