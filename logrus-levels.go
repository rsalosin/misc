package main

import (
	"flag"
	"os"
	log "github.com/sirupsen/logrus"
)

func main() {
	var kind string
	flag.StringVar(&kind,"K","I","what to show: (T)race, (D)ebug, (I)nfo, (W)arn, (E)rror, (F)atal, (P)anic")
	flag.Parse()
	log.Println("Log levels example",kind)
	//f you wish to add the calling method as a field, instruct the logger via:
	log.SetReportCaller(true)
	  // Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	log.Println("Output to stdout")
	switch kind {
	case "T":
		log.SetLevel(log.TraceLevel)
		log.Trace("Something very low level.")
	case "D":
		log.SetLevel(log.DebugLevel)
		log.Debug("Useful debugging information.")
	case "I":
		log.Info("Something noteworthy happened!")
	case "W":
		log.Warn("You should probably take a look at this.")
	case "E":
		log.Error("Something failed but I'm not quitting.")
	case "F":
		// Calls os.Exit(1) after logging
		log.Fatal("Bye.")
	case "P":
		// Calls panic() after logging
		log.Panic("I'm bailing.")
	default:
		log.Info("Something noteworthy happened by default!")
	}
	log.Println("Log example end")
}
