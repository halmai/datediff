package datecalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCalcDiffDays tests the test data explicitly given by the specification.
func TestCalcDiffDays(t *testing.T) {
	tests := []struct {
		name       string
		sendYMD1   [3]int
		sendYMD2   [3]int
		wantDiff   int
		wantErrStr string
	}{
		{
			name:       "GivenIdenticalDatesRequested_WhenCalcDiffDaysCalled_ThenErrorIsReturned",
			sendYMD1:   [3]int{2001, 1, 1},
			sendYMD2:   [3]int{2001, 1, 1},
			wantErrStr: "dates must be different",
		},
		{
			name:     "Given1stand3rdOfJanuaryRequested_WhenCalcDiffDaysCalled_Then1IsReturned",
			sendYMD1: [3]int{2001, 1, 1},
			sendYMD2: [3]int{2001, 1, 3},
			wantDiff: 1,
		},
		{
			name:     "Given20DaysLaterRequested_WhenCalcDiffDaysCalled_Then19IsReturned",
			sendYMD1: [3]int{1983, 6, 2},
			sendYMD2: [3]int{1983, 6, 22},
			wantDiff: 19,
		},
		{
			name:     "Given~5MonthsPlusLaterRequestedRequested_WhenCalcDiffDaysCalled_Then173IsReturned",
			sendYMD1: [3]int{1984, 7, 4},
			sendYMD2: [3]int{1984, 12, 25},
			wantDiff: 173,
		},
		{
			name:     "Given~6YearsEarlierRequested_WhenCalcDiffDaysCalled_Then1IsReturned",
			sendYMD1: [3]int{1989, 1, 3},
			sendYMD2: [3]int{1983, 8, 3},
			wantDiff: 1979, // todo: confirm if the specification is correct
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotDiff, gotErr := CalcDiffDays(
				test.sendYMD1[0],
				test.sendYMD1[1],
				test.sendYMD1[2],
				test.sendYMD2[0],
				test.sendYMD2[1],
				test.sendYMD2[2],
			)

			if test.wantErrStr == "" {
				assert.Nil(t, gotErr)
				assert.Equal(t, test.wantDiff, gotDiff)
			} else {
				assert.Equal(t, test.wantErrStr, gotErr.Error())
				assert.Equal(t, 0, gotDiff)
			}
		})
	}
}

func TestOffsetOfDay(t *testing.T) {
	tests := []struct {
		name      string
		giveYear  int
		giveMonth int
		giveDay   int
		wantIndex int
	}{
		{
			name:      "GivenFirstDayRequested_WhenOffsetOfDayCalled_ThenZeroIsReturned",
			giveYear:  1900,
			giveMonth: 1,
			giveDay:   1,
			wantIndex: 0,
		},
		{
			name:      "GivenSecondDayRequested_WhenOffsetOfDayCalled_ThenOneIsReturned",
			giveYear:  1900,
			giveMonth: 1,
			giveDay:   2,
			wantIndex: 1,
		},
		{
			name:      "GivenEndOfJanuaryInFirstYearRequested_WhenOffsetOfDayCalled_Then30IsReturned",
			giveYear:  1900,
			giveMonth: 1,
			giveDay:   31,
			wantIndex: 30,
		},
		{
			name:      "GivenBeginningOfFebruaryInFirstYearRequested_WhenOffsetOfDayCalled_Then31IsReturned",
			giveYear:  1900,
			giveMonth: 2,
			giveDay:   1,
			wantIndex: 31,
		},
		{
			name:      "GivenEndOfFebruaryInFirstYearRequested_WhenOffsetOfDayCalled_Then58IsReturned",
			giveYear:  1900,
			giveMonth: 2,
			giveDay:   28,
			wantIndex: 58,
		},
		{
			name:      "GivenBeginningOfMarchInFirstYearRequested_WhenOffsetOfDayCalled_Then59IsReturned",
			giveYear:  1900,
			giveMonth: 3,
			giveDay:   1,
			wantIndex: 59,
		},
		{
			name:      "Given28thOfFebruaryInFirstLeapYearRequested_WhenOffsetOfDayCalled_Then1518IsReturned",
			giveYear:  1904,
			giveMonth: 2,
			giveDay:   28,
			wantIndex: 1518,
		},
		{
			name:      "Given29thOfFebruaryInFirstLeapYearRequested_WhenOffsetOfDayCalled_Then1519IsReturned",
			giveYear:  1904,
			giveMonth: 2,
			giveDay:   29,
			wantIndex: 1519,
		},
		{
			name:      "Given1stOfMarchInFirstLeapYearRequested_WhenOffsetOfDayCalled_Then1520IsReturned",
			giveYear:  1904,
			giveMonth: 3,
			giveDay:   1,
			wantIndex: 1520,
		},
		{
			name:      "GivenEndOfTheLastYearRequested_WhenOffsetOfDayCalled_Then401766IsReturned",
			giveYear:  2999,
			giveMonth: 12,
			giveDay:   31,
			wantIndex: 401766,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotIndex := offsetOfDay(test.giveYear, test.giveMonth, test.giveDay)

			assert.Equal(t, test.wantIndex, gotIndex)
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name           string
		giveYear       int
		wantIsLeapYear bool
	}{
		{
			name:           "GivenNonFourthYearRequested_WhenIsLeapYearCalled_Then28IsReturned",
			giveYear:       2001,
			wantIsLeapYear: false,
		},
		{
			name:           "GivenFourthYearRequested_WhenIsLeapYearCalled_Then29IsReturned",
			giveYear:       1996,
			wantIsLeapYear: true,
		},
		{
			name:           "GivenHundredthYearRequested_WhenIsLeapYearCalled_Then29IsReturned",
			giveYear:       1900,
			wantIsLeapYear: false,
		},
		{
			name:           "GivenFourhundredthYearRequested_WhenIsLeapYearCalled_Then29IsReturned",
			giveYear:       2000,
			wantIsLeapYear: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotIsLeapYear := isLeapYear(test.giveYear)

			assert.Equal(t, test.wantIsLeapYear, gotIsLeapYear)
		})
	}
}

func TestLengthOfMonth(t *testing.T) {
	tests := []struct {
		name       string
		giveYM     [2]int
		wantLength int
	}{
		{
			name:       "GivenMonth01Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 1},
			wantLength: 31,
		},
		{
			name:       "GivenMonth02OfALeapYearRequested_WhenLengthOfMonthCalled_Then29IsReturned",
			giveYM:     [2]int{1996, 2},
			wantLength: 29,
		},
		{
			name:       "GivenMonth02OfANonLeapYearRequested_WhenLengthOfMonthCalled_Then28IsReturned",
			giveYM:     [2]int{1999, 2},
			wantLength: 28,
		},
		{
			name:       "GivenMonth03Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 3},
			wantLength: 31,
		},
		{
			name:       "GivenMonth04Requested_WhenLengthOfMonthCalled_Then30IsReturned",
			giveYM:     [2]int{1999, 4},
			wantLength: 30,
		},
		{
			name:       "GivenMonth05Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 5},
			wantLength: 31,
		},
		{
			name:       "GivenMonth06Requested_WhenLengthOfMonthCalled_Then30IsReturned",
			giveYM:     [2]int{1999, 6},
			wantLength: 30,
		},
		{
			name:       "GivenMonth07Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 7},
			wantLength: 31,
		},
		{
			name:       "GivenMonth08Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 8},
			wantLength: 31,
		},
		{
			name:       "GivenMonth09Requested_WhenLengthOfMonthCalled_Then30IsReturned",
			giveYM:     [2]int{1999, 9},
			wantLength: 30,
		},
		{
			name:       "GivenMonth10Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 10},
			wantLength: 31,
		},
		{
			name:       "GivenMonth11Requested_WhenLengthOfMonthCalled_Then30IsReturned",
			giveYM:     [2]int{1999, 11},
			wantLength: 30,
		},
		{
			name:       "GivenMonth12Requested_WhenLengthOfMonthCalled_Then31IsReturned",
			giveYM:     [2]int{1999, 12},
			wantLength: 31,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotLength := lengthOfMonth(test.giveYM[0], test.giveYM[1])

			assert.Equal(t, test.wantLength, gotLength)
		})
	}
}
