package commands

// import (
// 	"github.com/concourse/concourse/fly/rc"
// )

type JmctlCommand struct {
	Help HelpCommand `command:"help" description:"Print this help message"`

	// Target       rc.TargetName       `short:"t" long:"target" description:"Concourse target name"`
	// Targets      TargetsCommand      `command:"targets" alias:"ts" description:"List saved targets"`
	// DeleteTarget DeleteTargetCommand `command:"delete-target" alias:"dtg" description:"Delete target"`
	// EditTarget   EditTargetCommand   `command:"edit-target" alias:"etg" description:"Edit a target"`

	Version func() `short:"v" long:"version" description:"Print the version of jmctl and exit"`
	Verbose bool   `long:"verbose" description:"Print API requests and responses"`

	Templates TemplatesCommand `command:"templates"      alias:"ts" description:"List available templates"`
	// AbortBuild AbortBuildCommand `command:"abort-build" alias:"ab" description:"Abort a build"`
	// RerunBuild RerunBuildCommand `command:"rerun-build" alias:"rb" description:"Rerun a build"`

	// TriggerJob TriggerJobCommand `command:"trigger-job" alias:"tj" description:"Start a job in a pipeline"`

	// Volumes VolumesCommand `command:"volumes" alias:"vs" description:"List the active volumes"`

	// Description     WorkersCommand     `command:"workers" alias:"ws" description:"List the registered workers"`
	// LandWorker  LandWorkerCommand  `command:"land-worker" alias:"lw" description:"Land a worker"`
	// PruneWorker PruneWorkerCommand `command:"prune-worker" alias:"pw" description:"Prune a stalled, landing, landed, or retiring worker"`

	// Curl CurlCommand `command:"curl" alias:"c" description:"curl the api"`

	// Completion CompletionCommand `command:"completion" description:"generate shell completion code"`
}

var JmCtl JmctlCommand
