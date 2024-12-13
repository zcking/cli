package storage_credentials_test

import (
	"testing"

	"github.com/databricks/cli/internal/acc"
	"github.com/databricks/cli/internal/testcli"
	"github.com/databricks/cli/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestStorageCredentialsListRendersResponse(t *testing.T) {
	_, _ = acc.WorkspaceTest(t)

	// Check if metastore is assigned for the workspace, otherwise test will fail
	t.Log(testutil.GetEnvOrSkipTest(t, "TEST_METASTORE_ID"))

	stdout, stderr := testcli.RequireSuccessfulRun(t, "storage-credentials", "list")
	assert.NotEmpty(t, stdout)
	assert.Empty(t, stderr)
}