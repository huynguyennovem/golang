package main

import (
	"fmt"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/priorities"
)

type PriorityMetadataFactory struct {
}
func (pmf *PriorityMetadataFactory) PriorityMetadata() interface{} {
	return &priorityMetadata{
		totalNumNodes:           9,
	}
}

type priorityMetadata struct {
	totalNumNodes           int
}

type PriorityMetadataProducer func() interface{}

func NewPriorityMetadataFactory() PriorityMetadataProducer {
	factory := &PriorityMetadataFactory{}
	return factory.PriorityMetadata
}



func ImageLocalityPriorityMap(meta interface{}) int {
	priorityMeta := meta.(*priorityMetadata)
	return priorityMeta.totalNumNodes
}

func main() {

	NewPriorityMetadataFactory()

	// parse interface to struct
	temp := ImageLocalityPriorityMap
	fmt.Println(temp)


}
