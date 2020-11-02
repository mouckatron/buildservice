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
