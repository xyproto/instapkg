package instapkg

import (
	"testing"

	"github.com/bmizerany/assert"
)

// Only for Arch Linux, for now
func TestLinuxPackage(t *testing.T) {
	linuxPackage, err := NewArchPackage("linux")
	if err != nil {
		t.Fatalf("Error when examining package \"linux\": %s\n", err)
	}
	assert.Equal(t, linuxPackage.Name(), "linux")
	assert.Equal(t, linuxPackage.Installed(), true)
	filenames, err := linuxPackage.ListFiles()
	if err != nil {
		t.Fatalf("Error when listing files for the \"linux\" package: %s\n", err)
	}
	numFiles := len(filenames)
	if numFiles < 5 {
		t.Fatalf("Too few files in the linux package! Only: %d\n", numFiles)
	}
}
