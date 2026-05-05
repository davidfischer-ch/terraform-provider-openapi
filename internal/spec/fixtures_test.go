package spec

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pb33f/libopenapi"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// parseSpec parses an inline YAML string into a document model.
func parseSpec(t *testing.T, yaml string) *libopenapi.DocumentModel[v3high.Document] {
	t.Helper()
	doc, err := libopenapi.NewDocument([]byte(yaml))
	if err != nil {
		t.Fatalf("NewDocument: %v", err)
	}
	model, err := doc.BuildV3Model()
	if err != nil {
		t.Fatalf("BuildV3Model: %v", err)
	}
	return model
}

// mustParseFixture loads a YAML file from testdata/ and parses it into a document model.
func mustParseFixture(t *testing.T, name string) *libopenapi.DocumentModel[v3high.Document] {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("read testdata/%s: %v", name, err)
	}
	return parseSpec(t, string(b))
}

func fieldsByName(fields []*FieldSpec) map[string]*FieldSpec {
	m := make(map[string]*FieldSpec, len(fields))
	for _, f := range fields {
		m[f.Name] = f
	}
	return m
}
