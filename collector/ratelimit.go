package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab/api"
)

var (
	rateLimitUsedDesc = prometheus.NewDesc(
		"ynab_ratelimit_used",
		"The amount of the YNAB API rate limit that has been used.",
		nil, nil)

	rateLimitTotalDesc = prometheus.NewDesc(
		"ynab_ratelimit_total",
		"The total YNAB API rate limit.",
		nil, nil)
)

func newRateLimitUsedMetric(limit *api.RateLimit) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		rateLimitUsedDesc, prometheus.GaugeValue,
		float64(limit.Used()))
}

func newRateLimitTotalMetric(limit *api.RateLimit) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		rateLimitTotalDesc, prometheus.GaugeValue,
		float64(limit.Total()))
}
