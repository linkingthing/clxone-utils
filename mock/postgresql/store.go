package postgresql

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	restdb "github.com/linkingthing/gorest/db"
	restresource "github.com/linkingthing/gorest/resource"
)

func (m *RStoreMocker) Clean() {}

func (m *RStoreMocker) Close() {}

func (m *RStoreMocker) Begin() (restdb.Transaction, error) {
	return m, nil
}

func (m *RStoreMocker) Commit() error {
	return nil
}

func (m *RStoreMocker) Rollback() error {
	return nil
}

func (m *RStoreMocker) Insert(r restresource.Resource) (restresource.Resource, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			return r, msg.err
		} else if msg.expectType == expectTypeError {
			return nil, msg.err
		}

		return nil, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return nil, err
	}
}

func (m *RStoreMocker) Get(typ restdb.ResourceType, cond map[string]interface{}) (interface{}, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			return msg.result, msg.err
		} else if msg.expectType == expectTypeError {
			return nil, msg.err
		}

		return nil, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return nil, err
	}
}

func (m *RStoreMocker) GetOwned(owner restdb.ResourceType, ownerID string, owned restdb.ResourceType) (interface{}, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			return msg.result, msg.err
		} else if msg.expectType == expectTypeError {
			return nil, msg.err
		}

		return nil, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return nil, err
	}
}

func (m *RStoreMocker) Exists(typ restdb.ResourceType, cond map[string]interface{}) (bool, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeBool {
			return msg.result.(bool), msg.err
		} else if msg.expectType == expectTypeError {
			return false, msg.err
		}

		return false, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return false, err
	}
}

func (m *RStoreMocker) Count(typ restdb.ResourceType, cond map[string]interface{}) (int64, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeInt64 {
			return msg.result.(int64), msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func (m *RStoreMocker) Fill(cond map[string]interface{}, out interface{}) error {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			if msg.err != nil {
				return msg.err
			} else {
				return expectationToResources(msg.result, out)
			}
		} else if msg.expectType == expectTypeError {
			return msg.err
		}

		return fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return err
	}
}

func (m *RStoreMocker) Delete(typ restdb.ResourceType, cond map[string]interface{}) (int64, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeInt64 {
			return msg.result.(int64), msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func (m *RStoreMocker) Update(typ restdb.ResourceType, nv map[string]interface{}, cond map[string]interface{}) (int64, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeInt64 {
			return msg.result.(int64), msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func (m *RStoreMocker) FillOwned(owner restdb.ResourceType, ownerID string, out interface{}) error {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			if msg.err != nil {
				return msg.err
			}
			return expectationToResources(msg.result, out)
		} else if msg.expectType == expectTypeError {
			return msg.err
		}

		return fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return err
	}
}

func (m *RStoreMocker) GetEx(typ restdb.ResourceType, sql string, params ...interface{}) (interface{}, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			return msg.result, msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func (m *RStoreMocker) CountEx(typ restdb.ResourceType, sql string, params ...interface{}) (int64, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeInt64 {
			return msg.result.(int64), msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func (m *RStoreMocker) FillEx(out interface{}, sql string, params ...interface{}) error {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeResources {
			if msg.err != nil {
				return msg.err
			}
			return expectationToResources(msg.result, out)
		} else if msg.expectType == expectTypeError {
			return msg.err
		}

		return fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return err
	}
}

func (m *RStoreMocker) Exec(sql string, params ...interface{}) (int64, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeInt64 {
			return msg.result.(int64), msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func (m *RStoreMocker) CopyFromEx(typ restdb.ResourceType, columns []string, values [][]interface{}) (int64, error) {
	name := getMethodName()
	if msg, err := m.Mocker.AssertExpectation(name); err == nil {
		if msg.expectType == expectTypeInt64 {
			return msg.result.(int64), msg.err
		} else if msg.expectType == expectTypeError {
			return 0, msg.err
		}

		return 0, fmt.Errorf("not found expectation return type:%s in %s", msg.expectType, name)
	} else {
		return 0, err
	}
}

func getMethodName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	f := runtime.FuncForPC(pc)
	funcs := strings.Split(f.Name(), ".")
	var name string
	if len(funcs) > 0 {
		name = funcs[len(funcs)-1]
	}

	return name
}

func expectationToResources(in interface{}, out interface{}) error {
	outSlice := reflect.Indirect(reflect.ValueOf(out))
	inValue := reflect.Indirect(reflect.ValueOf(in))
	if inValue.Type() != outSlice.Type() {
		return fmt.Errorf("expect type %s not match out type %s", inValue.Type(), outSlice.Type())
	}

	for i := 0; i < inValue.Len(); i++ {
		outSlice.Set(reflect.Append(outSlice, inValue.Index(i)))
	}

	return nil
}
