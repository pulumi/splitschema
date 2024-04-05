package bschema

import (
	"fmt"
	"hash/crc32"
	"path/filepath"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func getPath(token string, kind string) (string, error) {
	t, err := tokens.ParseTypeToken(token)
	if err != nil {
		return "", err
	}

	modName := t.Module().Name().String()
	typeName := t.Name().String()

	if parts := strings.Split(modName, "/"); len(parts) > 0 {
		modName = parts[0]
	}
	filename := makeFileName(typeName)
	path := filepath.Join(modName, kind, filename)
	return path, nil
}

func makeFileName(s string) string {
	return strings.ToLower(s) + "-" + shortHash(s)
}

func shortHash(s string) string {
	return fmt.Sprintf("%08x", crc32.Checksum([]byte(s), crcTable))
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)
