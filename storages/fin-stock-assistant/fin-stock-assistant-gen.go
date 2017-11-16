package fin_stock_assistant

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/sql/wrap"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"strings"
	"time"
)

type BaseQuery struct {
	forUpdate     bool
	forShare      bool
	where         string
	limit         string
	order         string
	groupByFields []string
}

func (q *BaseQuery) buildQueryString() string {
	buf := bytes.NewBufferString("")

	if q.where != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(q.where)
	}

	if q.groupByFields != nil && len(q.groupByFields) > 0 {
		buf.WriteString(" GROUP BY ")
		buf.WriteString(strings.Join(q.groupByFields, ","))
	}

	if q.order != "" {
		buf.WriteString(" order by ")
		buf.WriteString(q.order)
	}

	if q.limit != "" {
		buf.WriteString(q.limit)
	}

	if q.forUpdate {
		buf.WriteString(" FOR UPDATE ")
	}

	if q.forShare {
		buf.WriteString(" LOCK IN SHARE MODE ")
	}

	return buf.String()
}

const STOCK_TABLE_NAME = "stock"

type STOCK_FIELD string

const STOCK_FIELD_ID = STOCK_FIELD("id")
const STOCK_FIELD_STOCK_ID = STOCK_FIELD("stock_id")
const STOCK_FIELD_EXCHANGE_ID = STOCK_FIELD("exchange_id")
const STOCK_FIELD_STOCK_CODE = STOCK_FIELD("stock_code")
const STOCK_FIELD_STOCK_NAME_CN = STOCK_FIELD("stock_name_cn")
const STOCK_FIELD_STOCK_NAME_EN = STOCK_FIELD("stock_name_en")
const STOCK_FIELD_LAUNCH_DATE = STOCK_FIELD("launch_date")
const STOCK_FIELD_COMPANY_NAME_CN = STOCK_FIELD("company_name_cn")
const STOCK_FIELD_COMPANY_NAME_EN = STOCK_FIELD("company_name_en")
const STOCK_FIELD_WEBSITE_URL = STOCK_FIELD("website_url")
const STOCK_FIELD_INDUSTRY_NAME = STOCK_FIELD("industry_name")
const STOCK_FIELD_CITY_NAME_CN = STOCK_FIELD("city_name_cn")
const STOCK_FIELD_CITY_NAME_EN = STOCK_FIELD("city_name_en")
const STOCK_FIELD_PROVINCE_NAME_CN = STOCK_FIELD("province_name_cn")
const STOCK_FIELD_PROVINCE_NAME_EN = STOCK_FIELD("province_name_en")
const STOCK_FIELD_CREATE_TIME = STOCK_FIELD("create_time")
const STOCK_FIELD_UPDATE_TIME = STOCK_FIELD("update_time")

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
	BaseQuery
	dao *StockDao
}

func NewStockQuery(dao *StockDao) *StockQuery {
	q := &StockQuery{}
	q.dao = dao

	return q
}

func (q *StockQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*Stock, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*Stock, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *StockQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *StockQuery) ForUpdate() *StockQuery {
	q.forUpdate = true
	return q
}

func (q *StockQuery) ForShare() *StockQuery {
	q.forShare = true
	return q
}

func (q *StockQuery) GroupBy(fields ...STOCK_FIELD) *StockQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *StockQuery) Limit(startIncluded int64, count int64) *StockQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *StockQuery) OrderBy(fieldName STOCK_FIELD, asc bool) *StockQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *StockQuery) OrderByGroupCount(asc bool) *StockQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *StockQuery) w(format string, a ...interface{}) *StockQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *StockQuery) Left() *StockQuery  { return q.w(" ( ") }
func (q *StockQuery) Right() *StockQuery { return q.w(" ) ") }
func (q *StockQuery) And() *StockQuery   { return q.w(" AND ") }
func (q *StockQuery) Or() *StockQuery    { return q.w(" OR ") }
func (q *StockQuery) Not() *StockQuery   { return q.w(" NOT ") }

func (q *StockQuery) Id_Equal(v int64) *StockQuery        { return q.w("id='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_NotEqual(v int64) *StockQuery     { return q.w("id<>'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_Less(v int64) *StockQuery         { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_LessEqual(v int64) *StockQuery    { return q.w("id<='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_Greater(v int64) *StockQuery      { return q.w("id>'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) Id_GreaterEqual(v int64) *StockQuery { return q.w("id>='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) StockId_Equal(v string) *StockQuery  { return q.w("stock_id='" + fmt.Sprint(v) + "'") }
func (q *StockQuery) StockId_NotEqual(v string) *StockQuery {
	return q.w("stock_id<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_Less(v string) *StockQuery { return q.w("stock_id<'" + fmt.Sprint(v) + "'") }
func (q *StockQuery) StockId_LessEqual(v string) *StockQuery {
	return q.w("stock_id<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_Greater(v string) *StockQuery {
	return q.w("stock_id>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockId_GreaterEqual(v string) *StockQuery {
	return q.w("stock_id>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_Equal(v string) *StockQuery {
	return q.w("exchange_id='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_NotEqual(v string) *StockQuery {
	return q.w("exchange_id<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_Less(v string) *StockQuery {
	return q.w("exchange_id<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_LessEqual(v string) *StockQuery {
	return q.w("exchange_id<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_Greater(v string) *StockQuery {
	return q.w("exchange_id>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ExchangeId_GreaterEqual(v string) *StockQuery {
	return q.w("exchange_id>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_Equal(v string) *StockQuery {
	return q.w("stock_code='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_NotEqual(v string) *StockQuery {
	return q.w("stock_code<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_Less(v string) *StockQuery {
	return q.w("stock_code<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_LessEqual(v string) *StockQuery {
	return q.w("stock_code<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_Greater(v string) *StockQuery {
	return q.w("stock_code>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockCode_GreaterEqual(v string) *StockQuery {
	return q.w("stock_code>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_Equal(v string) *StockQuery {
	return q.w("stock_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_NotEqual(v string) *StockQuery {
	return q.w("stock_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_Less(v string) *StockQuery {
	return q.w("stock_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_LessEqual(v string) *StockQuery {
	return q.w("stock_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_Greater(v string) *StockQuery {
	return q.w("stock_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameCn_GreaterEqual(v string) *StockQuery {
	return q.w("stock_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_Equal(v string) *StockQuery {
	return q.w("stock_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_NotEqual(v string) *StockQuery {
	return q.w("stock_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_Less(v string) *StockQuery {
	return q.w("stock_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_LessEqual(v string) *StockQuery {
	return q.w("stock_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_Greater(v string) *StockQuery {
	return q.w("stock_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) StockNameEn_GreaterEqual(v string) *StockQuery {
	return q.w("stock_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_Equal(v time.Time) *StockQuery {
	return q.w("launch_date='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_NotEqual(v time.Time) *StockQuery {
	return q.w("launch_date<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_Less(v time.Time) *StockQuery {
	return q.w("launch_date<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_LessEqual(v time.Time) *StockQuery {
	return q.w("launch_date<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_Greater(v time.Time) *StockQuery {
	return q.w("launch_date>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) LaunchDate_GreaterEqual(v time.Time) *StockQuery {
	return q.w("launch_date>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_Equal(v string) *StockQuery {
	return q.w("company_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_NotEqual(v string) *StockQuery {
	return q.w("company_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_Less(v string) *StockQuery {
	return q.w("company_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_LessEqual(v string) *StockQuery {
	return q.w("company_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_Greater(v string) *StockQuery {
	return q.w("company_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameCn_GreaterEqual(v string) *StockQuery {
	return q.w("company_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_Equal(v string) *StockQuery {
	return q.w("company_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_NotEqual(v string) *StockQuery {
	return q.w("company_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_Less(v string) *StockQuery {
	return q.w("company_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_LessEqual(v string) *StockQuery {
	return q.w("company_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_Greater(v string) *StockQuery {
	return q.w("company_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CompanyNameEn_GreaterEqual(v string) *StockQuery {
	return q.w("company_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_Equal(v string) *StockQuery {
	return q.w("website_url='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_NotEqual(v string) *StockQuery {
	return q.w("website_url<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_Less(v string) *StockQuery {
	return q.w("website_url<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_LessEqual(v string) *StockQuery {
	return q.w("website_url<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_Greater(v string) *StockQuery {
	return q.w("website_url>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) WebsiteUrl_GreaterEqual(v string) *StockQuery {
	return q.w("website_url>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_Equal(v string) *StockQuery {
	return q.w("industry_name='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_NotEqual(v string) *StockQuery {
	return q.w("industry_name<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_Less(v string) *StockQuery {
	return q.w("industry_name<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_LessEqual(v string) *StockQuery {
	return q.w("industry_name<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_Greater(v string) *StockQuery {
	return q.w("industry_name>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) IndustryName_GreaterEqual(v string) *StockQuery {
	return q.w("industry_name>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_Equal(v string) *StockQuery {
	return q.w("city_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_NotEqual(v string) *StockQuery {
	return q.w("city_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_Less(v string) *StockQuery {
	return q.w("city_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_LessEqual(v string) *StockQuery {
	return q.w("city_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_Greater(v string) *StockQuery {
	return q.w("city_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameCn_GreaterEqual(v string) *StockQuery {
	return q.w("city_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_Equal(v string) *StockQuery {
	return q.w("city_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_NotEqual(v string) *StockQuery {
	return q.w("city_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_Less(v string) *StockQuery {
	return q.w("city_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_LessEqual(v string) *StockQuery {
	return q.w("city_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_Greater(v string) *StockQuery {
	return q.w("city_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CityNameEn_GreaterEqual(v string) *StockQuery {
	return q.w("city_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_Equal(v string) *StockQuery {
	return q.w("province_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_NotEqual(v string) *StockQuery {
	return q.w("province_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_Less(v string) *StockQuery {
	return q.w("province_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_LessEqual(v string) *StockQuery {
	return q.w("province_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_Greater(v string) *StockQuery {
	return q.w("province_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameCn_GreaterEqual(v string) *StockQuery {
	return q.w("province_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_Equal(v string) *StockQuery {
	return q.w("province_name_en='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_NotEqual(v string) *StockQuery {
	return q.w("province_name_en<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_Less(v string) *StockQuery {
	return q.w("province_name_en<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_LessEqual(v string) *StockQuery {
	return q.w("province_name_en<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_Greater(v string) *StockQuery {
	return q.w("province_name_en>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) ProvinceNameEn_GreaterEqual(v string) *StockQuery {
	return q.w("province_name_en>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_Equal(v time.Time) *StockQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_NotEqual(v time.Time) *StockQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_Less(v time.Time) *StockQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_LessEqual(v time.Time) *StockQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_Greater(v time.Time) *StockQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) CreateTime_GreaterEqual(v time.Time) *StockQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_Equal(v time.Time) *StockQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_NotEqual(v time.Time) *StockQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_Less(v time.Time) *StockQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_LessEqual(v time.Time) *StockQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_Greater(v time.Time) *StockQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *StockQuery) UpdateTime_GreaterEqual(v time.Time) *StockQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
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

	_, err = stmt.Exec(ctx, e.StockId, e.ExchangeId, e.StockCode, e.StockNameCn, e.StockNameEn, e.LaunchDate, e.CompanyNameCn, e.CompanyNameEn, e.WebsiteUrl, e.IndustryName, e.CityNameCn, e.CityNameEn, e.ProvinceNameCn, e.ProvinceNameEn, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
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

func (dao *StockDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*Stock, error) {
	querySql := "SELECT " + STOCK_ALL_FIELDS_STRING + " FROM stock " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *StockDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*Stock, err error) {
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

func (dao *StockDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM stock " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *StockDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM stock " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *StockDao) GetQuery() *StockQuery {
	return NewStockQuery(dao)
}

const STOCK_INDEX_ADVICE_TABLE_NAME = "stock_index_advice"

type STOCK_INDEX_ADVICE_FIELD string

const STOCK_INDEX_ADVICE_FIELD_ID = STOCK_INDEX_ADVICE_FIELD("id")
const STOCK_INDEX_ADVICE_FIELD_INDEX_NAME = STOCK_INDEX_ADVICE_FIELD("index_name")
const STOCK_INDEX_ADVICE_FIELD_USED_COUNT = STOCK_INDEX_ADVICE_FIELD("used_count")
const STOCK_INDEX_ADVICE_FIELD_CREATE_TIME = STOCK_INDEX_ADVICE_FIELD("create_time")
const STOCK_INDEX_ADVICE_FIELD_UPDATE_TIME = STOCK_INDEX_ADVICE_FIELD("update_time")

const STOCK_INDEX_ADVICE_ALL_FIELDS_STRING = "id,index_name,used_count,create_time,update_time"

var STOCK_INDEX_ADVICE_ALL_FIELDS = []string{
	"id",
	"index_name",
	"used_count",
	"create_time",
	"update_time",
}

type StockIndexAdvice struct {
	Id         int64  //size=20
	IndexName  string //size=32
	UsedCount  int64  //size=20
	CreateTime time.Time
	UpdateTime time.Time
}

type StockIndexAdviceQuery struct {
	BaseQuery
	dao *StockIndexAdviceDao
}

func NewStockIndexAdviceQuery(dao *StockIndexAdviceDao) *StockIndexAdviceQuery {
	q := &StockIndexAdviceQuery{}
	q.dao = dao

	return q
}

func (q *StockIndexAdviceQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*StockIndexAdvice, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *StockIndexAdviceQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*StockIndexAdvice, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *StockIndexAdviceQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *StockIndexAdviceQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *StockIndexAdviceQuery) ForUpdate() *StockIndexAdviceQuery {
	q.forUpdate = true
	return q
}

func (q *StockIndexAdviceQuery) ForShare() *StockIndexAdviceQuery {
	q.forShare = true
	return q
}

func (q *StockIndexAdviceQuery) GroupBy(fields ...STOCK_INDEX_ADVICE_FIELD) *StockIndexAdviceQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *StockIndexAdviceQuery) Limit(startIncluded int64, count int64) *StockIndexAdviceQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *StockIndexAdviceQuery) OrderBy(fieldName STOCK_INDEX_ADVICE_FIELD, asc bool) *StockIndexAdviceQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *StockIndexAdviceQuery) OrderByGroupCount(asc bool) *StockIndexAdviceQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *StockIndexAdviceQuery) w(format string, a ...interface{}) *StockIndexAdviceQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *StockIndexAdviceQuery) Left() *StockIndexAdviceQuery  { return q.w(" ( ") }
func (q *StockIndexAdviceQuery) Right() *StockIndexAdviceQuery { return q.w(" ) ") }
func (q *StockIndexAdviceQuery) And() *StockIndexAdviceQuery   { return q.w(" AND ") }
func (q *StockIndexAdviceQuery) Or() *StockIndexAdviceQuery    { return q.w(" OR ") }
func (q *StockIndexAdviceQuery) Not() *StockIndexAdviceQuery   { return q.w(" NOT ") }

func (q *StockIndexAdviceQuery) Id_Equal(v int64) *StockIndexAdviceQuery {
	return q.w("id='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) Id_NotEqual(v int64) *StockIndexAdviceQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) Id_Less(v int64) *StockIndexAdviceQuery {
	return q.w("id<'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) Id_LessEqual(v int64) *StockIndexAdviceQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) Id_Greater(v int64) *StockIndexAdviceQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) Id_GreaterEqual(v int64) *StockIndexAdviceQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) IndexName_Equal(v string) *StockIndexAdviceQuery {
	return q.w("index_name='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) IndexName_NotEqual(v string) *StockIndexAdviceQuery {
	return q.w("index_name<>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) IndexName_Less(v string) *StockIndexAdviceQuery {
	return q.w("index_name<'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) IndexName_LessEqual(v string) *StockIndexAdviceQuery {
	return q.w("index_name<='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) IndexName_Greater(v string) *StockIndexAdviceQuery {
	return q.w("index_name>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) IndexName_GreaterEqual(v string) *StockIndexAdviceQuery {
	return q.w("index_name>='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UsedCount_Equal(v int64) *StockIndexAdviceQuery {
	return q.w("used_count='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UsedCount_NotEqual(v int64) *StockIndexAdviceQuery {
	return q.w("used_count<>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UsedCount_Less(v int64) *StockIndexAdviceQuery {
	return q.w("used_count<'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UsedCount_LessEqual(v int64) *StockIndexAdviceQuery {
	return q.w("used_count<='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UsedCount_Greater(v int64) *StockIndexAdviceQuery {
	return q.w("used_count>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UsedCount_GreaterEqual(v int64) *StockIndexAdviceQuery {
	return q.w("used_count>='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) CreateTime_Equal(v time.Time) *StockIndexAdviceQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) CreateTime_NotEqual(v time.Time) *StockIndexAdviceQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) CreateTime_Less(v time.Time) *StockIndexAdviceQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) CreateTime_LessEqual(v time.Time) *StockIndexAdviceQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) CreateTime_Greater(v time.Time) *StockIndexAdviceQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) CreateTime_GreaterEqual(v time.Time) *StockIndexAdviceQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UpdateTime_Equal(v time.Time) *StockIndexAdviceQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UpdateTime_NotEqual(v time.Time) *StockIndexAdviceQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UpdateTime_Less(v time.Time) *StockIndexAdviceQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UpdateTime_LessEqual(v time.Time) *StockIndexAdviceQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UpdateTime_Greater(v time.Time) *StockIndexAdviceQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *StockIndexAdviceQuery) UpdateTime_GreaterEqual(v time.Time) *StockIndexAdviceQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
}

type StockIndexAdviceDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewStockIndexAdviceDao(db *DB) (t *StockIndexAdviceDao, err error) {
	t = &StockIndexAdviceDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *StockIndexAdviceDao) init() (err error) {
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
func (dao *StockIndexAdviceDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock_index_advice (index_name,used_count,create_time,update_time) VALUES (?,?,?,?)")
	return err
}

func (dao *StockIndexAdviceDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock_index_advice SET index_name=?,used_count=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockIndexAdviceDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock_index_advice WHERE id=?")
	return err
}

func (dao *StockIndexAdviceDao) Insert(ctx context.Context, tx *wrap.Tx, e *StockIndexAdvice) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.IndexName, e.UsedCount, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockIndexAdviceDao) Update(ctx context.Context, tx *wrap.Tx, e *StockIndexAdvice) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, e.IndexName, e.UsedCount, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockIndexAdviceDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *StockIndexAdviceDao) scanRow(row *wrap.Row) (*StockIndexAdvice, error) {
	e := &StockIndexAdvice{}
	err := row.Scan(&e.Id, &e.IndexName, &e.UsedCount, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockIndexAdviceDao) scanRows(rows *wrap.Rows) (list []*StockIndexAdvice, err error) {
	list = make([]*StockIndexAdvice, 0)
	for rows.Next() {
		e := StockIndexAdvice{}
		err = rows.Scan(&e.Id, &e.IndexName, &e.UsedCount, &e.CreateTime, &e.UpdateTime)
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

func (dao *StockIndexAdviceDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*StockIndexAdvice, error) {
	querySql := "SELECT " + STOCK_INDEX_ADVICE_ALL_FIELDS_STRING + " FROM stock_index_advice " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *StockIndexAdviceDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*StockIndexAdvice, err error) {
	querySql := "SELECT " + STOCK_INDEX_ADVICE_ALL_FIELDS_STRING + " FROM stock_index_advice " + query
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

func (dao *StockIndexAdviceDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM stock_index_advice " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *StockIndexAdviceDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM stock_index_advice " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *StockIndexAdviceDao) GetQuery() *StockIndexAdviceQuery {
	return NewStockIndexAdviceQuery(dao)
}

const USER_INDEX_EVALUATE_TABLE_NAME = "user_index_evaluate"

type USER_INDEX_EVALUATE_FIELD string

const USER_INDEX_EVALUATE_FIELD_ID = USER_INDEX_EVALUATE_FIELD("id")
const USER_INDEX_EVALUATE_FIELD_USER_ID = USER_INDEX_EVALUATE_FIELD("user_id")
const USER_INDEX_EVALUATE_FIELD_STOCK_ID = USER_INDEX_EVALUATE_FIELD("stock_id")
const USER_INDEX_EVALUATE_FIELD_INDEX_NAME = USER_INDEX_EVALUATE_FIELD("index_name")
const USER_INDEX_EVALUATE_FIELD_EVAL_STARS = USER_INDEX_EVALUATE_FIELD("eval_stars")
const USER_INDEX_EVALUATE_FIELD_EVAL_REMARK = USER_INDEX_EVALUATE_FIELD("eval_remark")
const USER_INDEX_EVALUATE_FIELD_CREATE_TIME = USER_INDEX_EVALUATE_FIELD("create_time")
const USER_INDEX_EVALUATE_FIELD_UPDATE_TIME = USER_INDEX_EVALUATE_FIELD("update_time")

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
	BaseQuery
	dao *UserIndexEvaluateDao
}

func NewUserIndexEvaluateQuery(dao *UserIndexEvaluateDao) *UserIndexEvaluateQuery {
	q := &UserIndexEvaluateQuery{}
	q.dao = dao

	return q
}

func (q *UserIndexEvaluateQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*UserIndexEvaluate, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*UserIndexEvaluate, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *UserIndexEvaluateQuery) ForUpdate() *UserIndexEvaluateQuery {
	q.forUpdate = true
	return q
}

func (q *UserIndexEvaluateQuery) ForShare() *UserIndexEvaluateQuery {
	q.forShare = true
	return q
}

func (q *UserIndexEvaluateQuery) GroupBy(fields ...USER_INDEX_EVALUATE_FIELD) *UserIndexEvaluateQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *UserIndexEvaluateQuery) Limit(startIncluded int64, count int64) *UserIndexEvaluateQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *UserIndexEvaluateQuery) OrderBy(fieldName USER_INDEX_EVALUATE_FIELD, asc bool) *UserIndexEvaluateQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserIndexEvaluateQuery) OrderByGroupCount(asc bool) *UserIndexEvaluateQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserIndexEvaluateQuery) w(format string, a ...interface{}) *UserIndexEvaluateQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *UserIndexEvaluateQuery) Left() *UserIndexEvaluateQuery  { return q.w(" ( ") }
func (q *UserIndexEvaluateQuery) Right() *UserIndexEvaluateQuery { return q.w(" ) ") }
func (q *UserIndexEvaluateQuery) And() *UserIndexEvaluateQuery   { return q.w(" AND ") }
func (q *UserIndexEvaluateQuery) Or() *UserIndexEvaluateQuery    { return q.w(" OR ") }
func (q *UserIndexEvaluateQuery) Not() *UserIndexEvaluateQuery   { return q.w(" NOT ") }

func (q *UserIndexEvaluateQuery) Id_Equal(v int64) *UserIndexEvaluateQuery {
	return q.w("id='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_NotEqual(v int64) *UserIndexEvaluateQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_Less(v int64) *UserIndexEvaluateQuery {
	return q.w("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_LessEqual(v int64) *UserIndexEvaluateQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_Greater(v int64) *UserIndexEvaluateQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) Id_GreaterEqual(v int64) *UserIndexEvaluateQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_Equal(v string) *UserIndexEvaluateQuery {
	return q.w("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.w("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_Less(v string) *UserIndexEvaluateQuery {
	return q.w("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.w("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_Greater(v string) *UserIndexEvaluateQuery {
	return q.w("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UserId_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.w("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_Equal(v string) *UserIndexEvaluateQuery {
	return q.w("stock_id='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.w("stock_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_Less(v string) *UserIndexEvaluateQuery {
	return q.w("stock_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.w("stock_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_Greater(v string) *UserIndexEvaluateQuery {
	return q.w("stock_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) StockId_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.w("stock_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_Equal(v string) *UserIndexEvaluateQuery {
	return q.w("index_name='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.w("index_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_Less(v string) *UserIndexEvaluateQuery {
	return q.w("index_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.w("index_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_Greater(v string) *UserIndexEvaluateQuery {
	return q.w("index_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) IndexName_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.w("index_name>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_Equal(v int32) *UserIndexEvaluateQuery {
	return q.w("eval_stars='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_NotEqual(v int32) *UserIndexEvaluateQuery {
	return q.w("eval_stars<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_Less(v int32) *UserIndexEvaluateQuery {
	return q.w("eval_stars<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_LessEqual(v int32) *UserIndexEvaluateQuery {
	return q.w("eval_stars<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_Greater(v int32) *UserIndexEvaluateQuery {
	return q.w("eval_stars>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalStars_GreaterEqual(v int32) *UserIndexEvaluateQuery {
	return q.w("eval_stars>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_Equal(v string) *UserIndexEvaluateQuery {
	return q.w("eval_remark='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_NotEqual(v string) *UserIndexEvaluateQuery {
	return q.w("eval_remark<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_Less(v string) *UserIndexEvaluateQuery {
	return q.w("eval_remark<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_LessEqual(v string) *UserIndexEvaluateQuery {
	return q.w("eval_remark<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_Greater(v string) *UserIndexEvaluateQuery {
	return q.w("eval_remark>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) EvalRemark_GreaterEqual(v string) *UserIndexEvaluateQuery {
	return q.w("eval_remark>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_Equal(v time.Time) *UserIndexEvaluateQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_NotEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_Less(v time.Time) *UserIndexEvaluateQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_LessEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_Greater(v time.Time) *UserIndexEvaluateQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) CreateTime_GreaterEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_Equal(v time.Time) *UserIndexEvaluateQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_NotEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_Less(v time.Time) *UserIndexEvaluateQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_LessEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_Greater(v time.Time) *UserIndexEvaluateQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserIndexEvaluateQuery) UpdateTime_GreaterEqual(v time.Time) *UserIndexEvaluateQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
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

	_, err = stmt.Exec(ctx, e.UserId, e.StockId, e.IndexName, e.EvalStars, e.EvalRemark, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserIndexEvaluateDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
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

func (dao *UserIndexEvaluateDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*UserIndexEvaluate, error) {
	querySql := "SELECT " + USER_INDEX_EVALUATE_ALL_FIELDS_STRING + " FROM user_index_evaluate " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserIndexEvaluateDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserIndexEvaluate, err error) {
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

func (dao *UserIndexEvaluateDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM user_index_evaluate " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *UserIndexEvaluateDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM user_index_evaluate " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *UserIndexEvaluateDao) GetQuery() *UserIndexEvaluateQuery {
	return NewUserIndexEvaluateQuery(dao)
}

const USER_SETTING_TABLE_NAME = "user_setting"

type USER_SETTING_FIELD string

const USER_SETTING_FIELD_ID = USER_SETTING_FIELD("id")
const USER_SETTING_FIELD_USER_ID = USER_SETTING_FIELD("user_id")
const USER_SETTING_FIELD_CONFIG_KEY = USER_SETTING_FIELD("config_key")
const USER_SETTING_FIELD_CONFIG_VALUE = USER_SETTING_FIELD("config_value")
const USER_SETTING_FIELD_CREATE_TIME = USER_SETTING_FIELD("create_time")
const USER_SETTING_FIELD_UPDATE_TIME = USER_SETTING_FIELD("update_time")

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
	BaseQuery
	dao *UserSettingDao
}

func NewUserSettingQuery(dao *UserSettingDao) *UserSettingQuery {
	q := &UserSettingQuery{}
	q.dao = dao

	return q
}

func (q *UserSettingQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*UserSetting, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*UserSetting, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *UserSettingQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *UserSettingQuery) ForUpdate() *UserSettingQuery {
	q.forUpdate = true
	return q
}

func (q *UserSettingQuery) ForShare() *UserSettingQuery {
	q.forShare = true
	return q
}

func (q *UserSettingQuery) GroupBy(fields ...USER_SETTING_FIELD) *UserSettingQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *UserSettingQuery) Limit(startIncluded int64, count int64) *UserSettingQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *UserSettingQuery) OrderBy(fieldName USER_SETTING_FIELD, asc bool) *UserSettingQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserSettingQuery) OrderByGroupCount(asc bool) *UserSettingQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserSettingQuery) w(format string, a ...interface{}) *UserSettingQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *UserSettingQuery) Left() *UserSettingQuery  { return q.w(" ( ") }
func (q *UserSettingQuery) Right() *UserSettingQuery { return q.w(" ) ") }
func (q *UserSettingQuery) And() *UserSettingQuery   { return q.w(" AND ") }
func (q *UserSettingQuery) Or() *UserSettingQuery    { return q.w(" OR ") }
func (q *UserSettingQuery) Not() *UserSettingQuery   { return q.w(" NOT ") }

func (q *UserSettingQuery) Id_Equal(v int64) *UserSettingQuery { return q.w("id='" + fmt.Sprint(v) + "'") }
func (q *UserSettingQuery) Id_NotEqual(v int64) *UserSettingQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_Less(v int64) *UserSettingQuery { return q.w("id<'" + fmt.Sprint(v) + "'") }
func (q *UserSettingQuery) Id_LessEqual(v int64) *UserSettingQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_Greater(v int64) *UserSettingQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) Id_GreaterEqual(v int64) *UserSettingQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_Equal(v string) *UserSettingQuery {
	return q.w("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_NotEqual(v string) *UserSettingQuery {
	return q.w("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_Less(v string) *UserSettingQuery {
	return q.w("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_LessEqual(v string) *UserSettingQuery {
	return q.w("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_Greater(v string) *UserSettingQuery {
	return q.w("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UserId_GreaterEqual(v string) *UserSettingQuery {
	return q.w("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_Equal(v string) *UserSettingQuery {
	return q.w("config_key='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_NotEqual(v string) *UserSettingQuery {
	return q.w("config_key<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_Less(v string) *UserSettingQuery {
	return q.w("config_key<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_LessEqual(v string) *UserSettingQuery {
	return q.w("config_key<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_Greater(v string) *UserSettingQuery {
	return q.w("config_key>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigKey_GreaterEqual(v string) *UserSettingQuery {
	return q.w("config_key>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_Equal(v string) *UserSettingQuery {
	return q.w("config_value='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_NotEqual(v string) *UserSettingQuery {
	return q.w("config_value<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_Less(v string) *UserSettingQuery {
	return q.w("config_value<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_LessEqual(v string) *UserSettingQuery {
	return q.w("config_value<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_Greater(v string) *UserSettingQuery {
	return q.w("config_value>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) ConfigValue_GreaterEqual(v string) *UserSettingQuery {
	return q.w("config_value>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_Equal(v time.Time) *UserSettingQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_NotEqual(v time.Time) *UserSettingQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_Less(v time.Time) *UserSettingQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_LessEqual(v time.Time) *UserSettingQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_Greater(v time.Time) *UserSettingQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) CreateTime_GreaterEqual(v time.Time) *UserSettingQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_Equal(v time.Time) *UserSettingQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_NotEqual(v time.Time) *UserSettingQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_Less(v time.Time) *UserSettingQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_LessEqual(v time.Time) *UserSettingQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_Greater(v time.Time) *UserSettingQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserSettingQuery) UpdateTime_GreaterEqual(v time.Time) *UserSettingQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
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

	_, err = stmt.Exec(ctx, e.UserId, e.ConfigKey, e.ConfigValue, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserSettingDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
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

func (dao *UserSettingDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*UserSetting, error) {
	querySql := "SELECT " + USER_SETTING_ALL_FIELDS_STRING + " FROM user_setting " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserSettingDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserSetting, err error) {
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

func (dao *UserSettingDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM user_setting " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *UserSettingDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM user_setting " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *UserSettingDao) GetQuery() *UserSettingQuery {
	return NewUserSettingQuery(dao)
}

const USER_STOCK_EVALUATE_TABLE_NAME = "user_stock_evaluate"

type USER_STOCK_EVALUATE_FIELD string

const USER_STOCK_EVALUATE_FIELD_ID = USER_STOCK_EVALUATE_FIELD("id")
const USER_STOCK_EVALUATE_FIELD_USER_ID = USER_STOCK_EVALUATE_FIELD("user_id")
const USER_STOCK_EVALUATE_FIELD_STOCK_ID = USER_STOCK_EVALUATE_FIELD("stock_id")
const USER_STOCK_EVALUATE_FIELD_TOTAL_SCORE = USER_STOCK_EVALUATE_FIELD("total_score")
const USER_STOCK_EVALUATE_FIELD_INDEX_COUNT = USER_STOCK_EVALUATE_FIELD("index_count")
const USER_STOCK_EVALUATE_FIELD_EVAL_REMARK = USER_STOCK_EVALUATE_FIELD("eval_remark")
const USER_STOCK_EVALUATE_FIELD_CREATE_TIME = USER_STOCK_EVALUATE_FIELD("create_time")
const USER_STOCK_EVALUATE_FIELD_UPDATE_TIME = USER_STOCK_EVALUATE_FIELD("update_time")
const USER_STOCK_EVALUATE_FIELD_EXCHANGE_ID = USER_STOCK_EVALUATE_FIELD("exchange_id")
const USER_STOCK_EVALUATE_FIELD_STOCK_CODE = USER_STOCK_EVALUATE_FIELD("stock_code")
const USER_STOCK_EVALUATE_FIELD_STOCK_NAME_CN = USER_STOCK_EVALUATE_FIELD("stock_name_cn")
const USER_STOCK_EVALUATE_FIELD_LAUNCH_DATE = USER_STOCK_EVALUATE_FIELD("launch_date")
const USER_STOCK_EVALUATE_FIELD_INDUSTRY_NAME = USER_STOCK_EVALUATE_FIELD("industry_name")

const USER_STOCK_EVALUATE_ALL_FIELDS_STRING = "id,user_id,stock_id,total_score,index_count,eval_remark,create_time,update_time,exchange_id,stock_code,stock_name_cn,launch_date,industry_name"

var USER_STOCK_EVALUATE_ALL_FIELDS = []string{
	"id",
	"user_id",
	"stock_id",
	"total_score",
	"index_count",
	"eval_remark",
	"create_time",
	"update_time",
	"exchange_id",
	"stock_code",
	"stock_name_cn",
	"launch_date",
	"industry_name",
}

type UserStockEvaluate struct {
	Id           int64  //size=20
	UserId       string //size=32
	StockId      string //size=32
	TotalScore   float64
	IndexCount   int32  //size=11
	EvalRemark   string //size=256
	CreateTime   time.Time
	UpdateTime   time.Time
	ExchangeId   string //size=32
	StockCode    string //size=32
	StockNameCn  string //size=32
	LaunchDate   time.Time
	IndustryName string //size=32
}

type UserStockEvaluateQuery struct {
	BaseQuery
	dao *UserStockEvaluateDao
}

func NewUserStockEvaluateQuery(dao *UserStockEvaluateDao) *UserStockEvaluateQuery {
	q := &UserStockEvaluateQuery{}
	q.dao = dao

	return q
}

func (q *UserStockEvaluateQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*UserStockEvaluate, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*UserStockEvaluate, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *UserStockEvaluateQuery) ForUpdate() *UserStockEvaluateQuery {
	q.forUpdate = true
	return q
}

func (q *UserStockEvaluateQuery) ForShare() *UserStockEvaluateQuery {
	q.forShare = true
	return q
}

func (q *UserStockEvaluateQuery) GroupBy(fields ...USER_STOCK_EVALUATE_FIELD) *UserStockEvaluateQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *UserStockEvaluateQuery) Limit(startIncluded int64, count int64) *UserStockEvaluateQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *UserStockEvaluateQuery) OrderBy(fieldName USER_STOCK_EVALUATE_FIELD, asc bool) *UserStockEvaluateQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserStockEvaluateQuery) OrderByGroupCount(asc bool) *UserStockEvaluateQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserStockEvaluateQuery) w(format string, a ...interface{}) *UserStockEvaluateQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *UserStockEvaluateQuery) Left() *UserStockEvaluateQuery  { return q.w(" ( ") }
func (q *UserStockEvaluateQuery) Right() *UserStockEvaluateQuery { return q.w(" ) ") }
func (q *UserStockEvaluateQuery) And() *UserStockEvaluateQuery   { return q.w(" AND ") }
func (q *UserStockEvaluateQuery) Or() *UserStockEvaluateQuery    { return q.w(" OR ") }
func (q *UserStockEvaluateQuery) Not() *UserStockEvaluateQuery   { return q.w(" NOT ") }

func (q *UserStockEvaluateQuery) Id_Equal(v int64) *UserStockEvaluateQuery {
	return q.w("id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_NotEqual(v int64) *UserStockEvaluateQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_Less(v int64) *UserStockEvaluateQuery {
	return q.w("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_LessEqual(v int64) *UserStockEvaluateQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_Greater(v int64) *UserStockEvaluateQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) Id_GreaterEqual(v int64) *UserStockEvaluateQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_Equal(v string) *UserStockEvaluateQuery {
	return q.w("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_Less(v string) *UserStockEvaluateQuery {
	return q.w("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_Greater(v string) *UserStockEvaluateQuery {
	return q.w("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UserId_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_Equal(v string) *UserStockEvaluateQuery {
	return q.w("stock_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_Less(v string) *UserStockEvaluateQuery {
	return q.w("stock_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_Greater(v string) *UserStockEvaluateQuery {
	return q.w("stock_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockId_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_Equal(v float64) *UserStockEvaluateQuery {
	return q.w("total_score='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_NotEqual(v float64) *UserStockEvaluateQuery {
	return q.w("total_score<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_Less(v float64) *UserStockEvaluateQuery {
	return q.w("total_score<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_LessEqual(v float64) *UserStockEvaluateQuery {
	return q.w("total_score<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_Greater(v float64) *UserStockEvaluateQuery {
	return q.w("total_score>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) TotalScore_GreaterEqual(v float64) *UserStockEvaluateQuery {
	return q.w("total_score>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndexCount_Equal(v int32) *UserStockEvaluateQuery {
	return q.w("index_count='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndexCount_NotEqual(v int32) *UserStockEvaluateQuery {
	return q.w("index_count<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndexCount_Less(v int32) *UserStockEvaluateQuery {
	return q.w("index_count<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndexCount_LessEqual(v int32) *UserStockEvaluateQuery {
	return q.w("index_count<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndexCount_Greater(v int32) *UserStockEvaluateQuery {
	return q.w("index_count>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndexCount_GreaterEqual(v int32) *UserStockEvaluateQuery {
	return q.w("index_count>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_Equal(v string) *UserStockEvaluateQuery {
	return q.w("eval_remark='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("eval_remark<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_Less(v string) *UserStockEvaluateQuery {
	return q.w("eval_remark<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("eval_remark<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_Greater(v string) *UserStockEvaluateQuery {
	return q.w("eval_remark>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) EvalRemark_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("eval_remark>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_Equal(v time.Time) *UserStockEvaluateQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_NotEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_Less(v time.Time) *UserStockEvaluateQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_LessEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_Greater(v time.Time) *UserStockEvaluateQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) CreateTime_GreaterEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_Equal(v time.Time) *UserStockEvaluateQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_NotEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_Less(v time.Time) *UserStockEvaluateQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_LessEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_Greater(v time.Time) *UserStockEvaluateQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) UpdateTime_GreaterEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_Equal(v string) *UserStockEvaluateQuery {
	return q.w("exchange_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("exchange_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_Less(v string) *UserStockEvaluateQuery {
	return q.w("exchange_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("exchange_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_Greater(v string) *UserStockEvaluateQuery {
	return q.w("exchange_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) ExchangeId_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("exchange_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_Equal(v string) *UserStockEvaluateQuery {
	return q.w("stock_code='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_code<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_Less(v string) *UserStockEvaluateQuery {
	return q.w("stock_code<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_code<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_Greater(v string) *UserStockEvaluateQuery {
	return q.w("stock_code>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockCode_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_code>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockNameCn_Equal(v string) *UserStockEvaluateQuery {
	return q.w("stock_name_cn='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockNameCn_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_name_cn<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockNameCn_Less(v string) *UserStockEvaluateQuery {
	return q.w("stock_name_cn<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockNameCn_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_name_cn<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockNameCn_Greater(v string) *UserStockEvaluateQuery {
	return q.w("stock_name_cn>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) StockNameCn_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("stock_name_cn>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_Equal(v time.Time) *UserStockEvaluateQuery {
	return q.w("launch_date='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_NotEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("launch_date<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_Less(v time.Time) *UserStockEvaluateQuery {
	return q.w("launch_date<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_LessEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("launch_date<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_Greater(v time.Time) *UserStockEvaluateQuery {
	return q.w("launch_date>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) LaunchDate_GreaterEqual(v time.Time) *UserStockEvaluateQuery {
	return q.w("launch_date>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_Equal(v string) *UserStockEvaluateQuery {
	return q.w("industry_name='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_NotEqual(v string) *UserStockEvaluateQuery {
	return q.w("industry_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_Less(v string) *UserStockEvaluateQuery {
	return q.w("industry_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_LessEqual(v string) *UserStockEvaluateQuery {
	return q.w("industry_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_Greater(v string) *UserStockEvaluateQuery {
	return q.w("industry_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockEvaluateQuery) IndustryName_GreaterEqual(v string) *UserStockEvaluateQuery {
	return q.w("industry_name>='" + fmt.Sprint(v) + "'")
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
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_stock_evaluate (user_id,stock_id,total_score,index_count,eval_remark,create_time,update_time,exchange_id,stock_code,stock_name_cn,launch_date,industry_name) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *UserStockEvaluateDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_stock_evaluate SET user_id=?,stock_id=?,total_score=?,index_count=?,eval_remark=?,create_time=?,update_time=?,exchange_id=?,stock_code=?,stock_name_cn=?,launch_date=?,industry_name=? WHERE id=?")
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

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.IndexCount, e.EvalRemark, e.CreateTime, e.UpdateTime, e.ExchangeId, e.StockCode, e.StockNameCn, e.LaunchDate, e.IndustryName)
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

	_, err = stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.IndexCount, e.EvalRemark, e.CreateTime, e.UpdateTime, e.ExchangeId, e.StockCode, e.StockNameCn, e.LaunchDate, e.IndustryName, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserStockEvaluateDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserStockEvaluateDao) scanRow(row *wrap.Row) (*UserStockEvaluate, error) {
	e := &UserStockEvaluate{}
	err := row.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.IndexCount, &e.EvalRemark, &e.CreateTime, &e.UpdateTime, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.LaunchDate, &e.IndustryName)
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
		err = rows.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.IndexCount, &e.EvalRemark, &e.CreateTime, &e.UpdateTime, &e.ExchangeId, &e.StockCode, &e.StockNameCn, &e.LaunchDate, &e.IndustryName)
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

func (dao *UserStockEvaluateDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*UserStockEvaluate, error) {
	querySql := "SELECT " + USER_STOCK_EVALUATE_ALL_FIELDS_STRING + " FROM user_stock_evaluate " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserStockEvaluateDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserStockEvaluate, err error) {
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

func (dao *UserStockEvaluateDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM user_stock_evaluate " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *UserStockEvaluateDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM user_stock_evaluate " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *UserStockEvaluateDao) GetQuery() *UserStockEvaluateQuery {
	return NewUserStockEvaluateQuery(dao)
}

const USER_STOCK_INDEX_TABLE_NAME = "user_stock_index"

type USER_STOCK_INDEX_FIELD string

const USER_STOCK_INDEX_FIELD_ID = USER_STOCK_INDEX_FIELD("id")
const USER_STOCK_INDEX_FIELD_USER_ID = USER_STOCK_INDEX_FIELD("user_id")
const USER_STOCK_INDEX_FIELD_INDEX_NAME = USER_STOCK_INDEX_FIELD("index_name")
const USER_STOCK_INDEX_FIELD_UI_ORDER = USER_STOCK_INDEX_FIELD("ui_order")
const USER_STOCK_INDEX_FIELD_INDEX_DESC = USER_STOCK_INDEX_FIELD("index_desc")
const USER_STOCK_INDEX_FIELD_EVAL_WEIGHT = USER_STOCK_INDEX_FIELD("eval_weight")
const USER_STOCK_INDEX_FIELD_AI_WEIGHT = USER_STOCK_INDEX_FIELD("ai_weight")
const USER_STOCK_INDEX_FIELD_CREATE_TIME = USER_STOCK_INDEX_FIELD("create_time")
const USER_STOCK_INDEX_FIELD_UPDATE_TIME = USER_STOCK_INDEX_FIELD("update_time")

const USER_STOCK_INDEX_ALL_FIELDS_STRING = "id,user_id,index_name,ui_order,index_desc,eval_weight,ai_weight,create_time,update_time"

var USER_STOCK_INDEX_ALL_FIELDS = []string{
	"id",
	"user_id",
	"index_name",
	"ui_order",
	"index_desc",
	"eval_weight",
	"ai_weight",
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
	CreateTime time.Time
	UpdateTime time.Time
}

type UserStockIndexQuery struct {
	BaseQuery
	dao *UserStockIndexDao
}

func NewUserStockIndexQuery(dao *UserStockIndexDao) *UserStockIndexQuery {
	q := &UserStockIndexQuery{}
	q.dao = dao

	return q
}

func (q *UserStockIndexQuery) QueryOne(ctx context.Context, tx *wrap.Tx) (*UserStockIndex, error) {
	return q.dao.QueryOne(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) QueryList(ctx context.Context, tx *wrap.Tx) (list []*UserStockIndex, err error) {
	return q.dao.QueryList(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) QueryCount(ctx context.Context, tx *wrap.Tx) (count int64, err error) {
	return q.dao.QueryCount(ctx, tx, q.buildQueryString())
}

func (q *UserStockIndexQuery) QueryGroupBy(ctx context.Context, tx *wrap.Tx) (rows *wrap.Rows, err error) {
	return q.dao.QueryGroupBy(ctx, tx, q.groupByFields, q.buildQueryString())
}

func (q *UserStockIndexQuery) ForUpdate() *UserStockIndexQuery {
	q.forUpdate = true
	return q
}

func (q *UserStockIndexQuery) ForShare() *UserStockIndexQuery {
	q.forShare = true
	return q
}

func (q *UserStockIndexQuery) GroupBy(fields ...USER_STOCK_INDEX_FIELD) *UserStockIndexQuery {
	q.groupByFields = make([]string, len(fields))
	for i, v := range fields {
		q.groupByFields[i] = string(v)
	}
	return q
}

func (q *UserStockIndexQuery) Limit(startIncluded int64, count int64) *UserStockIndexQuery {
	q.limit = fmt.Sprintf(" limit %d,%d", startIncluded, count)
	return q
}

func (q *UserStockIndexQuery) OrderBy(fieldName USER_STOCK_INDEX_FIELD, asc bool) *UserStockIndexQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += string(fieldName) + " "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserStockIndexQuery) OrderByGroupCount(asc bool) *UserStockIndexQuery {
	if q.order != "" {
		q.order += ","
	}
	q.order += "count(1) "
	if asc {
		q.order += "asc"
	} else {
		q.order += "desc"
	}

	return q
}

func (q *UserStockIndexQuery) w(format string, a ...interface{}) *UserStockIndexQuery {
	q.where += fmt.Sprintf(format, a...)
	return q
}

func (q *UserStockIndexQuery) Left() *UserStockIndexQuery  { return q.w(" ( ") }
func (q *UserStockIndexQuery) Right() *UserStockIndexQuery { return q.w(" ) ") }
func (q *UserStockIndexQuery) And() *UserStockIndexQuery   { return q.w(" AND ") }
func (q *UserStockIndexQuery) Or() *UserStockIndexQuery    { return q.w(" OR ") }
func (q *UserStockIndexQuery) Not() *UserStockIndexQuery   { return q.w(" NOT ") }

func (q *UserStockIndexQuery) Id_Equal(v int64) *UserStockIndexQuery {
	return q.w("id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_NotEqual(v int64) *UserStockIndexQuery {
	return q.w("id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_Less(v int64) *UserStockIndexQuery {
	return q.w("id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_LessEqual(v int64) *UserStockIndexQuery {
	return q.w("id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_Greater(v int64) *UserStockIndexQuery {
	return q.w("id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) Id_GreaterEqual(v int64) *UserStockIndexQuery {
	return q.w("id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_Equal(v string) *UserStockIndexQuery {
	return q.w("user_id='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_NotEqual(v string) *UserStockIndexQuery {
	return q.w("user_id<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_Less(v string) *UserStockIndexQuery {
	return q.w("user_id<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_LessEqual(v string) *UserStockIndexQuery {
	return q.w("user_id<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_Greater(v string) *UserStockIndexQuery {
	return q.w("user_id>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UserId_GreaterEqual(v string) *UserStockIndexQuery {
	return q.w("user_id>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_Equal(v string) *UserStockIndexQuery {
	return q.w("index_name='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_NotEqual(v string) *UserStockIndexQuery {
	return q.w("index_name<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_Less(v string) *UserStockIndexQuery {
	return q.w("index_name<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_LessEqual(v string) *UserStockIndexQuery {
	return q.w("index_name<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_Greater(v string) *UserStockIndexQuery {
	return q.w("index_name>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexName_GreaterEqual(v string) *UserStockIndexQuery {
	return q.w("index_name>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_Equal(v int32) *UserStockIndexQuery {
	return q.w("ui_order='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_NotEqual(v int32) *UserStockIndexQuery {
	return q.w("ui_order<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_Less(v int32) *UserStockIndexQuery {
	return q.w("ui_order<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_LessEqual(v int32) *UserStockIndexQuery {
	return q.w("ui_order<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_Greater(v int32) *UserStockIndexQuery {
	return q.w("ui_order>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UiOrder_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.w("ui_order>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_Equal(v string) *UserStockIndexQuery {
	return q.w("index_desc='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_NotEqual(v string) *UserStockIndexQuery {
	return q.w("index_desc<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_Less(v string) *UserStockIndexQuery {
	return q.w("index_desc<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_LessEqual(v string) *UserStockIndexQuery {
	return q.w("index_desc<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_Greater(v string) *UserStockIndexQuery {
	return q.w("index_desc>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) IndexDesc_GreaterEqual(v string) *UserStockIndexQuery {
	return q.w("index_desc>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_Equal(v int32) *UserStockIndexQuery {
	return q.w("eval_weight='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_NotEqual(v int32) *UserStockIndexQuery {
	return q.w("eval_weight<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_Less(v int32) *UserStockIndexQuery {
	return q.w("eval_weight<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_LessEqual(v int32) *UserStockIndexQuery {
	return q.w("eval_weight<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_Greater(v int32) *UserStockIndexQuery {
	return q.w("eval_weight>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) EvalWeight_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.w("eval_weight>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_Equal(v int32) *UserStockIndexQuery {
	return q.w("ai_weight='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_NotEqual(v int32) *UserStockIndexQuery {
	return q.w("ai_weight<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_Less(v int32) *UserStockIndexQuery {
	return q.w("ai_weight<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_LessEqual(v int32) *UserStockIndexQuery {
	return q.w("ai_weight<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_Greater(v int32) *UserStockIndexQuery {
	return q.w("ai_weight>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) AiWeight_GreaterEqual(v int32) *UserStockIndexQuery {
	return q.w("ai_weight>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_Equal(v time.Time) *UserStockIndexQuery {
	return q.w("create_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_NotEqual(v time.Time) *UserStockIndexQuery {
	return q.w("create_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_Less(v time.Time) *UserStockIndexQuery {
	return q.w("create_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_LessEqual(v time.Time) *UserStockIndexQuery {
	return q.w("create_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_Greater(v time.Time) *UserStockIndexQuery {
	return q.w("create_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) CreateTime_GreaterEqual(v time.Time) *UserStockIndexQuery {
	return q.w("create_time>='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_Equal(v time.Time) *UserStockIndexQuery {
	return q.w("update_time='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_NotEqual(v time.Time) *UserStockIndexQuery {
	return q.w("update_time<>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_Less(v time.Time) *UserStockIndexQuery {
	return q.w("update_time<'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_LessEqual(v time.Time) *UserStockIndexQuery {
	return q.w("update_time<='" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_Greater(v time.Time) *UserStockIndexQuery {
	return q.w("update_time>'" + fmt.Sprint(v) + "'")
}
func (q *UserStockIndexQuery) UpdateTime_GreaterEqual(v time.Time) *UserStockIndexQuery {
	return q.w("update_time>='" + fmt.Sprint(v) + "'")
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
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO user_stock_index (user_id,index_name,ui_order,index_desc,eval_weight,ai_weight,create_time,update_time) VALUES (?,?,?,?,?,?,?,?)")
	return err
}

func (dao *UserStockIndexDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE user_stock_index SET user_id=?,index_name=?,ui_order=?,index_desc=?,eval_weight=?,ai_weight=?,create_time=?,update_time=? WHERE id=?")
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

	result, err := stmt.Exec(ctx, e.UserId, e.IndexName, e.UiOrder, e.IndexDesc, e.EvalWeight, e.AiWeight, e.CreateTime, e.UpdateTime)
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

	_, err = stmt.Exec(ctx, e.UserId, e.IndexName, e.UiOrder, e.IndexDesc, e.EvalWeight, e.AiWeight, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserStockIndexDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	_, err = stmt.Exec(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserStockIndexDao) scanRow(row *wrap.Row) (*UserStockIndex, error) {
	e := &UserStockIndex{}
	err := row.Scan(&e.Id, &e.UserId, &e.IndexName, &e.UiOrder, &e.IndexDesc, &e.EvalWeight, &e.AiWeight, &e.CreateTime, &e.UpdateTime)
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
		err = rows.Scan(&e.Id, &e.UserId, &e.IndexName, &e.UiOrder, &e.IndexDesc, &e.EvalWeight, &e.AiWeight, &e.CreateTime, &e.UpdateTime)
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

func (dao *UserStockIndexDao) QueryOne(ctx context.Context, tx *wrap.Tx, query string) (*UserStockIndex, error) {
	querySql := "SELECT " + USER_STOCK_INDEX_ALL_FIELDS_STRING + " FROM user_stock_index " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	return dao.scanRow(row)
}

func (dao *UserStockIndexDao) QueryList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserStockIndex, err error) {
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

func (dao *UserStockIndexDao) QueryCount(ctx context.Context, tx *wrap.Tx, query string) (count int64, err error) {
	querySql := "SELECT COUNT(1) FROM user_stock_index " + query
	var row *wrap.Row
	if tx == nil {
		row = dao.db.QueryRow(ctx, querySql)
	} else {
		row = tx.QueryRow(ctx, querySql)
	}
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return 0, err
	}

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (dao *UserStockIndexDao) QueryGroupBy(ctx context.Context, tx *wrap.Tx, groupByFields []string, query string) (rows *wrap.Rows, err error) {
	querySql := "SELECT " + strings.Join(groupByFields, ",") + ",count(1) FROM user_stock_index " + query
	if tx == nil {
		return dao.db.Query(ctx, querySql)
	} else {
		return tx.Query(ctx, querySql)
	}
}

func (dao *UserStockIndexDao) GetQuery() *UserStockIndexQuery {
	return NewUserStockIndexQuery(dao)
}

type DB struct {
	wrap.DB
	Stock             *StockDao
	StockIndexAdvice  *StockIndexAdviceDao
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

	d.StockIndexAdvice, err = NewStockIndexAdviceDao(d)
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
