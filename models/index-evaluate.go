package models

import "time"

type IndexEvaluate struct {
	IndexName  string
	EvalStars  int32
	EvalRemark string
	UpdateTime time.Time
}
