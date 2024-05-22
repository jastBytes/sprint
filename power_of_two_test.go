package sprint

import (
	"testing"
)

func TestDehumanize(t *testing.T) {
	type args struct {
		sizeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{sizeStr: "1"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "1Ki",
			args:    args{sizeStr: "1Ki"},
			want:    1024,
			wantErr: false,
		},
		{
			name:    "1Mi",
			args:    args{sizeStr: "1Mi"},
			want:    1024 * 1024,
			wantErr: false,
		},
		{
			name:    "1Gi",
			args:    args{sizeStr: "1Gi"},
			want:    1024 * 1024 * 1024,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Dehumanize(tt.args.sizeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dehumanize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dehumanize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanize(t *testing.T) {
	type args struct {
		size int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{size: 1},
			want: "1",
		},
		{
			name: "1Ki",
			args: args{size: 1024},
			want: "1Ki",
		},
		{
			name: "1Mi",
			args: args{size: 1024 * 1024},
			want: "1Mi",
		},
		{
			name: "1Gi",
			args: args{size: 1024 * 1024 * 1024},
			want: "1Gi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Humanize(tt.args.size)
			if got != tt.want {
				t.Errorf("Humanize() = %v, want %v", got, tt.want)
			}
		})
	}
}
