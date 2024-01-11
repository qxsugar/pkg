package alice

import (
	"errors"
	"testing"
)

func TestAlice(t *testing.T) {
	var err error

	// Test case where no error is expected
	a := New(
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
	)
	err = a.Error()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test case where an error is expected
	expectedError := errors.New("an error")
	a = New(
		func() error { return nil },
		func() error { return expectedError },
		func() error { return nil },
	)
	err = a.Error()
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	// Test case where an error is expected and should stop the execution of the next function
	var executed bool
	a = New(
		func() error { return expectedError },
		func() error { executed = true; return nil },
	)
	err = a.Error()
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}
	if executed {
		t.Errorf("expected function not to be executed")
	}
}

func TestAliceThen(t *testing.T) {
	var err error
	expectedError := errors.New("an error")

	// Test case where no error is expected
	a := NewAlice()
	a.Then(func() error { return nil }).
		Then(func() error { return nil }).
		Then(func() error { return nil })
	err = a.Error()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test case where an error is expected
	a = NewAlice()
	a.Then(func() error { return nil }).
		Then(func() error { return expectedError }).
		Then(func() error { return nil })
	err = a.Error()
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}

	// Test case where an error is expected and should stop the execution of the next function
	var executed bool
	a = NewAlice()
	a.Then(func() error { return expectedError }).
		Then(func() error { executed = true; return nil })
	err = a.Error()
	if err != expectedError {
		t.Errorf("expected %v, got %v", expectedError, err)
	}
	if executed {
		t.Errorf("expected function not to be executed")
	}
}
