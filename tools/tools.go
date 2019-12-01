// +build tools

package tools

// tool dependencies
import (
	_ "github.com/CircleCI-Public/circleci-cli"
	_ "github.com/frapposelli/wwhrd"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/sqs/goreturns"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
