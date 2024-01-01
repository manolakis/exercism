package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var result []Record

	for _, record := range in {
		if predicate(record) {
			result = append(result, record)
		}
	}

	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(daysPeriod DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return record.Day >= daysPeriod.From && record.Day <= daysPeriod.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(category string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, daysPeriod DaysPeriod) float64 {
	var amount float64

	for _, record := range Filter(in, ByDaysPeriod(daysPeriod)) {
		amount += record.Amount
	}

	return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	records := Filter(in, ByCategory(c))
	if len(records) == 0 {
		return 0, errors.New(fmt.Sprintf("unknown category %s", c))
	}

	return TotalByPeriod(records, p), nil
}
