package domain

type ReportRepository interface {
	Find(id ReportId) (Report, error)
	Save(i Report)
}
