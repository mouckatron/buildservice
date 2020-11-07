package buildrunner

import "testing"

func TestBuildSettingsCodeDir(t *testing.T) {
	b := BuildSettings{
		Name:       "myjob",
		ID:         "20060102-150405",
		WorkingDir: "/home/myjob/20060102-150405"}

	expected := "/home/myjob/20060102-150405/code"

	if b.CodeDir() != expected {
		t.Errorf("Wrong CodeDir: wanted %s, got %s",
			expected,
			b.CodeDir())
	}
}

func TestBuildSettingsArtifactDir(t *testing.T) {
	b := BuildSettings{
		Name:       "myjob",
		ID:         "20060102-150405",
		WorkingDir: "/home/myjob/20060102-150405"}

	expected := "/home/myjob/20060102-150405/artifacts"

	if b.ArtifactDir() != expected {
		t.Errorf("Wrong ArtifactDir: wanted %s, got %s",
			expected,
			b.ArtifactDir())
	}
}

func TestBuildSettingsBuildspecFile(t *testing.T) {
	b := BuildSettings{
		Name:          "myjob",
		ID:            "20060102-150405",
		WorkingDir:    "/home/myjob/20060102-150405",
		BuildspecPath: "buildspec-file.yml"}

	expected := "/home/myjob/20060102-150405/code/buildspec-file.yml"

	if b.BuildspecFile() != expected {
		t.Errorf("Wrong CodeDir: wanted %s, got %s",
			expected,
			b.BuildspecFile())
	}
}

func TestBuildSettingsToString(t *testing.T) {
	b := BuildSettings{
		Name:          "myjob",
		ID:            "20060102-150405",
		WorkingDir:    "/home/myjob/20060102-150405",
		BuildspecPath: "buildspec.yml",
		Branch:        "master",
		URL:           "https://github.com/mouckatron/buildservice.git"}

	expected := "Name: myjob, ID: 20060102-150405, URL: https://github.com/mouckatron/buildservice.git, Branch: master, BuildspecPath: buildspec.yml, WorkingDir: /home/myjob/20060102-150405"

	if b.ToString() != expected {
		t.Errorf("Wrong ToString: wanted %s, got %s",
			expected,
			b.ToString())
	}
}
