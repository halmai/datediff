package datecalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		name       string
		sendString string
		wantYMD    [3]int
		wantErrStr string
	}{
		{
			name:       "GivenStringWithNotThreeParts_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "this/is/not/a/three-component/date/string",
			wantErrStr: "date is not in format DD/MM/YYYY",
		},
		{
			name:       "GivenStringWithNotNumbers_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "this/is/nonnumerical",
			wantErrStr: "date is not in format DD/MM/YYYY",
		},
		{
			name:       "GivenTooSmallYear_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "1/2/1899",
			wantErrStr: "year part of date cannot be parsed: value \"1899\" is out of range [1900,2999]",
		},
		{
			name:       "GivenTooLargeYear_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "1/2/3000",
			wantErrStr: "year part of date cannot be parsed: value \"3000\" is out of range [1900,2999]",
		},
		{
			name:       "GivenTooSmallMonth_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "1/0/2022",
			wantErrStr: "month part of date cannot be parsed: value \"0\" is out of range [1,12]",
		},
		{
			name:       "GivenTooLargeMonth_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "1/13/2022",
			wantErrStr: "month part of date cannot be parsed: value \"13\" is out of range [1,12]",
		},
		{
			name:       "GivenTooSmallDay_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "0/12/2022",
			wantErrStr: "day part of date cannot be parsed: value \"0\" is out of range [1,31]",
		},
		{
			name:       "GivenTooLargeDay_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "32/12/2022",
			wantErrStr: "day part of date cannot be parsed: value \"32\" is out of range [1,31]",
		},
		{
			name:       "GivenTooLargeDay_WhenParseDateCalled_ThenErrorIsReturned",
			sendString: "31/12/2022",
			wantYMD:    [3]int{2022, 12, 31},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			y, m, d, gotErr := ParseDate(test.sendString)

			if test.wantErrStr == "" {
				assert.Nil(t, gotErr)
				assert.Equal(t, test.wantYMD, [3]int{y, m, d})
			} else {
				assert.Equal(t, test.wantErrStr, gotErr.Error())
				assert.Equal(t, [3]int{0, 0, 0}, [3]int{y, m, d})
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name       string
		sendString string
		sendMin    int
		sendMax    int
		wantErrStr string
		wantInt    int
	}{
		{
			name:       "GivenNumberHas3DigitsAndMinIs1AndMaxIs1000_WhenParseIntCalled_ThenCorrectValueIsReturned",
			sendString: "123",
			sendMin:    1,
			sendMax:    1000,
			wantInt:    123,
		},
		{
			name:       "GivenNumberHas3DigitsAndMinAndMaxAreTheSame_WhenParseIntCalled_ThenCorrectValueIsReturned",
			sendString: "123",
			sendMin:    123,
			sendMax:    123,
			wantInt:    123,
		},
		{
			name:       "GivenNumberHas3DigitsAndMinAndMaxAreAboveIt_WhenParseIntCalled_ThenErrorIsReturned",
			sendString: "123",
			sendMin:    888,
			sendMax:    999,
			wantErrStr: "value \"123\" is out of range [888,999]",
		},
		{
			name:       "GivenNumberHas3DigitsAndMinAndMaxAreUnderIt_WhenParseIntCalled_ThenErrorIsReturned",
			sendString: "123",
			sendMin:    111,
			sendMax:    122,
			wantErrStr: "value \"123\" is out of range [111,122]",
		},
		{
			name:       "GivenNumberIsNegativeAndMinAndMaxSurroundIt_WhenParseIntCalled_ThenCorrectValueIsReturned",
			sendString: "-123",
			sendMin:    -333,
			sendMax:    1,
			wantInt:    -123,
		},
		{
			name:       "GivenNumberIsNegativeAndMinAndMaxArePositive_WhenParseIntCalled_ThenErrorIsReturned",
			sendString: "-123",
			sendMin:    0,
			sendMax:    100,
			wantErrStr: "value \"-123\" is out of range [0,100]",
		},
		{
			name:       "GivenNumberContainsNonNumericalCharacter_WhenParseIntCalled_ThenErrorIsReturned",
			sendString: "This is not a number",
			sendMin:    0,
			sendMax:    100,
			wantErrStr: "parsing error: strconv.ParseInt: parsing \"This is not a number\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotInt, gotErr := parseInt(test.sendString, test.sendMin, test.sendMax)

			if test.wantErrStr != "" {
				assert.Equal(t, 0, gotInt)
				assert.Equal(t, test.wantErrStr, gotErr.Error())
			} else {
				assert.Nil(t, gotErr)
				assert.Equal(t, test.wantInt, gotInt)
			}
		})
	}
}
