// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"test_demo/db/model"
)

func newOrderTable(db *gorm.DB, opts ...gen.DOOption) orderTable {
	_orderTable := orderTable{}

	_orderTable.orderTableDo.UseDB(db, opts...)
	_orderTable.orderTableDo.UseModel(&model.OrderTable{})

	tableName := _orderTable.orderTableDo.TableName()
	_orderTable.ALL = field.NewAsterisk(tableName)
	_orderTable.OrderID = field.NewInt64(tableName, "order_id")
	_orderTable.UserID = field.NewInt64(tableName, "user_id")
	_orderTable.ProductID = field.NewInt64(tableName, "product_id")
	_orderTable.OrderTime = field.NewInt64(tableName, "order_time")
	_orderTable.OrderStatus = field.NewInt64(tableName, "order_status")
	_orderTable.OrderAmount = field.NewInt64(tableName, "order_amount")

	_orderTable.fillFieldMap()

	return _orderTable
}

type orderTable struct {
	orderTableDo orderTableDo

	ALL         field.Asterisk
	OrderID     field.Int64 // 订单ID
	UserID      field.Int64 // 用户ID
	ProductID   field.Int64 // 商品ID
	OrderTime   field.Int64 // 下单时间
	OrderStatus field.Int64 // 订单状态
	OrderAmount field.Int64 // 下单数量

	fieldMap map[string]field.Expr
}

func (o orderTable) Table(newTableName string) *orderTable {
	o.orderTableDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orderTable) As(alias string) *orderTable {
	o.orderTableDo.DO = *(o.orderTableDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orderTable) updateTableName(table string) *orderTable {
	o.ALL = field.NewAsterisk(table)
	o.OrderID = field.NewInt64(table, "order_id")
	o.UserID = field.NewInt64(table, "user_id")
	o.ProductID = field.NewInt64(table, "product_id")
	o.OrderTime = field.NewInt64(table, "order_time")
	o.OrderStatus = field.NewInt64(table, "order_status")
	o.OrderAmount = field.NewInt64(table, "order_amount")

	o.fillFieldMap()

	return o
}

func (o *orderTable) WithContext(ctx context.Context) IOrderTableDo {
	return o.orderTableDo.WithContext(ctx)
}

func (o orderTable) TableName() string { return o.orderTableDo.TableName() }

func (o orderTable) Alias() string { return o.orderTableDo.Alias() }

func (o orderTable) Columns(cols ...field.Expr) gen.Columns { return o.orderTableDo.Columns(cols...) }

func (o *orderTable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orderTable) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["user_id"] = o.UserID
	o.fieldMap["product_id"] = o.ProductID
	o.fieldMap["order_time"] = o.OrderTime
	o.fieldMap["order_status"] = o.OrderStatus
	o.fieldMap["order_amount"] = o.OrderAmount
}

func (o orderTable) clone(db *gorm.DB) orderTable {
	o.orderTableDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orderTable) replaceDB(db *gorm.DB) orderTable {
	o.orderTableDo.ReplaceDB(db)
	return o
}

type orderTableDo struct{ gen.DO }

type IOrderTableDo interface {
	gen.SubQuery
	Debug() IOrderTableDo
	WithContext(ctx context.Context) IOrderTableDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrderTableDo
	WriteDB() IOrderTableDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrderTableDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrderTableDo
	Not(conds ...gen.Condition) IOrderTableDo
	Or(conds ...gen.Condition) IOrderTableDo
	Select(conds ...field.Expr) IOrderTableDo
	Where(conds ...gen.Condition) IOrderTableDo
	Order(conds ...field.Expr) IOrderTableDo
	Distinct(cols ...field.Expr) IOrderTableDo
	Omit(cols ...field.Expr) IOrderTableDo
	Join(table schema.Tabler, on ...field.Expr) IOrderTableDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrderTableDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrderTableDo
	Group(cols ...field.Expr) IOrderTableDo
	Having(conds ...gen.Condition) IOrderTableDo
	Limit(limit int) IOrderTableDo
	Offset(offset int) IOrderTableDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderTableDo
	Unscoped() IOrderTableDo
	Create(values ...*model.OrderTable) error
	CreateInBatches(values []*model.OrderTable, batchSize int) error
	Save(values ...*model.OrderTable) error
	First() (*model.OrderTable, error)
	Take() (*model.OrderTable, error)
	Last() (*model.OrderTable, error)
	Find() ([]*model.OrderTable, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderTable, err error)
	FindInBatches(result *[]*model.OrderTable, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OrderTable) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrderTableDo
	Assign(attrs ...field.AssignExpr) IOrderTableDo
	Joins(fields ...field.RelationField) IOrderTableDo
	Preload(fields ...field.RelationField) IOrderTableDo
	FirstOrInit() (*model.OrderTable, error)
	FirstOrCreate() (*model.OrderTable, error)
	FindByPage(offset int, limit int) (result []*model.OrderTable, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrderTableDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orderTableDo) Debug() IOrderTableDo {
	return o.withDO(o.DO.Debug())
}

func (o orderTableDo) WithContext(ctx context.Context) IOrderTableDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderTableDo) ReadDB() IOrderTableDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderTableDo) WriteDB() IOrderTableDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderTableDo) Session(config *gorm.Session) IOrderTableDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderTableDo) Clauses(conds ...clause.Expression) IOrderTableDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderTableDo) Returning(value interface{}, columns ...string) IOrderTableDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderTableDo) Not(conds ...gen.Condition) IOrderTableDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderTableDo) Or(conds ...gen.Condition) IOrderTableDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderTableDo) Select(conds ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderTableDo) Where(conds ...gen.Condition) IOrderTableDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderTableDo) Order(conds ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderTableDo) Distinct(cols ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderTableDo) Omit(cols ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderTableDo) Join(table schema.Tabler, on ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderTableDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderTableDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderTableDo) Group(cols ...field.Expr) IOrderTableDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderTableDo) Having(conds ...gen.Condition) IOrderTableDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderTableDo) Limit(limit int) IOrderTableDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderTableDo) Offset(offset int) IOrderTableDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderTableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderTableDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderTableDo) Unscoped() IOrderTableDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderTableDo) Create(values ...*model.OrderTable) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderTableDo) CreateInBatches(values []*model.OrderTable, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderTableDo) Save(values ...*model.OrderTable) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderTableDo) First() (*model.OrderTable, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderTable), nil
	}
}

func (o orderTableDo) Take() (*model.OrderTable, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderTable), nil
	}
}

func (o orderTableDo) Last() (*model.OrderTable, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderTable), nil
	}
}

func (o orderTableDo) Find() ([]*model.OrderTable, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrderTable), err
}

func (o orderTableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderTable, err error) {
	buf := make([]*model.OrderTable, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderTableDo) FindInBatches(result *[]*model.OrderTable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderTableDo) Attrs(attrs ...field.AssignExpr) IOrderTableDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderTableDo) Assign(attrs ...field.AssignExpr) IOrderTableDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderTableDo) Joins(fields ...field.RelationField) IOrderTableDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderTableDo) Preload(fields ...field.RelationField) IOrderTableDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderTableDo) FirstOrInit() (*model.OrderTable, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderTable), nil
	}
}

func (o orderTableDo) FirstOrCreate() (*model.OrderTable, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderTable), nil
	}
}

func (o orderTableDo) FindByPage(offset int, limit int) (result []*model.OrderTable, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orderTableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderTableDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderTableDo) Delete(models ...*model.OrderTable) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderTableDo) withDO(do gen.Dao) *orderTableDo {
	o.DO = *do.(*gen.DO)
	return o
}