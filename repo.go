package instapkg

type Repo interface {
	HasPackage(name string) bool
	FindPackageByName(name string) (Package, error)
	FindPackagesWithFilename(filename string, pathMatters bool) ([]Package, error)
}

type ArchRepo struct {
}

func (ar *ArchRepo) HasPackage(name string) bool {
	return false
}

func (ar *ArchRepo) FindPackageByName(name string) (Package, error) {
	return nil, nil
}

func (ar *ArchRepo) FindPackageWithFilename(filename string, pathMatters bool) ([]Package, error) {
	return nil, nil
}
