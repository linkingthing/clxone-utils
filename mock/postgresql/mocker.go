package postgresql

import (
	"fmt"

	restdb "github.com/linkingthing/gorest/db"
)

type RStoreMocker struct {
	*Mocker
}

type Mocker struct {
	expectMsgs []*expectMsg
}

type expectMsg struct {
	method     ExpectMethod
	result     interface{}
	err        error
	expectType expectType
}

type expectType string

const (
	expectTypeError     expectType = "error"
	expectTypeBool      expectType = "bool"
	expectTypeInt64     expectType = "int64"
	expectTypeResources expectType = "resources"
)

type ExpectMethod string

const (
	ExpectMethodInsert     ExpectMethod = "Insert"
	ExpectMethodFill       ExpectMethod = "Fill"
	ExpectMethodGet        ExpectMethod = "Get"
	ExpectMethodGetOwned   ExpectMethod = "GetOwned"
	ExpectMethodExists     ExpectMethod = "Exists"
	ExpectMethodCount      ExpectMethod = "Count"
	ExpectMethodDelete     ExpectMethod = "Delete"
	ExpectMethodUpdate     ExpectMethod = "Update"
	ExpectMethodFillOwned  ExpectMethod = "FillOwned"
	ExpectMethodGetEx      ExpectMethod = "GetEx"
	ExpectMethodCountEx    ExpectMethod = "CountEx"
	ExpectMethodFillEx     ExpectMethod = "FillEx"
	ExpectMethodExec       ExpectMethod = "Exec"
	ExpectMethodCopyFromEx ExpectMethod = "CopyFromEx"
	ExpectMethodCopyFrom 	ExpectMethod = "CopyFrom"
	ExpectMethodBegin      ExpectMethod = "Begin"
	ExpectMethodCommit     ExpectMethod = "Commit"
	ExpectMethodRollback   ExpectMethod = "Rollback"
)

func NewMocker() (restdb.ResourceStore, *Mocker, error) {
	m := &Mocker{}
	return &RStoreMocker{m}, m, nil
}

func (m *Mocker) ExpectExec(method ExpectMethod) *expectMsg {
	msg := &expectMsg{method: method}
	m.expectMsgs = append(m.expectMsgs, msg)
	return msg
}

func (m *Mocker) AssertExpectation(method string) (*expectMsg, error) {
	if len(m.expectMsgs) > 0 {
		if msg := m.expectMsgs[0]; string(msg.method) == method {
			m.expectMsgs = m.expectMsgs[1:]
			return msg, nil
		}
	}

	return nil, fmt.Errorf("not found expectation method: %s", method)
}

func (m *Mocker) Clean() {
	m.expectMsgs = m.expectMsgs[:0]
}

func (e *expectMsg) ReturnError(err error) *expectMsg {
	if e.expectType == "" {
		e.expectType = expectTypeError
	}
	e.err = err
	return e
}

func (e *expectMsg) ReturnResources(rs interface{}) *expectMsg {
	e.result = rs
	e.expectType = expectTypeResources
	return e
}

func (e *expectMsg) ReturnBool(ok bool) *expectMsg {
	e.result = ok
	e.expectType = expectTypeBool
	return e
}

func (e *expectMsg) ReturnInt64(value int64) *expectMsg {
	e.result = value
	e.expectType = expectTypeInt64
	return e
}
