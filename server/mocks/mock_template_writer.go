// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server (interfaces: TemplateWriter)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	io "io"
	"reflect"
	"time"
)

type MockTemplateWriter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockTemplateWriter() *MockTemplateWriter {
	return &MockTemplateWriter{fail: pegomock.GlobalFailHandler}
}

func (mock *MockTemplateWriter) Execute(wr io.Writer, data interface{}) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockTemplateWriter().")
	}
	params := []pegomock.Param{wr, data}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Execute", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockTemplateWriter) VerifyWasCalledOnce() *VerifierTemplateWriter {
	return &VerifierTemplateWriter{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockTemplateWriter) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierTemplateWriter {
	return &VerifierTemplateWriter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockTemplateWriter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierTemplateWriter {
	return &VerifierTemplateWriter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockTemplateWriter) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierTemplateWriter {
	return &VerifierTemplateWriter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierTemplateWriter struct {
	mock                   *MockTemplateWriter
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierTemplateWriter) Execute(wr io.Writer, data interface{}) *TemplateWriter_Execute_OngoingVerification {
	params := []pegomock.Param{wr, data}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Execute", params, verifier.timeout)
	return &TemplateWriter_Execute_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type TemplateWriter_Execute_OngoingVerification struct {
	mock              *MockTemplateWriter
	methodInvocations []pegomock.MethodInvocation
}

func (c *TemplateWriter_Execute_OngoingVerification) GetCapturedArguments() (io.Writer, interface{}) {
	wr, data := c.GetAllCapturedArguments()
	return wr[len(wr)-1], data[len(data)-1]
}

func (c *TemplateWriter_Execute_OngoingVerification) GetAllCapturedArguments() (_param0 []io.Writer, _param1 []interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]io.Writer, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(io.Writer)
		}
		_param1 = make([]interface{}, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(interface{})
		}
	}
	return
}
