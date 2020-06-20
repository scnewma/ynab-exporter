package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab/api/category"
)

var (
	categoryInfoDesc = prometheus.NewDesc(
		"ynab_category_info",
		"Information about the YNAB category",
		[]string{"budget_id", "category_id", "name", "category_group_id", "hidden", "deleted"}, nil)

	categoryBudgetedDesc = prometheus.NewDesc(
		"ynab_category_budgeted",
		"The budgeted amount for this category in the current month",
		[]string{"category_id"}, nil)

	categoryActivityDesc = prometheus.NewDesc(
		"ynab_category_activity",
		"The activity amount for this category in the current month",
		[]string{"category_id"}, nil)

	categoryBalanceDesc = prometheus.NewDesc(
		"ynab_category_balance",
		"The balance for this category in the current month",
		[]string{"category_id"}, nil)
)

func newCategoryInfoMetric(budgetID string, cat *category.Category) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		categoryInfoDesc,
		prometheus.GaugeValue,
		1,
		budgetID, cat.ID, cat.Name, cat.CategoryGroupID, bool2str(cat.Hidden), bool2str(cat.Deleted))
}

func newCategoryBudgetedMetric(cat *category.Category) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		categoryBudgetedDesc, prometheus.GaugeValue,
		dollars(cat.Budgeted), cat.ID)
}

func newCategoryActivityMetric(cat *category.Category) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		categoryActivityDesc, prometheus.GaugeValue,
		dollars(cat.Activity), cat.ID)
}

func newCategoryBalanceMetric(cat *category.Category) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		categoryBalanceDesc, prometheus.GaugeValue,
		dollars(cat.Balance), cat.ID)
}
