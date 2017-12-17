package models

import "time"

type UserIndexEvaluate struct {
	IndexName  string
	EvalStars  int32
	EvalRemark string
	UpdateTime time.Time
}
