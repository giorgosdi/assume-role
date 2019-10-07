package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstractSessionAndArn(t *testing.T) {

	for name, tc := range map[string]struct {
		State    map[string]string
		Resp     []string
		Expected []string
	}{
		"testcase1": {
			State: map[string]string{
				"profile": "testprofile1",
				"account": "123456789",
				"role":    "testrole1",
			},
			Expected: []string{
				"arn:aws:iam::123456789:role/testrole1",
				"role-testprofile1-testrole1",
			},
		},
		"testcase2": {
			State: map[string]string{
				"profile": "testprofile2",
				"account": "987654321",
				"role":    "testrole2",
			},
			Expected: []string{
				"arn:aws:iam::987654321:role/testrole2",
				"role-testprofile2-testrole2",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if len(tc.Resp) == 0 {
				tc.Resp = make([]string, 2)
			}
			roleArn, roleSessionName := constractSessionAndArn(tc.State)
			tc.Resp = []string{
				roleArn,
				roleSessionName,
			}
			assert.EqualValues(t, tc.Resp, tc.Expected)
		})
	}
}
