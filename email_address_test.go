package obfuscate_test

import (
	"testing"

	"github.com/networkteam/obfuscate"
)

func TestEmailAddressPartially(t *testing.T) {
	tests := []struct {
		name         string
		emailAddress string
		want         string
	}{
		{
			name:         "empty string",
			emailAddress: "",
			want:         "",
		},
		{
			name:         "no @",
			emailAddress: "test",
			want:         "t*t",
		},
		{
			name:         "shortest local",
			emailAddress: "t@t",
			want:         "*@t",
		},
		{
			name:         "longer local",
			emailAddress: "test@t",
			want:         "t*t@t",
		},
		{
			name:         "normal address",
			emailAddress: "john.doe@example.com",
			want:         "j*n.d*e@e*e.com",
		},
		{
			name:         "short local",
			emailAddress: "jd@example.com",
			want:         "j*@e*e.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obfuscate.EmailAddressPartially(tt.emailAddress); got != tt.want {
				t.Errorf("EmailAddressPartially(%s) = %v, want %v", tt.emailAddress, got, tt.want)
			}
		})
	}
}
