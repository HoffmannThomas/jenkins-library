package cmd

import (
	"strings"

	"github.com/HoffmannThomas/jenkins-library/pkg/command"
	"github.com/HoffmannThomas/jenkins-library/pkg/log"
	"github.com/HoffmannThomas/jenkins-library/pkg/telemetry"
)

func karmaExecuteTests(config karmaExecuteTestsOptions, telemetryData *telemetry.CustomData) {
	c := command.Command{}
	// reroute command output to loging framework
	// also log stdout as Karma reports into it
	c.Stdout(log.Writer())
	c.Stderr(log.Writer())
	runKarma(config, &c)
}

func runKarma(config karmaExecuteTestsOptions, command command.ExecRunner) {
	installCommandTokens := tokenize(config.InstallCommand)
	command.SetDir(config.ModulePath)
	err := command.RunExecutable(installCommandTokens[0], installCommandTokens[1:]...)
	if err != nil {
		log.Entry().
			WithError(err).
			WithField("command", config.InstallCommand).
			Fatal("failed to execute install command")
	}

	runCommandTokens := tokenize(config.RunCommand)
	command.SetDir(config.ModulePath)
	err = command.RunExecutable(runCommandTokens[0], runCommandTokens[1:]...)
	if err != nil {
		log.Entry().
			WithError(err).
			WithField("command", config.RunCommand).
			Fatal("failed to execute run command")
	}
}

func tokenize(command string) []string {
	return strings.Split(command, " ")
}
