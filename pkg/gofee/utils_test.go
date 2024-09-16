package gofee

import "testing"

// TestContains tests the Contains function
// It tests if the function returns true if the character is in the set
func TestContains(t *testing.T) {
	type args struct {
		set  string
		char rune
	}
	
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Character is in the set",
			args: args{
				set:  "abc",
				char: 'a',
			},
			want: true,
		},
		{
			name: "Character is not in the set",
			args: args{
				set:  "abc",
				char: 'd',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.set, tt.args.char); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}