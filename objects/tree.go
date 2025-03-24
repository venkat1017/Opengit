package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

type TreeEntry struct {
	Mode string
	Name string
	Hash string
}

type Tree struct {
	Entries []TreeEntry
}

func ParseTree(data []byte) *Tree {
	tree := &Tree{Entries: make([]TreeEntry, 0)}
	pos := 0

	for pos < len(data) {
		// Find space after mode
		spacePos := pos
		for spacePos < len(data) && data[spacePos] != ' ' {
			spacePos++
		}
		mode := string(data[pos:spacePos])

		// Find null byte after name
		pos = spacePos + 1
		nullPos := pos
		for nullPos < len(data) && data[nullPos] != 0 {
			nullPos++
		}
		name := string(data[pos:nullPos])

		// Read 20 bytes for hash
		pos = nullPos + 1
		if pos+20 > len(data) {
			break
		}
		hash := hex.EncodeToString(data[pos : pos+20])
		pos += 20

		tree.Entries = append(tree.Entries, TreeEntry{
			Mode: mode,
			Name: name,
			Hash: hash,
		})
	}

	return tree
}

func (t *Tree) Serialize() []byte {
	// Sort entries by name for consistent hashing
	sort.Slice(t.Entries, func(i, j int) bool {
		return t.Entries[i].Name < t.Entries[j].Name
	})

	// Calculate total size for the header
	size := 0
	for _, entry := range t.Entries {
		// Format: mode<space>name<null>hash(20 bytes)
		size += len(entry.Mode) + 1 + len(entry.Name) + 1 + 20
	}

	// Create header
	header := fmt.Sprintf("tree %d\x00", size)
	result := []byte(header)

	// Add entries
	for _, entry := range t.Entries {
		// Convert hash from hex to bytes
		hashBytes := make([]byte, 20)
		fmt.Sscanf(entry.Hash, "%x", &hashBytes)

		// Append entry data
		result = append(result, []byte(entry.Mode)...)
		result = append(result, ' ')
		result = append(result, []byte(entry.Name)...)
		result = append(result, 0)
		result = append(result, hashBytes...)
	}

	return result
}

func (t *Tree) Hash() string {
	data := t.Serialize()
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash)
}
