package index

import (
	"Opengit/repo"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Mode     string    // e.g., "100644" for regular file
	Hash     string    // SHA-1 hash of the blob
	ModTime  time.Time // Last modification time
	FileSize int64     // File size in bytes
}

type Index struct {
	Entries map[string]*Entry
}

func ReadIndex(r *repo.Repository) (*Index, error) {
	indexPath := r.GitFile("index")
	data, err := os.ReadFile(indexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Index{Entries: make(map[string]*Entry)}, nil
		}
		return nil, fmt.Errorf("reading index: %v", err)
	}

	buf := bytes.NewReader(data)
	idx := &Index{Entries: make(map[string]*Entry)}
	for buf.Len() > 0 {
		var pathLen uint32
		binary.Read(buf, binary.BigEndian, &pathLen)
		pathBytes := make([]byte, pathLen)
		buf.Read(pathBytes)
		path := string(pathBytes)

		var mode [6]byte
		buf.Read(mode[:])
		hashBytes := make([]byte, 40)
		buf.Read(hashBytes)
		var mtimeSec, mtimeNsec int64
		binary.Read(buf, binary.BigEndian, &mtimeSec)
		binary.Read(buf, binary.BigEndian, &mtimeNsec)
		var fileSize int64
		binary.Read(buf, binary.BigEndian, &fileSize)

		idx.Entries[path] = &Entry{
			Mode:     string(mode[:]),
			Hash:     string(hashBytes),
			ModTime:  time.Unix(mtimeSec, mtimeNsec),
			FileSize: fileSize,
		}
	}
	return idx, nil
}

func WriteIndex(r *repo.Repository, idx *Index) error {
	indexPath := r.GitFile("index")
	var buf bytes.Buffer
	for path, entry := range idx.Entries {
		pathLen := uint32(len(path))
		binary.Write(&buf, binary.BigEndian, pathLen)
		buf.WriteString(path)
		buf.WriteString(entry.Mode)
		buf.WriteString(entry.Hash)
		binary.Write(&buf, binary.BigEndian, entry.ModTime.Unix())
		binary.Write(&buf, binary.BigEndian, int64(entry.ModTime.Nanosecond()))
		binary.Write(&buf, binary.BigEndian, entry.FileSize)
	}
	return os.WriteFile(indexPath, buf.Bytes(), 0644)
}
