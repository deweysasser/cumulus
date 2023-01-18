package program

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"runtime"
)

// Options is the structure of program options
type Options struct {
	Version version `cmd:"" help:"Show program version"`
	// VersionCmd VersionCmd `name:"version" cmd:"" help:"show program version"`

	Account  Accounts `cmd:""`
	Instance struct {
		List InstanceList `cmd:""`
	} `cmd:""`
	Snapshot struct {
		List SnapshotList `cmd:""`
	} `cmd:""`
	MachineImage struct {
		List MachineImageList `cmd:""`
	} `cmd:""`
	Volume struct {
		List VolumeList `cmd:""`
	} `cmd:""`

	DNS struct {
		Zone struct {
			List ZoneList `cmd:""`
		} `cmd:""`
		Record struct {
			List RecordList `cmd:""`
		} `cmd:""`
	} `cmd:""`

	RDS struct {
		Cluster struct {
			List ClusterList `cmd:""`
		} `cmd:""`
	} `cmd:""`

	SNS struct {
		Topic struct {
			List TopicList `cmd:""`
		} `cmd:""`
		Subscription struct {
			List SubscriptionList `cmd:""`
		} `cmd:""`
	} `cmd:""`

	Debug        bool   `group:"Output" help:"Show debugging information"`
	OutputFormat string `group:"Output" enum:"auto,jsonl,terminal" default:"auto" help:"How to show program output (auto|terminal|jsonl)"`
	Quiet        bool   `group:"Output" short:"q" help:"Be less verbose than usual"`
	Verbose      bool   `group:"Output" short:"v" help:"Be more verbose than usual"`
	Profile      struct {
		CPU    bool `group:"Profile" help:"profile the CPU" hidden:""`
		Memory bool `group:"Profile" help:"profile the Memory usage" hidden:""`
	} `embed:"" prefix:"profile." hidden:""`
}

type version struct{}

func (version version) Run() error {
	fmt.Println(Version)
	return nil
}

// Parse calls the CLI parsing routines
func (program *Options) Parse(args []string) (*kong.Context, error) {
	parser, err := kong.New(program,
		kong.ShortUsageOnError(),
		// kong.Description("Brief Program Summary"),
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return parser.Parse(args)
}

// Run runs the program
func (program *Options) Run() error {
	return nil
}

// AfterApply runs after the options are parsed but before anything runs
func (program *Options) AfterApply() error {
	program.initLogging()
	return nil
}

func (program *Options) initLogging() {

	switch {
	case program.Debug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case program.Verbose:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case program.Quiet:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	var out io.Writer = os.Stderr

	if os.Getenv("TERM") == "" && runtime.GOOS == "windows" {
		out = colorable.NewColorableStdout()
	}

	if program.OutputFormat == "terminal" ||
		(program.OutputFormat == "auto" && isTerminal(os.Stderr)) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: out})
	} else {
		log.Logger = log.Output(out)
	}

	zerolog.DefaultContextLogger = &log.Logger

	log.Logger.Debug().
		Str("version", Version).
		Str("program", os.Args[0]).
		Msg("Starting")
}

// isTerminal returns true if the file given points to a character device (i.e. a terminal)
func isTerminal(file *os.File) bool {
	if fileInfo, err := file.Stat(); err != nil {
		log.Err(err).Msg("Error running stat")
		return false
	} else {
		return (fileInfo.Mode() & os.ModeCharDevice) != 0
	}
}
