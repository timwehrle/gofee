package gofe

import "testing"

func TestMapToCharset(t *testing.T) {
	type args struct {
		length int
		config PasswordConfig
	}

	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedLen int
		expectedSet string
	}{
		{
			name: "Empty charset",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  false,
					IncludeUppers:  false,
					IncludeDigits:  false,
					IncludeSymbols: false,
				},
			},
			wantErr:     true,
			expectedLen: 0,
			expectedSet: "",
		},
		{
			name: "Valid charset",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  true,
					IncludeUppers:  true,
					IncludeDigits:  true,
					IncludeSymbols: true,
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: All,
		},
		{
			name: "Valid charset with no symbols",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  true,
					IncludeUppers:  true,
					IncludeDigits:  true,
					IncludeSymbols: false,
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: Lowers + Uppers + Digits,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToCharset(tt.args.length, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToCharset() error = %v, wanted error = %v", err, tt.wantErr)
			}

			if len(got) != tt.expectedLen {
				t.Errorf("MapToCharset() length = %v, wanted length = %v", len(got), tt.expectedLen)
			}

			for _, c := range got {
				if !contains(tt.expectedSet, c) {
					t.Errorf("MapToCharset() character = %v, not found in expected set", c)
				}
			}
		})
	}
}

func contains(set string, char rune) bool {
	for _, c := range set {
		if c == char {
			return true
		}
	}
	return false
}
