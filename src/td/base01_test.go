package td

import (
	"slices"
	"testing"
)

// 基本测试
func TestSplit00(t *testing.T) {
	data := "a-b-c"
	got := Split(data, "-")
	want := []string{"a", "b", "c"}

	if !slices.Equal[[]string](got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSplit02(t *testing.T) {
	data := "abcdef"
	got := Split(data, "cd")
	want := []string{"ab", "ef"}

	if !slices.Equal[[]string](got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSplit03(t *testing.T) {
	data := "你-好-呀-ab-cd"
	got := Split(data, "-")
	want := []string{"你", "好", "呀", "ab", "cd"}

	if !slices.Equal[[]string](got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// 子测试
func TestSplit04(t *testing.T) {
	t.Run("test01", func(t *testing.T) {
		data := "你-好-呀-ab-cd"
		got := Split(data, "-")
		want := []string{"你", "好", "呀", "ab", "cd"}

		if !slices.Equal[[]string](got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test02", func(t *testing.T) {
		data := "你-好-呀-ab-cd"
		got := Split(data, "-")
		want := []string{"你", "好", "呀", "ab", "cd"}

		if !slices.Equal[[]string](got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

// 表格驱动测试
func TestSplit05(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{
			name:  "test01",
			input: "a-b-c",
			sep:   "-",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "test02",
			input: "你-好-呀-ab-cd",
			sep:   "-",
			want:  []string{"你", "好", "呀", "ab", "cd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.input, tt.sep)
			if !slices.Equal[[]string](got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
