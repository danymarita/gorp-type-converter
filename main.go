package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/danymarita/gorp-type-converter/config"
	"github.com/danymarita/gorp-type-converter/model"
	"github.com/danymarita/gorp-type-converter/repo"
)

func insertData(repo *repo.PlanRepo) {
	data := model.Plan{
		UserID:    60,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Campaigns: model.Campaigns{
			Data: []model.CampaignItem{
				{
					CategoryID: 2,
					NetAmount:  100000,
				},
				{
					CategoryID: 4,
					NetAmount:  50000,
				},
			},
		},
	}
	err := repo.Insert(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func printData(repo *repo.PlanRepo) {
	plans, err := repo.Get()
	if err != nil {
		log.Fatalln(err.Error())
	}
	data, err := json.MarshalIndent(plans, "", " ")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(string(data))
}

func main() {
	dbMap, err := config.InitDB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer dbMap.Db.Close()

	repo := repo.NewPlanRepo(dbMap)
	// insertData(repo)
	printData(repo)
}
