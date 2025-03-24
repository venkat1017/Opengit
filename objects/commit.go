package objects

import (
	"crypto/sha1"
	"fmt"
	"strings"
	"time"
)

type Commit struct {
	Tree      string
	Parents   []string
	Author    string
	Committer string
	Message   string
	Timestamp time.Time
}

func ParseCommit(data []byte) *Commit {
	commit := &Commit{
		Parents: make([]string, 0),
	}

	// Split into headers and message
	parts := strings.SplitN(string(data), "\n\n", 2)
	if len(parts) != 2 {
		return commit
	}

	headers := strings.Split(parts[0], "\n")
	commit.Message = parts[1]

	// Parse headers
	for _, header := range headers {
		if header == "" {
			continue
		}

		key := header[:strings.Index(header, " ")]
		value := header[strings.Index(header, " ")+1:]

		switch key {
		case "tree":
			commit.Tree = value
		case "parent":
			commit.Parents = append(commit.Parents, value)
		case "author":
			commit.Author = value
		case "committer":
			commit.Committer = value
		}
	}

	return commit
}

func (c *Commit) Serialize() []byte {
	// Format: tree <hash>\n
	//         parent <hash>\n (optional, multiple)
	//         author <name> <timestamp>\n
	//         committer <name> <timestamp>\n
	//         \n
	//         <message>
	content := fmt.Sprintf("tree %s\n", c.Tree)

	for _, parent := range c.Parents {
		content += fmt.Sprintf("parent %s\n", parent)
	}

	timestamp := c.Timestamp.Unix()
	content += fmt.Sprintf("author %s %d +0000\n", c.Author, timestamp)
	content += fmt.Sprintf("committer %s %d +0000\n", c.Committer, timestamp)
	content += "\n" + c.Message

	header := fmt.Sprintf("commit %d\x00", len(content))
	return append([]byte(header), []byte(content)...)
}

func (c *Commit) Hash() string {
	data := c.Serialize()
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash)
}
