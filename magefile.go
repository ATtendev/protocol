//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/livekit/mageutil"
)

var Default = Proto

func Bootstrap() error {
	return mageutil.Run(context.Background(),
		"go install github.com/twitchtv/twirp/protoc-gen-twirp@latest",
		"go install google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	)
}

// regenerate protobuf
func Proto() error {
	twirpProtoFiles := []string{
		"vc_egress.proto",
		"vc_ingress.proto",
		"vc_room.proto",
		"vc_sip.proto",
		"vc_webhook.proto",
	}
	grpcProtoFiles := []string{
		"vc_agent.proto",
		"vc_analytics.proto",
		"vc_internal.proto",
		"vc_models.proto",
		"vc_rtc.proto",
		"vc_webhook.proto",
	}

	fmt.Println("generating protobuf")
	target := "vc"
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}
	protocGoPath, err := getToolPath("protoc-gen-go")
	if err != nil {
		return err
	}
	twirpPath, err := getToolPath("protoc-gen-twirp")
	if err != nil {
		return err
	}
	protocGrpcGoPath, err := getToolPath("protoc-gen-go-grpc")
	if err != nil {
		return err
	}
	fmt.Println("generating twirp protobuf")
	args := append([]string{
		"--go_out", target,
		"--twirp_out", target,
		"--go_opt=paths=source_relative",
		"--twirp_opt=paths=source_relative",
		"--plugin=go=" + protocGoPath,
		"--plugin=twirp=" + twirpPath,
		"-I=./protobufs",
	}, twirpProtoFiles...)
	cmd := exec.Command(protoc, args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("generating grpc protobuf")
	args = append([]string{
		"--go_out", target,
		"--go-grpc_out", target,
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		"--plugin=go=" + protocGoPath,
		"--plugin=go-grpc=" + protocGrpcGoPath,
		"-I=./protobufs",
	}, grpcProtoFiles...)
	cmd = exec.Command(protoc, args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// run tests
func Test() error {
	cmd := exec.Command("go", "test", "-race", "./...")
	connectStd(cmd)
	return cmd.Run()
}

// helpers

func getToolPath(name string) (string, error) {
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}
	// check under gopath
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	p := filepath.Join(gopath, "bin", name)
	if _, err := os.Stat(p); err != nil {
		return "", err
	}
	return p, nil
}

func connectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
