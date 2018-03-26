package main

import (
	"time"

	"bitbucket.org/440-labs/influxdb-comparisons/query"
)

// TimescaleDBDevopsGroupby produces TimescaleDB-specific queries for the devops groupby case.
type TimescaleDBDevopsGroupby struct {
	TimescaleDBDevops
	numMetrics int
}

// NewTimescaleDBDevopsGroupBy produces a function that produces a new TimescaleDBDevopsGroupby for the given parameters
func NewTimescaleDBDevopsGroupBy(numMetrics int) QueryGeneratorMaker {
	return func(start, end time.Time) QueryGenerator {
		underlying := newTimescaleDBDevopsCommon(start, end)
		return &TimescaleDBDevopsGroupby{
			TimescaleDBDevops: *underlying,
			numMetrics:        numMetrics,
		}
	}
}

// Dispatch fills in the query.Query
func (d *TimescaleDBDevopsGroupby) Dispatch(scaleVar int) query.Query {
	q := query.NewTimescaleDB() // from pool
	d.MeanCPUMetricsDayByHourAllHostsGroupbyHost(q, d.numMetrics)
	return q
}