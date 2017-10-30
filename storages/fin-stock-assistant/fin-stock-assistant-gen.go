package fin_stock_assistant

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/sql/wrap"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

const STOCK_TABLE_NAME = "stock"

const STOCK_FIELD_ID = "id"
const STOCK_FIELD_STOCK_ID = "stock_id"
const STOCK_FIELD_EXCHANGE_ID = "exchange_id"
const STOCK_FIELD_STOCK_CODE = "stock_code"
const STOCK_FIELD_STOCK_NAME_CN = "stock_name_cn"
const STOCK_FIELD_STOCK_NAME_EN = "stock_name_en"
const STOCK_FIELD_LAUNCH_DATE = "launch_date"
const STOCK_FIELD_COMPANY_NAME_CN = "company_name_cn"
const STOCK_FIELD_COMPANY_NAME_EN = "company_name_en"
const STOCK_FIELD_WEBSITE_URL = "website_url"
const STOCK_FIELD_INDUSTRY_NAME = "industry_name"
const STOCK_FIELD_CITY_NAME_CN = "city_name_cn"
const STOCK_FIELD_CITY_NAME_EN = "city_name_en"
const STOCK_FIELD_PROVINCE_NAME_CN = "province_name_cn"
const STOCK_FIELD_PROVINCE_NAME_EN = "province_name_en"
const STOCK_FIELD_CREATE_TIME = "create_time"
const STOCK_FIELD_UPDATE_TIME = "update_time"

const STOCK_ALL_FIELDS_STRING = "id,stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time"

var STOCK_ALL_FIELDS = []string{
	"id",
	"stock_id",
	"exchange_id",
	"stock_code",
	"stock_name_cn",
	"stock_name_en",
	"launch_date",
	"company_name_cn",
	"company_name_en",
	"website_url",
	"industry_name",
	"city_name_cn",
	"city_name_en",
	"province_name_cn",
	"province_name_en",
	"create_time",
	"update_time",
}

type Stock struct {
	Id             int64  //size=20
	StockId        string //size=32
	ExchangeId     string //size=32
	StockCode      string //size=32
	StockNameCn    string //size=32
	StockNameEn    string //size=32
	LaunchDate     time.Time
	CompanyNameCn  string //size=128
	CompanyNameEn  string //size=128
	WebsiteUrl     string //size=128
	IndustryName   string //size=32
	CityNameCn     string //size=128
	CityNameEn     string //size=128
	ProvinceNameCn string //size=128
	ProvinceNameEn string //size=128
	CreateTime     time.Time
	UpdateTime     time.Time
}

type StockQuery struct {
	dao         *StockDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewStockQuery(dao *StockDao) *StockQuery {
	q := &StockQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	return buf.String()
}

func (q *StockQuery) Select(ctx context.Context) (*Stock, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *StockQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*Stock, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*Stock, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) SelectList(ctx context.Context) (list []*Stock, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *StockQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*Stock, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*Stock, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) Limit(startIncluded int64, count int64) *StockQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockQuery) Sort(fieldName string, asc bool) *StockQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}

func (q *StockQuery) where(format string, a ...interface{}) *StockQuery {
	q.whereBuffer.WriteString(fmt.Sprintf(format, a...))
	return q
}

func (q *StockQuery) Left() *StockQuery  { return q.where(" ( ") }
func (q *StockQuery) Right() *StockQuery { return q.where(" ) ") }
func (q *StockQuery) And() *StockQuery   { return q.where(" AND ") }
func (q *StockQuery) Or() *StockQuery    { return q.where(" OR ") }
func (q *StockQuery) Not() *StockQuery   { return q.where(" NOT ") }

func (q *StockQuery) Id_Equal(v int64) *StockQuery        { return q.where("id='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_NotEqual(v int64) *StockQuery     { return q.where("id<>'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_Less(v int64) *StockQuery         { return q.where("id<'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_LessEqual(v int64) *StockQuery    { return q.where("id<='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_Greater(v int64) *StockQuery      { return q.where("id>'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_GreaterEqual(v int64) *StockQuery { return q.where("id>='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) StockId_Equal(v string) *StockQuery {
	return q.where("stock_id='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_NotEqual(v string) *StockQuery {
	return q.where("stock_id<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_Less(v string) *StockQuery {
	return q.where("stock_id<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_LessEqual(v string) *StockQuery {
	return q.where("stock_id<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_Greater(v string) *StockQuery {
	return q.where("stock_id>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_GreaterEqual(v string) *StockQuery {
	return q.where("stock_id>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_Equal(v string) *StockQuery {
	return q.where("exchange_id='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_NotEqual(v string) *StockQuery {
	return q.where("exchange_id<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_Less(v string) *StockQuery {
	return q.where("exchange_id<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_LessEqual(v string) *StockQuery {
	return q.where("exchange_id<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_Greater(v string) *StockQuery {
	return q.where("exchange_id>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_GreaterEqual(v string) *StockQuery {
	return q.where("exchange_id>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_Equal(v string) *StockQuery {
	return q.where("stock_code='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_NotEqual(v string) *StockQuery {
	return q.where("stock_code<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_Less(v string) *StockQuery {
	return q.where("stock_code<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_LessEqual(v string) *StockQuery {
	return q.where("stock_code<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_Greater(v string) *StockQuery {
	return q.where("stock_code>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_GreaterEqual(v string) *StockQuery {
	return q.where("stock_code>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_Equal(v string) *StockQuery {
	return q.where("stock_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_NotEqual(v string) *StockQuery {
	return q.where("stock_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_Less(v string) *StockQuery {
	return q.where("stock_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_LessEqual(v string) *StockQuery {
	return q.where("stock_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_Greater(v string) *StockQuery {
	return q.where("stock_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_GreaterEqual(v string) *StockQuery {
	return q.where("stock_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_Equal(v string) *StockQuery {
	return q.where("stock_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_NotEqual(v string) *StockQuery {
	return q.where("stock_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_Less(v string) *StockQuery {
	return q.where("stock_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_LessEqual(v string) *StockQuery {
	return q.where("stock_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_Greater(v string) *StockQuery {
	return q.where("stock_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_GreaterEqual(v string) *StockQuery {
	return q.where("stock_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_Equal(v time.Time) *StockQuery {
	return q.where("launch_date='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_NotEqual(v time.Time) *StockQuery {
	return q.where("launch_date<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_Less(v time.Time) *StockQuery {
	return q.where("launch_date<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_LessEqual(v time.Time) *StockQuery {
	return q.where("launch_date<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_Greater(v time.Time) *StockQuery {
	return q.where("launch_date>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_GreaterEqual(v time.Time) *StockQuery {
	return q.where("launch_date>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_Equal(v string) *StockQuery {
	return q.where("company_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_NotEqual(v string) *StockQuery {
	return q.where("company_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_Less(v string) *StockQuery {
	return q.where("company_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_LessEqual(v string) *StockQuery {
	return q.where("company_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_Greater(v string) *StockQuery {
	return q.where("company_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_GreaterEqual(v string) *StockQuery {
	return q.where("company_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_Equal(v string) *StockQuery {
	return q.where("company_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_NotEqual(v string) *StockQuery {
	return q.where("company_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_Less(v string) *StockQuery {
	return q.where("company_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_LessEqual(v string) *StockQuery {
	return q.where("company_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_Greater(v string) *StockQuery {
	return q.where("company_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_GreaterEqual(v string) *StockQuery {
	return q.where("company_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_Equal(v string) *StockQuery {
	return q.where("website_url='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_NotEqual(v string) *StockQuery {
	return q.where("website_url<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_Less(v string) *StockQuery {
	return q.where("website_url<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_LessEqual(v string) *StockQuery {
	return q.where("website_url<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_Greater(v string) *StockQuery {
	return q.where("website_url>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_GreaterEqual(v string) *StockQuery {
	return q.where("website_url>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_Equal(v string) *StockQuery {
	return q.where("industry_name='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_NotEqual(v string) *StockQuery {
	return q.where("industry_name<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_Less(v string) *StockQuery {
	return q.where("industry_name<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_LessEqual(v string) *StockQuery {
	return q.where("industry_name<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_Greater(v string) *StockQuery {
	return q.where("industry_name>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_GreaterEqual(v string) *StockQuery {
	return q.where("industry_name>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_Equal(v string) *StockQuery {
	return q.where("city_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_NotEqual(v string) *StockQuery {
	return q.where("city_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_Less(v string) *StockQuery {
	return q.where("city_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_LessEqual(v string) *StockQuery {
	return q.where("city_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_Greater(v string) *StockQuery {
	return q.where("city_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_GreaterEqual(v string) *StockQuery {
	return q.where("city_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_Equal(v string) *StockQuery {
	return q.where("city_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_NotEqual(v string) *StockQuery {
	return q.where("city_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_Less(v string) *StockQuery {
	return q.where("city_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_LessEqual(v string) *StockQuery {
	return q.where("city_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_Greater(v string) *StockQuery {
	return q.where("city_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_GreaterEqual(v string) *StockQuery {
	return q.where("city_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_Equal(v string) *StockQuery {
	return q.where("province_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_NotEqual(v string) *StockQuery {
	return q.where("province_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_Less(v string) *StockQuery {
	return q.where("province_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_LessEqual(v string) *StockQuery {
	return q.where("province_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_Greater(v string) *StockQuery {
	return q.where("province_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_GreaterEqual(v string) *StockQuery {
	return q.where("province_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_Equal(v string) *StockQuery {
	return q.where("province_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_NotEqual(v string) *StockQuery {
	return q.where("province_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_Less(v string) *StockQuery {
	return q.where("province_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_LessEqual(v string) *StockQuery {
	return q.where("province_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_Greater(v string) *StockQuery {
	return q.where("province_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_GreaterEqual(v string) *StockQuery {
	return q.where("province_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_Equal(v time.Time) *StockQuery {
	return q.where("create_time='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_NotEqual(v time.Time) *StockQuery {
	return q.where("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_Less(v time.Time) *StockQuery {
	return q.where("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_LessEqual(v time.Time) *StockQuery {
	return q.where("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_Greater(v time.Time) *StockQuery {
	return q.where("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_GreaterEqual(v time.Time) *StockQuery {
	return q.where("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_Equal(v time.Time) *StockQuery {
	return q.where("update_time='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_NotEqual(v time.Time) *StockQuery {
	return q.where("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_Less(v time.Time) *StockQuery {
	return q.where("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_LessEqual(v time.Time) *StockQuery {
	return q.where("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_Greater(v time.Time) *StockQuery {
	return q.where("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_GreaterEqual(v time.Time) *StockQuery {
	return q.where("update_time>='" + fmt.Sprint(v) + "'")
}

type StockDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewStockDao(db *DB) (t *StockDao, err error) {
	t = &StockDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *StockDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}
func (dao *StockDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock (stock_id,exchange_id,stock_code,stock_name_cn,stock_name_en,launch_date,company_name_cn,company_name_en,website_url,industry_name,city_name_cn,city_name_en,province_name_cn,province_name_en,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *StockDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock SET stock_id=?,exchange_id=?,stock_code=?,stock_name_cn=?,stock_name_en=?,launch_date=?,company_name_cn=?,company_name_en=?,website_url=?,industry_name=?,city_name_cn=?,city_name_en=?,province_name_cn=?,province_name_en=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock WHERE id=?")
	return err
}

func (dao *StockDao) Insert(ctx context.Context, tx *wrap.Tx, e *Stock) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockDao) Update(ctx context.Context, tx *wrap.Tx, e *Stock) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *StockDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *StockDao) scanRow(row *wrap.Row) (*Stock, error) {
	e := &Stock{}
	err := row.Scan(&e.Id, &e.StockId, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.StockNameEn, &e.LaunchDate, &e.CompanyNameCn, &e.CompanyNameEn, &e.WebsiteUrl, &e.IndustryName, &e.CityNameCn, &e.CityNameEn, &e.ProvinceNameCn, &e.ProvinceNameEn, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockDao) scanRows(rows *wrap.Rows) (list []*Stock, err error) {
	list = make([]*Stock, 0)
	for rows.Next() {
		e := Stock{}
		err = rows.Scan(&e.Id, &e.StockId, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.StockNameEn, &e.LaunchDate, &e.CompanyNameCn, &e.CompanyNameEn, &e.WebsiteUrl, &e.IndustryName, &e.CityNameCn, &e.CityNameEn, &e.ProvinceNameCn, &e.ProvinceNameEn, &e.CreateTime, &e.UpdateTime)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *StockDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*Stock, error) {
	querySql := "SELECT " + STOCK_ALL_FIELDS_STRING + " FROM stock " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *StockDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*Stock, err error) {
	querySql := "SELECT " + STOCK_ALL_FIELDS_STRING + " FROM stock " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *StockDao) GetQuery() *StockQuery {
	return NewStockQuery(dao)
}

const USER_INDEX_EVALUATE_TABLE_NAME = "user_index_evaluate"

const USER_INDEX_EVALUATE_FIELD_ID = "id"
const USER_INDEX_EVALUATE_FIELD_USER_ID = "user_id"
const USER_INDEX_EVALUATE_FIELD_STOCK_ID = "stock_id"
const USER_INDEX_EVALUATE_FIELD_INDEX_NAME = "index_name"
const USER_INDEX_EVALUATE_FIELD_EVAL_STARS = "eval_stars"
const USER_INDEX_EVALUATE_FIELD_EVAL_REMARK = "eval_remark"
const USER_INDEX_EVALUATE_FIELD_CREATE_TIME = "create_time"
const USER_INDEX_EVALUATE_FIELD_UPDATE_TIME = "update_time"

const USER_INDEX_EVALUATE_ALL_FIELDS_STRING = "id,user_id,stock_id,index_name,eval_stars,eval_remark,create_time,update_time"

var USER_INDEX_EVALUATE_ALL_FIELDS = []string{
	"id",
	"user_id",
	"stock_id",
	"index_name",
	"eval_stars",
	"eval_remark",
	"create_time",
	"update_time",
}

type UserIndexEvaluate struct {
	Id         int64  //size=20
	UserId     string //size=32
	StockId    string //size=32
	IndexName  string //size=32
	EvalStars  int32  //size=10
	EvalRemark string //size=256
	CreateTime time.Time
	UpdateTime time.Time
}

type UserIndexEvaluateQuery struct {
	dao         *UserIndexEvaluateDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewUserIndexEvaluateQuery(dao *UserIndexEvaluateDao) *UserIndexEvaluateQuery {
	q := &UserIndexEvaluateQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *UserIndexEvaluateQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	return buf.String()
}

func (q *UserIndexEvaluateQuery) Select(ctx context.Context) (*UserIndexEvaluate, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*UserIndexEvaluate, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*UserIndexEvaluate, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) SelectList(ctx context.Context) (list []*UserIndexEvaluate, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*UserIndexEvaluate, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*UserIndexEvaluate, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) Limit(startIncluded int64, count int64) *UserIndexEvaluateQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *UserIndexEvaluateQuery) Sort(fieldName string, asc bool) *UserIndexEvaluateQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}

func (q *UserIndexEvaluateQuery) where(format string, a ...interface{}) *UserIndexEvaluateQuery {
	q.whereBuffer.WriteString(fmt.Sprintf(format, a...))
	return q
}

func (q *UserIndexEvaluateQuery) Left() *UserIndexEvaluateQuery  { return q.where(" ( ") }
func (q *UserIndexEvaluateQuery) Right() *UserIndexEvaluateQuery { return q.where(" ) ") }
func (q *UserIndexEvaluateQuery) And() *UserIndexEvaluateQuery   { return q.where(" AND ") }
func (q *UserIndexEvaluateQuery) Or() *UserIndexEvaluateQuery    { return q.where(" OR ") }
func (q *UserIndexEvaluateQuery) Not() *UserIndexEvaluateQuery   { return q.where(" NOT ") }

func (q *UserIndexEvaluateQuery) Id_Equal(v int64) *UserIndexEvaluateQuery {
	return q.where("id='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_NotEqual(v int64) *UserIndexEvaluateQuery {
	return q.where("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_Less(v int64) *UserIndexEvaluateQuery {
	return q.where("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_LessEqual(v int64) *UserIndexEvaluateQuery {
	return q.where("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_Greater(v int64) *UserIndexEvaluateQuery {
	return q.where("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_GreaterEqual(v int64) *UserIndexEvaluateQuery {
	return q.where("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_Equal(v string) *UserIndexEvaluateQuery {
	return q.where("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.where("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_Less(v string) *UserIndexEvaluateQuery {
	return q.where("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.where("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_Greater(v string) *UserIndexEvaluateQuery {
	return q.where("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.where("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_Equal(v string) *UserIndexEvaluateQuery {
	return q.where("stock_id='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.where("stock_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_Less(v string) *UserIndexEvaluateQuery {
	return q.where("stock_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.where("stock_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_Greater(v string) *UserIndexEvaluateQuery {
	return q.where("stock_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.where("stock_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_Equal(v string) *UserIndexEvaluateQuery {
	return q.where("index_name='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.where("index_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_Less(v string) *UserIndexEvaluateQuery {
	return q.where("index_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.where("index_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_Greater(v string) *UserIndexEvaluateQuery {
	return q.where("index_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.where("index_name>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_Equal(v int32) *UserIndexEvaluateQuery {
	return q.where("eval_stars='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_NotEqual(v int32) *UserIndexEvaluateQuery {
	return q.where("eval_stars<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_Less(v int32) *UserIndexEvaluateQuery {
	return q.where("eval_stars<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_LessEqual(v int32) *UserIndexEvaluateQuery {
	return q.where("eval_stars<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_Greater(v int32) *UserIndexEvaluateQuery {
	return q.where("eval_stars>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_GreaterEqual(v int32) *UserIndexEvaluateQuery {
	return q.where("eval_stars>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_Equal(v string) *UserIndexEvaluateQuery {
	return q.where("eval_remark='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.where("eval_remark<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_Less(v string) *UserIndexEvaluateQuery {
	return q.where("eval_remark<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.where("eval_remark<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_Greater(v string) *UserIndexEvaluateQuery {
	return q.where("eval_remark>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.where("eval_remark>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_Equal(v time.Time) *UserIndexEvaluateQuery {
	return q.where("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_NotEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.where("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_Less(v time.Time) *UserIndexEvaluateQuery {
	return q.where("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_LessEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.where("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_Greater(v time.Time) *UserIndexEvaluateQuery {
	return q.where("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_GreaterEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.where("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_Equal(v time.Time) *UserIndexEvaluateQuery {
	return q.where("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_NotEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.where("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_Less(v time.Time) *UserIndexEvaluateQuery {
	return q.where("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_LessEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.where("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_Greater(v time.Time) *UserIndexEvaluateQuery {
	return q.where("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_GreaterEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.where("update_time>='" + fmt.Sprint(v) + "'")
}

type UserIndexEvaluateDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewUserIndexEvaluateDao(db *DB) (t *UserIndexEvaluateDao, err error) {
	t = &UserIndexEvaluateDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *UserIndexEvaluateDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}
func (dao *UserIndexEvaluateDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_index_evaluate (user_id,stock_id,index_name,eval_stars,eval_remark,create_time,update_time) VALUES (?,?,?,?,?,?,?)")
	return err
}

func (dao *UserIndexEvaluateDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_index_evaluate SET user_id=?,stock_id=?,index_name=?,eval_stars=?,eval_remark=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *UserIndexEvaluateDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM user_index_evaluate WHERE id=?")
	return err
}

func (dao *UserIndexEvaluateDao) Insert(ctx context.Context, tx *wrap.Tx, e *UserIndexEvaluate) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.IndexName, e.EvalStars, e.EvalRemark, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *UserIndexEvaluateDao) Update(ctx context.Context, tx *wrap.Tx, e *UserIndexEvaluate) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.IndexName, e.EvalStars, e.EvalRemark, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserIndexEvaluateDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserIndexEvaluateDao) scanRow(row *wrap.Row) (*UserIndexEvaluate, error) {
	e := &UserIndexEvaluate{}
	err := row.Scan(&e.Id, &e.UserId, &e.StockId, &e.IndexName, &e.EvalStars, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *UserIndexEvaluateDao) scanRows(rows *wrap.Rows) (list []*UserIndexEvaluate, err error) {
	list = make([]*UserIndexEvaluate, 0)
	for rows.Next() {
		e := UserIndexEvaluate{}
		err = rows.Scan(&e.Id, &e.UserId, &e.StockId, &e.IndexName, &e.EvalStars, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *UserIndexEvaluateDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*UserIndexEvaluate, error) {
	querySql := "SELECT " + USER_INDEX_EVALUATE_ALL_FIELDS_STRING + " FROM user_index_evaluate " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserIndexEvaluateDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserIndexEvaluate, err error) {
	querySql := "SELECT " + USER_INDEX_EVALUATE_ALL_FIELDS_STRING + " FROM user_index_evaluate " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *UserIndexEvaluateDao) GetQuery() *UserIndexEvaluateQuery {
	return NewUserIndexEvaluateQuery(dao)
}

const USER_SETTING_TABLE_NAME = "user_setting"

const USER_SETTING_FIELD_ID = "id"
const USER_SETTING_FIELD_USER_ID = "user_id"
const USER_SETTING_FIELD_CONFIG_KEY = "config_key"
const USER_SETTING_FIELD_CONFIG_VALUE = "config_value"
const USER_SETTING_FIELD_CREATE_TIME = "create_time"
const USER_SETTING_FIELD_UPDATE_TIME = "update_time"

const USER_SETTING_ALL_FIELDS_STRING = "id,user_id,config_key,config_value,create_time,update_time"

var USER_SETTING_ALL_FIELDS = []string{
	"id",
	"user_id",
	"config_key",
	"config_value",
	"create_time",
	"update_time",
}

type UserSetting struct {
	Id          int64  //size=20
	UserId      string //size=32
	ConfigKey   string //size=32
	ConfigValue string //size=1024
	CreateTime  time.Time
	UpdateTime  time.Time
}

type UserSettingQuery struct {
	dao         *UserSettingDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewUserSettingQuery(dao *UserSettingDao) *UserSettingQuery {
	q := &UserSettingQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *UserSettingQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	return buf.String()
}

func (q *UserSettingQuery) Select(ctx context.Context) (*UserSetting, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *UserSettingQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*UserSetting, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*UserSetting, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) SelectList(ctx context.Context) (list []*UserSetting, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *UserSettingQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*UserSetting, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*UserSetting, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) Limit(startIncluded int64, count int64) *UserSettingQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *UserSettingQuery) Sort(fieldName string, asc bool) *UserSettingQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}

func (q *UserSettingQuery) where(format string, a ...interface{}) *UserSettingQuery {
	q.whereBuffer.WriteString(fmt.Sprintf(format, a...))
	return q
}

func (q *UserSettingQuery) Left() *UserSettingQuery  { return q.where(" ( ") }
func (q *UserSettingQuery) Right() *UserSettingQuery { return q.where(" ) ") }
func (q *UserSettingQuery) And() *UserSettingQuery   { return q.where(" AND ") }
func (q *UserSettingQuery) Or() *UserSettingQuery    { return q.where(" OR ") }
func (q *UserSettingQuery) Not() *UserSettingQuery   { return q.where(" NOT ") }

func (q *UserSettingQuery) Id_Equal(v int64) *UserSettingQuery {
	return q.where("id='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_NotEqual(v int64) *UserSettingQuery {
	return q.where("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_Less(v int64) *UserSettingQuery {
	return q.where("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_LessEqual(v int64) *UserSettingQuery {
	return q.where("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_Greater(v int64) *UserSettingQuery {
	return q.where("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_GreaterEqual(v int64) *UserSettingQuery {
	return q.where("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_Equal(v string) *UserSettingQuery {
	return q.where("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_NotEqual(v string) *UserSettingQuery {
	return q.where("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_Less(v string) *UserSettingQuery {
	return q.where("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_LessEqual(v string) *UserSettingQuery {
	return q.where("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_Greater(v string) *UserSettingQuery {
	return q.where("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_GreaterEqual(v string) *UserSettingQuery {
	return q.where("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_Equal(v string) *UserSettingQuery {
	return q.where("config_key='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_NotEqual(v string) *UserSettingQuery {
	return q.where("config_key<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_Less(v string) *UserSettingQuery {
	return q.where("config_key<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_LessEqual(v string) *UserSettingQuery {
	return q.where("config_key<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_Greater(v string) *UserSettingQuery {
	return q.where("config_key>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_GreaterEqual(v string) *UserSettingQuery {
	return q.where("config_key>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_Equal(v string) *UserSettingQuery {
	return q.where("config_value='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_NotEqual(v string) *UserSettingQuery {
	return q.where("config_value<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_Less(v string) *UserSettingQuery {
	return q.where("config_value<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_LessEqual(v string) *UserSettingQuery {
	return q.where("config_value<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_Greater(v string) *UserSettingQuery {
	return q.where("config_value>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_GreaterEqual(v string) *UserSettingQuery {
	return q.where("config_value>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_Equal(v time.Time) *UserSettingQuery {
	return q.where("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_NotEqual(v time.Time) *UserSettingQuery {
	return q.where("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_Less(v time.Time) *UserSettingQuery {
	return q.where("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_LessEqual(v time.Time) *UserSettingQuery {
	return q.where("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_Greater(v time.Time) *UserSettingQuery {
	return q.where("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_GreaterEqual(v time.Time) *UserSettingQuery {
	return q.where("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_Equal(v time.Time) *UserSettingQuery {
	return q.where("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_NotEqual(v time.Time) *UserSettingQuery {
	return q.where("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_Less(v time.Time) *UserSettingQuery {
	return q.where("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_LessEqual(v time.Time) *UserSettingQuery {
	return q.where("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_Greater(v time.Time) *UserSettingQuery {
	return q.where("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_GreaterEqual(v time.Time) *UserSettingQuery {
	return q.where("update_time>='" + fmt.Sprint(v) + "'")
}

type UserSettingDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewUserSettingDao(db *DB) (t *UserSettingDao, err error) {
	t = &UserSettingDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *UserSettingDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}
func (dao *UserSettingDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_setting (user_id,config_key,config_value,create_time,update_time) VALUES (?,?,?,?,?)")
	return err
}

func (dao *UserSettingDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_setting SET user_id=?,config_key=?,config_value=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *UserSettingDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM user_setting WHERE id=?")
	return err
}

func (dao *UserSettingDao) Insert(ctx context.Context, tx *wrap.Tx, e *UserSetting) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.ConfigKey, e.ConfigValue, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *UserSettingDao) Update(ctx context.Context, tx *wrap.Tx, e *UserSetting) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.ConfigKey, e.ConfigValue, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserSettingDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserSettingDao) scanRow(row *wrap.Row) (*UserSetting, error) {
	e := &UserSetting{}
	err := row.Scan(&e.Id, &e.UserId, &e.ConfigKey, &e.ConfigValue, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *UserSettingDao) scanRows(rows *wrap.Rows) (list []*UserSetting, err error) {
	list = make([]*UserSetting, 0)
	for rows.Next() {
		e := UserSetting{}
		err = rows.Scan(&e.Id, &e.UserId, &e.ConfigKey, &e.ConfigValue, &e.CreateTime, &e.UpdateTime)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *UserSettingDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*UserSetting, error) {
	querySql := "SELECT " + USER_SETTING_ALL_FIELDS_STRING + " FROM user_setting " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserSettingDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserSetting, err error) {
	querySql := "SELECT " + USER_SETTING_ALL_FIELDS_STRING + " FROM user_setting " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *UserSettingDao) GetQuery() *UserSettingQuery {
	return NewUserSettingQuery(dao)
}

const USER_STOCK_EVALUATE_TABLE_NAME = "user_stock_evaluate"

const USER_STOCK_EVALUATE_FIELD_ID = "id"
const USER_STOCK_EVALUATE_FIELD_USER_ID = "user_id"
const USER_STOCK_EVALUATE_FIELD_STOCK_ID = "stock_id"
const USER_STOCK_EVALUATE_FIELD_TOTAL_SCORE = "total_score"
const USER_STOCK_EVALUATE_FIELD_EVAL_REMARK = "eval_remark"
const USER_STOCK_EVALUATE_FIELD_CREATE_TIME = "create_time"
const USER_STOCK_EVALUATE_FIELD_UPDATE_TIME = "update_time"
const USER_STOCK_EVALUATE_FIELD_EXCHANGE_ID = "exchange_id"
const USER_STOCK_EVALUATE_FIELD_STOCK_CODE = "stock_code"
const USER_STOCK_EVALUATE_FIELD_LAUNCH_DATE = "launch_date"
const USER_STOCK_EVALUATE_FIELD_INDUSTRY_NAME = "industry_name"

const USER_STOCK_EVALUATE_ALL_FIELDS_STRING = "id,user_id,stock_id,total_score,eval_remark,create_time,update_time,exchange_id,stock_code,launch_date,industry_name"

var USER_STOCK_EVALUATE_ALL_FIELDS = []string{
	"id",
	"user_id",
	"stock_id",
	"total_score",
	"eval_remark",
	"create_time",
	"update_time",
	"exchange_id",
	"stock_code",
	"launch_date",
	"industry_name",
}

type UserStockEvaluate struct {
	Id           int64  //size=20
	UserId       string //size=32
	StockId      string //size=32
	TotalScore   float64
	EvalRemark   string //size=256
	CreateTime   time.Time
	UpdateTime   time.Time
	ExchangeId   string //size=32
	StockCode    string //size=32
	LaunchDate   time.Time
	IndustryName string //size=32
}

type UserStockEvaluateQuery struct {
	dao         *UserStockEvaluateDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewUserStockEvaluateQuery(dao *UserStockEvaluateDao) *UserStockEvaluateQuery {
	q := &UserStockEvaluateQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *UserStockEvaluateQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	return buf.String()
}

func (q *UserStockEvaluateQuery) Select(ctx context.Context) (*UserStockEvaluate, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*UserStockEvaluate, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*UserStockEvaluate, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) SelectList(ctx context.Context) (list []*UserStockEvaluate, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*UserStockEvaluate, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*UserStockEvaluate, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) Limit(startIncluded int64, count int64) *UserStockEvaluateQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *UserStockEvaluateQuery) Sort(fieldName string, asc bool) *UserStockEvaluateQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}

func (q *UserStockEvaluateQuery) where(format string, a ...interface{}) *UserStockEvaluateQuery {
	q.whereBuffer.WriteString(fmt.Sprintf(format, a...))
	return q
}

func (q *UserStockEvaluateQuery) Left() *UserStockEvaluateQuery  { return q.where(" ( ") }
func (q *UserStockEvaluateQuery) Right() *UserStockEvaluateQuery { return q.where(" ) ") }
func (q *UserStockEvaluateQuery) And() *UserStockEvaluateQuery   { return q.where(" AND ") }
func (q *UserStockEvaluateQuery) Or() *UserStockEvaluateQuery    { return q.where(" OR ") }
func (q *UserStockEvaluateQuery) Not() *UserStockEvaluateQuery   { return q.where(" NOT ") }

func (q *UserStockEvaluateQuery) Id_Equal(v int64) *UserStockEvaluateQuery {
	return q.where("id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_NotEqual(v int64) *UserStockEvaluateQuery {
	return q.where("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_Less(v int64) *UserStockEvaluateQuery {
	return q.where("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_LessEqual(v int64) *UserStockEvaluateQuery {
	return q.where("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_Greater(v int64) *UserStockEvaluateQuery {
	return q.where("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_GreaterEqual(v int64) *UserStockEvaluateQuery {
	return q.where("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_Equal(v string) *UserStockEvaluateQuery {
	return q.where("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_NotEqual(v string) *UserStockEvaluateQuery {
	return q.where("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_Less(v string) *UserStockEvaluateQuery {
	return q.where("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_LessEqual(v string) *UserStockEvaluateQuery {
	return q.where("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_Greater(v string) *UserStockEvaluateQuery {
	return q.where("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.where("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_Equal(v string) *UserStockEvaluateQuery {
	return q.where("stock_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_NotEqual(v string) *UserStockEvaluateQuery {
	return q.where("stock_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_Less(v string) *UserStockEvaluateQuery {
	return q.where("stock_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_LessEqual(v string) *UserStockEvaluateQuery {
	return q.where("stock_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_Greater(v string) *UserStockEvaluateQuery {
	return q.where("stock_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.where("stock_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_Equal(v float64) *UserStockEvaluateQuery {
	return q.where("total_score='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_NotEqual(v float64) *UserStockEvaluateQuery {
	return q.where("total_score<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_Less(v float64) *UserStockEvaluateQuery {
	return q.where("total_score<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_LessEqual(v float64) *UserStockEvaluateQuery {
	return q.where("total_score<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_Greater(v float64) *UserStockEvaluateQuery {
	return q.where("total_score>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_GreaterEqual(v float64) *UserStockEvaluateQuery {
	return q.where("total_score>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_Equal(v string) *UserStockEvaluateQuery {
	return q.where("eval_remark='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_NotEqual(v string) *UserStockEvaluateQuery {
	return q.where("eval_remark<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_Less(v string) *UserStockEvaluateQuery {
	return q.where("eval_remark<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_LessEqual(v string) *UserStockEvaluateQuery {
	return q.where("eval_remark<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_Greater(v string) *UserStockEvaluateQuery {
	return q.where("eval_remark>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.where("eval_remark>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_Equal(v time.Time) *UserStockEvaluateQuery {
	return q.where("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_NotEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_Less(v time.Time) *UserStockEvaluateQuery {
	return q.where("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_LessEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_Greater(v time.Time) *UserStockEvaluateQuery {
	return q.where("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_GreaterEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_Equal(v time.Time) *UserStockEvaluateQuery {
	return q.where("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_NotEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_Less(v time.Time) *UserStockEvaluateQuery {
	return q.where("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_LessEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_Greater(v time.Time) *UserStockEvaluateQuery {
	return q.where("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_GreaterEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("update_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_Equal(v string) *UserStockEvaluateQuery {
	return q.where("exchange_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_NotEqual(v string) *UserStockEvaluateQuery {
	return q.where("exchange_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_Less(v string) *UserStockEvaluateQuery {
	return q.where("exchange_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_LessEqual(v string) *UserStockEvaluateQuery {
	return q.where("exchange_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_Greater(v string) *UserStockEvaluateQuery {
	return q.where("exchange_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.where("exchange_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_Equal(v string) *UserStockEvaluateQuery {
	return q.where("stock_code='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_NotEqual(v string) *UserStockEvaluateQuery {
	return q.where("stock_code<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_Less(v string) *UserStockEvaluateQuery {
	return q.where("stock_code<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_LessEqual(v string) *UserStockEvaluateQuery {
	return q.where("stock_code<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_Greater(v string) *UserStockEvaluateQuery {
	return q.where("stock_code>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.where("stock_code>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_Equal(v time.Time) *UserStockEvaluateQuery {
	return q.where("launch_date='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_NotEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("launch_date<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_Less(v time.Time) *UserStockEvaluateQuery {
	return q.where("launch_date<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_LessEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("launch_date<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_Greater(v time.Time) *UserStockEvaluateQuery {
	return q.where("launch_date>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_GreaterEqual(v time.Time) *UserStockEvaluateQuery {
	return q.where("launch_date>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_Equal(v string) *UserStockEvaluateQuery {
	return q.where("industry_name='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_NotEqual(v string) *UserStockEvaluateQuery {
	return q.where("industry_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_Less(v string) *UserStockEvaluateQuery {
	return q.where("industry_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_LessEqual(v string) *UserStockEvaluateQuery {
	return q.where("industry_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_Greater(v string) *UserStockEvaluateQuery {
	return q.where("industry_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.where("industry_name>='" + fmt.Sprint(v) + "'")
}

type UserStockEvaluateDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewUserStockEvaluateDao(db *DB) (t *UserStockEvaluateDao, err error) {
	t = &UserStockEvaluateDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *UserStockEvaluateDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}
func (dao *UserStockEvaluateDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_stock_evaluate (user_id,stock_id,total_score,eval_remark,create_time,update_time,exchange_id,stock_code,launch_date,industry_name) VALUES (?,?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *UserStockEvaluateDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_stock_evaluate SET user_id=?,stock_id=?,total_score=?,eval_remark=?,create_time=?,update_time=?,exchange_id=?,stock_code=?,launch_date=?,industry_name=? WHERE id=?")
	return err
}

func (dao *UserStockEvaluateDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM user_stock_evaluate WHERE id=?")
	return err
}

func (dao *UserStockEvaluateDao) Insert(ctx context.Context, tx *wrap.Tx, e *UserStockEvaluate) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.EvalRemark, e.CreateTime, e.UpdateTime, e.ExchangeId, e.StockCode, e.LaunchDate, e.IndustryName)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *UserStockEvaluateDao) Update(ctx context.Context, tx *wrap.Tx, e *UserStockEvaluate) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.EvalRemark, e.CreateTime, e.UpdateTime, e.ExchangeId, e.StockCode, e.LaunchDate, e.IndustryName, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserStockEvaluateDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserStockEvaluateDao) scanRow(row *wrap.Row) (*UserStockEvaluate, error) {
	e := &UserStockEvaluate{}
	err := row.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.EvalRemark, &e.CreateTime, &e.UpdateTime, &e.ExchangeId, &e.StockCode, &e.LaunchDate, &e.IndustryName)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *UserStockEvaluateDao) scanRows(rows *wrap.Rows) (list []*UserStockEvaluate, err error) {
	list = make([]*UserStockEvaluate, 0)
	for rows.Next() {
		e := UserStockEvaluate{}
		err = rows.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.EvalRemark, &e.CreateTime, &e.UpdateTime, &e.ExchangeId, &e.StockCode, &e.LaunchDate, &e.IndustryName)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *UserStockEvaluateDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*UserStockEvaluate, error) {
	querySql := "SELECT " + USER_STOCK_EVALUATE_ALL_FIELDS_STRING + " FROM user_stock_evaluate " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserStockEvaluateDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserStockEvaluate, err error) {
	querySql := "SELECT " + USER_STOCK_EVALUATE_ALL_FIELDS_STRING + " FROM user_stock_evaluate " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *UserStockEvaluateDao) GetQuery() *UserStockEvaluateQuery {
	return NewUserStockEvaluateQuery(dao)
}

const USER_STOCK_INDEX_TABLE_NAME = "user_stock_index"

const USER_STOCK_INDEX_FIELD_ID = "id"
const USER_STOCK_INDEX_FIELD_USER_ID = "user_id"
const USER_STOCK_INDEX_FIELD_INDEX_NAME = "index_name"
const USER_STOCK_INDEX_FIELD_UI_ORDER = "ui_order"
const USER_STOCK_INDEX_FIELD_INDEX_DESC = "index_desc"
const USER_STOCK_INDEX_FIELD_EVAL_WEIGHT = "eval_weight"
const USER_STOCK_INDEX_FIELD_AI_WEIGHT = "ai_weight"
const USER_STOCK_INDEX_FIELD_NI_WEIGHT = "ni_weight"
const USER_STOCK_INDEX_FIELD_CREATE_TIME = "create_time"
const USER_STOCK_INDEX_FIELD_UPDATE_TIME = "update_time"

const USER_STOCK_INDEX_ALL_FIELDS_STRING = "id,user_id,index_name,ui_order,index_desc,eval_weight,ai_weight,ni_weight,create_time,update_time"

var USER_STOCK_INDEX_ALL_FIELDS = []string{
	"id",
	"user_id",
	"index_name",
	"ui_order",
	"index_desc",
	"eval_weight",
	"ai_weight",
	"ni_weight",
	"create_time",
	"update_time",
}

type UserStockIndex struct {
	Id         int64  //size=20
	UserId     string //size=32
	IndexName  string //size=32
	UiOrder    int32  //size=11
	IndexDesc  string //size=256
	EvalWeight int32  //size=11
	AiWeight   int32  //size=11
	NiWeight   int32  //size=11
	CreateTime time.Time
	UpdateTime time.Time
}

type UserStockIndexQuery struct {
	dao         *UserStockIndexDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewUserStockIndexQuery(dao *UserStockIndexDao) *UserStockIndexQuery {
	q := &UserStockIndexQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *UserStockIndexQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.forShare {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forUpdate {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	whereSql := q.whereBuffer.String()
	if whereSql != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(whereSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	return buf.String()
}

func (q *UserStockIndexQuery) Select(ctx context.Context) (*UserStockIndex, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *UserStockIndexQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*UserStockIndex, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*UserStockIndex, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) SelectList(ctx context.Context) (list []*UserStockIndex, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *UserStockIndexQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*UserStockIndex, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*UserStockIndex, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) Limit(startIncluded int64, count int64) *UserStockIndexQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *UserStockIndexQuery) Sort(fieldName string, asc bool) *UserStockIndexQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}

func (q *UserStockIndexQuery) where(format string, a ...interface{}) *UserStockIndexQuery {
	q.whereBuffer.WriteString(fmt.Sprintf(format, a...))
	return q
}

func (q *UserStockIndexQuery) Left() *UserStockIndexQuery  { return q.where(" ( ") }
func (q *UserStockIndexQuery) Right() *UserStockIndexQuery { return q.where(" ) ") }
func (q *UserStockIndexQuery) And() *UserStockIndexQuery   { return q.where(" AND ") }
func (q *UserStockIndexQuery) Or() *UserStockIndexQuery    { return q.where(" OR ") }
func (q *UserStockIndexQuery) Not() *UserStockIndexQuery   { return q.where(" NOT ") }

func (q *UserStockIndexQuery) Id_Equal(v int64) *UserStockIndexQuery {
	return q.where("id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_NotEqual(v int64) *UserStockIndexQuery {
	return q.where("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_Less(v int64) *UserStockIndexQuery {
	return q.where("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_LessEqual(v int64) *UserStockIndexQuery {
	return q.where("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_Greater(v int64) *UserStockIndexQuery {
	return q.where("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_GreaterEqual(v int64) *UserStockIndexQuery {
	return q.where("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_Equal(v string) *UserStockIndexQuery {
	return q.where("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_NotEqual(v string) *UserStockIndexQuery {
	return q.where("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_Less(v string) *UserStockIndexQuery {
	return q.where("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_LessEqual(v string) *UserStockIndexQuery {
	return q.where("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_Greater(v string) *UserStockIndexQuery {
	return q.where("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_GreaterEqual(v string) *UserStockIndexQuery {
	return q.where("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_Equal(v string) *UserStockIndexQuery {
	return q.where("index_name='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_NotEqual(v string) *UserStockIndexQuery {
	return q.where("index_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_Less(v string) *UserStockIndexQuery {
	return q.where("index_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_LessEqual(v string) *UserStockIndexQuery {
	return q.where("index_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_Greater(v string) *UserStockIndexQuery {
	return q.where("index_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_GreaterEqual(v string) *UserStockIndexQuery {
	return q.where("index_name>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_Equal(v int32) *UserStockIndexQuery {
	return q.where("ui_order='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_NotEqual(v int32) *UserStockIndexQuery {
	return q.where("ui_order<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_Less(v int32) *UserStockIndexQuery {
	return q.where("ui_order<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_LessEqual(v int32) *UserStockIndexQuery {
	return q.where("ui_order<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_Greater(v int32) *UserStockIndexQuery {
	return q.where("ui_order>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.where("ui_order>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_Equal(v string) *UserStockIndexQuery {
	return q.where("index_desc='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_NotEqual(v string) *UserStockIndexQuery {
	return q.where("index_desc<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_Less(v string) *UserStockIndexQuery {
	return q.where("index_desc<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_LessEqual(v string) *UserStockIndexQuery {
	return q.where("index_desc<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_Greater(v string) *UserStockIndexQuery {
	return q.where("index_desc>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_GreaterEqual(v string) *UserStockIndexQuery {
	return q.where("index_desc>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_Equal(v int32) *UserStockIndexQuery {
	return q.where("eval_weight='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_NotEqual(v int32) *UserStockIndexQuery {
	return q.where("eval_weight<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_Less(v int32) *UserStockIndexQuery {
	return q.where("eval_weight<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_LessEqual(v int32) *UserStockIndexQuery {
	return q.where("eval_weight<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_Greater(v int32) *UserStockIndexQuery {
	return q.where("eval_weight>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.where("eval_weight>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_Equal(v int32) *UserStockIndexQuery {
	return q.where("ai_weight='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_NotEqual(v int32) *UserStockIndexQuery {
	return q.where("ai_weight<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_Less(v int32) *UserStockIndexQuery {
	return q.where("ai_weight<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_LessEqual(v int32) *UserStockIndexQuery {
	return q.where("ai_weight<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_Greater(v int32) *UserStockIndexQuery {
	return q.where("ai_weight>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.where("ai_weight>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) NiWeight_Equal(v int32) *UserStockIndexQuery {
	return q.where("ni_weight='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) NiWeight_NotEqual(v int32) *UserStockIndexQuery {
	return q.where("ni_weight<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) NiWeight_Less(v int32) *UserStockIndexQuery {
	return q.where("ni_weight<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) NiWeight_LessEqual(v int32) *UserStockIndexQuery {
	return q.where("ni_weight<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) NiWeight_Greater(v int32) *UserStockIndexQuery {
	return q.where("ni_weight>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) NiWeight_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.where("ni_weight>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_Equal(v time.Time) *UserStockIndexQuery {
	return q.where("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_NotEqual(v time.Time) *UserStockIndexQuery {
	return q.where("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_Less(v time.Time) *UserStockIndexQuery {
	return q.where("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_LessEqual(v time.Time) *UserStockIndexQuery {
	return q.where("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_Greater(v time.Time) *UserStockIndexQuery {
	return q.where("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_GreaterEqual(v time.Time) *UserStockIndexQuery {
	return q.where("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_Equal(v time.Time) *UserStockIndexQuery {
	return q.where("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_NotEqual(v time.Time) *UserStockIndexQuery {
	return q.where("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_Less(v time.Time) *UserStockIndexQuery {
	return q.where("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_LessEqual(v time.Time) *UserStockIndexQuery {
	return q.where("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_Greater(v time.Time) *UserStockIndexQuery {
	return q.where("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_GreaterEqual(v time.Time) *UserStockIndexQuery {
	return q.where("update_time>='" + fmt.Sprint(v) + "'")
}

type UserStockIndexDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewUserStockIndexDao(db *DB) (t *UserStockIndexDao, err error) {
	t = &UserStockIndexDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *UserStockIndexDao) init() (err error) {
	err = dao.prepareInsertStmt()
	if err != nil {
		return err
	}

	err = dao.prepareUpdateStmt()
	if err != nil {
		return err
	}

	err = dao.prepareDeleteStmt()
	if err != nil {
		return err
	}

	return nil
}
func (dao *UserStockIndexDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_stock_index (user_id,index_name,ui_order,index_desc,eval_weight,ai_weight,ni_weight,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *UserStockIndexDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_stock_index SET user_id=?,index_name=?,ui_order=?,index_desc=?,eval_weight=?,ai_weight=?,ni_weight=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *UserStockIndexDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM user_stock_index WHERE id=?")
	return err
}

func (dao *UserStockIndexDao) Insert(ctx context.Context, tx *wrap.Tx, e *UserStockIndex) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.IndexName, e.UiOrder, e.IndexDesc, e.EvalWeight, e.AiWeight, e.NiWeight, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *UserStockIndexDao) Update(ctx context.Context, tx *wrap.Tx, e *UserStockIndex) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.IndexName, e.UiOrder, e.IndexDesc, e.EvalWeight, e.AiWeight, e.NiWeight, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserStockIndexDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("rowsAffected:%s", rowsAffected)
	}

	return nil
}

func (dao *UserStockIndexDao) scanRow(row *wrap.Row) (*UserStockIndex, error) {
	e := &UserStockIndex{}
	err := row.Scan(&e.Id, &e.UserId, &e.IndexName, &e.UiOrder, &e.IndexDesc, &e.EvalWeight, &e.AiWeight, &e.NiWeight, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *UserStockIndexDao) scanRows(rows *wrap.Rows) (list []*UserStockIndex, err error) {
	list = make([]*UserStockIndex, 0)
	for rows.Next() {
		e := UserStockIndex{}
		err = rows.Scan(&e.Id, &e.UserId, &e.IndexName, &e.UiOrder, &e.IndexDesc, &e.EvalWeight, &e.AiWeight, &e.NiWeight, &e.CreateTime, &e.UpdateTime)
		if err != nil {
			return nil, err
		}
		list = append(list, &e)
	}
	if rows.Err() != nil {
		err = rows.Err()
		return nil, err
	}

	return list, nil
}

func (dao *UserStockIndexDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*UserStockIndex, error) {
	querySql := "SELECT " + USER_STOCK_INDEX_ALL_FIELDS_STRING + " FROM user_stock_index " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserStockIndexDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserStockIndex, err error) {
	querySql := "SELECT " + USER_STOCK_INDEX_ALL_FIELDS_STRING + " FROM user_stock_index " + query
	var rows *wrap.Rows
	if tx == nil {
		rows, err = dao.db.Query(ctx, querySql)
	} else {
		rows, err = tx.Query(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *UserStockIndexDao) GetQuery() *UserStockIndexQuery {
	return NewUserStockIndexQuery(dao)
}

type DB struct {
	wrap.DB
	Stock             *StockDao
	UserIndexEvaluate *UserIndexEvaluateDao
	UserSetting       *UserSettingDao
	UserStockEvaluate *UserStockEvaluateDao
	UserStockIndex    *UserStockIndexDao
}

func NewDB(connectionString string) (d *DB, err error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connectionString nil")
	}

	d = &DB{}

	db, err := wrap.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	d.DB = *db

	err = d.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	d.Stock, err = NewStockDao(d)
	if err != nil {
		return nil, err
	}

	d.UserIndexEvaluate, err = NewUserIndexEvaluateDao(d)
	if err != nil {
		return nil, err
	}

	d.UserSetting, err = NewUserSettingDao(d)
	if err != nil {
		return nil, err
	}

	d.UserStockEvaluate, err = NewUserStockEvaluateDao(d)
	if err != nil {
		return nil, err
	}

	d.UserStockIndex, err = NewUserStockIndexDao(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
