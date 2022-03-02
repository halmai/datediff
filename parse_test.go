package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
