package fin_stock_assistant

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/NeuronEvolution/StockAssistant/models"
	"github.com/NeuronEvolution/sql/runtime"
)

func (dao *StockIndexIdGenDao) NextStockIndexId() (id int64, err error) {
	tx, err := dao.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	dbGen, err := dao.GetQuery().
		Id_Column(runtime.RELATION_EQUAL, 1).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return 0, err
	}

	if dbGen == nil {
		return 0, err
	}

	dbGen.CurrentIndexId++

	err = dao.Update(context.Background(), tx, dbGen)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return dbGen.CurrentIndexId, nil
}

func (dao *StockIndexDao) InsertStockIndex(userId string, index *models.StockIndex) (indexAdded *models.StockIndex, err error) {
	tx, err := dao.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndexOld, err := dao.GetQuery().
		UserId_Column(runtime.RELATION_EQUAL, userId).
		IndexName_Column(runtime.RELATION_EQUAL, index.IndexName).
		SelectForUpdate(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	if dbIndexOld != nil {
		return nil, fmt.Errorf("name exist")
	}

	indexId, err := dao.db.StockIndexIdGen.NextStockIndexId()
	if err != nil {
		return nil, err
	}

	dbIndex := &StockIndex{}
	dbIndex.IndexId = fmt.Sprint(indexId)
	dbIndex.IndexName = index.IndexName
	dbIndex.IndexDesc = index.IndexDesc
	dbIndex.EvalWeight = index.EvalWeight
	dbIndex.AiWeight = index.AIWeight
	dbIndex.NiWeight = index.NIWeight

	rowsAffected, err := dao.Insert(context.Background(), tx, dbIndex)
	if err != nil {
		return nil, err
	}

	if rowsAffected != 1 {
		return nil, fmt.Errorf("rowsAffected: %s", rowsAffected)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	indexAdded = &models.StockIndex{}
	indexAdded.IndexId = dbIndex.IndexId
	indexAdded.IndexName = index.IndexName
	indexAdded.IndexDesc = index.IndexDesc
	indexAdded.EvalWeight = index.EvalWeight
	indexAdded.AIWeight = index.AIWeight
	indexAdded.NIWeight = index.NIWeight

	return indexAdded, nil
}

func (dao *StockIndexDao) UpdateStockIndex(userId string, index *models.StockIndex) (indexUpdated *models.StockIndex, err error) {
	tx, err := dao.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dbIndex, err := dao.SelectByIndexId(context.Background(), tx, index.IndexId)
	if err != nil {
		return nil, err
	}

	if dbIndex == nil {
		return nil, fmt.Errorf("index not exist")
	}

	if dbIndex.UserId != userId {
		return nil, fmt.Errorf("user have no this index")
	}

	//check name
	if dbIndex.IndexName != index.IndexName {
		dbIndexNameOld, err := dao.SelectByUserIdAndIndexName(context.Background(), tx, userId, index.IndexName)
		if err != nil {
			return nil, err
		}

		if dbIndexNameOld != nil {
			return nil, fmt.Errorf("index name used")
		}
	}

	dbIndex.IndexName = index.IndexName
	dbIndex.IndexDesc = index.IndexDesc
	dbIndex.EvalWeight = index.EvalWeight
	dbIndex.AiWeight = index.AIWeight
	dbIndex.NiWeight = index.NIWeight

	err = dao.Update(context.Background(), tx, dbIndex)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	indexUpdated = &models.StockIndex{}
	indexUpdated.IndexId = dbIndex.IndexId
	indexUpdated.IndexName = dbIndex.IndexName
	indexUpdated.IndexDesc = dbIndex.IndexDesc
	indexUpdated.EvalWeight = dbIndex.EvalWeight
	index.AIWeight = dbIndex.AiWeight
	index.NIWeight = dbIndex.NiWeight

	return indexUpdated, nil
}

func (dao *StockIndexDao) DeleteStockIndex(userId string, indexId string) (err error) {
	dbIndex, err := dao.SelectByIndexId(context.Background(), nil, indexId)
	if err != nil {
		return err
	}

	if dbIndex == nil {
		return fmt.Errorf("already deleted")
	}

	if dbIndex.UserId != userId {
		return fmt.Errorf("user not owner")
	}

	err = dao.Delete(context.Background(), nil, dbIndex.Id)
	if err != nil {
		return err
	}

	return nil
}
