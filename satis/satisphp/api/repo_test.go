package api

import (
	"reflect"
	"testing"
)

func TestNewRepo(t *testing.T) {
	type args struct {
		repoType string
		url      string
	}
	tests := []struct {
		name string
		args args
		want *Repo
	}{
		{
			"valid",
			args{
				repoType: "repository",
				url:      "http://github.com/",
			},
			&Repo{
				ID:   "1680316628",
				Type: "repository",
				URL:  "http://github.com/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepo(tt.args.repoType, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
