package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		file     string
		encoding string
		noDash   bool
		eq       bool
		k        keys
		kv       keys
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"yml: check", args{file: "config.yml", k: keys{"example.text"}}, []string{"textFlag"}},
		{"json: check", args{file: "config.json", k: keys{"example.text"}}, []string{"textFlag"}},
		{"common: kv single dash", args{file: "config.yml", kv: keys{"t=example.text"}}, []string{"-t textFlag"}},
		{"common: kv double dash", args{file: "config.yml", kv: keys{"text=example.text"}}, []string{"--text textFlag"}},
		{"yml: keyed dynamic array", args{file: "config.yml", kv: keys{"arr=example.deep.array"}}, []string{"--arr true", "--arr false", "--arr 1", "--arr yes"}},
		{"json: keyed dynamic array", args{file: "config.json", kv: keys{"arr=example.deep.array"}}, []string{"--arr true", "--arr false", "--arr 1", "--arr yes"}},
		{"common: unkeyed dynamic array", args{file: "config.json", k: keys{"example.array"}}, []string{"one", "two"}},
		{"common: no-dash",
			args{
				file:   "config.json",
				noDash: true,
				kv: keys{
					"e=example.text",
					"-e=example.text",
					"--e=example.text",
				},
			},
			[]string{"e textFlag", "-e textFlag", "--e textFlag"},
		},
		{"common: space separator",
			args{
				file: "config.json",
				eq:   false,
				kv: keys{
					"e=example.text",
				},
			},
			[]string{"-e textFlag"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.file, tt.args.encoding, tt.args.noDash, tt.args.eq, tt.args.k, tt.args.kv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
