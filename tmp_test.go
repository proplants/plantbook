package main

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"ok", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := hello(); (err != nil) != tt.wantErr {
				t.Errorf("hello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
