package docker

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/strangelove-ventures/heighliner/builder"

	"github.com/stretchr/testify/require"
)

const HyperlaneImageName = "hyperlane-simd"

// Build a docker image where 'buildDir' is the local path to the source files you want to build from,
// and dockerfilePath is where the docker file is.
func BuildHeighlinerHyperlaneImage(imageName, tarFilePath, buildDir, goModPath, dockerfilePath string) {
	local := true
	buildConfig := builder.HeighlinerDockerBuildConfig{
		SkipPush: true,
	}
	chainNodeConfig := builder.ChainNodeConfig{
		GoModPath:          goModPath, // This is used BEFORE the docker build starts. Its used to look up some ARGS that will be used during the build.
		Name:               imageName,
		GithubOrganization: "strangelove-ventures",
		GithubRepo:         "hyperlane-cosmos",
		BuildDir:           buildDir,       // This path should be relative to the TarFilePath. (normally should just be '.')
		DockerfilePath:     dockerfilePath, // This path should be relative to the TarFilePath. (E.g. depends on the docker context)
		Dockerfile:         "cosmos",       // This MUST be set to choose a BASE_VERSION (of the image)
		BuildTarget:        "make build",
		Binaries:           []string{"build/simd"},
		BuildEnv:           []string{"BUILD_TAGS=muslc"},
		TarFilePath:        tarFilePath, // This is used to determine which directory to tarball for the docker build.
	}
	chainConfig := builder.ChainNodeDockerBuildConfig{
		Build: chainNodeConfig,
	}
	chainQueuedBuilds := builder.HeighlinerQueuedChainBuilds{ChainConfigs: []builder.ChainNodeDockerBuildConfig{chainConfig}}
	heighlinerBuilder := builder.NewHeighlinerBuilder(buildConfig, 1, local, false)
	heighlinerBuilder.AddToQueue(chainQueuedBuilds)
	heighlinerBuilder.BuildImages()
}

type dockerLogLine struct {
	Stream      string            `json:"stream"`
	Aux         any               `json:"aux"`
	Error       string            `json:"error"`
	ErrorDetail dockerErrorDetail `json:"errorDetail"`
}

type dockerErrorDetail struct {
	Message string `json:"message"`
}

func handleDockerBuildOutput(t *testing.T, body io.Reader) {
	var logLine dockerLogLine

	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		logLine.Stream = ""
		logLine.Aux = nil
		logLine.Error = ""
		logLine.ErrorDetail = dockerErrorDetail{}

		line := scanner.Text()

		_ = json.Unmarshal([]byte(line), &logLine)
		if logLine.Stream != "" {
			fmt.Print(logLine.Stream)
		}
		if logLine.Aux != nil {
			fmt.Print(logLine.Aux)
		}
		if logLine.Error != "" {
			fmt.Print(logLine.Error)
		}
		if logLine.ErrorDetail.Message != "" {
			fmt.Print(logLine.ErrorDetail.Message)
		}
	}

	require.Equalf(t, "", logLine.Error, "docker image build error: %s", logLine.ErrorDetail.Message)
}
