// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2019-2021, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle)
// +build linux,386 linux,arm linux,mips linux,mipsle

package lock

import "golang.org/x/sys/unix"

func init() {
	setLk = unix.F_SETLK64
}
