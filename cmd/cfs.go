package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"datanode"
	"metanode"
	"volmgr"

	"util"
)

const (
	Version = "0.1"
)

var (
	configFile = flag.String("c", "", "config file path")
)

func main() {
	log.Println("*The Container Filesystem*")

	flag.Parse()
	cfg, e := util.NewConfig(*configFile)
	if e != nil {
		log.Println("Bad config file, ", e)
		return
	}
	role := cfg.GetString("role")

	//for multi-cpu scheduling
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	profPort := cfg.GetString("prof")
	if len(profPort) != 0 {
		go func() {
			log.Println(http.ListenAndServe(":"+profPort, nil))
		}()
	}

	switch role {
	case "volmgr":
		server, err := volmgr.NewServer(cfg)
		if err == nil {
			server.Start()
		}

	case "metanode":
		server, err := metanode.NewServer(cfg)
		if err == nil {
			server.Start()
		}

	case "datanode":
		server, err := datanode.NewServer(cfg)
		if err == nil {
			server.Start()
		}

	default:
		log.Println("Invalid Server Role")
	}

	return
}
