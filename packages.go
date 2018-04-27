package instapkg

import (
	"bytes"
	"os/exec"
)

// TODO: Decide if Package should rather be a universal data-container struct or not.
// TODO: Decide where the actual package repo querying should be happening, in methods, at construction of structs or in the repo structs.
// TODO: Decide where and when caching should happen.

type Package interface {
	Name() string
	ListFiles() []string
	HasFile(filename string) bool
	Installed() bool
}

// --- Arch Linux ---

type ArchPackage struct {
	name              string
	filenamesWithPath []string
	installed         bool
}

func (ap *ArchPackage) Name() string {
	return ap.name
}

func (ap *ArchPackage) ListFiles() ([]string, error) {
	return ap.filenamesWithPath, nil
}

func (ap *ArchPackage) Installed() bool {
	return ap.installed

}

func NewArchPackage(name string) (*ArchPackage, error) {
	// TODO: Find out how to set LC_ALL=C before executing
	out, err := exec.Command("/usr/bin/pacman", "-Ql", name).CombinedOutput()
	if err != nil {
		return nil, err
	}
	filenames := []string{}
	installed := true
	if bytes.HasPrefix(out, []byte("error: ")) {
		installed = false
	} else {
		for _, line := range bytes.Split(out, []byte("\n")) {
			if len(line) <= len(name) {
				continue
			}
			filenames = append(filenames, string(line[len(name)+1:]))
		}
	}
	return &ArchPackage{name, filenames, installed}, nil
}

// --- Ubuntu ---

type UbuntuPackage struct {
	name              string
	filenamesWithPath []string
	installed         bool
}

func (ap *UbuntuPackage) Name() string {
	return ap.name
}

func (ap *UbuntuPackage) ListFiles() []string {
	return ap.filenamesWithPath
}

func (ap *UbuntuPackage) Installed() bool {
	return ap.installed
}

func NewUbuntuPackage(name string, filenamesWithPath []string, installed bool) *UbuntuPackage {
	return &UbuntuPackage{name, filenamesWithPath, installed}
}
