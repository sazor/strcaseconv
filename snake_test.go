package strcaseconv

import "testing"

func TestKebabToSnake(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"default", "foo-bar", "foo_bar"},
		{"long", "foo-bar-lorem-ipsum", "foo_bar_lorem_ipsum"},
		{"without delimiter", "foo", "foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabToSnake(tt.arg); got != tt.want {
				t.Errorf("KebabToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"lower camel", "fooBar", "foo_bar"},
		{"upper camel", "FooBar", "foo_bar"},
		{"acronym", "FooBarID", "foo_bar_id"},
		{"acronym", "FOOBar", "foo_bar"},
		{"long", "fooBarLoremIpsum", "foo_bar_lorem_ipsum"},
		{"lower camel word", "foo", "foo"},
		{"upper camel word", "Foo", "foo"},
		{"blank", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToSnake(tt.arg); got != tt.want {
				t.Errorf("CamelToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

var kebabToSnakeResult string

func BenchmarkKebabToSnake(b *testing.B) {
	original := "foo-bar-lorem-ipsum"
	for i := 0; i < b.N; i++ {
		kebabToSnakeResult = KebabToSnake(original)
	}
}

var camelToSnakeResult string

func BenchmarkCamelToSnake(b *testing.B) {
	original := "fooBarLoremIpsum"
	for i := 0; i < b.N; i++ {
		camelToSnakeResult = CamelToSnake(original)
	}
}
