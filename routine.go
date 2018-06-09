package routine

import (
	"math"
)

// Routine - the struct of goroutines
type Routine struct {
	GoNum int // the number of goroutines
	Size  int // how many records for each goroutine to process
	Total int // total number of records
}

// Processor - the struct of processor
type Processor struct {
	Index  int // the index of goroutine
	Size   int // how many records for each goroutine to process
	Offset int // offset = index * size
	Total  int // total number of records
}

// ProcessResult - the process result
type ProcessResult interface{}

// ProcessFunc - the process function
type ProcessFunc func(p Processor) ProcessResult

// New - get a new routine, note that the goroutine number and record number per goroutine
// is calculated dynamically, and probably will not equal to the input
func New(totalRecords, maxGoNum, maxSizePerGo int) *Routine {
	num := math.Ceil(float64(totalRecords) / float64(maxSizePerGo))
	goNum := int(math.Min(num, float64(maxGoNum)))
	sizePerGo := int(math.Ceil(float64(totalRecords) / float64(goNum)))
	return &Routine{GoNum: goNum, Size: sizePerGo, Total: totalRecords}
}

// Wait - run the processors, wait and gather all the results from them
func (r *Routine) Wait(f ProcessFunc) []ProcessResult {

	var result []ProcessResult

	var ch = make(chan ProcessResult, r.GoNum)

	for i := 0; i < r.GoNum; i++ {
		go func(i int) {
			p := Processor{
				Index:  i,
				Size:   r.Size,
				Offset: i * r.Size,
				Total:  r.Total,
			}
			ch <- f(p)
		}(i)
	}

	for i := 0; i < r.GoNum; i++ {
		result = append(result, <-ch)
	}

	return result
}
