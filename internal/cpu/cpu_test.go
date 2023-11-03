package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCPU_getMemoAndPayeeName(t *testing.T) {
	type args struct {
		description string
		comment     string
	}
	tests := []struct {
		name          string
		args          args
		wantMemo      *string
		wantPayeeName *string
	}{
		{
			name: "description(from smb) and comment",
			args: args{
				description: "from: description ",
				comment:     " comment ",
			},
			wantMemo:      stringPtr("comment"),
			wantPayeeName: stringPtr("description"),
		},
		{
			name: "description(from smb) and no comment",
			args: args{
				description: "from:  description ",
				comment:     "",
			},
			wantMemo:      nil,
			wantPayeeName: stringPtr("description"),
		},
		{
			name: "description and no comment",
			args: args{
				description: "description",
				comment:     "",
			},
			wantMemo:      stringPtr("description"),
			wantPayeeName: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMemo, gotPayeeName := getMemoAndPayeeName(tt.args.description, tt.args.comment)

			// Simplify the comparison of pointers
			if tt.wantMemo != nil {
				require.Equal(t, *tt.wantMemo, *gotMemo)
			} else {
				require.Nil(t, gotMemo)
			}

			if tt.wantPayeeName != nil {
				require.Equal(t, *tt.wantPayeeName, *gotPayeeName)
			} else {
				require.Nil(t, gotPayeeName)
			}
		})

	}
}
