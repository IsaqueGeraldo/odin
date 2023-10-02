package odin

import (
	"reflect"
	"testing"
)

func TestBootstrap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Bootstrap()
		})
	}
}

func TestGetenv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Getenv(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getenv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getenv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []Environment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetenv(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setenv(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Setenv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnsetenv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unsetenv(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Unsetenv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnviron(t *testing.T) {
	tests := []struct {
		name    string
		want    []Environment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Environ()
			if (err != nil) != tt.wantErr {
				t.Errorf("Environ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Environ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearenv(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Clearenv(); (err != nil) != tt.wantErr {
				t.Errorf("Clearenv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRenameKey(t *testing.T) {
	type args struct {
		oldKey string
		newKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RenameKey(tt.args.oldKey, tt.args.newKey); (err != nil) != tt.wantErr {
				t.Errorf("RenameKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sanitizeKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeKey(tt.args.key); got != tt.want {
				t.Errorf("sanitizeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
