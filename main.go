package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "webhook plugin"
	app.Usage = "webhook plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:   "urls",
			Usage:  "list of urls",
			EnvVar: "PLUGIN_URLS",
		},
		cli.StringFlag{
			Name:   "method",
			Usage:  "webhook method",
			EnvVar: "PLUGIN_METHOD",
			Value:  "POST",
		},
		cli.StringFlag{
			Name:   "bearer",
			Usage:  "bearer token",
			EnvVar: "PLUGIN_BEARER",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "username for basic auth",
			EnvVar: "PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "password for basic auth",
			EnvVar: "PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "content-type",
			Usage:  "content type",
			EnvVar: "PLUGIN_CONTENT_TYPE",
			Value:  "application/json",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug information",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.BoolFlag{
			Name:   "skip-verify",
			Usage:  "skip ssl verification",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:     c.String("build.tag"),
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Branch:  c.String("commit.branch"),
			Author:  c.String("commit.author"),
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
			Started: c.Int64("build.started"),
			Created: c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Method:      c.String("method"),
			Username:    c.String("username"),
			Password:    c.String("password"),
			ContentType: c.String("content-type"),
			Headers:     c.StringSlice("headers"),
			URLs:        c.StringSlice("urls"),
			SkipVerify:  c.Bool("skip-verify"),
			BearerToken: c.String("bearer"),
			Debug:       c.Bool("debug"),
		},
	}
	return plugin.Exec()
}
