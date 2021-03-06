package command

import (
	"fmt"
	"sort"

	"github.com/quintilesims/layer0/setup/aws"
	"github.com/quintilesims/layer0/setup/instance"
	"github.com/urfave/cli"
)

type status struct {
	Local  bool
	Remote bool
}

func (f *CommandFactory) List() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "list local and/or remote Layer0 instances",
		Flags: append(awsFlags,
			cli.BoolTFlag{
				Name:  "l, local",
				Usage: "show local Layer0 instances, denoted by 'l' (default: true)",
			},
			cli.BoolTFlag{
				Name:  "r, remote",
				Usage: "show remote Layer0 instances, denoted by 'r' (default: true)",
			}),
		Action: func(c *cli.Context) error {
			instances := map[string]status{}
			if c.Bool("remote") {
				if err := f.addRemoteInstances(c, instances); err != nil {
					return err
				}
			}

			if c.Bool("local") {
				if err := f.addLocalInstances(instances); err != nil {
					return err
				}
			}

			fmt.Println("STATUS \t NAME")
			sortAndIterate(instances, func(name string, status status) {
				switch {
				case status.Local && !status.Remote:
					fmt.Printf("l \t %s\n", name)
				case !status.Local && status.Remote:
					fmt.Printf("r \t %s\n", name)
				default:
					fmt.Printf("lr \t %s\n", name)
				}
			})

			return nil
		},
	}
}

func (f *CommandFactory) addRemoteInstances(c *cli.Context, current map[string]status) error {
	// The default AWS region is passed here (as opposed to other l0-setup
	// operations) because listing buckets from S3 is a region-agnostic
	// operation. All S3 buckets in the AWS account will be retrieved
	// regardless of what region is provided.
	provider, err := f.newAWSProviderHelper(c, aws.DEFAULT_AWS_REGION)
	if err != nil {
		return err
	}

	remote, err := instance.ListRemoteInstances(provider.S3)
	if err != nil {
		return err
	}

	for _, r := range remote {
		v, ok := current[r]
		if !ok {
			current[r] = status{Remote: true}
			continue
		}

		current[r] = status{Remote: true, Local: v.Local}
	}

	return nil
}

func (f *CommandFactory) addLocalInstances(current map[string]status) error {
	local, err := instance.ListLocalInstances()
	if err != nil {
		return err
	}

	for _, l := range local {
		v, ok := current[l]
		if !ok {
			current[l] = status{Local: true}
			continue
		}

		current[l] = status{Remote: v.Remote, Local: true}
	}

	return nil
}

func sortAndIterate(instances map[string]status, fn func(string, status)) {
	names := []string{}
	for name := range instances {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fn(name, instances[name])
	}
}
