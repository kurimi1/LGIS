// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"LGIS/model"
)

func newMctz(db *gorm.DB) mctz {
	_mctz := mctz{}

	_mctz.mctzDo.UseDB(db)
	_mctz.mctzDo.UseModel(&model.Mctz{})

	tableName := _mctz.mctzDo.TableName()
	_mctz.ALL = field.NewField(tableName, "*")
	_mctz.KCAAA = field.NewString(tableName, "KCAAA")
	_mctz.MDABA = field.NewString(tableName, "MDABA")
	_mctz.MDAFA = field.NewInt32(tableName, "MDAFA")
	_mctz.MDAFB = field.NewInt32(tableName, "MDAFB")
	_mctz.MDAFH = field.NewFloat64(tableName, "MDAFH")
	_mctz.MDBFNG = field.NewFloat64(tableName, "MDBFNG")
	_mctz.MDBEB = field.NewString(tableName, "MDBEB")
	_mctz.MDDABA = field.NewString(tableName, "MDDABA")
	_mctz.MDEG = field.NewString(tableName, "MDEG")
	_mctz.MDAGI = field.NewString(tableName, "MDAGI")
	_mctz.MDAGO = field.NewString(tableName, "MDAGO")
	_mctz.MDABC = field.NewString(tableName, "MDABC")
	_mctz.MDAGD = field.NewString(tableName, "MDAGD")
	_mctz.MDAGF = field.NewString(tableName, "MDAGF")
	_mctz.MDBLIA = field.NewString(tableName, "MDBLIA")
	_mctz.MDBK = field.NewFloat32(tableName, "MDBK")
	_mctz.MDBMB = field.NewString(tableName, "MDBMB")
	_mctz.MDBMF = field.NewString(tableName, "MDBMF")
	_mctz.数据 = field.NewString(tableName, "数据")

	_mctz.fillFieldMap()

	return _mctz
}

type mctz struct {
	mctzDo mctzDo

	ALL    field.Field
	KCAAA  field.String
	MDABA  field.String
	MDAFA  field.Int32
	MDAFB  field.Int32
	MDAFH  field.Float64
	MDBFNG field.Float64
	MDBEB  field.String
	MDDABA field.String
	MDEG   field.String
	MDAGI  field.String
	MDAGO  field.String
	MDABC  field.String
	MDAGD  field.String
	MDAGF  field.String
	MDBLIA field.String
	MDBK   field.Float32
	MDBMB  field.String
	MDBMF  field.String
	数据     field.String

	fieldMap map[string]field.Expr
}

func (m mctz) Table(newTableName string) *mctz {
	m.mctzDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mctz) As(alias string) *mctz {
	m.mctzDo.DO = *(m.mctzDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mctz) updateTableName(table string) *mctz {
	m.ALL = field.NewField(table, "*")
	m.KCAAA = field.NewString(table, "KCAAA")
	m.MDABA = field.NewString(table, "MDABA")
	m.MDAFA = field.NewInt32(table, "MDAFA")
	m.MDAFB = field.NewInt32(table, "MDAFB")
	m.MDAFH = field.NewFloat64(table, "MDAFH")
	m.MDBFNG = field.NewFloat64(table, "MDBFNG")
	m.MDBEB = field.NewString(table, "MDBEB")
	m.MDDABA = field.NewString(table, "MDDABA")
	m.MDEG = field.NewString(table, "MDEG")
	m.MDAGI = field.NewString(table, "MDAGI")
	m.MDAGO = field.NewString(table, "MDAGO")
	m.MDABC = field.NewString(table, "MDABC")
	m.MDAGD = field.NewString(table, "MDAGD")
	m.MDAGF = field.NewString(table, "MDAGF")
	m.MDBLIA = field.NewString(table, "MDBLIA")
	m.MDBK = field.NewFloat32(table, "MDBK")
	m.MDBMB = field.NewString(table, "MDBMB")
	m.MDBMF = field.NewString(table, "MDBMF")
	m.数据 = field.NewString(table, "数据")

	m.fillFieldMap()

	return m
}

func (m *mctz) WithContext(ctx context.Context) *mctzDo { return m.mctzDo.WithContext(ctx) }

func (m mctz) TableName() string { return m.mctzDo.TableName() }

func (m mctz) Alias() string { return m.mctzDo.Alias() }

func (m *mctz) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mctz) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 19)
	m.fieldMap["KCAAA"] = m.KCAAA
	m.fieldMap["MDABA"] = m.MDABA
	m.fieldMap["MDAFA"] = m.MDAFA
	m.fieldMap["MDAFB"] = m.MDAFB
	m.fieldMap["MDAFH"] = m.MDAFH
	m.fieldMap["MDBFNG"] = m.MDBFNG
	m.fieldMap["MDBEB"] = m.MDBEB
	m.fieldMap["MDDABA"] = m.MDDABA
	m.fieldMap["MDEG"] = m.MDEG
	m.fieldMap["MDAGI"] = m.MDAGI
	m.fieldMap["MDAGO"] = m.MDAGO
	m.fieldMap["MDABC"] = m.MDABC
	m.fieldMap["MDAGD"] = m.MDAGD
	m.fieldMap["MDAGF"] = m.MDAGF
	m.fieldMap["MDBLIA"] = m.MDBLIA
	m.fieldMap["MDBK"] = m.MDBK
	m.fieldMap["MDBMB"] = m.MDBMB
	m.fieldMap["MDBMF"] = m.MDBMF
	m.fieldMap["数据"] = m.数据
}

func (m mctz) clone(db *gorm.DB) mctz {
	m.mctzDo.ReplaceDB(db)
	return m
}

type mctzDo struct{ gen.DO }

func (m mctzDo) Debug() *mctzDo {
	return m.withDO(m.DO.Debug())
}

func (m mctzDo) WithContext(ctx context.Context) *mctzDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mctzDo) ReadDB(ctx context.Context) *mctzDo {
	return m.WithContext(ctx).Clauses(dbresolver.Read)
}

func (m mctzDo) WriteDB(ctx context.Context) *mctzDo {
	return m.WithContext(ctx).Clauses(dbresolver.Write)
}

func (m mctzDo) Clauses(conds ...clause.Expression) *mctzDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mctzDo) Returning(value interface{}, columns ...string) *mctzDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mctzDo) Not(conds ...gen.Condition) *mctzDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mctzDo) Or(conds ...gen.Condition) *mctzDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mctzDo) Select(conds ...field.Expr) *mctzDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mctzDo) Where(conds ...gen.Condition) *mctzDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mctzDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *mctzDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m mctzDo) Order(conds ...field.Expr) *mctzDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mctzDo) Distinct(cols ...field.Expr) *mctzDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mctzDo) Omit(cols ...field.Expr) *mctzDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mctzDo) Join(table schema.Tabler, on ...field.Expr) *mctzDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mctzDo) LeftJoin(table schema.Tabler, on ...field.Expr) *mctzDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mctzDo) RightJoin(table schema.Tabler, on ...field.Expr) *mctzDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mctzDo) Group(cols ...field.Expr) *mctzDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mctzDo) Having(conds ...gen.Condition) *mctzDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mctzDo) Limit(limit int) *mctzDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mctzDo) Offset(offset int) *mctzDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mctzDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *mctzDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mctzDo) Unscoped() *mctzDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mctzDo) Create(values ...*model.Mctz) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mctzDo) CreateInBatches(values []*model.Mctz, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mctzDo) Save(values ...*model.Mctz) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mctzDo) First() (*model.Mctz, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mctz), nil
	}
}

func (m mctzDo) Take() (*model.Mctz, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mctz), nil
	}
}

func (m mctzDo) Last() (*model.Mctz, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mctz), nil
	}
}

func (m mctzDo) Find() ([]*model.Mctz, error) {
	result, err := m.DO.Find()
	return result.([]*model.Mctz), err
}

func (m mctzDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Mctz, err error) {
	buf := make([]*model.Mctz, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mctzDo) FindInBatches(result *[]*model.Mctz, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mctzDo) Attrs(attrs ...field.AssignExpr) *mctzDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mctzDo) Assign(attrs ...field.AssignExpr) *mctzDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mctzDo) Joins(fields ...field.RelationField) *mctzDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mctzDo) Preload(fields ...field.RelationField) *mctzDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mctzDo) FirstOrInit() (*model.Mctz, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mctz), nil
	}
}

func (m mctzDo) FirstOrCreate() (*model.Mctz, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Mctz), nil
	}
}

func (m mctzDo) FindByPage(offset int, limit int) (result []*model.Mctz, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m mctzDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m *mctzDo) withDO(do gen.Dao) *mctzDo {
	m.DO = *do.(*gen.DO)
	return m
}
