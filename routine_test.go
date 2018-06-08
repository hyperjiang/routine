package routine

import (
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
