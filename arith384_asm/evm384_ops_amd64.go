// Copyright Supranational LLC
// Licensed under the Apache License, Version 2.0, see LICENSE for details.
// SPDX-License-Identifier: Apache-2.0

package arith384_asm

//go:noescape
func AddMod384(ret, a, b, p *[6]uint64)

//go:noescape
func SubMod384(ret, a, b, p *[6]uint64)

//go:noescape
func MulMod384(ret, a, b, p *[6]uint64, inv uint64)
