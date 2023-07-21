// Copyright (c) 2023 BlockDev AG
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/svc/eventlog"
	"golang.zx2c4.com/wireguard/windows/conf"
	"golang.zx2c4.com/wireguard/windows/tunnel"
)

func setupEventLog(loggerName string) (*eventlog.Log, error) {
	err := eventlog.InstallAsEventCreate(loggerName, eventlog.Info|eventlog.Error|eventlog.Warning)
	if err != nil && !strings.Contains(err.Error(), "registry key already exists") {
		return nil, err
	}
	return eventlog.Open(loggerName)
}

func main() {
	wlog, err := setupEventLog("wireguard_svc")
	if err != nil {
		log.Fatalf("Could not setup event logger: %v", err)
	}

	runAsService := flag.Bool("service", false, "run as windows service")
	configFile := flag.String("config-file", "", "path to wireguard config file")
	flag.Parse()
	if !*runAsService || *configFile == "" {
		flag.Usage()
		log.Fatalln()
	}

	conf.PresetRootDirectory(filepath.Dir(*configFile))
	if err := tunnel.Run(*configFile); err != nil {
		_ = wlog.Error(1, fmt.Sprintf("Service run error: %v", err))
	}
}
