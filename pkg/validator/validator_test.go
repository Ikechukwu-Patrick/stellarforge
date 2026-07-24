package validator

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{
			name:    "Empty address",
			address: "",
			want:    false,
		},
		{
			name:    "Wrong prefix",
			address: "X1234567890123456789012345678901234567890123456789012345",
			want:    false,
		},
		{
			name:    "Wrong length",
			address: "G123",
			want:    false,
		},
		{
			name:    "Looks valid",
			address: "GAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			want:    true,
		},
	}

	for _, tt := range tests {
		got, _ := Validate(tt.address)

		if got != tt.want {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.want, got)
		}
	}
}