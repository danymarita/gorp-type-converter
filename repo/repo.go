package repo

import (
	"github.com/danymarita/gorp-type-converter/model"
	"gopkg.in/gorp.v2"
)

type PlanRepo struct {
	DBMap *gorp.DbMap
}

func NewPlanRepo(d *gorp.DbMap) *PlanRepo {
	return &PlanRepo{
		DBMap: d,
	}
}

func (p PlanRepo) Insert(data model.Plan) (err error) {
	return p.DBMap.Insert(&data)
}

func (p PlanRepo) Get() (data []model.Plan, err error) {
	_, err = p.DBMap.Select(&data, "SELECT * FROM plans")
	return
}
