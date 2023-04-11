package cmd

import (
	"ftpctl/pkg"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:                   "ftpctl",
		UseShortOptionHandling: true,
		Usage:                  "it is a ftp server command line configuration tool",
		Version:                "0.1",
	}

	app.Commands = []*cli.Command{
		{
			Name:        "upload",
			Description: "upload file to the remote ftp server of the dirctory, the default directory is /",
			Flags:       getFlag(),
			Action:      action,
		},
		{
			Name:        "download",
			Description: "from the remote ftp server download file to the current dirctory",
			Flags:       getFlag(),
			Action:      action,
		},
		{
			Name:        "delete",
			Description: "from the remote ftp server delete file",
			Flags:       getFlag(),
			Action:      action,
		},
		{
			Name:        "mkdir",
			Description: "from the remote ftp server create the new dirctory",
			Flags:       getFlag(),
			Action:      action,
		},
		{
			Name:        "rmdir",
			Description: "recursive deletion of directories from the remote ftp server, if file exists in the directory it will also be cleared",
			Flags:       getFlag(),
			Action:      action,
		},
		{
			Name:        "get",
			Description: "get all the files in the current directory from the remote ftp server, the default directory is /",
			Flags:       getFlag(),
			Action:      action,
		},
		{
			Name:        "rename",
			Description: "rename files from the remote ftp server",
			Flags:       getFlag(),
			Action:      action,
		},
	}
	return app
}

func action(c *cli.Context) error {
	pkg.Host = c.String("host")
	pkg.Username = c.String("username")
	pkg.Password = c.String("password")
	pkg.SrcFile = c.String("srcfile")
	pkg.DesFile = c.String("desfile")
	pkg.Path = c.String("path")
	switch c.Command.Name {
	case "upload":
		if err := upload(); err != nil {
			return err
		}
	case "download":
		if err := download(); err != nil {
			return err
		}
	case "delete":
		if err := delete(); err != nil {
			return err
		}
	case "mkdir":
		if err := mkdir(); err != nil {
			return err
		}
	case "get":
		if err := get(); err != nil {
			return err
		}
	case "rmdir":
		if err := rmdir(); err != nil {
			return err
		}
	case "rename":
		if err := rename(); err != nil {
			return err
		}
	}
	return nil
}

func upload() error {
	return pkg.Upload()
}

func download() error {
	return pkg.Download()
}

func delete() error {
	return pkg.Delete()
}

func mkdir() error {
	return pkg.MakeDir()
}

func rmdir() error {
	return pkg.RemoveDir()
}

func get() error {
	return pkg.Get()
}
func rename() error {
	return pkg.Rename()
}

func getFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "host",
			Value:   "127.0.0.1",
			Usage:   "Host ip address for remote connection to ftp server.",
			Aliases: []string{"a"},
		},
		&cli.StringFlag{
			Name:    "username",
			Value:   "admin",
			Usage:   "The username of login ftp server.",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    "password",
			Value:   "admin",
			Usage:   "The password of login ftp server.",
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:    "srcfile",
			Value:   "",
			Usage:   "This file is used for upload or download.",
			Aliases: []string{"s"},
		},
		&cli.StringFlag{
			Name:    "desfile",
			Value:   "",
			Usage:   "Rename the file to a new name.",
			Aliases: []string{"d"},
		},
		&cli.StringFlag{
			Name:    "path",
			Value:   "",
			Usage:   "The dirctory is used for upload or download.",
			Aliases: []string{"P"},
		},
	}
}
