package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Githubactions test" {
		t.Errorf("Greet() = %s; want Githubactions test", result)
	}
}
