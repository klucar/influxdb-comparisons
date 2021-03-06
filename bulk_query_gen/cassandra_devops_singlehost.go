package main

import "time"

// CassandraDevopsSingleHost produces Cassandra-specific queries for the devops single-host case.
type CassandraDevopsSingleHost struct {
	CassandraDevops
}

func NewCassandraDevopsSingleHost(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newCassandraDevopsCommon(dbConfig, start, end).(*CassandraDevops)
	return &CassandraDevopsSingleHost{
		CassandraDevops: *underlying,
	}
}

func (d *CassandraDevopsSingleHost) Dispatch(_, scaleVar int) Query {
	q := NewCassandraQuery() // from pool
	d.MaxCPUUsageHourByMinuteOneHost(q, scaleVar)
	return q
}
