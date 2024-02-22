package td

import (
	"reflect"
	"testing"
)

// vscode自动生产测试TestSplit
func TestSplit(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "test11",
			args: args{
				str: "a-b-c",
				sep: "-",
			},
			wantResult: []string{"a", "b", "c"},
		},
		{
			name: "test22",
			args: args{
				str: "你-好-呀-ab-cd",
				sep: "-",
			},
			wantResult: []string{"你", "好", "呀", "ab", "cd"},
		},

		{
			name: "test33",
			args: args{
				str: "",
				sep: "-",
			},
			wantResult: []string{""},
		},

		{
			name: "test44",
			args: args{
				str: "abcd",
				sep: "",
			},
			wantResult: []string{"abcd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Split(tt.args.str, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
