package fin_stock_assistant

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NeuronEvolution/log"
	"github.com/NeuronEvolution/sql/runtime"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

const INDEX_EVALUATE_TABLE_NAME = "index_evaluate"

const INDEX_EVALUATE_FIELD_ID = "id"
const INDEX_EVALUATE_FIELD_USER_ID = "user_id"
const INDEX_EVALUATE_FIELD_STOCK_ID = "stock_id"
const INDEX_EVALUATE_FIELD_INDEX_ID = "index_id"
const INDEX_EVALUATE_FIELD_EVAL_STARS = "eval_stars"
const INDEX_EVALUATE_FIELD_EVAL_REMARK = "eval_remark"
const INDEX_EVALUATE_FIELD_CREATE_TIME = "create_time"
const INDEX_EVALUATE_FIELD_UPDATE_TIME = "update_time"

const INDEX_EVALUATE_ALL_FIELDS_STRING = "id,user_id,stock_id,index_id,eval_stars,eval_remark,create_time,update_time"

var INDEX_EVALUATE_ALL_FIELDS = []string{
	"id",
	"user_id",
	"stock_id",
	"index_id",
	"eval_stars",
	"eval_remark",
	"create_time",
	"update_time",
}

type IndexEvaluate struct {
	Id         int64  //size=20
	UserId     string //size=32
	StockId    string //size=32
	IndexId    string //size=32
	EvalStars  int32  //size=10
	EvalRemark string //size=256
	CreateTime time.Time
	UpdateTime time.Time
}

type IndexEvaluateQuery struct {
	dao *IndexEvaluateDao
	runtime.Query
}

func NewIndexEvaluateQuery(dao *IndexEvaluateDao) *IndexEvaluateQuery {
	q := &IndexEvaluateQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *IndexEvaluateQuery) Select(ctx context.Context) (*IndexEvaluate, error) {
	return q.dao.Select(ctx, nil, q.BuildQueryString())
}

func (q *IndexEvaluateQuery) SelectForUpdate(ctx context.Context, tx *runtime.Tx) (*IndexEvaluate, error) {
	q.ForUpdate = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *IndexEvaluateQuery) SelectForShare(ctx context.Context, tx *runtime.Tx) (*IndexEvaluate, error) {
	q.ForShare = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *IndexEvaluateQuery) SelectList(ctx context.Context) (list []*IndexEvaluate, err error) {
	return q.dao.SelectList(ctx, nil, q.BuildQueryString())
}

func (q *IndexEvaluateQuery) SelectListForUpdate(ctx context.Context, tx *runtime.Tx) (list []*IndexEvaluate, err error) {
	q.ForUpdate = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *IndexEvaluateQuery) SelectListForShare(ctx context.Context, tx *runtime.Tx) (list []*IndexEvaluate, err error) {
	q.ForShare = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *IndexEvaluateQuery) Left() *IndexEvaluateQuery {
	q.WhereBuffer.WriteString(" ( ")
	return q
}

func (q *IndexEvaluateQuery) Right() *IndexEvaluateQuery {
	q.WhereBuffer.WriteString(" ) ")
	return q
}

func (q *IndexEvaluateQuery) And() *IndexEvaluateQuery {
	q.WhereBuffer.WriteString(" AND ")
	return q
}

func (q *IndexEvaluateQuery) Or() *IndexEvaluateQuery {
	q.WhereBuffer.WriteString(" OR ")
	return q
}

func (q *IndexEvaluateQuery) Not() *IndexEvaluateQuery {
	q.WhereBuffer.WriteString(" NOT ")
	return q
}

func (q *IndexEvaluateQuery) Limit(startIncluded int64, count int64) *IndexEvaluateQuery {
	q.LimitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *IndexEvaluateQuery) Sort(fieldName string, asc bool) *IndexEvaluateQuery {
	if asc {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *IndexEvaluateQuery) Id_Column(r runtime.Relation, v int64) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UserId_Column(r runtime.Relation, v string) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("user_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) StockId_Column(r runtime.Relation, v string) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("stock_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) IndexId_Column(r runtime.Relation, v string) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("index_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalStars_Column(r runtime.Relation, v int32) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("eval_stars" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) EvalRemark_Column(r runtime.Relation, v string) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("eval_remark" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) CreateTime_Column(r runtime.Relation, v time.Time) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *IndexEvaluateQuery) UpdateTime_Column(r runtime.Relation, v time.Time) *IndexEvaluateQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type IndexEvaluateDao struct {
	logger                                 *zap.Logger
	db                                     *DB
	insertStmt                             *runtime.Stmt
	updateStmt                             *runtime.Stmt
	deleteStmt                             *runtime.Stmt
	selectStmtAll                          *runtime.Stmt
	selectStmtById                         *runtime.Stmt
	selectStmtByUpdateTime                 *runtime.Stmt
	selectStmtByUserId                     *runtime.Stmt
	selectStmtByUserIdAndStockId           *runtime.Stmt
	selectStmtByUserIdAndStockIdAndIndexId *runtime.Stmt
}

func NewIndexEvaluateDao(db *DB) (t *IndexEvaluateDao) {
	t = &IndexEvaluateDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *IndexEvaluateDao) Init() (err error) {
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

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUpdateTime()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserIdAndStockId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserIdAndStockIdAndIndexId()
	if err != nil {
		return err
	}

	return nil
}
func (dao *IndexEvaluateDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO index_evaluate (user_id,stock_id,index_id,eval_stars,eval_remark,create_time,update_time) VALUES (?,?,?,?,?,?,?)")
	return err
}

func (dao *IndexEvaluateDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE index_evaluate SET user_id=?,stock_id=?,index_id=?,eval_stars=?,eval_remark=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *IndexEvaluateDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM index_evaluate WHERE id=?")
	return err
}

func (dao *IndexEvaluateDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate")
	return err
}

func (dao *IndexEvaluateDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate WHERE id=?")
	return err
}

func (dao *IndexEvaluateDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare(context.Background(), "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate WHERE update_time=?")
	return err
}

func (dao *IndexEvaluateDao) prepareSelectStmtByUserId() (err error) {
	dao.selectStmtByUserId, err = dao.db.Prepare(context.Background(), "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate WHERE user_id=?")
	return err
}

func (dao *IndexEvaluateDao) prepareSelectStmtByUserIdAndStockId() (err error) {
	dao.selectStmtByUserIdAndStockId, err = dao.db.Prepare(context.Background(), "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate WHERE user_id=? AND stock_id=?")
	return err
}

func (dao *IndexEvaluateDao) prepareSelectStmtByUserIdAndStockIdAndIndexId() (err error) {
	dao.selectStmtByUserIdAndStockIdAndIndexId, err = dao.db.Prepare(context.Background(), "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate WHERE user_id=? AND stock_id=? AND index_id=?")
	return err
}

func (dao *IndexEvaluateDao) Insert(ctx context.Context, tx *runtime.Tx, e *IndexEvaluate) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.IndexId, e.EvalStars, e.EvalRemark, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *IndexEvaluateDao) Update(ctx context.Context, tx *runtime.Tx, e *IndexEvaluate) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.IndexId, e.EvalStars, e.EvalRemark, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *IndexEvaluateDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *IndexEvaluateDao) ScanRow(row *runtime.Row) (*IndexEvaluate, error) {
	e := &IndexEvaluate{}
	err := row.Scan(&e.Id, &e.UserId, &e.StockId, &e.IndexId, &e.EvalStars, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *IndexEvaluateDao) ScanRows(rows *runtime.Rows) (list []*IndexEvaluate, err error) {
	list = make([]*IndexEvaluate, 0)
	for rows.Next() {
		e := IndexEvaluate{}
		err = rows.Scan(&e.Id, &e.UserId, &e.StockId, &e.IndexId, &e.EvalStars, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
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

func (dao *IndexEvaluateDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*IndexEvaluate, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *IndexEvaluateDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*IndexEvaluate, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate "+query)
	return dao.ScanRow(row)
}

func (dao *IndexEvaluateDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*IndexEvaluate, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+INDEX_EVALUATE_ALL_FIELDS_STRING+" FROM index_evaluate "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *IndexEvaluateDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*IndexEvaluate, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *IndexEvaluateDao) SelectByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (*IndexEvaluate, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UpdateTime))
}

func (dao *IndexEvaluateDao) SelectListByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (list []*IndexEvaluate, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UpdateTime)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *IndexEvaluateDao) SelectListByUserId(ctx context.Context, tx *runtime.Tx, UserId string) (list []*IndexEvaluate, err error) {
	stmt := dao.selectStmtByUserId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *IndexEvaluateDao) SelectListByUserIdAndStockId(ctx context.Context, tx *runtime.Tx, UserId string, StockId string) (list []*IndexEvaluate, err error) {
	stmt := dao.selectStmtByUserIdAndStockId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId, StockId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *IndexEvaluateDao) SelectListByUserIdAndStockIdAndIndexId(ctx context.Context, tx *runtime.Tx, UserId string, StockId string, IndexId string) (list []*IndexEvaluate, err error) {
	stmt := dao.selectStmtByUserIdAndStockIdAndIndexId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId, StockId, IndexId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
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
	dao *StockEvaluateDao
	runtime.Query
}

func NewStockEvaluateQuery(dao *StockEvaluateDao) *StockEvaluateQuery {
	q := &StockEvaluateQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockEvaluateQuery) Select(ctx context.Context) (*StockEvaluate, error) {
	return q.dao.Select(ctx, nil, q.BuildQueryString())
}

func (q *StockEvaluateQuery) SelectForUpdate(ctx context.Context, tx *runtime.Tx) (*StockEvaluate, error) {
	q.ForUpdate = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockEvaluateQuery) SelectForShare(ctx context.Context, tx *runtime.Tx) (*StockEvaluate, error) {
	q.ForShare = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockEvaluateQuery) SelectList(ctx context.Context) (list []*StockEvaluate, err error) {
	return q.dao.SelectList(ctx, nil, q.BuildQueryString())
}

func (q *StockEvaluateQuery) SelectListForUpdate(ctx context.Context, tx *runtime.Tx) (list []*StockEvaluate, err error) {
	q.ForUpdate = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockEvaluateQuery) SelectListForShare(ctx context.Context, tx *runtime.Tx) (list []*StockEvaluate, err error) {
	q.ForShare = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockEvaluateQuery) Left() *StockEvaluateQuery {
	q.WhereBuffer.WriteString(" ( ")
	return q
}

func (q *StockEvaluateQuery) Right() *StockEvaluateQuery {
	q.WhereBuffer.WriteString(" ) ")
	return q
}

func (q *StockEvaluateQuery) And() *StockEvaluateQuery {
	q.WhereBuffer.WriteString(" AND ")
	return q
}

func (q *StockEvaluateQuery) Or() *StockEvaluateQuery {
	q.WhereBuffer.WriteString(" OR ")
	return q
}

func (q *StockEvaluateQuery) Not() *StockEvaluateQuery {
	q.WhereBuffer.WriteString(" NOT ")
	return q
}

func (q *StockEvaluateQuery) Limit(startIncluded int64, count int64) *StockEvaluateQuery {
	q.LimitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockEvaluateQuery) Sort(fieldName string, asc bool) *StockEvaluateQuery {
	if asc {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *StockEvaluateQuery) Id_Column(r runtime.Relation, v int64) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UserId_Column(r runtime.Relation, v string) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("user_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) StockId_Column(r runtime.Relation, v string) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("stock_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) TotalScore_Column(r runtime.Relation, v float64) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("total_score" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) EvalRemark_Column(r runtime.Relation, v string) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("eval_remark" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) CreateTime_Column(r runtime.Relation, v time.Time) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockEvaluateQuery) UpdateTime_Column(r runtime.Relation, v time.Time) *StockEvaluateQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type StockEvaluateDao struct {
	logger                       *zap.Logger
	db                           *DB
	insertStmt                   *runtime.Stmt
	updateStmt                   *runtime.Stmt
	deleteStmt                   *runtime.Stmt
	selectStmtAll                *runtime.Stmt
	selectStmtById               *runtime.Stmt
	selectStmtByUpdateTime       *runtime.Stmt
	selectStmtByUserId           *runtime.Stmt
	selectStmtByUserIdAndStockId *runtime.Stmt
}

func NewStockEvaluateDao(db *DB) (t *StockEvaluateDao) {
	t = &StockEvaluateDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *StockEvaluateDao) Init() (err error) {
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

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUpdateTime()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserIdAndStockId()
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

func (dao *StockEvaluateDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate")
	return err
}

func (dao *StockEvaluateDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate WHERE id=?")
	return err
}

func (dao *StockEvaluateDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate WHERE update_time=?")
	return err
}

func (dao *StockEvaluateDao) prepareSelectStmtByUserId() (err error) {
	dao.selectStmtByUserId, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate WHERE user_id=?")
	return err
}

func (dao *StockEvaluateDao) prepareSelectStmtByUserIdAndStockId() (err error) {
	dao.selectStmtByUserIdAndStockId, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate WHERE user_id=? AND stock_id=?")
	return err
}

func (dao *StockEvaluateDao) Insert(ctx context.Context, tx *runtime.Tx, e *StockEvaluate) (id int64, err error) {
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

func (dao *StockEvaluateDao) Update(ctx context.Context, tx *runtime.Tx, e *StockEvaluate) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.StockId, e.TotalScore, e.EvalRemark, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockEvaluateDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockEvaluateDao) ScanRow(row *runtime.Row) (*StockEvaluate, error) {
	e := &StockEvaluate{}
	err := row.Scan(&e.Id, &e.UserId, &e.StockId, &e.TotalScore, &e.EvalRemark, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockEvaluateDao) ScanRows(rows *runtime.Rows) (list []*StockEvaluate, err error) {
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

func (dao *StockEvaluateDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*StockEvaluate, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockEvaluateDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*StockEvaluate, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate "+query)
	return dao.ScanRow(row)
}

func (dao *StockEvaluateDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*StockEvaluate, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_EVALUATE_ALL_FIELDS_STRING+" FROM stock_evaluate "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockEvaluateDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*StockEvaluate, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *StockEvaluateDao) SelectByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (*StockEvaluate, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UpdateTime))
}

func (dao *StockEvaluateDao) SelectListByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (list []*StockEvaluate, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UpdateTime)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockEvaluateDao) SelectByUserIdAndStockId(ctx context.Context, tx *runtime.Tx, UserId string, StockId string) (*StockEvaluate, error) {
	stmt := dao.selectStmtByUserIdAndStockId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UserId, StockId))
}

func (dao *StockEvaluateDao) SelectListByUserId(ctx context.Context, tx *runtime.Tx, UserId string) (list []*StockEvaluate, err error) {
	stmt := dao.selectStmtByUserId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockEvaluateDao) GetQuery() *StockEvaluateQuery {
	return NewStockEvaluateQuery(dao)
}

const STOCK_INDEX_TABLE_NAME = "stock_index"

const STOCK_INDEX_FIELD_ID = "id"
const STOCK_INDEX_FIELD_INDEX_ID = "index_id"
const STOCK_INDEX_FIELD_USER_ID = "user_id"
const STOCK_INDEX_FIELD_INDEX_NAME = "index_name"
const STOCK_INDEX_FIELD_INDEX_DESC = "index_desc"
const STOCK_INDEX_FIELD_EVAL_WEIGHT = "eval_weight"
const STOCK_INDEX_FIELD_AI_WEIGHT = "ai_weight"
const STOCK_INDEX_FIELD_NI_WEIGHT = "ni_weight"
const STOCK_INDEX_FIELD_CREATE_TIME = "create_time"
const STOCK_INDEX_FIELD_UPDATE_TIME = "update_time"

const STOCK_INDEX_ALL_FIELDS_STRING = "id,index_id,user_id,index_name,index_desc,eval_weight,ai_weight,ni_weight,create_time,update_time"

var STOCK_INDEX_ALL_FIELDS = []string{
	"id",
	"index_id",
	"user_id",
	"index_name",
	"index_desc",
	"eval_weight",
	"ai_weight",
	"ni_weight",
	"create_time",
	"update_time",
}

type StockIndex struct {
	Id         int64  //size=20
	IndexId    string //size=32
	UserId     string //size=32
	IndexName  string //size=32
	IndexDesc  string //size=256
	EvalWeight int32  //size=11
	AiWeight   int32  //size=11
	NiWeight   int32  //size=11
	CreateTime time.Time
	UpdateTime time.Time
}

type StockIndexQuery struct {
	dao *StockIndexDao
	runtime.Query
}

func NewStockIndexQuery(dao *StockIndexDao) *StockIndexQuery {
	q := &StockIndexQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockIndexQuery) Select(ctx context.Context) (*StockIndex, error) {
	return q.dao.Select(ctx, nil, q.BuildQueryString())
}

func (q *StockIndexQuery) SelectForUpdate(ctx context.Context, tx *runtime.Tx) (*StockIndex, error) {
	q.ForUpdate = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexQuery) SelectForShare(ctx context.Context, tx *runtime.Tx) (*StockIndex, error) {
	q.ForShare = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexQuery) SelectList(ctx context.Context) (list []*StockIndex, err error) {
	return q.dao.SelectList(ctx, nil, q.BuildQueryString())
}

func (q *StockIndexQuery) SelectListForUpdate(ctx context.Context, tx *runtime.Tx) (list []*StockIndex, err error) {
	q.ForUpdate = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexQuery) SelectListForShare(ctx context.Context, tx *runtime.Tx) (list []*StockIndex, err error) {
	q.ForShare = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexQuery) Left() *StockIndexQuery {
	q.WhereBuffer.WriteString(" ( ")
	return q
}

func (q *StockIndexQuery) Right() *StockIndexQuery {
	q.WhereBuffer.WriteString(" ) ")
	return q
}

func (q *StockIndexQuery) And() *StockIndexQuery {
	q.WhereBuffer.WriteString(" AND ")
	return q
}

func (q *StockIndexQuery) Or() *StockIndexQuery {
	q.WhereBuffer.WriteString(" OR ")
	return q
}

func (q *StockIndexQuery) Not() *StockIndexQuery {
	q.WhereBuffer.WriteString(" NOT ")
	return q
}

func (q *StockIndexQuery) Limit(startIncluded int64, count int64) *StockIndexQuery {
	q.LimitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockIndexQuery) Sort(fieldName string, asc bool) *StockIndexQuery {
	if asc {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *StockIndexQuery) Id_Column(r runtime.Relation, v int64) *StockIndexQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexId_Column(r runtime.Relation, v string) *StockIndexQuery {
	q.WhereBuffer.WriteString("index_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UserId_Column(r runtime.Relation, v string) *StockIndexQuery {
	q.WhereBuffer.WriteString("user_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexName_Column(r runtime.Relation, v string) *StockIndexQuery {
	q.WhereBuffer.WriteString("index_name" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) IndexDesc_Column(r runtime.Relation, v string) *StockIndexQuery {
	q.WhereBuffer.WriteString("index_desc" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) EvalWeight_Column(r runtime.Relation, v int32) *StockIndexQuery {
	q.WhereBuffer.WriteString("eval_weight" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) AiWeight_Column(r runtime.Relation, v int32) *StockIndexQuery {
	q.WhereBuffer.WriteString("ai_weight" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) NiWeight_Column(r runtime.Relation, v int32) *StockIndexQuery {
	q.WhereBuffer.WriteString("ni_weight" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) CreateTime_Column(r runtime.Relation, v time.Time) *StockIndexQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexQuery) UpdateTime_Column(r runtime.Relation, v time.Time) *StockIndexQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type StockIndexDao struct {
	logger                         *zap.Logger
	db                             *DB
	insertStmt                     *runtime.Stmt
	updateStmt                     *runtime.Stmt
	deleteStmt                     *runtime.Stmt
	selectStmtAll                  *runtime.Stmt
	selectStmtById                 *runtime.Stmt
	selectStmtByUpdateTime         *runtime.Stmt
	selectStmtByIndexName          *runtime.Stmt
	selectStmtByIndexId            *runtime.Stmt
	selectStmtByUserId             *runtime.Stmt
	selectStmtByUserIdAndIndexName *runtime.Stmt
}

func NewStockIndexDao(db *DB) (t *StockIndexDao) {
	t = &StockIndexDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *StockIndexDao) Init() (err error) {
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

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUpdateTime()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByIndexName()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByIndexId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserIdAndIndexName()
	if err != nil {
		return err
	}

	return nil
}
func (dao *StockIndexDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock_index (index_id,user_id,index_name,index_desc,eval_weight,ai_weight,ni_weight,create_time,update_time) VALUES (?,?,?,?,?,?,?,?,?)")
	return err
}

func (dao *StockIndexDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock_index SET index_id=?,user_id=?,index_name=?,index_desc=?,eval_weight=?,ai_weight=?,ni_weight=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockIndexDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock_index WHERE id=?")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index WHERE id=?")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index WHERE update_time=?")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtByIndexName() (err error) {
	dao.selectStmtByIndexName, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index WHERE index_name=?")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtByIndexId() (err error) {
	dao.selectStmtByIndexId, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index WHERE index_id=?")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtByUserId() (err error) {
	dao.selectStmtByUserId, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index WHERE user_id=?")
	return err
}

func (dao *StockIndexDao) prepareSelectStmtByUserIdAndIndexName() (err error) {
	dao.selectStmtByUserIdAndIndexName, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index WHERE user_id=? AND index_name=?")
	return err
}

func (dao *StockIndexDao) Insert(ctx context.Context, tx *runtime.Tx, e *StockIndex) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.IndexId, e.UserId, e.IndexName, e.IndexDesc, e.EvalWeight, e.AiWeight, e.NiWeight, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockIndexDao) Update(ctx context.Context, tx *runtime.Tx, e *StockIndex) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.IndexId, e.UserId, e.IndexName, e.IndexDesc, e.EvalWeight, e.AiWeight, e.NiWeight, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockIndexDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockIndexDao) ScanRow(row *runtime.Row) (*StockIndex, error) {
	e := &StockIndex{}
	err := row.Scan(&e.Id, &e.IndexId, &e.UserId, &e.IndexName, &e.IndexDesc, &e.EvalWeight, &e.AiWeight, &e.NiWeight, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockIndexDao) ScanRows(rows *runtime.Rows) (list []*StockIndex, err error) {
	list = make([]*StockIndex, 0)
	for rows.Next() {
		e := StockIndex{}
		err = rows.Scan(&e.Id, &e.IndexId, &e.UserId, &e.IndexName, &e.IndexDesc, &e.EvalWeight, &e.AiWeight, &e.NiWeight, &e.CreateTime, &e.UpdateTime)
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

func (dao *StockIndexDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*StockIndex, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*StockIndex, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index "+query)
	return dao.ScanRow(row)
}

func (dao *StockIndexDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*StockIndex, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_INDEX_ALL_FIELDS_STRING+" FROM stock_index "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*StockIndex, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *StockIndexDao) SelectByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (*StockIndex, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UpdateTime))
}

func (dao *StockIndexDao) SelectListByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (list []*StockIndex, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UpdateTime)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexDao) SelectByIndexName(ctx context.Context, tx *runtime.Tx, IndexName string) (*StockIndex, error) {
	stmt := dao.selectStmtByIndexName
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, IndexName))
}

func (dao *StockIndexDao) SelectListByIndexName(ctx context.Context, tx *runtime.Tx, IndexName string) (list []*StockIndex, err error) {
	stmt := dao.selectStmtByIndexName
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, IndexName)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexDao) SelectByIndexId(ctx context.Context, tx *runtime.Tx, IndexId string) (*StockIndex, error) {
	stmt := dao.selectStmtByIndexId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, IndexId))
}

func (dao *StockIndexDao) SelectByUserIdAndIndexName(ctx context.Context, tx *runtime.Tx, UserId string, IndexName string) (*StockIndex, error) {
	stmt := dao.selectStmtByUserIdAndIndexName
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UserId, IndexName))
}

func (dao *StockIndexDao) SelectListByUserId(ctx context.Context, tx *runtime.Tx, UserId string) (list []*StockIndex, err error) {
	stmt := dao.selectStmtByUserId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexDao) GetQuery() *StockIndexQuery {
	return NewStockIndexQuery(dao)
}

const STOCK_INDEX_ID_GEN_TABLE_NAME = "stock_index_id_gen"

const STOCK_INDEX_ID_GEN_FIELD_ID = "id"
const STOCK_INDEX_ID_GEN_FIELD_CURRENT_INDEX_ID = "current_index_id"
const STOCK_INDEX_ID_GEN_FIELD_CREATE_TIME = "create_time"
const STOCK_INDEX_ID_GEN_FIELD_UPDATE_TIME = "update_time"

const STOCK_INDEX_ID_GEN_ALL_FIELDS_STRING = "id,current_index_id,create_time,update_time"

var STOCK_INDEX_ID_GEN_ALL_FIELDS = []string{
	"id",
	"current_index_id",
	"create_time",
	"update_time",
}

type StockIndexIdGen struct {
	Id             int64 //size=20
	CurrentIndexId int64 //size=20
	CreateTime     time.Time
	UpdateTime     time.Time
}

type StockIndexIdGenQuery struct {
	dao *StockIndexIdGenDao
	runtime.Query
}

func NewStockIndexIdGenQuery(dao *StockIndexIdGenDao) *StockIndexIdGenQuery {
	q := &StockIndexIdGenQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *StockIndexIdGenQuery) Select(ctx context.Context) (*StockIndexIdGen, error) {
	return q.dao.Select(ctx, nil, q.BuildQueryString())
}

func (q *StockIndexIdGenQuery) SelectForUpdate(ctx context.Context, tx *runtime.Tx) (*StockIndexIdGen, error) {
	q.ForUpdate = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexIdGenQuery) SelectForShare(ctx context.Context, tx *runtime.Tx) (*StockIndexIdGen, error) {
	q.ForShare = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexIdGenQuery) SelectList(ctx context.Context) (list []*StockIndexIdGen, err error) {
	return q.dao.SelectList(ctx, nil, q.BuildQueryString())
}

func (q *StockIndexIdGenQuery) SelectListForUpdate(ctx context.Context, tx *runtime.Tx) (list []*StockIndexIdGen, err error) {
	q.ForUpdate = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexIdGenQuery) SelectListForShare(ctx context.Context, tx *runtime.Tx) (list []*StockIndexIdGen, err error) {
	q.ForShare = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *StockIndexIdGenQuery) Left() *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString(" ( ")
	return q
}

func (q *StockIndexIdGenQuery) Right() *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString(" ) ")
	return q
}

func (q *StockIndexIdGenQuery) And() *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString(" AND ")
	return q
}

func (q *StockIndexIdGenQuery) Or() *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString(" OR ")
	return q
}

func (q *StockIndexIdGenQuery) Not() *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString(" NOT ")
	return q
}

func (q *StockIndexIdGenQuery) Limit(startIncluded int64, count int64) *StockIndexIdGenQuery {
	q.LimitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *StockIndexIdGenQuery) Sort(fieldName string, asc bool) *StockIndexIdGenQuery {
	if asc {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *StockIndexIdGenQuery) Id_Column(r runtime.Relation, v int64) *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexIdGenQuery) CurrentIndexId_Column(r runtime.Relation, v int64) *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString("current_index_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexIdGenQuery) CreateTime_Column(r runtime.Relation, v time.Time) *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *StockIndexIdGenQuery) UpdateTime_Column(r runtime.Relation, v time.Time) *StockIndexIdGenQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type StockIndexIdGenDao struct {
	logger         *zap.Logger
	db             *DB
	insertStmt     *runtime.Stmt
	updateStmt     *runtime.Stmt
	deleteStmt     *runtime.Stmt
	selectStmtAll  *runtime.Stmt
	selectStmtById *runtime.Stmt
}

func NewStockIndexIdGenDao(db *DB) (t *StockIndexIdGenDao) {
	t = &StockIndexIdGenDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *StockIndexIdGenDao) Init() (err error) {
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

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	return nil
}
func (dao *StockIndexIdGenDao) prepareInsertStmt() (err error) {
	dao.insertStmt, err = dao.db.Prepare(context.Background(), "INSERT INTO stock_index_id_gen (current_index_id,create_time,update_time) VALUES (?,?,?)")
	return err
}

func (dao *StockIndexIdGenDao) prepareUpdateStmt() (err error) {
	dao.updateStmt, err = dao.db.Prepare(context.Background(), "UPDATE stock_index_id_gen SET current_index_id=?,create_time=?,update_time=? WHERE id=?")
	return err
}

func (dao *StockIndexIdGenDao) prepareDeleteStmt() (err error) {
	dao.deleteStmt, err = dao.db.Prepare(context.Background(), "DELETE FROM stock_index_id_gen WHERE id=?")
	return err
}

func (dao *StockIndexIdGenDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ID_GEN_ALL_FIELDS_STRING+" FROM stock_index_id_gen")
	return err
}

func (dao *StockIndexIdGenDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+STOCK_INDEX_ID_GEN_ALL_FIELDS_STRING+" FROM stock_index_id_gen WHERE id=?")
	return err
}

func (dao *StockIndexIdGenDao) Insert(ctx context.Context, tx *runtime.Tx, e *StockIndexIdGen) (id int64, err error) {
	stmt := dao.insertStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.CurrentIndexId, e.CreateTime, e.UpdateTime)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *StockIndexIdGenDao) Update(ctx context.Context, tx *runtime.Tx, e *StockIndexIdGen) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.CurrentIndexId, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockIndexIdGenDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *StockIndexIdGenDao) ScanRow(row *runtime.Row) (*StockIndexIdGen, error) {
	e := &StockIndexIdGen{}
	err := row.Scan(&e.Id, &e.CurrentIndexId, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *StockIndexIdGenDao) ScanRows(rows *runtime.Rows) (list []*StockIndexIdGen, err error) {
	list = make([]*StockIndexIdGen, 0)
	for rows.Next() {
		e := StockIndexIdGen{}
		err = rows.Scan(&e.Id, &e.CurrentIndexId, &e.CreateTime, &e.UpdateTime)
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

func (dao *StockIndexIdGenDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*StockIndexIdGen, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexIdGenDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*StockIndexIdGen, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+STOCK_INDEX_ID_GEN_ALL_FIELDS_STRING+" FROM stock_index_id_gen "+query)
	return dao.ScanRow(row)
}

func (dao *StockIndexIdGenDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*StockIndexIdGen, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+STOCK_INDEX_ID_GEN_ALL_FIELDS_STRING+" FROM stock_index_id_gen "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *StockIndexIdGenDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*StockIndexIdGen, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *StockIndexIdGenDao) GetQuery() *StockIndexIdGenQuery {
	return NewStockIndexIdGenQuery(dao)
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
	dao *UserSettingDao
	runtime.Query
}

func NewUserSettingQuery(dao *UserSettingDao) *UserSettingQuery {
	q := &UserSettingQuery{}
	q.dao = dao
	q.WhereBuffer = bytes.NewBufferString("")
	q.LimitBuffer = bytes.NewBufferString("")
	q.OrderBuffer = bytes.NewBufferString("")

	return q
}

func (q *UserSettingQuery) Select(ctx context.Context) (*UserSetting, error) {
	return q.dao.Select(ctx, nil, q.BuildQueryString())
}

func (q *UserSettingQuery) SelectForUpdate(ctx context.Context, tx *runtime.Tx) (*UserSetting, error) {
	q.ForUpdate = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *UserSettingQuery) SelectForShare(ctx context.Context, tx *runtime.Tx) (*UserSetting, error) {
	q.ForShare = true
	return q.dao.Select(ctx, tx, q.BuildQueryString())
}

func (q *UserSettingQuery) SelectList(ctx context.Context) (list []*UserSetting, err error) {
	return q.dao.SelectList(ctx, nil, q.BuildQueryString())
}

func (q *UserSettingQuery) SelectListForUpdate(ctx context.Context, tx *runtime.Tx) (list []*UserSetting, err error) {
	q.ForUpdate = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *UserSettingQuery) SelectListForShare(ctx context.Context, tx *runtime.Tx) (list []*UserSetting, err error) {
	q.ForShare = true
	return q.dao.SelectList(ctx, tx, q.BuildQueryString())
}

func (q *UserSettingQuery) Left() *UserSettingQuery {
	q.WhereBuffer.WriteString(" ( ")
	return q
}

func (q *UserSettingQuery) Right() *UserSettingQuery {
	q.WhereBuffer.WriteString(" ) ")
	return q
}

func (q *UserSettingQuery) And() *UserSettingQuery {
	q.WhereBuffer.WriteString(" AND ")
	return q
}

func (q *UserSettingQuery) Or() *UserSettingQuery {
	q.WhereBuffer.WriteString(" OR ")
	return q
}

func (q *UserSettingQuery) Not() *UserSettingQuery {
	q.WhereBuffer.WriteString(" NOT ")
	return q
}

func (q *UserSettingQuery) Limit(startIncluded int64, count int64) *UserSettingQuery {
	q.LimitBuffer.WriteString(fmt.Sprintf(" limit %d,%d", startIncluded, count))
	return q
}

func (q *UserSettingQuery) Sort(fieldName string, asc bool) *UserSettingQuery {
	if asc {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s asc", fieldName))
	} else {
		q.OrderBuffer.WriteString(fmt.Sprintf(" order by %s desc", fieldName))
	}

	return q
}
func (q *UserSettingQuery) Id_Column(r runtime.Relation, v int64) *UserSettingQuery {
	q.WhereBuffer.WriteString("id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UserId_Column(r runtime.Relation, v string) *UserSettingQuery {
	q.WhereBuffer.WriteString("user_id" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigKey_Column(r runtime.Relation, v string) *UserSettingQuery {
	q.WhereBuffer.WriteString("config_key" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) ConfigValue_Column(r runtime.Relation, v string) *UserSettingQuery {
	q.WhereBuffer.WriteString("config_value" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) CreateTime_Column(r runtime.Relation, v time.Time) *UserSettingQuery {
	q.WhereBuffer.WriteString("create_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

func (q *UserSettingQuery) UpdateTime_Column(r runtime.Relation, v time.Time) *UserSettingQuery {
	q.WhereBuffer.WriteString("update_time" + string(r) + "'" + fmt.Sprint(v) + "'")
	return q
}

type UserSettingDao struct {
	logger                         *zap.Logger
	db                             *DB
	insertStmt                     *runtime.Stmt
	updateStmt                     *runtime.Stmt
	deleteStmt                     *runtime.Stmt
	selectStmtAll                  *runtime.Stmt
	selectStmtById                 *runtime.Stmt
	selectStmtByUpdateTime         *runtime.Stmt
	selectStmtByUserId             *runtime.Stmt
	selectStmtByUserIdAndConfigKey *runtime.Stmt
}

func NewUserSettingDao(db *DB) (t *UserSettingDao) {
	t = &UserSettingDao{}
	t.logger = log.TypedLogger(t)
	t.db = db

	return t
}

func (dao *UserSettingDao) Init() (err error) {
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

	err = dao.prepareSelectStmtAll()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtById()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUpdateTime()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserId()
	if err != nil {
		return err
	}

	err = dao.prepareSelectStmtByUserIdAndConfigKey()
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

func (dao *UserSettingDao) prepareSelectStmtAll() (err error) {
	dao.selectStmtAll, err = dao.db.Prepare(context.Background(), "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting")
	return err
}

func (dao *UserSettingDao) prepareSelectStmtById() (err error) {
	dao.selectStmtById, err = dao.db.Prepare(context.Background(), "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting WHERE id=?")
	return err
}

func (dao *UserSettingDao) prepareSelectStmtByUpdateTime() (err error) {
	dao.selectStmtByUpdateTime, err = dao.db.Prepare(context.Background(), "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting WHERE update_time=?")
	return err
}

func (dao *UserSettingDao) prepareSelectStmtByUserId() (err error) {
	dao.selectStmtByUserId, err = dao.db.Prepare(context.Background(), "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting WHERE user_id=?")
	return err
}

func (dao *UserSettingDao) prepareSelectStmtByUserIdAndConfigKey() (err error) {
	dao.selectStmtByUserIdAndConfigKey, err = dao.db.Prepare(context.Background(), "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting WHERE user_id=? AND config_key=?")
	return err
}

func (dao *UserSettingDao) Insert(ctx context.Context, tx *runtime.Tx, e *UserSetting) (id int64, err error) {
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

func (dao *UserSettingDao) Update(ctx context.Context, tx *runtime.Tx, e *UserSetting) (rowsAffected int64, err error) {
	stmt := dao.updateStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, e.UserId, e.ConfigKey, e.ConfigValue, e.CreateTime, e.UpdateTime, e.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *UserSettingDao) Delete(ctx context.Context, tx *runtime.Tx, id int64) (rowsAffected int64, err error) {
	stmt := dao.deleteStmt
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	result, err := stmt.Exec(ctx, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (dao *UserSettingDao) ScanRow(row *runtime.Row) (*UserSetting, error) {
	e := &UserSetting{}
	err := row.Scan(&e.Id, &e.UserId, &e.ConfigKey, &e.ConfigValue, &e.CreateTime, &e.UpdateTime)
	if err != nil {
		if err == runtime.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return e, nil
}

func (dao *UserSettingDao) ScanRows(rows *runtime.Rows) (list []*UserSetting, err error) {
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

func (dao *UserSettingDao) SelectAll(ctx context.Context, tx *runtime.Tx) (list []*UserSetting, err error) {
	stmt := dao.selectStmtAll
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *UserSettingDao) Select(ctx context.Context, tx *runtime.Tx, query string) (*UserSetting, error) {
	row := dao.db.QueryRow(ctx, "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting "+query)
	return dao.ScanRow(row)
}

func (dao *UserSettingDao) SelectList(ctx context.Context, tx *runtime.Tx, query string) (list []*UserSetting, err error) {
	rows, err := dao.db.Query(ctx, "SELECT "+USER_SETTING_ALL_FIELDS_STRING+" FROM user_setting "+query)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *UserSettingDao) SelectById(ctx context.Context, tx *runtime.Tx, Id int64) (*UserSetting, error) {
	stmt := dao.selectStmtById
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, Id))
}

func (dao *UserSettingDao) SelectByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (*UserSetting, error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	return dao.ScanRow(stmt.QueryRow(ctx, UpdateTime))
}

func (dao *UserSettingDao) SelectListByUpdateTime(ctx context.Context, tx *runtime.Tx, UpdateTime time.Time) (list []*UserSetting, err error) {
	stmt := dao.selectStmtByUpdateTime
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UpdateTime)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *UserSettingDao) SelectListByUserId(ctx context.Context, tx *runtime.Tx, UserId string) (list []*UserSetting, err error) {
	stmt := dao.selectStmtByUserId
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *UserSettingDao) SelectListByUserIdAndConfigKey(ctx context.Context, tx *runtime.Tx, UserId string, ConfigKey string) (list []*UserSetting, err error) {
	stmt := dao.selectStmtByUserIdAndConfigKey
	if tx != nil {
		stmt = tx.Stmt(ctx, stmt)
	}

	rows, err := stmt.Query(ctx, UserId, ConfigKey)
	if err != nil {
		dao.logger.Error("sqlDriver", zap.Error(err))
		return nil, err
	}

	return dao.ScanRows(rows)
}

func (dao *UserSettingDao) GetQuery() *UserSettingQuery {
	return NewUserSettingQuery(dao)
}

type DB struct {
	runtime.DB
	IndexEvaluate   *IndexEvaluateDao
	StockEvaluate   *StockEvaluateDao
	StockIndex      *StockIndexDao
	StockIndexIdGen *StockIndexIdGenDao
	UserSetting     *UserSettingDao
}

func NewDB(connectionString string) (d *DB, err error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connectionString nil")
	}

	d = &DB{}

	db, err := runtime.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	d.DB = *db

	err = d.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	d.IndexEvaluate = NewIndexEvaluateDao(d)
	err = d.IndexEvaluate.Init()
	if err != nil {
		return nil, err
	}

	d.StockEvaluate = NewStockEvaluateDao(d)
	err = d.StockEvaluate.Init()
	if err != nil {
		return nil, err
	}

	d.StockIndex = NewStockIndexDao(d)
	err = d.StockIndex.Init()
	if err != nil {
		return nil, err
	}

	d.StockIndexIdGen = NewStockIndexIdGenDao(d)
	err = d.StockIndexIdGen.Init()
	if err != nil {
		return nil, err
	}

	d.UserSetting = NewUserSettingDao(d)
	err = d.UserSetting.Init()
	if err != nil {
		return nil, err
	}

	return d, nil
}
