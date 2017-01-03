package logic

import (
	log "github.com/Sirupsen/logrus"
	"github.com/golang/mock/gomock"
	"github.com/quintilesims/layer0/api/backend/mock_backend"
	"github.com/quintilesims/layer0/common/config"
	"github.com/quintilesims/layer0/common/db/job_store"
	"github.com/quintilesims/layer0/common/db/tag_store"
	"github.com/quintilesims/layer0/common/models"
	"github.com/quintilesims/layer0/common/testutils"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.SetTestConfig()
	log.SetLevel(log.FatalLevel)
	jobLogger.Level = log.FatalLevel
	retCode := m.Run()
	os.Exit(retCode)
}

type TestLogic struct {
	Backend  *mock_backend.MockBackend
	JobStore *job_store.MysqlJobStore
	TagStore *tag_store.MysqlTagStore
}

func NewTestLogic(t *testing.T) (*TestLogic, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	config := testutils.GetDBConfig()
	tagStore := tag_store.NewMysqlTagStore(config)
	if err := tagStore.Init(); err != nil {
		t.Fatal(err)
	}

	if err := tagStore.Clear(); err != nil {
		t.Fatal(err)
	}

	jobStore := job_store.NewMysqlJobStore(config)
	if err := jobStore.Init(); err != nil {
		t.Fatal(err)
	}

	if err := jobStore.Clear(); err != nil {
		t.Fatal(err)
	}

	logic := &TestLogic{
		Backend:  mock_backend.NewMockBackend(ctrl),
		JobStore: jobStore,
		TagStore: tagStore,
	}

	return logic, ctrl
}

func (l *TestLogic) AddTags(t *testing.T, tags []*models.Tag) {
	for _, tag := range tags {
		if err := l.TagStore.Insert(tag); err != nil {
			t.Fatal(err)
		}
	}
}

func (l *TestLogic) Logic() Logic {
	return *NewLogic(l.TagStore, l.JobStore, l.Backend)
}
