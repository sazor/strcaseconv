package strcaseconv

import "testing"

func TestSnakeToKebab(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"default", "foo_bar", "foo-bar"},
		{"long", "foo_bar_lorem_ipsum", "foo-bar-lorem-ipsum"},
		{"without delimiter", "foo", "foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToKebab(tt.arg); got != tt.want {
				t.Errorf("SnakeToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelToKebab(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"lower camel", "fooBar", "foo-bar"},
		{"upper camel", "FooBar", "foo-bar"},
		{"acronym", "FooBarID", "foo-bar-id"},
		{"acronym", "FOOBar", "foo-bar"},
		{"long", "fooBarLoremIpsum", "foo-bar-lorem-ipsum"},
		{"lower camel word", "foo", "foo"},
		{"upper camel word", "Foo", "foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToKebab(tt.arg); got != tt.want {
				t.Errorf("CamelToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

var snakeToKebabResult string

func BenchmarkSnakeToKebab(b *testing.B) {
	original := "foo_bar_lorem_ipsum"
	for i := 0; i < b.N; i++ {
		snakeToKebabResult = SnakeToKebab(original)
	}
}

var camelToKebabResult string

func BenchmarkCamelToKebab(b *testing.B) {
	original := "FooBarLoremIpsum"
	for i := 0; i < b.N; i++ {
		camelToKebabResult = CamelToKebab(original)
	}
}
