package sanitize

import "testing"

func TestStringToAsciiBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{"Test 1", args{"Sławomir Jasiński"}, "Slawomir Jasinski"},
		{"Test 2", args{"Gżegżółka"}, "Gzegzolka"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Accents(tt.args.s); gotResult != tt.wantResult {
				t.Errorf("Accents() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
