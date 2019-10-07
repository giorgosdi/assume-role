package api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func createSession(profile string) (*sts.STS, error) {
	stsSession, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
	})
	if err != nil {
		return nil, err
	}
	stsSess := sts.New(stsSession)
	return stsSess, nil
}

func constractSessionAndArn(state map[string]string) (string, string) {
	roleArn := fmt.Sprintf("arn:aws:iam::%s:role/%s", state["account"], state["role"])
	roleSessionName := fmt.Sprintf("role-%s-%s", state["profile"], state["role"])
	return roleArn, roleSessionName
}

func AssumeRole(state map[string]string) (*sts.AssumeRoleOutput, error) {

	stsSess, err := createSession(state["profile"])
	if err != nil {
		return nil, err
	}
	roleArn, roleSessionName := constractSessionAndArn(state)

	role := &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(roleSessionName),
	}
	resp, err := stsSess.AssumeRole(role)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
