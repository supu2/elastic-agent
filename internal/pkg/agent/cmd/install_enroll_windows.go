// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build windows

package cmd

import (
	"os/exec"

	"github.com/elastic/elastic-agent/pkg/utils"
)

func enrollCmdExtras(cmd *exec.Cmd, ownership utils.FileOwner) error {
	// TODO: Add ability to call enroll as non-Administrator on Windows.
	return nil
}
