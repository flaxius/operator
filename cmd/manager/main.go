package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"runtime"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	sdkVersion "github.com/operator-framework/operator-sdk/version"
)

func printVersion() {
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
	log.Info(fmt.Sprintf("Version of operator-sdk: %v", sdkVersion.Version))

}

func main() {
	logrus.Info("Go Version: %s", runtime.Version())
	logrus.Info("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Info("Version of operator-sdk: %v", sdkVersion.Version)
}