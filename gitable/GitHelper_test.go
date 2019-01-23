package gitable

import "testing"

type testGit struct {
}

func (c testGit) RepoType() int {
	return 1
}

func (c testGit) URL() string {
	return "git@github.com:derrick-zhu/PandaBuilder.git"
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
	return "git@github.com:derrick-zhu/FlutterPandaShop.git"
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
		{"", args{testGitPandaShop{}, ""}, "ee7247984c2932d0b265f6980df562a1cfb465c4"},
		{"", args{testGitPandaShop{}, "develop"}, "ee7247984c2932d0b265f6980df562a1cfb465c4"},
		{"", args{testGitPandaShop{}, "master"}, "c23daf4e64959eb79bc54c4283a13b14c9c20346"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GitRetrieveCommitHash(tt.args.aGit, tt.args.ref); got != tt.want {
				t.Errorf("GitRetrieveCommitHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
