package refs

import (
	"Opengit/repo"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Refs struct {
	r *repo.Repository
}

func NewRefs(r *repo.Repository) *Refs {
	return &Refs{r: r}
}

func (refs *Refs) ReadRef(ref string) (string, error) {
	path := refs.r.GitFile(ref)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // Ref doesn’t exist
		}
		return "", fmt.Errorf("reading ref %s: %v", ref, err)
	}
	return strings.TrimSpace(string(data)), nil
}

func (refs *Refs) WriteRef(ref, value string) error {
	path := refs.r.GitFile(ref)
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("creating ref dir: %v", err)
	}
	return os.WriteFile(path, []byte(value+"\n"), 0644)
}

func (refs *Refs) ReadHEAD() (string, error) {
	head, err := refs.ReadRef("HEAD")
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(head, "ref: ") {
		return refs.ReadRef(strings.TrimPrefix(head, "ref: "))
	}
	return head, nil // Direct SHA-1 (detached HEAD)
}

func (refs *Refs) WriteHEAD(value string) error {
	return refs.WriteRef("HEAD", value)
}
