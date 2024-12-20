package game

import (
	"reflect"
	"testing"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
)

func TestCalculateStats(t *testing.T) {
	type args struct {
		base  app.Stats
		ivs   app.Stats
		evs   app.Stats
		level int
	}
	tests := []struct {
		name string
		args args
		want app.Stats
	}{
		{
			name: "Garchomp test",
			args: args{
				base: app.Stats{
					HP:      108,
					Attack:  130,
					Defense: 95,
					SpAtk:   80,
					SpDef:   85,
					Speed:   102,
				},
				ivs: app.Stats{
					HP:      24,
					Attack:  12,
					Defense: 30,
					SpAtk:   16,
					SpDef:   23,
					Speed:   5,
				},
				evs: app.Stats{
					HP:      74,
					Attack:  190,
					Defense: 91,
					SpAtk:   48,
					SpDef:   84,
					Speed:   23,
				},
				level: 78,
			},
			want: app.Stats{
				HP:      289,
				Attack:  253,
				Defense: 193,
				SpAtk:   151,
				SpDef:   171,
				Speed:   171,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateStats(tt.args.base, tt.args.ivs, tt.args.evs, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
