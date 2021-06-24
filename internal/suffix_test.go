package internal

import "testing"

func TestHeadSuffix(t *testing.T) {
	type args struct {
		s string
	}

	type want struct {
		head   string
		suffix string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: want{
				head:   "",
				suffix: "",
			},
		},
		{
			name: "single char",
			args: args{
				s: "S",
			},
			want: want{
				head:   "",
				suffix: "S",
			},
		},
		{
			name: "2 char",
			args: args{
				s: "1S",
			},
			want: want{
				head:   "1",
				suffix: "S",
			},
		},
		{
			name: "3 char",
			args: args{
				s: "10H",
			},
			want: want{
				head:   "10",
				suffix: "H",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead, gotSuffix := HeadSuffix(tt.args.s)
			if gotHead != tt.want.head {
				t.Errorf("HeadSuffix() gotHead = %v, want %v", gotHead, tt.want.head)
			}

			if gotSuffix != tt.want.suffix {
				t.Errorf("HeadSuffix() gotSuffix = %v, want %v", gotSuffix, tt.want.suffix)
			}
		})
	}
}
