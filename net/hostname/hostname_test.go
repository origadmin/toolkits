package hostname

import (
	"net"
	"testing"
)

func TestHostname_Replace(t *testing.T) {
	type fields struct {
		lists           []string
		hosts           map[string]net.IP
		prefix          string
		suffix          string
		useCustomSuffix bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "replace_1",
			fields: fields{
				lists: []string{
					"host1:192.168.1.1",
					"host2:192.168.1.2",
				},
				hosts: map[string]net.IP{
					"host3": net.ParseIP("192.168.1.3"),
				},
			},
			args: args{
				name: "@host1",
			},
			want: "192.168.1.1",
		},
		{
			name: "replace_2",
			fields: fields{
				lists: []string{
					"host1:192.168.1.1",
					"host2:192.168.1.2",
				},
				hosts: map[string]net.IP{
					"host3": net.ParseIP("192.168.1.3"),
				},
			},
			args: args{
				name: "@host1:9874",
			},
			want: "192.168.1.1:9874",
		},
		{
			name: "replace_3",
			fields: fields{
				lists: []string{
					"host1:192.168.1.1",
					"host2:192.168.1.2",
				},
				hosts: map[string]net.IP{
					"host3": net.ParseIP("192.168.1.3"),
				},
			},
			args: args{
				name: "@host2:9874",
			},
			want: "192.168.1.2:9874",
		},
		{
			name: "replace_4",
			fields: fields{
				lists: []string{
					"host1:192.168.1.1",
					"host2:192.168.1.2",
				},
				hosts: map[string]net.IP{
					"host3": net.ParseIP("192.168.1.3"),
				},
			},
			args: args{
				name: "@host3:9874",
			},
			want: "192.168.1.3:9874",
		},
		{
			name: "replace_5",
			fields: fields{
				lists: []string{
					"host1:192.168.1.1",
					"host2:192.168.1.2",
				},
				hosts: map[string]net.IP{
					"host3": net.ParseIP("192.168.1.3"),
				},
				suffix: ":",
			},
			args: args{
				name: "@host1:",
			},
			want: "192.168.1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts []Setting
			if len(tt.fields.lists) > 0 {
				opts = append(opts, WithHosts(tt.fields.lists, ":"))
			}
			if len(tt.fields.hosts) > 0 {
				opts = append(opts, WithHostMap(tt.fields.hosts))
			}
			if tt.fields.prefix != "" {
				opts = append(opts, UseCustomPrefix(tt.fields.prefix))
			}
			if tt.fields.suffix != "" {
				opts = append(opts, UseCustomSuffix(tt.fields.suffix))
			}
			r := New(opts...)
			//t.Logf("r.hosts: %+v", r.hosts)
			if got := r.Replace(tt.args.name); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
