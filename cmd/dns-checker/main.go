package main

import (
	"net"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/coreos/etcd-operator/pkg/util/constants"
	"github.com/coreos/etcd-operator/version"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("etcd-backup-operator Version: %v", version.Version)
	logrus.Infof("Git SHA: %s", version.GitSHA)

	args := os.Args[1:]

	dnsLookupIntervalStr := os.Getenv(constants.EnvDnsLookupInterval)
	dnsLookupInterval := constants.DefaultDnsInterval
	if len(dnsLookupIntervalStr) != 0 {
		tmp, err := strconv.Atoi(dnsLookupIntervalStr)
		if err == nil && tmp > 0 {
			dnsLookupInterval = time.Duration(tmp) * time.Second
		}
	}

	logrus.Info("DNS entries: ", args)
	logrus.Info("DNS lookup interval: ", dnsLookupInterval)

	for _, dnsEntry := range args {
		for {
			_, err := net.LookupHost(dnsEntry)
			if err == nil {
				logrus.Info("DNS lookup succeed: ", dnsEntry)
				break
			} else {
				logrus.Warn("DNS lookup failed: ", dnsEntry)
				time.Sleep(dnsLookupInterval)
			}
		}
	}

	logrus.Info("DNS lookup succeeded for all entries. Exiting...")
}
