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

func (cl *Cluster) DeleteBackupJob(ctx context.Context, id string) error {
	if err := cl.client.Delete(ctx, fmt.Sprintf("/cluster/backup/%s", id), nil); err != nil {
		return fmt.Errorf("delete cluster backup job: %w", err)
	}

	return nil
}

func (cl *Cluster) CreateBackupJob(ctx context.Context, job BackupJob) error {
	if err := cl.client.Post(ctx, "/cluster/backup", &job, nil); err != nil {
		return fmt.Errorf("create backup job: %w", err)
	}

	return nil
}

func (cl *Cluster) UpdateBackupJob(ctx context.Context, id string, req UpdateBackupJobRequest) error {
	if err := cl.client.Put(ctx, fmt.Sprintf("/cluster/backup/%s", id), &req, nil); err != nil {
		return fmt.Errorf("update cluster backup job: %w", err)
	}

	return nil
}
