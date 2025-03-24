package repo

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"os"
	"path/filepath"
)

type Repository struct {
	Worktree string // Path to the working directory
	Gitdir   string // Path to the .git directory
}

func NewRepository(path string) (*Repository, error) {
	gitdir := filepath.Join(path, ".opengit")
	info, err := os.Stat(gitdir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(".opengit directory not found")
		}
		return nil, fmt.Errorf("accessing .opengit directory: %v", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf(".opengit is not a directory")
	}
	return &Repository{
		Worktree: path,
		Gitdir:   gitdir,
	}, nil
}

func (r *Repository) GitFile(parts ...string) string {
	return filepath.Join(append([]string{r.Gitdir}, parts...)...)
}

func (r *Repository) WriteObject(data []byte, hash string) error {
	var buf bytes.Buffer
	w, err := zlib.NewWriterLevel(&buf, zlib.BestCompression)
	if err != nil {
		return fmt.Errorf("creating zlib writer: %v", err)
	}
	_, err = w.Write(data)
	w.Close()
	if err != nil {
		return fmt.Errorf("compressing data: %v", err)
	}

	objPath := r.GitFile("objects", hash[:2], hash[2:])
	err = os.MkdirAll(filepath.Dir(objPath), 0755)
	if err != nil {
		return fmt.Errorf("creating object dir: %v", err)
	}
	return os.WriteFile(objPath, buf.Bytes(), 0644)
}

func (r *Repository) ReadObject(hash string) ([]byte, error) {
	objPath := r.GitFile("objects", hash[:2], hash[2:])
	compressed, err := os.ReadFile(objPath)
	if err != nil {
		return nil, fmt.Errorf("reading object %s: %v", hash, err)
	}

	rdr, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, fmt.Errorf("decompressing object %s: %v", hash, err)
	}
	defer rdr.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(rdr)
	if err != nil {
		return nil, fmt.Errorf("reading decompressed data: %v", err)
	}
	return buf.Bytes(), nil
}
