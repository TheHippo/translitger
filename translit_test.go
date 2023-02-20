package translitger

import "testing"

func TestTranslit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Frühstück",
			args: args{
				input: "Frühstück",
			},
			want: "Fruehstueck",
		},
		{
			name: "Süßspeise",
			args: args{
				input: "Süßspeise",
			},
			want: "Suessspeise",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Translit(tt.args.input); got != tt.want {
				t.Errorf("Translit() = %v, want %v", got, tt.want)
			}
		})
	}
}
