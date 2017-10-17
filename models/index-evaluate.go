package models

import "time"

type IndexEvaluate struct {
	IndexId    string
	EvalStars  int32
	EvalRemark string
	UpdateTime time.Time
}
