package main

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   int
		wantErr bool
	}{
		{"Empty string", args{s: ""}, false, 0, true},
		{"[]", args{s: "[]"}, true, 0, false},
		{"a[b]c", args{s: "a[b]c"}, true, 0, false},
		{"{}[]", args{s: "{}[]"}, true, 0, false},
		{"[()]", args{s: "[()]"}, true, 0, false},
		{"(())", args{s: "(())"}, true, 0, false},
		{"{[]}()", args{s: "{[]}()"}, true, 0, false},
		{"{", args{s: "{"}, false, 1, false},
		{"ac{df", args{s: "ac{df"}, false, 3, false},
		{"{}{", args{s: "{}{"}, false, 3, false},
		{"}", args{s: "}"}, false, 1, false},
		{"{[}", args{s: "{[}"}, false, 3, false},
		{"foo(bar);", args{s: "foo(bar);"}, true, 0, false},
		{"foo(bar[i);", args{s: "foo(bar[i);"}, false, 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := isBalanced(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("isBalanced() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isBalanced() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isBalanced() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
