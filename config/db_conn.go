package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/danymarita/gorp-type-converter/model"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

const (
	host     = "localhost"
	port     = 5432
	username = "dany"
	password = "passw0rd"
	dbName   = "gorp_type_converter"
)

type CustomTypeConverter struct{}

func (c CustomTypeConverter) ToDb(val interface{}) (interface{}, error) {
	switch t := val.(type) {
	case model.Campaigns:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return val, nil
}

func (c CustomTypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *model.Campaigns:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDb: Unable to convert Plan entry to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	}
	return gorp.CustomScanner{}, false
}

func InitDB() (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbName, password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{
		Db:            db,
		Dialect:       gorp.PostgresDialect{},
		TypeConverter: CustomTypeConverter{},
	}

	dbMap.AddTableWithName(model.Plan{}, "plans").SetKeys(true, "ID")
	return dbMap, nil
}
