package gitable

import "testing"

type testGit struct {
}

func (c testGit) URL() string {
	return "git@github.com:derrick-zhu/PandaBuilder.git"
}

func (c testGit) REF() string {
	return "develop"
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
		{"", args{testGit{}, "0.0.0"}, "d0447eca2774780b749c4f630db804961fc056bc"},
		{"", args{testGit{}, "branch_test"}, "45e6ba9431d54a5b362c1bb3535d6636c247d575"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GitRetrieveCommitHash(tt.args.aGit, tt.args.ref); got != tt.want {
				t.Errorf("GitRetrieveCommitHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
