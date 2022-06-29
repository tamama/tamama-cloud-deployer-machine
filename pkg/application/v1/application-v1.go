package applicationv1

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
)

type ApplicationV1 struct {
	Application *cli.App
}

func NewApplicationV1(ctx context.Context) (instance *ApplicationV1) {
	instance = &ApplicationV1{}

	instance.Application = instance._NewApp()
	
	return instance
}

func (_self *ApplicationV1) Run(ctx context.Context, err chan error) {
	err <- _self.Application.Run(os.Args)
}
