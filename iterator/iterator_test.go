package iterator

import "testing"

func TestIterator(t *testing.T){
	var aggregate Aggregate
	aggregate = NewNumbers(10,100)
	IteratorPrint(aggregate.Iterator())
}
