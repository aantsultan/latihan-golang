package test

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"latihan-golang-restful-api-nkw/app"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/model"
	"testing"
)

func TestDatabase(t *testing.T) {
	db := app.NewDB()
	defer db.Close()
}

func TestQuerySelect(t *testing.T) {
	db := app.NewDB()
	defer db.Close()

	ctx := context.Background()

	SQL := "select factory_id, factory_name from s_factory_master "
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer db.Close()

	for rows.Next() {
		factoryMaster := model.FactoryMaster{}
		err := rows.Scan(&factoryMaster.Id, &factoryMaster.Name)
		helper.PanicIfError(err)
		fmt.Println(factoryMaster)
	}

}
