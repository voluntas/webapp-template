package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/voluntas/webapp"
	"golang.org/x/sync/errgroup"
)

var (
	revision string = "air"
	version  string = ""
)

var (
	g errgroup.Group
)

func main() {
	configFilePath := flag.String("C", "config.yaml", "webapp の設定ファイルへのパス")
	v := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *v {
		fmt.Printf(version, revision)
		os.Exit(0)
	}

	buf, err := os.ReadFile(*configFilePath)
	if err != nil {
		// 読み込めない場合 Fatal で終了
		log.Fatal("cannot open config file, err=", err)
	}

	var config webapp.Config
	if err := webapp.InitConfig(buf, &config); err != nil {
		// パースに失敗した場合 Fatal で終了
		log.Fatal("cannot parse config file, err=", err)
	}

	if err := webapp.InitLogger(config); err != nil {
		log.Fatal("cannot parse config file, err=", err)
	}

	conn, err := webapp.NewConn(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	defer conn.Close()

	server, err := webapp.NewServer(revision, config, conn)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	g.Go(func() error {
		return server.Start(config.ServerIPAddress, config.ServerPort)
	})

	g.Go(func() error {
		return server.StartExporter(config.ExporterIPAddress, config.ExporterPort)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
