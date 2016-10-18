package main

import (
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/jondot/runs"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// VERSION gets overwritten by release target
var VERSION = "dev"
var _ = VERSION

var (
	distance       = kingpin.Flag("distance", "Distance between groups.").Short('d').Required().Int()
	gpat           = kingpin.Flag("glob", "Glob for files with a pattern.").Short('g').Default("*.jpg").String()
	outdir         = kingpin.Flag("out", "Where to build groups.").Short('o').Default("out").String()
	prefix         = kingpin.Flag("prefix", "Group directory prefix.").Short('p').Default("group_").String()
	format         = kingpin.Flag("format", "Use file names and specify a regular expression to capture ordinals.").Short('f').Default("").String()
	layout         = kingpin.Flag("layout", "If provided, treat captures as time and this is the string to parse it.").Short('l').Default("").String()
	command        = kingpin.Flag("group-command", "Run a command on each group.").String()
	shouldCreate   = kingpin.Flag("create", "Create the groups on disk by COPYING files.").Short('c').Bool()
	shouldMove     = kingpin.Flag("move", "Create the groups on disk by MOVING files.").Short('m').Bool()
	isExif         = kingpin.Flag("exif", "Use EXIF date.").Short('e').Bool()
	isFiletime     = kingpin.Flag("filetime", "Use file time.").Short('t').Bool()
	shouldSequence = kingpin.Flag("sequence", "Add sequence number to file names per group.").Short('s').Bool()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(VERSION).Author("Dotan Nahum (@jondot)")
	kingpin.CommandLine.Help = `Magical file grouping.

Examples:

	Sort media:
	$ sq -d 5000 -g '*.avi' -m

	Group photos by shooting session:
	$ sq -d 1800 -g '*.*' --exif -c

	Stitch security camera frames:
	$ sq -d 300 -g '*.*' -t -c -s --group-command 'ffmpeg -framerate 2 -i "{{GROUP}}/%d.jpg" -c:v libx264 {{GROUP}}.avi'
	`
	kingpin.Parse()

	matches, err := filepath.Glob(*gpat)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	total := len(matches)
	log.Printf("Found %d file(s).", total)
	if total == 0 {
		os.Exit(0)
	}

	c := ResolveConverter(*format, *layout, *isExif, *isFiletime)
	dist := float64(*distance)

	things := []interface{}{}
	for _, file := range matches {
		things = append(things, file)
	}
	grouped := runs.Detect(things, func(thing interface{}) uint64 {
		return c.Convert(thing.(string))
	}, func(a, b uint64) bool {
		return math.Abs(float64(a-b)) < dist
	})

	walker := &runs.GroupWalker{}
	walker.Walk(grouped, &GroupPrinter{})

	// Following duplicates code, can DRY this, but for now prefer verbosity.
	if *shouldCreate {
		walker.Walk(grouped, &GroupFileOp{
			OutputDir:      *outdir,
			GroupPrefix:    *prefix,
			ShouldSequence: *shouldSequence,
			FileOp:         OpCopy,
		})
		log.Printf("Created groups.")
	} else if *shouldMove {
		walker.Walk(grouped, &GroupFileOp{
			OutputDir:      *outdir,
			GroupPrefix:    *prefix,
			ShouldSequence: *shouldSequence,
			FileOp:         OpMove,
		})
		log.Printf("Moved into groups.")
	}
	if *command != "" {
		walker.Walk(grouped, &GroupCommandRunner{OutputDir: *outdir, GroupPrefix: *prefix, Command: *command})
		log.Printf("Ran all commands.")
	}
}
