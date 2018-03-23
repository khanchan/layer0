package lock

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type LockSchema struct {
	LockID   string
	Acquired int64
}

type DynamoLock struct {
	table  dynamo.Table
	expiry time.Duration
}

func NewDynamoLock(session *session.Session, table string, expiry time.Duration) *DynamoLock {
	return &DynamoLock{
		table:  dynamo.New(session).Table(table),
		expiry: expiry,
	}
}

func (d *DynamoLock) Acquire(lockID string) (bool, error) {
	lock := LockSchema{
		LockID:   lockID,
		Acquired: time.Now().UnixNano(),
	}

	if err := d.table.Put(lock).If("attribute_not_exists(LockID)").Run(); err != nil {
		if err, ok := err.(awserr.Error); ok && err.Code() == "ConditionalCheckFailedException" {
			return d.acquireIfExpired(lockID)
		}

		return false, err
	}

	return true, nil
}

func (d *DynamoLock) acquireIfExpired(lockID string) (bool, error) {
	oldestPossibleAcquiredTime := time.Now().Add(-d.expiry)

	if err := d.table.Update("LockID", lockID).
		Set("Acquired", time.Now().UnixNano()).
		If("'Acquired' <= ?", oldestPossibleAcquiredTime.UnixNano()).
		Run(); err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (d *DynamoLock) Release(lockID string) error {
	return d.table.Delete("LockID", lockID).Run()
}

func (d *DynamoLock) List() ([]LockSchema, error) {
	locks := []LockSchema{}
	if err := d.table.Scan().
		Consistent(false).
		All(&locks); err != nil {
		return nil, err
	}

	return locks, nil
}

func (d *DynamoLock) clear() error {
	locks, err := d.List()
	if err != nil {
		return err
	}

	for _, lock := range locks {
		if err := d.Release(lock.LockID); err != nil {
			return err
		}
	}

	return nil
}