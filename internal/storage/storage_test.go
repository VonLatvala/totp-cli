package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	s "github.com/yitsushi/totp-cli/internal/storage"
)

func TestFindNamespace(t *testing.T) {
	storage := &s.Storage{
		Namespaces: []*s.Namespace{
			{Name: "Namespace1"},
			{Name: "Namespace2"},
			{Name: "Namespace3"},
		},
	}

	namespace, err := storage.FindNamespace("Namespace1")

	assert.Equal(t, err, nil, "Error should be nil")
	assert.Equal(t, namespace.Name, "Namespace1", "Found namespace name should be Namespace1")
}

func TestFindNamespace_NotFound(t *testing.T) {
	storage := &s.Storage{
		Namespaces: []*s.Namespace{
			{Name: "Namespace1"},
			{Name: "Namespace2"},
			{Name: "Namespace3"},
		},
	}

	namespace, err := storage.FindNamespace("NamespaceNotFound")

	assert.EqualError(
		t,
		err,
		"namespace not found: NamespaceNotFound",
		"Error should be 'namespace not found: NamespaceNotFound'",
	)
	assert.Equal(t, namespace, &s.Namespace{}, "Namespace should be nil")
}

func TestDeleteNamespace(t *testing.T) {
	var (
		namespace *s.Namespace
		err       error
	)

	storage := &s.Storage{
		Namespaces: []*s.Namespace{
			{Name: "Namespace1"},
			{Name: "Namespace2"},
			{Name: "Namespace3"},
		},
	}

	assert.Equal(t, len(storage.Namespaces), 3)
	namespace, err = storage.FindNamespace("Namespace1")
	assert.Equal(t, err, nil, "Error should be nil")

	storage.DeleteNamespace(namespace)
	assert.Equal(t, len(storage.Namespaces), 2)
	namespace, err = storage.FindNamespace("Namespace1")
	assert.EqualError(
		t,
		err,
		"namespace not found: Namespace1",
		"Error should be 'namespace not found: Namespace1'")
	// Delete again :D
	storage.DeleteNamespace(namespace)
	assert.Equal(t, len(storage.Namespaces), 2)
}
