package proxmox

import (
	"context"
	"fmt"
)

func (cl *Cluster) ListBackupJobs(ctx context.Context) ([]BackupJob, error) {
	var backupJobs []BackupJob
	if err := cl.client.Get(ctx, "/cluster/backup", &backupJobs); err != nil {
		return nil, fmt.Errorf("list cluster backup jobs: %w", err)
	}

	return backupJobs, nil
}
