package maps

import (
	"testing"
)

type TestStruct struct {
	Name string `json:"name" yml:"name2"`
	Age  int    `json:"age" yml:"age2"`
	Sex  bool   `json:"sex"`
}

func TestMap_Convert(t *testing.T) {
	type fields struct {
	}
	type args struct {
		bind any
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test",
			args: args{
				bind: new(TestStruct),
				name: "name2",
			},
			want: "name",
		},
		{
			name: "test",
			args: args{
				bind: new(TestStruct),
				name: "age2",
			},
			want: "age",
		},
		{
			name: "test",
			args: args{
				bind: new(TestStruct),
				name: "age",
			},
			want: "age",
		},
		{
			name: "test",
			args: args{
				bind: new(TestStruct),
				name: "Sex",
			},
			want: "sex",
		},
		{
			name: "test",
			args: args{
				bind: TestStruct{},
				name: "Sex",
			},
			want: "sex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			m.Bind(tt.args.bind, "yml")
			b := m.Get(tt.args.bind)
			if b == nil {
				t.Fatal("b is nil")
			}
			if got := b.Convert(tt.args.name); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
