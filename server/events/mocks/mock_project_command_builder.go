// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: ProjectCommandBuilder)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockProjectCommandBuilder struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectCommandBuilder() *MockProjectCommandBuilder {
	return &MockProjectCommandBuilder{fail: pegomock.GlobalFailHandler}
}

func (mock *MockProjectCommandBuilder) BuildAutoplanCommands(ctx *events.CommandContext) ([]models.ProjectCommandContext, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandBuilder().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("BuildAutoplanCommands", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectCommandContext)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectCommandContext
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectCommandContext)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockProjectCommandBuilder) BuildPlanCommands(ctx *events.CommandContext, commentCommand *events.CommentCommand) ([]models.ProjectCommandContext, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandBuilder().")
	}
	params := []pegomock.Param{ctx, commentCommand}
	result := pegomock.GetGenericMockFrom(mock).Invoke("BuildPlanCommands", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectCommandContext)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectCommandContext
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectCommandContext)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockProjectCommandBuilder) BuildApplyCommands(ctx *events.CommandContext, commentCommand *events.CommentCommand) ([]models.ProjectCommandContext, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectCommandBuilder().")
	}
	params := []pegomock.Param{ctx, commentCommand}
	result := pegomock.GetGenericMockFrom(mock).Invoke("BuildApplyCommands", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectCommandContext)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectCommandContext
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectCommandContext)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockProjectCommandBuilder) VerifyWasCalledOnce() *VerifierProjectCommandBuilder {
	return &VerifierProjectCommandBuilder{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectCommandBuilder) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierProjectCommandBuilder {
	return &VerifierProjectCommandBuilder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectCommandBuilder) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierProjectCommandBuilder {
	return &VerifierProjectCommandBuilder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectCommandBuilder) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierProjectCommandBuilder {
	return &VerifierProjectCommandBuilder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierProjectCommandBuilder struct {
	mock                   *MockProjectCommandBuilder
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierProjectCommandBuilder) BuildAutoplanCommands(ctx *events.CommandContext) *ProjectCommandBuilder_BuildAutoplanCommands_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "BuildAutoplanCommands", params, verifier.timeout)
	return &ProjectCommandBuilder_BuildAutoplanCommands_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectCommandBuilder_BuildAutoplanCommands_OngoingVerification struct {
	mock              *MockProjectCommandBuilder
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectCommandBuilder_BuildAutoplanCommands_OngoingVerification) GetCapturedArguments() *events.CommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *ProjectCommandBuilder_BuildAutoplanCommands_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
	}
	return
}

func (verifier *VerifierProjectCommandBuilder) BuildPlanCommands(ctx *events.CommandContext, commentCommand *events.CommentCommand) *ProjectCommandBuilder_BuildPlanCommands_OngoingVerification {
	params := []pegomock.Param{ctx, commentCommand}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "BuildPlanCommands", params, verifier.timeout)
	return &ProjectCommandBuilder_BuildPlanCommands_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectCommandBuilder_BuildPlanCommands_OngoingVerification struct {
	mock              *MockProjectCommandBuilder
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectCommandBuilder_BuildPlanCommands_OngoingVerification) GetCapturedArguments() (*events.CommandContext, *events.CommentCommand) {
	ctx, commentCommand := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], commentCommand[len(commentCommand)-1]
}

func (c *ProjectCommandBuilder_BuildPlanCommands_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext, _param1 []*events.CommentCommand) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
		_param1 = make([]*events.CommentCommand, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*events.CommentCommand)
		}
	}
	return
}

func (verifier *VerifierProjectCommandBuilder) BuildApplyCommands(ctx *events.CommandContext, commentCommand *events.CommentCommand) *ProjectCommandBuilder_BuildApplyCommands_OngoingVerification {
	params := []pegomock.Param{ctx, commentCommand}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "BuildApplyCommands", params, verifier.timeout)
	return &ProjectCommandBuilder_BuildApplyCommands_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectCommandBuilder_BuildApplyCommands_OngoingVerification struct {
	mock              *MockProjectCommandBuilder
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectCommandBuilder_BuildApplyCommands_OngoingVerification) GetCapturedArguments() (*events.CommandContext, *events.CommentCommand) {
	ctx, commentCommand := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], commentCommand[len(commentCommand)-1]
}

func (c *ProjectCommandBuilder_BuildApplyCommands_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext, _param1 []*events.CommentCommand) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
		_param1 = make([]*events.CommentCommand, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*events.CommentCommand)
		}
	}
	return
}
