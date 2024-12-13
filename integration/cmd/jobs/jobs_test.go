package jobs_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/databricks/cli/internal/testcli"
	"github.com/databricks/cli/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateJob(t *testing.T) {
	testutil.Require(t, testutil.Azure)
	stdout, stderr := testcli.RequireSuccessfulRun(t, "jobs", "create", "--json", "@testdata/create_job_without_workers.json", "--log-level=debug")
	assert.Empty(t, stderr.String())
	var output map[string]int
	err := json.Unmarshal(stdout.Bytes(), &output)
	require.NoError(t, err)
	testcli.RequireSuccessfulRun(t, "jobs", "delete", fmt.Sprint(output["job_id"]), "--log-level=debug")
}
