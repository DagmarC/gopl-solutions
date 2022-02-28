package github

import (
	"fmt"
	"time"
)

type DateCateg string

var (
	BeforeDate  DateCateg = "created:<%d-%.2d-%.2d"              // created on or before  the date (more than 1yr old)
	AfterDate   DateCateg = "created:>%d-%.2d-%.2d"              //created after the date
	BetweenDate DateCateg = "created:%d-%.2d-%.2d..%d-%.2d-%.2d" // In between
)

func (d *DateCateg) String() string {
	return string(*d)
}

func (d *DateCateg) AddDateFromNow(pYear, pMonth, pDay int) string {
	t := time.Now().AddDate(pYear, pMonth, pDay)
	return fmt.Sprintf(d.String(), t.Year(), t.Month(), t.Day())
}

func (d *DateCateg) AddDatesFromNow(pYear, pMonth, pDay, qYear, qMonth, qDay int) string {
	t1 := time.Now().AddDate(pYear, pMonth, pDay)
	t2 := time.Now().AddDate(qYear, qMonth, qDay)

	return fmt.Sprintf(d.String(), t1.Year(), t1.Month(), t1.Day(), t2.Year(), t2.Month(), t2.Day())
}

func InitCategories() map[string]*IssuesSearchResult {
	categories := make(map[string]*IssuesSearchResult, 4)

	categories[BeforeDate.AddDateFromNow(-1, 0, 0)] = &IssuesSearchResult{} // Before one year.
	categories[AfterDate.AddDateFromNow(0, -1, 0)] = &IssuesSearchResult{}  // Less than month.

	// Newer than 1 year but Older than 1 Month
	categories[BetweenDate.AddDatesFromNow(-1, 0, 0, 0, -1, 0)] = &IssuesSearchResult{}

	return categories
}
