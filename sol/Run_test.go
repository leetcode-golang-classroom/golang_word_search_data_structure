package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	commands := &[]string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"}
	values := &[][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}}
	for idx := 0; idx < b.N; idx++ {
		Run(commands, values)
	}
}
func TestRun(t *testing.T) {
	type args struct {
		commands *[]string
		values   *[][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{commands: &[]string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
				values: &[][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}},
			},
			want: []string{"null", "null", "null", "null", "false", "true", "true", "true"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.commands, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
