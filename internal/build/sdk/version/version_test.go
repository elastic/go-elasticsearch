package version

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		versionStr string
	}
	tests := []struct {
		name    string
		args    args
		want    Version
		wantErr bool
	}{
		{
			name:    "9.0.1-SNAPSHOT",
			args:    args{versionStr: "9.0.1-SNAPSHOT"},
			want:    Version{9, 0, 1, "SNAPSHOT"},
			wantErr: false,
		},
		{
			name:    "9.0.1-rc1",
			args:    args{versionStr: "9.0.1-rc1"},
			want:    Version{9, 0, 1, "rc1"},
			wantErr: false,
		},
		{
			name:    "9.0.1",
			args:    args{versionStr: "9.0.1"},
			want:    Version{9, 0, 1, ""},
			wantErr: false,
		},
		{
			name:    "9.0.1.0",
			args:    args{versionStr: "9.0.1.0"},
			wantErr: true,
		},
		{
			name:    "9.0.1.0-SNAPSHOT",
			args:    args{versionStr: "9.0.1.0-SNAPSHOT"},
			wantErr: true,
		},
		{
			name:    "9.0-SNAPSHOT",
			args:    args{versionStr: "9.0-SNAPSHOT"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.versionStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
