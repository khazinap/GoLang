package main

import (
	"errors"
	"testing"
)

func TestSuccess(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{name: "test1", tokens: []string{"2", "3", "+"}, want: 5},
		{name: "test2", tokens: []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"}, want: 14},
		{name: "test3", tokens: []string{"42"}, want: 42},
		{name: "test4", tokens: []string{"10", "4", "-"}, want: 6},
		{name: "test5", tokens: []string{"6", "7", "*"}, want: 42},
		{name: "test6", tokens: []string{"9", "2", "/"}, want: 4},
		{name: "test7", tokens: []string{"-3", "5", "+"}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalPostfix(tt.tokens)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("there is no expected value %d, got %d", tt.want, got)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []string
		wantErr error
	}{
		{name: "division by zero", tokens: []string{"1", "0", "/"}, wantErr: ErrDivByZero},
		{name: "test1", tokens: []string{"1", "+"}},
		{name: "test2", tokens: []string{"2", "3", "&"}},
		{name: "test3", tokens: []string{}},
		{name: "test4", tokens: []string{"1", "2", "3", "+"}},
		{name: "test5", tokens: []string{"+"}},
		{name: "test6", tokens: []string{"2", "3.5", "+"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalPostfix(tt.tokens)
			if err == nil {
				t.Fatalf("there is no error")
			}
			if got != 0 {
				t.Errorf("expected 0 on error")
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("there is no expected eeror %v, got %v", tt.wantErr, err)
			}
		})
	}
}
