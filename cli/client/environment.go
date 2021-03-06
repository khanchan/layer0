package client

import (
	"github.com/quintilesims/layer0/common/models"
)

func (c *APIClient) CreateEnvironment(name, instanceSize string, minCount int, userData []byte, os, amiID string) (*models.Environment, error) {
	req := models.CreateEnvironmentRequest{
		EnvironmentName:  name,
		InstanceSize:     instanceSize,
		MinClusterCount:  minCount,
		UserDataTemplate: userData,
		OperatingSystem:  os,
		AMIID:            amiID,
	}

	var environment *models.Environment
	if err := c.Execute(c.Sling("environment/").Post("").BodyJSON(req), &environment); err != nil {
		return nil, err
	}

	return environment, nil
}

func (c *APIClient) DeleteEnvironment(id string) (string, error) {
	jobID, err := c.ExecuteWithJob(c.Sling("environment/").Delete(id))
	if err != nil {
		return "", err
	}

	return jobID, nil
}

func (c *APIClient) GetEnvironment(id string) (*models.Environment, error) {
	var environment *models.Environment
	if err := c.Execute(c.Sling("environment/").Get(id), &environment); err != nil {
		return nil, err
	}

	return environment, nil
}

func (c *APIClient) ListEnvironments() ([]*models.EnvironmentSummary, error) {
	var environments []*models.EnvironmentSummary
	if err := c.Execute(c.Sling("environment/").Get(""), &environments); err != nil {
		return nil, err
	}

	return environments, nil
}

func (c *APIClient) UpdateEnvironment(id string, minCount int) (*models.Environment, error) {
	req := models.UpdateEnvironmentRequest{
		MinClusterCount: minCount,
	}

	var environment *models.Environment
	if err := c.Execute(c.Sling("environment/").Put(id).BodyJSON(req), &environment); err != nil {
		return nil, err
	}

	return environment, nil
}

func (c *APIClient) CreateLink(sourceID string, destinationID string) error {
	req := models.CreateEnvironmentLinkRequest{
		EnvironmentID: destinationID,
	}

	var resp string
	if err := c.Execute(c.Sling("environment/").Post(sourceID+"/link").BodyJSON(req), &resp); err != nil {
		return err
	}

	return nil
}

func (c *APIClient) DeleteLink(sourceID string, destinationID string) error {
	var resp string
	if err := c.Execute(c.Sling("environment/").Delete(sourceID+"/link/"+destinationID), &resp); err != nil {
		return err
	}

	return nil
}
