//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package compressionhelpers

import (
	"github.com/weaviate/weaviate/adapters/repos/db/vector/hnsw/distancer/asm"
	"golang.org/x/sys/cpu"
)

func init() {
	if cpu.X86.HasAVX2 {
		l2SquaredByteImpl = asm.L2ByteAVX256
		dotByteImpl = asm.DotByteAVX256
	}
}
