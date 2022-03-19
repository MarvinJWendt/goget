package remotes

import (
	"testing"
)

func Test_getRemote(t *testing.T) {
	type args struct {
		arg string
	}
	type repo struct {
		fullName string
		htmlUrl  string
	}
	tests := []struct {
		name    string
		args    args
		want    *repo
		wantErr bool
	}{{
		"fyne", args{"fyne"}, &repo{"fyne-io/fyne", "https://github.com/fyne-io/fyne"}, false,
	}, {
		"pterm", args{"pterm"}, &repo{"pterm/pterm", "https://github.com/pterm/pterm"}, false,
	}, {
		"Nonexistent package", args{"Nonexistent_package"}, nil, true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRemote(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			result := true
			if got == nil && tt.want == nil {
				return
			}
			for _, repo := range *got {
				if repo.FullName == tt.want.fullName && repo.HtmlUrl == tt.want.htmlUrl {
					result = false
					break
				}
			}
			if result {
				t.Errorf("getRemote() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVersion(t *testing.T) {
	type args struct {
		repoName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{{
		"fyne", args{"fyne-io/fyne"}, "v2", false,
	}, {
		"pterm", args{"pterm/pterm"}, "", false,
	}, {
		"Nonexistent package", args{"Nonexistent_package"}, "", true,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getVersion(tt.args.repoName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
