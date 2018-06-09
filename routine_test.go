package routine

import (
	"fmt"
	"testing"
)

func TestRoutine(t *testing.T) {

	var process = func(p Processor) ProcessResult {
		processed := p.Size
		if (p.Index+1)*p.Size > p.Total {
			processed = p.Total - p.Index*p.Size
		}

		return processed
	}

	total := 10000
	r := New(total, 10, 1000)
	sum := 0
	res := r.Wait(process)
	for _, v := range res {
		sum += v.(int)
	}
	if sum != total {
		t.Errorf("sum = %v, want %v", sum, total)
	}

	total = 5001
	r = New(total, 10, 1000)
	sum = 0
	res = r.Wait(process)
	for _, v := range res {
		sum += v.(int)
	}
	if sum != total {
		t.Errorf("sum = %v, want %v", sum, total)
	}

	total = 999999999999999
	r = New(total, 10, 1000)
	sum = 0
	res = r.Wait(process)
	for _, v := range res {
		sum += v.(int)
	}
	if sum != total {
		t.Errorf("sum = %v, want %v", sum, total)
	}
}

func TestRoutineDemo(t *testing.T) {
	total, goNum, size := 90741, 10, 1000

	rt := New(total, goNum, size)
	fmt.Printf("%+v\n", rt)

	type processResult struct {
		Index      int
		ProcessNum int
	}
	var process = func(p Processor) ProcessResult {
		fmt.Printf("index: %d, offset: %d, size: %d, total: %d\n", p.Index, p.Offset, p.Size, p.Total)

		processNum := 0
		offset := p.Offset
		limit := 1000
		maxOffset := p.Offset + p.Size
		if maxOffset > p.Total {
			maxOffset = p.Total
		}

		for offset < maxOffset {
			if offset+limit > maxOffset {
				limit = maxOffset - offset
			}

			sql := "select * from some_table limit %d, offset %d\n"
			fmt.Printf(sql, limit, offset)

			offset += limit
			processNum += limit
		}

		result := &processResult{
			Index:      p.Index,
			ProcessNum: processNum,
		}

		return result
	}

	var count int
	res := rt.Wait(process)
	for _, v := range res {
		r := v.(*processResult)
		count += r.ProcessNum

		if r.Index == 9 && r.ProcessNum != 9066 { // the last goroutine
			t.Errorf("processNum = %v, want %v", r.ProcessNum, 9066)
		}
	}

	if count != total {
		t.Errorf("count = %v, want %v", count, total)
	}
}
