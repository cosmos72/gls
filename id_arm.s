// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "textflag.h" // for NOSPLIT
#include "../../src/runtime/go_tls.h"

TEXT ·GoId(SB),NOSPLIT,$0-4
	MOVW g, goid+0(FP)
	RET
