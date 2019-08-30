package chargify

import "testing"

func Test_getAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "test get key",
			want: "DYUjq1rrx3E1tutWCBx1ND0VlZ6G92P4BjbJ2dNP9A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := apiKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
