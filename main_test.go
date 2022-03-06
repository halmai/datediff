package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockReaderResponse struct {
	line string
	err  error
}

type mockReader struct {
	invocationIndex *int
	responses       []mockReaderResponse
}

type mockOutputCollector struct{}

func (moc mockOutputCollector) Printf(string) {}

func (m mockReader) Read(p []byte) (n int, err error) {
	if errX := m.responses[*m.invocationIndex].err; errX != nil {
		return 0, errX
	}

	line := m.responses[*m.invocationIndex].line + "\n"

	*m.invocationIndex++

	l := copy(p, line)

	return l, nil
}

func TestProcessDates(t *testing.T) {
	tests := []struct {
		name                string
		mockReaderResponses []mockReaderResponse
		wantDiffDays        int
		wantErrStr          string
	}{
		{
			name: "GivenReadingFirstDateFails_WhenProcessDatesCalled_ThenFailureIsReturned",
			mockReaderResponses: []mockReaderResponse{
				{
					err: errors.New("failure in Reader"),
				},
			},
			wantErrStr: "failure in Reader",
		},
		{
			name: "GivenReadingSecondDateFails_WhenProcessDatesCalled_ThenFailureIsReturned",
			mockReaderResponses: []mockReaderResponse{
				{
					line: "31/12/1999",
				},
				{
					err: errors.New("failure in Reader"),
				},
			},
			wantErrStr: "failure in Reader",
		},
		{
			name: "GivenFirstDateIsInvalid_WhenProcessDatesCalled_ThenErrorIsReturned",
			mockReaderResponses: []mockReaderResponse{
				{
					line: "2/March/2022",
				},
			},
			wantErrStr: "date is not in format DD/MM/YYYY",
		},
		{
			name: "GivenSecondDateIsInvalid_WhenProcessDatesCalled_ThenErrorIsReturned",
			mockReaderResponses: []mockReaderResponse{
				{
					line: "2/1/2022",
				},
				{
					line: "2/March/2022",
				},
			},
			wantErrStr: "date is not in format DD/MM/YYYY",
		},
		{
			name: "GivenHappyPathDatesProvided_WhenProcessDatesCalled_ThenCorrectIntIsReturned",
			mockReaderResponses: []mockReaderResponse{
				{
					line: "2/6/1983",
				},
				{
					line: "22/6/1983",
				},
			},
			wantDiffDays: 19,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var counter int

			var moc mockOutputCollector

			gotDiffDays, gotErr := processDates(mockReader{
				invocationIndex: &counter,
				responses:       test.mockReaderResponses,
			}, moc)

			if test.wantErrStr == "" {
				assert.Nil(t, gotErr)
				assert.Equal(t, test.wantDiffDays, gotDiffDays)
			} else {
				assert.Equal(t, test.wantErrStr, gotErr.Error())
				assert.Equal(t, 0, gotDiffDays)
			}
		})
	}
}
