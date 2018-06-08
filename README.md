# Read flags from config file

[![GoDoc](https://godoc.org/github.com/hyperjiang/routine?status.svg)](https://godoc.org/github.com/hyperjiang/routine)
[![Build Status](https://travis-ci.org/hyperjiang/routine.svg?branch=master)](https://travis-ci.org/hyperjiang/routine)
[![](https://goreportcard.com/badge/github.com/hyperjiang/routine)](https://goreportcard.com/report/github.com/hyperjiang/routine)
[![codecov](https://codecov.io/gh/hyperjiang/routine/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/routine)


## Install

```
go get -u github.com/hyperjiang/routine
```

## Usage

```
package main

import (
	"fmt"

	"github.com/hyperjiang/routine"
)

func main() {

	total, max, size := 10000, 10, 1000

	r := routine.New(total, max, size)

	fmt.Println(r)

	var process = func(p routine.Processor) routine.ProcessResult {
		fmt.Printf("index: %d, offset: %d, size: %d\n", p.Index, p.Offset, p.Size)

		type rs struct {
			Index     int
			Processed int
		}

		processed := p.Size
		if (p.Index+1)*p.Size > p.Total {
			processed = p.Total - p.Index*p.Size
		}

		return rs{Index: p.Index, Processed: processed}
	}

	res := r.Wait(process)
	fmt.Println(res)
}
```
