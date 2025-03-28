// Copyright 2012-2024 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trafficctl

import (
	"fmt"
	"io"

	"github.com/florianl/go-tc"
)

const (
	BasicHelp = `tc filter ... basic [ match EMATCH_TREE ] [ action ACTION_SPEC ] [ classid CLASSID ]`
)

// ParseBasicParams parses the cmdline arguments for `tc filter ... basic ...`
// and returns a *tc.Object.
func ParseBasicParams(out io.Writer, params []string) (*tc.Object, error) {
	b := &tc.Basic{}
	var err error

	for i := 0; i < len(params); i = i + 2 {
		switch params[i] {
		case "match":
			return nil, ErrNotImplemented
		case "action":
			// Only generic actions allowed here
			b.Actions, err = ParseActionGAT(out, params[1:])
			if err != nil {
				return nil, err
			}
		case "classid", "flowid":
			id, err := ParseClassID(params[1])
			if err != nil {
				return nil, err
			}
			indirect := uint32(id)
			b.ClassID = &indirect
		case "help":
			fmt.Fprint(out, BasicHelp)
			return nil, nil
		default:
			return nil, ErrInvalidArg
		}
	}

	ret := &tc.Object{}
	ret.Kind = "basic"
	ret.Basic = b

	return ret, nil
}
