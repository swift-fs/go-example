package td

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	convey.Convey("test add", t, func() {
		type args struct {
			a int
			b int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			{
				name: "正整数相加",
				args: args{
					a: 10,
					b: 20,
				},
				want: 30,
			},
			{
				name: "一正一负相加",
				args: args{
					a: -10,
					b: 20,
				},
				want: 10,
			},
		}
		for _, tt := range tests {
			convey.Convey(tt.name, func() {
				got := Add(tt.args.a, tt.args.b)
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		}
	})

}
