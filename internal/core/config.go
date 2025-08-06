package core

type Config struct {
	RegionQPath string
}
type Exam struct {
	ExamId int64
	QA
}
type QA struct {
	ExamId       int64
	QCode        string
	Capability   string
	Q            string
	Opts         string
	QOpts        string
	Ans          string
	DetraderAns  string
	Level        string
	Analysis     string
	RatingPoints string
	QCreateAt    string
	ACreateAt    string
}
