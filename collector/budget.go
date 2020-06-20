package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab/api/budget"
)

var (
	budgetInfoDesc = prometheus.NewDesc(
		"ynab_budget_info",
		"Information about the YNAB budget",
		[]string{"budget_id", "budget_name"}, nil)
)

func newBudgetInfoMetric(bud *budget.Budget) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		budgetInfoDesc, prometheus.GaugeValue,
		1, bud.ID, bud.Name)
}
