package kit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlice(t *testing.T) {
	t.Run("no error chain", func(t *testing.T) {
		a := New(
			func() error { return nil },
			func() error { return nil },
			func() error { return nil },
		)
		err := a.Error()
		assert.NoError(t, err)
	})

	t.Run("error in middle stops execution", func(t *testing.T) {
		expectedError := errors.New("an error")
		executionOrder := []int{}

		a := New(
			func() error { executionOrder = append(executionOrder, 1); return nil },
			func() error { executionOrder = append(executionOrder, 2); return expectedError },
			func() error { executionOrder = append(executionOrder, 3); return nil },
		)

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, []int{1, 2}, executionOrder) // Third function should not execute
	})

	t.Run("first error stops execution", func(t *testing.T) {
		expectedError := errors.New("first error")
		var executed bool

		a := New(
			func() error { return expectedError },
			func() error { executed = true; return nil },
		)

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.False(t, executed, "second function should not execute")
	})

	t.Run("empty chain", func(t *testing.T) {
		a := New()
		err := a.Error()
		assert.NoError(t, err)
	})

	t.Run("single function success", func(t *testing.T) {
		called := false
		a := New(func() error {
			called = true
			return nil
		})

		err := a.Error()
		assert.NoError(t, err)
		assert.True(t, called)
	})

	t.Run("single function error", func(t *testing.T) {
		expectedError := errors.New("single error")
		a := New(func() error {
			return expectedError
		})

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestAliceThen(t *testing.T) {
	t.Run("chained Then calls success", func(t *testing.T) {
		executionOrder := []int{}

		a := NewAlice()
		a.Then(func() error { executionOrder = append(executionOrder, 1); return nil }).
			Then(func() error { executionOrder = append(executionOrder, 2); return nil }).
			Then(func() error { executionOrder = append(executionOrder, 3); return nil })

		err := a.Error()
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, executionOrder)
	})

	t.Run("chained Then calls with error", func(t *testing.T) {
		expectedError := errors.New("an error")
		executionOrder := []int{}

		a := NewAlice()
		a.Then(func() error { executionOrder = append(executionOrder, 1); return nil }).
			Then(func() error { executionOrder = append(executionOrder, 2); return expectedError }).
			Then(func() error { executionOrder = append(executionOrder, 3); return nil })

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, []int{1, 2}, executionOrder)
	})

	t.Run("Then on already errored Alice", func(t *testing.T) {
		firstError := errors.New("first error")
		var shouldNotExecute bool

		a := NewAlice()
		a.Then(func() error { return firstError }).
			Then(func() error { shouldNotExecute = true; return nil })

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, firstError, err)
		assert.False(t, shouldNotExecute)
	})

	t.Run("Then returns same Alice instance", func(t *testing.T) {
		a := NewAlice()
		result := a.Then(func() error { return nil })
		assert.Equal(t, a, result, "Then should return the same Alice instance")
	})

	t.Run("calling Then multiple times after error", func(t *testing.T) {
		firstError := errors.New("first error")
		executionCount := 0

		a := NewAlice()
		a.Then(func() error { return firstError })

		// All subsequent Then calls should be no-ops
		a.Then(func() error { executionCount++; return nil })
		a.Then(func() error { executionCount++; return nil })
		a.Then(func() error { executionCount++; return nil })

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, firstError, err)
		assert.Equal(t, 0, executionCount)
	})
}

func TestAliceEdgeCases(t *testing.T) {
	t.Run("nil function in New", func(t *testing.T) {
		// Note: This would panic in real usage, but we test the expected behavior
		// In practice, this test documents the behavior rather than testing safety
		assert.Panics(t, func() {
			New(nil)
		})
	})

	t.Run("nil function in Then", func(t *testing.T) {
		a := NewAlice()
		assert.Panics(t, func() {
			a.Then(nil)
		})
	})

	t.Run("different error types", func(t *testing.T) {
		customError := &Exception{code: ErrNotFound, info: "not found"}

		a := New(func() error { return customError })
		err := a.Error()

		assert.Error(t, err)
		assert.Equal(t, customError, err)
		assert.IsType(t, &Exception{}, err)
	})

	t.Run("panic in function", func(t *testing.T) {
		assert.Panics(t, func() {
			a := New(func() error {
				panic("test panic")
			})
			a.Error()
		})
	})
}

func TestAliceUseCases(t *testing.T) {
	const insertedData = "inserted"
	
	t.Run("database transaction simulation", func(t *testing.T) {
		var (
			connected   bool
			transaction bool
			data        string
			committed   bool
		)

		a := New(
			func() error { // Connect to database
				connected = true
				return nil
			},
			func() error { // Begin transaction
				if !connected {
					return errors.New("not connected")
				}
				transaction = true
				return nil
			},
			func() error { // Insert data
				if !transaction {
					return errors.New("no transaction")
				}
				data = insertedData
				return nil
			},
			func() error { // Commit transaction
				if data != insertedData {
					return errors.New("no data to commit")
				}
				committed = true
				return nil
			},
		)

		err := a.Error()
		assert.NoError(t, err)
		assert.True(t, connected)
		assert.True(t, transaction)
		assert.Equal(t, insertedData, data)
		assert.True(t, committed)
	})

	t.Run("validation chain", func(t *testing.T) {
		type User struct {
			Name  string
			Email string
			Age   int
		}

		user := User{Name: "John", Email: "john@example.com", Age: 25}

		a := New(
			func() error { // Validate name
				if user.Name == "" {
					return errors.New("name is required")
				}
				return nil
			},
			func() error { // Validate email
				if user.Email == "" {
					return errors.New("email is required")
				}
				return nil
			},
			func() error { // Validate age
				if user.Age < 18 {
					return errors.New("age must be at least 18")
				}
				return nil
			},
		)

		err := a.Error()
		assert.NoError(t, err)
	})

	t.Run("validation chain with error", func(t *testing.T) {
		type User struct {
			Name  string
			Email string
			Age   int
		}

		user := User{Name: "John", Email: "", Age: 25} // Missing email

		a := New(
			func() error { // Validate name
				if user.Name == "" {
					return errors.New("name is required")
				}
				return nil
			},
			func() error { // Validate email
				if user.Email == "" {
					return errors.New("email is required")
				}
				return nil
			},
			func() error { // Validate age (should not be called)
				if user.Age < 18 {
					return errors.New("age must be at least 18")
				}
				return nil
			},
		)

		err := a.Error()
		assert.Error(t, err)
		assert.Equal(t, "email is required", err.Error())
	})
}

func TestAliceComparison(t *testing.T) {
	t.Run("New vs NewAlice with Then", func(t *testing.T) {
		expectedError := errors.New("test error")

		// Using New
		a1 := New(
			func() error { return nil },
			func() error { return expectedError },
		)

		// Using NewAlice with Then
		a2 := NewAlice()
		a2.Then(func() error { return nil }).
			Then(func() error { return expectedError })

		err1 := a1.Error()
		err2 := a2.Error()

		assert.Equal(t, err1, err2)
		assert.Equal(t, expectedError, err1)
		assert.Equal(t, expectedError, err2)
	})
}
