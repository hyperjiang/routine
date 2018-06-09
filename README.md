# Helper for run goroutines

[![GoDoc](https://godoc.org/github.com/hyperjiang/routine?status.svg)](https://godoc.org/github.com/hyperjiang/routine)
[![Build Status](https://travis-ci.org/hyperjiang/routine.svg?branch=master)](https://travis-ci.org/hyperjiang/routine)
[![](https://goreportcard.com/badge/github.com/hyperjiang/routine)](https://goreportcard.com/report/github.com/hyperjiang/routine)
[![codecov](https://codecov.io/gh/hyperjiang/routine/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/routine)


## Install

```
go get -u github.com/hyperjiang/routine
```

## Usage example

```
package main

import (
	"fmt"

	"github.com/hyperjiang/routine"
)

func main() {

	total, goNum, size := 90741, 10, 1000

	rt := routine.New(total, goNum, size)
	fmt.Printf("%+v\n", rt)

	type processResult struct {
		Index      int
		ProcessNum int
	}
	var process = func(p routine.Processor) routine.ProcessResult {
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
	}

	fmt.Printf("total: %d\n", count)
}
```
