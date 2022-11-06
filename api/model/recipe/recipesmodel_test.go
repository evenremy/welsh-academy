package recipe

import "testing"

func Test_arrayToSqlString(t *testing.T) {
	type args struct {
		array []int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty array",
			args: args{array: []int64{}},
			want: "()",
		},
		{
			name: "array with 1 number",
			args: args{array: []int64{13}},
			want: "(13)",
		},
		{
			name: "array with n numbers",
			args: args{array: []int64{10, 13}},
			want: "(10,13)",
		},
		{
			name: "Nil",
			args: args{array: nil},
			want: "()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayToSqlString(tt.args.array); got != tt.want {
				t.Errorf("arrayToSqlString() = %v, want %v", got, tt.want)
			}
		})
	}
}
