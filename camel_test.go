package strcaseconv

import (
	"testing"
)

func TestSnakeToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"default", "foo_bar", "FooBar"},
		{"long", "foo_bar_lorem_ipsum", "FooBarLoremIpsum"},
		{"without delimiter", "foo", "Foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToUpperCamel(tt.arg); got != tt.want {
				t.Errorf("SnakeToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"default", "foo_bar", "fooBar"},
		{"long", "foo_bar_lorem_ipsum", "fooBarLoremIpsum"},
		{"without delimiter", "foo", "foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToLowerCamel(tt.arg); got != tt.want {
				t.Errorf("SnakeToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKebabToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"default", "foo-bar", "FooBar"},
		{"long", "foo-bar-lorem-ipsum", "FooBarLoremIpsum"},
		{"without delimiter", "foo", "Foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabToUpperCamel(tt.arg); got != tt.want {
				t.Errorf("KebabToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKebabToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"default", "foo-bar", "fooBar"},
		{"long", "foo-bar-lorem-ipsum", "fooBarLoremIpsum"},
		{"without delimiter", "foo", "foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabToLowerCamel(tt.arg); got != tt.want {
				t.Errorf("KebabToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

var snakeToUpperCamelResult string

func BenchmarkSnakeToUpperCamel(b *testing.B) {
	original := "foo_bar_lorem_ipsum"
	for i := 0; i < b.N; i++ {
		snakeToUpperCamelResult = SnakeToUpperCamel(original)
	}
}
