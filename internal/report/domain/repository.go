package domain

type ReportRepository interface {
	Find(id int) (Report, error)
	Save(i Report)
}
