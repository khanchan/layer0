package command

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/quintilesims/layer0/client"
	"github.com/quintilesims/layer0/common/models"
	"github.com/urfave/cli"
)

const (
	FLAG_BI_DIRECTIONAL = "bi-directional"
)

type EnvironmentCommand struct {
	*CommandBase
}

func NewEnvironmentCommand(b *CommandBase) *EnvironmentCommand {
	return &EnvironmentCommand{b}
}

func (e *EnvironmentCommand) Command() cli.Command {
	return cli.Command{
		Name:  "environment",
		Usage: "manage layer0 environments",
		Subcommands: []cli.Command{
			{
				Name:      "create",
				Usage:     "create a new environment",
				Action:    e.create,
				ArgsUsage: "ENVIRONMENT_NAME",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "size",
						Value: "m3.medium",
						Usage: "size of the ec2 instances to use in the environment cluster",
					},
					cli.IntFlag{
						Name:  "min-count",
						Value: 0,
						Usage: "minimum number of instances allowed in the environment cluster",
					},
					cli.StringFlag{
						Name:  "user-data",
						Usage: "path to user data file",
					},
					cli.StringFlag{
						Name:  "os",
						Value: "linux",
						Usage: "specifies if the environment will run windows or linux containers",
					},
					cli.StringFlag{
						Name:  "ami",
						Usage: "specifies a custom AMI ID to use in the environment",
					},
				},
			},
			{
				Name:      "delete",
				Usage:     "delete an environment",
				ArgsUsage: "ENVIRONMENT_NAME",
				Action:    e.delete,
			},
			{
				Name:      "list",
				Usage:     "list all environments",
				Action:    e.list,
				ArgsUsage: " ",
			},
			{
				Name:      "get",
				Usage:     "describe an environment",
				Action:    e.read,
				ArgsUsage: "ENVIRONMENT_NAME",
			},
			{
				Name:      "setmincount",
				Usage:     "set the minimum instance count for an environment cluster",
				Action:    e.update,
				ArgsUsage: "ENVIRONMENT_NAME COUNT",
			},
			{
				Name:      "link",
				Usage:     "links two environments together",
				Action:    e.link,
				ArgsUsage: "SOURCE_ENVIRONMENT_NAME DESTINATION_ENVIRONMENT_NAME",
				Flags: []cli.Flag{
					cli.BoolTFlag{
						Name:  FLAG_BI_DIRECTIONAL,
						Usage: "specifies whether the link should be direcional",
					},
				},
			},
			{
				Name:      "unlink",
				Usage:     "unlinks two previously linked environments",
				Action:    e.unlink,
				ArgsUsage: "SOURCE_ENVIRONMENT_NAME DESTINATION_ENVIRONMENT_NAME",
				Flags: []cli.Flag{
					cli.BoolTFlag{
						Name:  FLAG_BI_DIRECTIONAL,
						Usage: "specifies whether the link should be direcional",
					},
				},
			},
		},
	}
}

func (e *EnvironmentCommand) create(c *cli.Context) error {
	args, err := extractArgs(c.Args(), "ENVIRONMENT_NAME")
	if err != nil {
		return err
	}

	var userData []byte
	if path := c.String("user-data"); path != "" {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		userData = content
	}

	req := models.CreateEnvironmentRequest{
		EnvironmentName:  args["ENVIRONMENT_NAME"],
		InstanceSize:     c.String("size"),
		MinClusterCount:  c.Int("min-count"),
		UserDataTemplate: userData,
		OperatingSystem:  c.String("os"),
		AMIID:            c.String("ami"),
	}

	if err := req.Validate(); err != nil {
		return err
	}

	jobID, err := e.client.CreateEnvironment(req)
	if err != nil {
		return err
	}

	return e.waitOnJobHelper(c, jobID, "creating", func(environmentID string) error {
		environment, err := e.client.ReadEnvironment(environmentID)
		if err != nil {
			return err
		}

		return e.printer.PrintEnvironments(environment)
	})
}

func (e *EnvironmentCommand) delete(c *cli.Context) error {
	return e.deleteHelper(c, "environment", func(environmentID string) (string, error) {
		return e.client.DeleteEnvironment(environmentID)
	})
}

func (e *EnvironmentCommand) list(c *cli.Context) error {
	environmentSummaries, err := e.client.ListEnvironments()
	if err != nil {
		return err
	}

	return e.printer.PrintEnvironmentSummaries(environmentSummaries...)
}

func (e *EnvironmentCommand) read(c *cli.Context) error {
	args, err := extractArgs(c.Args(), "ENVIRONMENT_NAME")
	if err != nil {
		return err
	}

	environmentIDs, err := e.resolver.Resolve("environment", args["ENVIRONMENT_NAME"])
	if err != nil {
		return err
	}

	environments := make([]*models.Environment, len(environmentIDs))
	for i, environmentID := range environmentIDs {
		environment, err := e.client.ReadEnvironment(environmentID)
		if err != nil {
			return err
		}

		environments[i] = environment
	}

	return e.printer.PrintEnvironments(environments...)
}

func (e *EnvironmentCommand) update(c *cli.Context) error {
	args, err := extractArgs(c.Args(), "ENVIRONMENT_NAME", "COUNT")
	if err != nil {
		return err
	}

	minClusterCount, err := strconv.Atoi(args["COUNT"])
	if err != nil {
		return fmt.Errorf("'%s' is not a valid integer", args["COUNT"])
	}

	id, err := e.resolveSingleEntityIDHelper("environment", args["ENVIRONMENT_NAME"])
	if err != nil {
		return err
	}

	req := models.UpdateEnvironmentRequest{MinClusterCount: &minClusterCount}

	if err := req.Validate(); err != nil {
		return err
	}

	jobID, err := e.client.UpdateEnvironment(id, req)
	if err != nil {
		return err
	}

	return e.waitOnJobHelper(c, jobID, "updating", func(environmentID string) error {
		environment, err := e.client.ReadEnvironment(environmentID)
		if err != nil {
			return err
		}

		return e.printer.PrintEnvironments(environment)
	})
}

func (e *EnvironmentCommand) link(c *cli.Context) error {
	generateAddLinkRequest := func(sourceEnvironmentID, destEnvironmentID string) (models.UpdateEnvironmentRequest, error) {
		env, err := e.client.ReadEnvironment(sourceEnvironmentID)
		if err != nil {
			return models.UpdateEnvironmentRequest{}, err
		}

		if !client.Contains(destEnvironmentID, env.Links) {
			appendedLinks := append(env.Links, destEnvironmentID)
		}

		req := models.UpdateEnvironmentRequest{Links: &appendedLinks}
		return req, nil
	}

	return e.updateEnvironmentLinksHelper(c, generateAddLinkRequest)
}

func (e *EnvironmentCommand) unlink(c *cli.Context) error {
	generateRemoveLinkRequest := func(sourceEnvironmentID, destEnvironmentID string) (models.UpdateEnvironmentRequest, error) {
		env, err := e.client.ReadEnvironment(sourceEnvironmentID)
		if err != nil {
			return models.UpdateEnvironmentRequest{}, err
		}

		updatedLinks := []string{}
		for _, link := range env.Links {
			if link != destEnvironmentID {
				updatedLinks = append(updatedLinks, link)
			}
		}

		req := models.UpdateEnvironmentRequest{Links: &updatedLinks}

		return req, nil
	}

	return e.updateEnvironmentLinksHelper(c, generateRemoveLinkRequest)
}

func (e *EnvironmentCommand) updateEnvironmentLinksHelper(c *cli.Context,
	generateReq func(string, string) (models.UpdateEnvironmentRequest, error)) error {
	args, err := extractArgs(c.Args(), "SOURCE_ENVIRONMENT_NAME", "DESTINATION_ENVIRONMENT_NAME")
	if err != nil {
		return err
	}

	sourceEnvironmentID, err := e.resolveSingleEntityIDHelper("environment", args["SOURCE_ENVIRONMENT_NAME"])
	if err != nil {
		return err
	}

	destEnvironmentID, err := e.resolveSingleEntityIDHelper("environment", args["DESTINATION_ENVIRONMENT_NAME"])
	if err != nil {
		return err
	}

	if sourceEnvironmentID == destEnvironmentID {
		return fmt.Errorf("Cannot link an environment to itself")
	}

	addLinkFN := func(sourceEnvID, destEnvID string) error {
		updateEnvReq, err := generateReq(sourceEnvID, destEnvID)
		if err != nil {
			return err
		}

		jobID1, err := e.client.UpdateEnvironment(sourceEnvID, updateEnvReq)
		if err != nil {
			return err
		}

		return e.waitOnJobHelper(c, jobID1, "updating", func(environmentID string) error {
			e.printer.Printf("Environment update successfull")
			return nil
		})
	}

	addLinkFN(sourceEnvironmentID, destEnvironmentID)

	if !c.Bool(FLAG_BI_DIRECTIONAL) {
		return nil
	}

	return addLinkFN(destEnvironmentID, sourceEnvironmentID)
}