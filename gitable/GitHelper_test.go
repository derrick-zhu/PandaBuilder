package gitable

import (
	"testing"
)

type testGit struct {
}

func (c testGit) RepoType() int {
	return 1
}

func (c testGit) URL() string {
	return "https://github.com/derrick-zhu/PandaBuilder.git"
}

func (c testGit) REF() string {
	return "develop"
}

type testGitPandaShop struct {
}

func (c testGitPandaShop) RepoType() int {
	return 1
}

func (c testGitPandaShop) URL() string {
	return "https://github.com/derrick-zhu/FlutterPandaShop.git"
}

func (c testGitPandaShop) REF() string {
	return "develop"
}

func TestGetRetrieveCurrentGitCommitHash(t *testing.T) {
	type args struct {
		ref string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"0.0.0"}, "d0447eca2774780b749c4f630db804961fc056bc"},
		{"", args{"branch_test"}, "45e6ba9431d54a5b362c1bb3535d6636c247d575"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRetrieveCurrentGitCommitHash(tt.args.ref); got != tt.want {
				t.Errorf("GitRetrieveCommitHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitRetrieveCommitHash(t *testing.T) {
	type args struct {
		aGit GitProtocol
		ref  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{testGit{}, ""}, "287339ee86da414212a29917a46069b5b752e1fb"},
		{"", args{testGit{}, "develop"}, "287339ee86da414212a29917a46069b5b752e1fb"},
		{"", args{testGit{}, "master"}, "84233b81ae6aac83bd9f6124fcdaf06e2e0fddf4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GitRetrieveCommitHash(tt.args.aGit, tt.args.ref); got != tt.want {
				t.Errorf("GitRetrieveCommitHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitClone(t *testing.T) {
	type args struct {
		aGit  GitProtocol
		local string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{testGit{}, "./fooRepo"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GitClone(tt.args.aGit, tt.args.local); got != tt.want {
				t.Errorf("GitClone() = %v, want %v", got, tt.want)
			}
		})
	}
}
