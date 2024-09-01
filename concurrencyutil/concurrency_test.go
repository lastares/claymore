package concurrencyutil

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func mockFunc() error {
	return nil
}

func mockFuncDo1() error {
	fmt.Println("mockFuncDo1")
	return nil
}

func mockFuncDo2() error {
	fmt.Println("mockFuncDo2")
	return nil
}

func mockFuncDo3() error {
	fmt.Println("mockFuncDo3")
	return nil
}

func mockFuncDo4() error {
	fmt.Println("mockFuncDo4")
	return nil
}

func mockFuncDo5() error {
	fmt.Println("mockFuncDo5")
	return nil
}

func mockFuncDoWithError() error {
	fmt.Println("mockFuncDo3")
	return errors.New("mockFuncDo3 failed")
}

// Mock function to simulate error
func mockFuncWithError() error {
	return errors.New("mock error")
}

// TestNewWg_Success tests if NewWg successfully executes a slice of functions without error
func TestNewWg_Success(t *testing.T) {
	// Create a slice of mock functions
	funcs := []func() error{mockFunc, mockFunc}

	// Call NewWg with the mock functions
	err := NewWg(funcs)
	if err != nil {
		t.Errorf("NewWg() error = %v, wantErr %v", err, false)
	}
}

func TestNewWg_Error(t *testing.T) {
	// Create a slice of mock functions, one of which returns an error
	funcs := []func() error{mockFunc, mockFuncWithError}
	err := NewWg(funcs)
	if err != nil {
		t.Logf("NewWg() error != nil, wantErr true")
	}
}

// TestNewWgWithContext tests if NewWg can be called with a context
func TestNewWgWithContext(t *testing.T) {
	// Create a slice of mock functions
	funcs := []func() error{mockFunc, mockFunc}

	// Create a background context
	ctx := context.Background()

	// Call NewWg with the mock functions and a context option
	err := NewWg(funcs, WithContext(ctx))
	if err != nil {
		t.Errorf("NewWgWithContext() error = %v, wantErr %v", err, false)
	}
}

func TestNewWgDoSomething(t *testing.T) {
	// Create a slice of mock functions
	funcs := []func() error{mockFuncDo1, mockFuncDo2, mockFuncDo3, mockFuncDo4, mockFuncDo5}

	// Create a background context
	ctx := context.Background()

	// Call NewWg with the mock functions and a context option
	err := NewWg(funcs, WithContext(ctx))
	if err != nil {
		t.Errorf("NewWgWithContext() error = %v, wantErr %v", err, false)
	}
}

func TestNewWgDoSomethingWithLimit(t *testing.T) {
	// Create a slice of mock functions
	funcs := []func() error{mockFuncDo1, mockFuncDo2}

	// Create a background context
	ctx := context.Background()
	limit := 2
	// Call NewWg with the mock functions and a context option
	err := NewWg(funcs, WithContext(ctx), WithLimit(limit))
	if err != nil {
		t.Errorf("NewWgWithContext() error = %v, wantErr %v", err, false)
	}
}

func TestNewWgDoSomethingWithError(t *testing.T) {
	// Create a slice of mock functions
	funcs := []func() error{mockFuncDo1, mockFuncDo2, mockFuncDoWithError}

	// Create a background context
	ctx := context.Background()

	// Call NewWg with the mock functions and a context option
	err := NewWg(funcs, WithContext(ctx), WithLimit(2))
	if err != nil {
		t.Logf("NewWgWithContext() error = %v, wantErr %v", err, true)
	}
}
