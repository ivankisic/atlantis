// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: GithubPullGetter)

package mocks

import (
	github "github.com/google/go-github/github"
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockGithubPullGetter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGithubPullGetter() *MockGithubPullGetter {
	return &MockGithubPullGetter{fail: pegomock.GlobalFailHandler}
}

func (mock *MockGithubPullGetter) GetPullRequest(repo models.Repo, pullNum int) (*github.PullRequest, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockGithubPullGetter().")
	}
	params := []pegomock.Param{repo, pullNum}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullRequest", params, []reflect.Type{reflect.TypeOf((**github.PullRequest)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *github.PullRequest
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*github.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGithubPullGetter) VerifyWasCalledOnce() *VerifierGithubPullGetter {
	return &VerifierGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockGithubPullGetter) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierGithubPullGetter {
	return &VerifierGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockGithubPullGetter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierGithubPullGetter {
	return &VerifierGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockGithubPullGetter) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierGithubPullGetter {
	return &VerifierGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierGithubPullGetter struct {
	mock                   *MockGithubPullGetter
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierGithubPullGetter) GetPullRequest(repo models.Repo, pullNum int) *GithubPullGetter_GetPullRequest_OngoingVerification {
	params := []pegomock.Param{repo, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullRequest", params, verifier.timeout)
	return &GithubPullGetter_GetPullRequest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type GithubPullGetter_GetPullRequest_OngoingVerification struct {
	mock              *MockGithubPullGetter
	methodInvocations []pegomock.MethodInvocation
}

func (c *GithubPullGetter_GetPullRequest_OngoingVerification) GetCapturedArguments() (models.Repo, int) {
	repo, pullNum := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pullNum[len(pullNum)-1]
}

func (c *GithubPullGetter_GetPullRequest_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}
