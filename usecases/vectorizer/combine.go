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

package vectorizer

import "github.com/weaviate/weaviate/entities/dto"

// CombineVectors combines all of the vector into sum of their parts
func CombineVectors[T dto.Embedding](vectors []T) T {
	switch any(vectors).(type) {
	case [][]float32:
		return any(CombineVectorsWithWeights(any(vectors).([][]float32), nil)).(T)
	default:
		return nil
	}
}

func CombineVectorsWithWeights(vectors [][]float32, weights []float32) []float32 {
	maxVectorLength := 0
	for i := range vectors {
		if len(vectors[i]) > maxVectorLength {
			maxVectorLength = len(vectors[i])
		}
	}
	sums := make([]float32, maxVectorLength)
	dividers := make([]int32, maxVectorLength)
	for indx, vector := range vectors {
		for i := 0; i < len(vector); i++ {
			if weights != nil {
				// apply weight to vector value
				sums[i] += vector[i] * weights[indx]
			} else {
				sums[i] += vector[i]
			}
			dividers[i]++
		}
	}
	for i := 0; i < len(sums); i++ {
		sums[i] /= float32(dividers[i])
	}

	return sums
}
