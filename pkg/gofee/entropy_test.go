package gofee

import (
	"math"
	"testing"
)

func TestCalculateEntropy(t *testing.T) {
	type args struct {
		charsetSize    int
		passwordLength int
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Valid entropy calculation",
			args: args{
				charsetSize:    64,
				passwordLength: 16,
			},
			want:    96.0,
			wantErr: false,
		},
		{
			name: "Zero charset size",
			args: args{
				charsetSize:    0,
				passwordLength: 16,
			},
			want:    0.0,
			wantErr: true,
		},
		{
			name: "Zero password length",
			args: args{
				charsetSize:    64,
				passwordLength: 0,
			},
			want:    0.0,
			wantErr: true,
		},
		{
			name: "Negative charset size",
			args: args{
				charsetSize:    -64,
				passwordLength: 16,
			},
			want:    0.0,
			wantErr: true,
		},
		{
			name: "Negative password length",
			args: args{
				charsetSize:    64,
				passwordLength: -16,
			},
			want:    0.0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateEntropy(tt.args.charsetSize, tt.args.passwordLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateEntropy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !math.IsNaN(got) {
				t.Errorf("CalculateEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}
