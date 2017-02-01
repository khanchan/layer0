package framework

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/quintilesims/sts/models"
	"testing"
)

type SystemTestService struct {
	T   *testing.T
	URL string
}

func NewSystemTestService(t *testing.T, url string) *SystemTestService {
	return &SystemTestService{
		T:   t,
		URL: fmt.Sprintf("http://%s", url),
	}
}

func (s *SystemTestService) Die() {
	req := models.SetHealthRequest{
		Mode: "die",
	}

	var health models.Health
	s.execute(s.sling().Post("/health").BodyJSON(req), &health)
}

func (s *SystemTestService) sling() *sling.Sling {
	return sling.New().Base(s.URL)
}

func (s *SystemTestService) execute(sling *sling.Sling, v interface{}) {
	var jsonerr struct{ Error string }
	resp, err := sling.Receive(v, &jsonerr)
	if err != nil {
		s.T.Fatalf("Error executing STS: %v", err)
	}

	if err := jsonerr.Error; err != "" {
		s.T.Fatalf("Error from STS: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		s.T.Fatalf("STS returned invalid status code :%s", resp.Status)
	}
}