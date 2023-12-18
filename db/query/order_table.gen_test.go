// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"test_demo/db/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.OrderTable{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.OrderTable{}) fail: %s", err)
	}
}

func Test_orderTableQuery(t *testing.T) {
	orderTable := newOrderTable(db)
	orderTable = *orderTable.As(orderTable.TableName())
	_do := orderTable.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(orderTable.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <order_table> fail:", err)
		return
	}

	_, ok := orderTable.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from orderTable success")
	}

	err = _do.Create(&model.OrderTable{})
	if err != nil {
		t.Error("create item in table <order_table> fail:", err)
	}

	err = _do.Save(&model.OrderTable{})
	if err != nil {
		t.Error("create item in table <order_table> fail:", err)
	}

	err = _do.CreateInBatches([]*model.OrderTable{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <order_table> fail:", err)
	}

	_, err = _do.Select(orderTable.ALL).Take()
	if err != nil {
		t.Error("Take() on table <order_table> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <order_table> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <order_table> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <order_table> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.OrderTable{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <order_table> fail:", err)
	}

	_, err = _do.Select(orderTable.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <order_table> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <order_table> fail:", err)
	}

	_, err = _do.Select(orderTable.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <order_table> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <order_table> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <order_table> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <order_table> fail:", err)
	}

	_, err = _do.ScanByPage(&model.OrderTable{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <order_table> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <order_table> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <order_table> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <order_table> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <order_table> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <order_table> fail:", err)
	}
}
