package fin_stock_assistant

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/wrap"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

const INDEX_EVALUATE_TABLE_NAME = "index_evaluate"

const INDEX_EVALUATE_FIELD_ID = "id"
const INDEX_EVALUATE_FIELD_USER_ID = "user_id"
const INDEX_EVALUATE_FIELD_STOCK_ID = "stock_id"
const INDEX_EVALUATE_FIELD_INDEX_NAME = "index_name"
const INDEX_EVALUATE_FIELD_EVAL_STARS = "eval_stars"
const INDEX_EVALUATE_FIELD_EVAL_REMARK = "eval_remark"
const INDEX_EVALUATE_FIELD_CREATE_TIME = "create_time"
const INDEX_EVALUATE_FIELD_UPDATE_TIME = "update_time"

const INDEX_EVALUATE_ALL_FIELDS_STRING = "id,user_id,stock_id,index_name,eval_stars,eval_remark,create_time,update_time"

var INDEX_EVALUATE_ALL_FIELDS = []string{
	"id",
	"user_id",
	"stock_id",
	"index_name",
	"eval_stars",
	"eval_remark",
	"create_time",
	"update_time",
}

type IndexEvaluate struct {
	Id         int64  //size=20
	UserId     string //size=32
	StockId    string //size=32
	IndexName  string //size=32
	EvalStars  int32  //size=10
	EvalRemark string //size=256
	CreateTime time.Time
	UpdateTime time.Time
}

type IndexEvaluateQuery struct {
	dao         *IndexEvaluateDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewIndexEvaluateQuery(dao *IndexEvaluateDao) *IndexEvaluateQuery {
	q := &IndexEvaluateQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *IndexEvaluateQuery) buildQueryString() string {
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

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	return buf.String()
}

func (q *IndexEvaluateQuery) Select(ctx context.Context) (*IndexEvaluate, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *IndexEvaluateQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*IndexEvaluate, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *IndexEvaluateQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*IndexEvaluate, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *IndexEvaluateQuery) SelectList(ctx context.Context) (list []*IndexEvaluate, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *IndexEvaluateQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*IndexEvaluate, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *IndexEvaluateQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*IndexEvaluate, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *IndexEvaluateQuery) Left() *IndexEvaluateQuery {
	q.whereBuffer.WriteString(" ( ")
	return q
}

func (q *IndexEvaluateQuery) Right() *IndexEvaluateQuery {
	q.whereBuffer.WriteString(" ) ")
	return q
}

func (q *IndexEvaluateQuery) And() *IndexEvaluateQuery {
	q.whereBuffer.WriteString(" AND ")
	return q
}

func (q *IndexEvaluateQuery) Or() *IndexEvaluateQuery {
	q.whereBuffer.WriteString(" OR ")
	return q
}

func (q *IndexEvaluateQuery) Not() *IndexEvaluateQuery {
	q.whereBuffer.WriteString(" NOT ")
	return q
}

func (q *IndexEvaluateQuery) Limit(startIncluded int64, count int64) *IndexEvaluateQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *IndexEvaluateQuery) Sort(fieldName string, asc bool) *IndexEvaluateQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *IndexEvaluateQuery) Id_Equal(v int64) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) Id_NotEqual(v int64) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) Id_Less(v int64) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) Id_LessEqual(v int64) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) Id_Greater(v int64) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) Id_GreaterEqual(v int64) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_Equal(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_NotEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("user_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_Less(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("user_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_LessEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("user_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_Greater(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("user_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_GreaterEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_Equal(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("stock_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_NotEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("stock_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_Less(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("stock_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_LessEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("stock_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_Greater(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("stock_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_GreaterEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("stock_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexName_Equal(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("index_name='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexName_NotEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("index_name<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexName_Less(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("index_name<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexName_LessEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("index_name<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexName_Greater(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("index_name>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexName_GreaterEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("index_name='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_Equal(v int32) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_stars='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_NotEqual(v int32) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_stars<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_Less(v int32) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_stars<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_LessEqual(v int32) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_stars<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_Greater(v int32) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_stars>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_GreaterEqual(v int32) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_stars='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_Equal(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_NotEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_Less(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_LessEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_Greater(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_GreaterEqual(v string) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_Equal(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_NotEqual(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("create_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_Less(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("create_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_LessEqual(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("create_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_Greater(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("create_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_GreaterEqual(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_Equal(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_NotEqual(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("update_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_Less(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("update_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_LessEqual(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("update_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_Greater(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("update_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_GreaterEqual(v time.Time) *IndexEvaluateQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

type IndexEvaluateDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewIndexEvaluateDao(db *DB) (t *IndexEvaluateDao, err error) {
	t = &IndexEvaluateDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *IndexEvaluateDao) init() (err error) {
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
func (dao *IndexEvaluateDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO index_evaluate (user_id,stock_id,index_name,eval_stars,eval_remark,create_time,update_time) VALUES (?,?,?,?,?,?,?)")
	return err
}

func (dao *IndexEvaluateDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE index_evaluate SET user_id=?,stock_id=?,index_name=?,eval_stars=?,eval_remark=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *IndexEvaluateDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM index_evaluate WHERE id=?")
	return err
}

func (dao *IndexEvaluateDao) Insert(ctx context.Context, tx *wrap.Tx, e *IndexEvaluate) (id int64, err error) {
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

func (dao *IndexEvaluateDao) Update(ctx context.Context, tx *wrap.Tx, e *IndexEvaluate) (err error) {
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

func (dao *IndexEvaluateDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
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

func (dao *IndexEvaluateDao) scanRow(row *wrap.Row) (*IndexEvaluate, error) {
	e := &IndexEvaluate{}
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

func (dao *IndexEvaluateDao) scanRows(rows *wrap.Rows) (list []*IndexEvaluate, err error) {
	list = make([]*IndexEvaluate, 0)
	for rows.Next() {
		e := IndexEvaluate{}
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

func (dao *IndexEvaluateDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*IndexEvaluate, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate "+query)
	return dao.scanRow(row)
}

func (dao *IndexEvaluateDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*IndexEvaluate, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *IndexEvaluateDao) GetQuery() *IndexEvaluateQuery {
	return NewIndexEvaluateQuery(dao)
}

const STOCK_EVALUATE_TABLE_NAME = "stock_evaluate"

const STOCK_EVALUATE_FIELD_ID = "id"
const STOCK_EVALUATE_FIELD_USER_ID = "user_id"
const STOCK_EVALUATE_FIELD_STOCK_ID = "stock_id"
const STOCK_EVALUATE_FIELD_TOTAL_SCORE = "total_score"
const STOCK_EVALUATE_FIELD_EVAL_REMARK = "eval_remark"
const STOCK_EVALUATE_FIELD_CREATE_TIME = "create_time"
const STOCK_EVALUATE_FIELD_UPDATE_TIME = "update_time"

const STOCK_EVALUATE_ALL_FIELDS_STRING = "id,user_id,stock_id,total_score,eval_remark,create_time,update_time"

var STOCK_EVALUATE_ALL_FIELDS = []string{
	"id",
	"user_id",
	"stock_id",
	"total_score",
	"eval_remark",
	"create_time",
	"update_time",
}

type StockEvaluate struct {
	Id         int64  //size=20
	UserId     string //size=32
	StockId    string //size=32
	TotalScore float64
	EvalRemark string //size=256
	CreateTime time.Time
	UpdateTime time.Time
}

type StockEvaluateQuery struct {
	dao         *StockEvaluateDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewStockEvaluateQuery(dao *StockEvaluateDao) *StockEvaluateQuery {
	q := &StockEvaluateQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockEvaluateQuery) buildQueryString() string {
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

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	return buf.String()
}

func (q *StockEvaluateQuery) Select(ctx context.Context) (*StockEvaluate, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *StockEvaluateQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*StockEvaluate, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockEvaluateQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*StockEvaluate, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockEvaluateQuery) SelectList(ctx context.Context) (list []*StockEvaluate, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *StockEvaluateQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*StockEvaluate, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockEvaluateQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*StockEvaluate, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockEvaluateQuery) Left() *StockEvaluateQuery {
	q.whereBuffer.WriteString(" ( ")
	return q
}

func (q *StockEvaluateQuery) Right() *StockEvaluateQuery {
	q.whereBuffer.WriteString(" ) ")
	return q
}

func (q *StockEvaluateQuery) And() *StockEvaluateQuery {
	q.whereBuffer.WriteString(" AND ")
	return q
}

func (q *StockEvaluateQuery) Or() *StockEvaluateQuery {
	q.whereBuffer.WriteString(" OR ")
	return q
}

func (q *StockEvaluateQuery) Not() *StockEvaluateQuery {
	q.whereBuffer.WriteString(" NOT ")
	return q
}

func (q *StockEvaluateQuery) Limit(startIncluded int64, count int64) *StockEvaluateQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockEvaluateQuery) Sort(fieldName string, asc bool) *StockEvaluateQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *StockEvaluateQuery) Id_Equal(v int64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) Id_NotEqual(v int64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) Id_Less(v int64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) Id_LessEqual(v int64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) Id_Greater(v int64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) Id_GreaterEqual(v int64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_Equal(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_NotEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("user_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_Less(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("user_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_LessEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("user_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_Greater(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("user_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_GreaterEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_Equal(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("stock_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_NotEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("stock_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_Less(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("stock_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_LessEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("stock_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_Greater(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("stock_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_GreaterEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("stock_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_Equal(v float64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("total_score='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_NotEqual(v float64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("total_score<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_Less(v float64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("total_score<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_LessEqual(v float64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("total_score<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_Greater(v float64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("total_score>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_GreaterEqual(v float64) *StockEvaluateQuery {
	q.whereBuffer.WriteString("total_score='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_Equal(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_NotEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_Less(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_LessEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_Greater(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_GreaterEqual(v string) *StockEvaluateQuery {
	q.whereBuffer.WriteString("eval_remark='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_Equal(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_NotEqual(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("create_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_Less(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("create_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_LessEqual(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("create_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_Greater(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("create_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_GreaterEqual(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_Equal(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_NotEqual(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("update_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_Less(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("update_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_LessEqual(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("update_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_Greater(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("update_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_GreaterEqual(v time.Time) *StockEvaluateQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

type StockEvaluateDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewStockEvaluateDao(db *DB) (t *StockEvaluateDao, err error) {
	t = &StockEvaluateDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *StockEvaluateDao) init() (err error) {
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
func (dao *StockEvaluateDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock_evaluate (user_id,stock_id,total_score,eval_remark,create_time,update_time) VALUES (?,?,?,?,?,?)")
	return err
}

func (dao *StockEvaluateDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock_evaluate SET user_id=?,stock_id=?,total_score=?,eval_remark=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockEvaluateDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock_evaluate WHERE id=?")
	return err
}

func (dao *StockEvaluateDao) Insert(ctx context.Context, tx *wrap.Tx, e *StockEvaluate) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.EvalRemark, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockEvaluateDao) Update(ctx context.Context, tx *wrap.Tx, e *StockEvaluate) (err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.EvalRemark, e.CreateTime, e.UpdateTime, e.Id)
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

func (dao *StockEvaluateDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
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

func (dao *StockEvaluateDao) scanRow(row *wrap.Row) (*StockEvaluate, error) {
	e := &StockEvaluate{}
	err := row.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == wrap.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockEvaluateDao) scanRows(rows *wrap.Rows) (list []*StockEvaluate, err error) {
	list = make([]*StockEvaluate, 0)
	for rows.Next() {
		e := StockEvaluate{}
		err = rows.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
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

func (dao *StockEvaluateDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*StockEvaluate, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate "+query)
	return dao.scanRow(row)
}

func (dao *StockEvaluateDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*StockEvaluate, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *StockEvaluateDao) GetQuery() *StockEvaluateQuery {
	return NewStockEvaluateQuery(dao)
}

const STOCK_INDEX_TABLE_NAME = "stock_index"

const STOCK_INDEX_FIELD_ID = "id"
const STOCK_INDEX_FIELD_USER_ID = "user_id"
const STOCK_INDEX_FIELD_INDEX_NAME = "index_name"
const STOCK_INDEX_FIELD_UI_ORDER = "ui_order"
const STOCK_INDEX_FIELD_INDEX_DESC = "index_desc"
const STOCK_INDEX_FIELD_EVAL_WEIGHT = "eval_weight"
const STOCK_INDEX_FIELD_AI_WEIGHT = "ai_weight"
const STOCK_INDEX_FIELD_NI_WEIGHT = "ni_weight"
const STOCK_INDEX_FIELD_CREATE_TIME = "create_time"
const STOCK_INDEX_FIELD_UPDATE_TIME = "update_time"

const STOCK_INDEX_ALL_FIELDS_STRING = "id,user_id,index_name,ui_order,index_desc,eval_weight,ai_weight,ni_weight,create_time,update_time"

var STOCK_INDEX_ALL_FIELDS = []string{
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

type StockIndex struct {
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

type StockIndexQuery struct {
	dao         *StockIndexDao
	forUpdate   bool
	forShare    bool
	whereBuffer *bytes.Buffer
	limitBuffer *bytes.Buffer
	orderBuffer *bytes.Buffer
}

func NewStockIndexQuery(dao *StockIndexDao) *StockIndexQuery {
	q := &StockIndexQuery{}
	q.dao = dao
	q.whereBuffer = bytes.NewBufferString("")
	q.limitBuffer = bytes.NewBufferString("")
	q.orderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockIndexQuery) buildQueryString() string {
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

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
	}

	return buf.String()
}

func (q *StockIndexQuery) Select(ctx context.Context) (*StockIndex, error) {
	return q.dao.doSelect(ctx, nil, q.buildQueryString())
}

func (q *StockIndexQuery) SelectForUpdate(ctx context.Context, tx *wrap.Tx) (*StockIndex, error) {
	q.forUpdate = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockIndexQuery) SelectForShare(ctx context.Context, tx *wrap.Tx) (*StockIndex, error) {
	q.forShare = true
	return q.dao.doSelect(ctx, tx, q.buildQueryString())
}

func (q *StockIndexQuery) SelectList(ctx context.Context) (list []*StockIndex, err error) {
	return q.dao.doSelectList(ctx, nil, q.buildQueryString())
}

func (q *StockIndexQuery) SelectListForUpdate(ctx context.Context, tx *wrap.Tx) (list []*StockIndex, err error) {
	q.forUpdate = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockIndexQuery) SelectListForShare(ctx context.Context, tx *wrap.Tx) (list []*StockIndex, err error) {
	q.forShare = true
	return q.dao.doSelectList(ctx, tx, q.buildQueryString())
}

func (q *StockIndexQuery) Left() *StockIndexQuery {
	q.whereBuffer.WriteString(" ( ")
	return q
}

func (q *StockIndexQuery) Right() *StockIndexQuery {
	q.whereBuffer.WriteString(" ) ")
	return q
}

func (q *StockIndexQuery) And() *StockIndexQuery {
	q.whereBuffer.WriteString(" AND ")
	return q
}

func (q *StockIndexQuery) Or() *StockIndexQuery {
	q.whereBuffer.WriteString(" OR ")
	return q
}

func (q *StockIndexQuery) Not() *StockIndexQuery {
	q.whereBuffer.WriteString(" NOT ")
	return q
}

func (q *StockIndexQuery) Limit(startIncluded int64, count int64) *StockIndexQuery {
	q.limitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockIndexQuery) Sort(fieldName string, asc bool) *StockIndexQuery {
	if asc {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.orderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *StockIndexQuery) Id_Equal(v int64) *StockIndexQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) Id_NotEqual(v int64) *StockIndexQuery {
	q.whereBuffer.WriteString("id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) Id_Less(v int64) *StockIndexQuery {
	q.whereBuffer.WriteString("id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) Id_LessEqual(v int64) *StockIndexQuery {
	q.whereBuffer.WriteString("id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) Id_Greater(v int64) *StockIndexQuery {
	q.whereBuffer.WriteString("id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) Id_GreaterEqual(v int64) *StockIndexQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_Equal(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_NotEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("user_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_Less(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("user_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_LessEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("user_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_Greater(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("user_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_GreaterEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_Equal(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_name='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_NotEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_name<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_Less(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_name<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_LessEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_name<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_Greater(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_name>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_GreaterEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_name='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UiOrder_Equal(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ui_order='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UiOrder_NotEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ui_order<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UiOrder_Less(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ui_order<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UiOrder_LessEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ui_order<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UiOrder_Greater(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ui_order>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UiOrder_GreaterEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ui_order='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_Equal(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_desc='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_NotEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_desc<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_Less(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_desc<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_LessEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_desc<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_Greater(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_desc>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_GreaterEqual(v string) *StockIndexQuery {
	q.whereBuffer.WriteString("index_desc='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_Equal(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("eval_weight='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_NotEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("eval_weight<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_Less(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("eval_weight<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_LessEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("eval_weight<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_Greater(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("eval_weight>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_GreaterEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("eval_weight='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_Equal(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ai_weight='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_NotEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ai_weight<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_Less(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ai_weight<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_LessEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ai_weight<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_Greater(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ai_weight>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_GreaterEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ai_weight='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_Equal(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ni_weight='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_NotEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ni_weight<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_Less(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ni_weight<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_LessEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ni_weight<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_Greater(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ni_weight>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_GreaterEqual(v int32) *StockIndexQuery {
	q.whereBuffer.WriteString("ni_weight='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_Equal(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_NotEqual(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("create_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_Less(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("create_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_LessEqual(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("create_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_Greater(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("create_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_GreaterEqual(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_Equal(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_NotEqual(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("update_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_Less(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("update_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_LessEqual(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("update_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_Greater(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("update_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_GreaterEqual(v time.Time) *StockIndexQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

type StockIndexDao struct {
	logger     *zap.Logger
	db         *DB
	insertStmt *wrap.Stmt
	updateStmt *wrap.Stmt
	deleteStmt *wrap.Stmt
}

func NewStockIndexDao(db *DB) (t *StockIndexDao, err error) {
	t = &StockIndexDao{}
	t.logger = log.TypedLogger(t)
	t.db = db
	err = t.init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (dao *StockIndexDao) init() (err error) {
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
func (dao *StockIndexDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock_index (user_id,index_name,ui_order,index_desc,eval_weight,ai_weight,ni_weight,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *StockIndexDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock_index SET user_id=?,index_name=?,ui_order=?,index_desc=?,eval_weight=?,ai_weight=?,ni_weight=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockIndexDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock_index WHERE id=?")
	return err
}

func (dao *StockIndexDao) Insert(ctx context.Context, tx *wrap.Tx, e *StockIndex) (id int64, err error) {
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

func (dao *StockIndexDao) Update(ctx context.Context, tx *wrap.Tx, e *StockIndex) (err error) {
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

func (dao *StockIndexDao) Delete(ctx context.Context, tx *wrap.Tx, id int64) (err error) {
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

func (dao *StockIndexDao) scanRow(row *wrap.Row) (*StockIndex, error) {
	e := &StockIndex{}
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

func (dao *StockIndexDao) scanRows(rows *wrap.Rows) (list []*StockIndex, err error) {
	list = make([]*StockIndex, 0)
	for rows.Next() {
		e := StockIndex{}
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

func (dao *StockIndexDao) doSelect(ctx context.Context, tx *wrap.Tx, query string) (*StockIndex, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index "+query)
	return dao.scanRow(row)
}

func (dao *StockIndexDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*StockIndex, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *StockIndexDao) GetQuery() *StockIndexQuery {
	return NewStockIndexQuery(dao)
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

	limitSql := q.limitBuffer.String()
	if limitSql != "" {
		buf.WriteString(limitSql)
	}

	orderSql := q.orderBuffer.String()
	if orderSql != "" {
		buf.WriteString(orderSql)
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

func (q *UserSettingQuery) Left() *UserSettingQuery {
	q.whereBuffer.WriteString(" ( ")
	return q
}

func (q *UserSettingQuery) Right() *UserSettingQuery {
	q.whereBuffer.WriteString(" ) ")
	return q
}

func (q *UserSettingQuery) And() *UserSettingQuery {
	q.whereBuffer.WriteString(" AND ")
	return q
}

func (q *UserSettingQuery) Or() *UserSettingQuery {
	q.whereBuffer.WriteString(" OR ")
	return q
}

func (q *UserSettingQuery) Not() *UserSettingQuery {
	q.whereBuffer.WriteString(" NOT ")
	return q
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
func (q *UserSettingQuery) Id_Equal(v int64) *UserSettingQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) Id_NotEqual(v int64) *UserSettingQuery {
	q.whereBuffer.WriteString("id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) Id_Less(v int64) *UserSettingQuery {
	q.whereBuffer.WriteString("id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) Id_LessEqual(v int64) *UserSettingQuery {
	q.whereBuffer.WriteString("id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) Id_Greater(v int64) *UserSettingQuery {
	q.whereBuffer.WriteString("id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) Id_GreaterEqual(v int64) *UserSettingQuery {
	q.whereBuffer.WriteString("id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_Equal(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_NotEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("user_id<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_Less(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("user_id<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_LessEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("user_id<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_Greater(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("user_id>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_GreaterEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("user_id='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_Equal(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_key='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_NotEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_key<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_Less(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_key<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_LessEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_key<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_Greater(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_key>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_GreaterEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_key='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_Equal(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_value='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_NotEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_value<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_Less(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_value<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_LessEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_value<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_Greater(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_value>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_GreaterEqual(v string) *UserSettingQuery {
	q.whereBuffer.WriteString("config_value='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_Equal(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_NotEqual(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("create_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_Less(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("create_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_LessEqual(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("create_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_Greater(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("create_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_GreaterEqual(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("create_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_Equal(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_NotEqual(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("update_time<>'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_Less(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("update_time<'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_LessEqual(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("update_time<='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_Greater(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("update_time>='" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_GreaterEqual(v time.Time) *UserSettingQuery {
	q.whereBuffer.WriteString("update_time='" + fmt.Sprint(v) + "'")
	return q
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
	row := dao.db.QueryRow(ctx, "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting "+query)
	return dao.scanRow(row)
}

func (dao *UserSettingDao) doSelectList(ctx context.Context, tx *wrap.Tx, query string) (list []*UserSetting, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.scanRows(rows)
}

func (dao *UserSettingDao) GetQuery() *UserSettingQuery {
	return NewUserSettingQuery(dao)
}

type DB struct {
	wrap.DB
	IndexEvaluate *IndexEvaluateDao
	StockEvaluate *StockEvaluateDao
	StockIndex    *StockIndexDao
	UserSetting   *UserSettingDao
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

	d.IndexEvaluate, err = NewIndexEvaluateDao(d)
	if err != nil {
		return nil, err
	}

	d.StockEvaluate, err = NewStockEvaluateDao(d)
	if err != nil {
		return nil, err
	}

	d.StockIndex, err = NewStockIndexDao(d)
	if err != nil {
		return nil, err
	}

	d.UserSetting, err = NewUserSettingDao(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
