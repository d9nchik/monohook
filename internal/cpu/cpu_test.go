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
			name: "description(from smb) and comment",
			args: args{
				description: "Вiд: Петро Максим Олександрович",
				comment:     "Для себе",
			},
			wantMemo:      stringPtr("Для себе"),
			wantPayeeName: stringPtr("Петро Максим Олександрович"),
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

func TestShortenString(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		maxLen int
		want   string
	}{
		{name: "shorter than max", str: "hello", maxLen: 10, want: "hello"},
		{name: "equal to max", str: "hello", maxLen: 5, want: "hello"},
		{name: "longer than max", str: "hello world", maxLen: 5, want: "hello"},
		{name: "empty string", str: "", maxLen: 5, want: ""},
		{name: "unicode truncation", str: "Привіт світ", maxLen: 6, want: "Привіт"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, shortenString(tt.str, tt.maxLen))
		})
	}
}
