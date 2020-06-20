package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab/api/category"
)

var (
	categoryGroupInfoDesc = prometheus.NewDesc(
		"ynab_category_group_info",
		"Information about the YNAB category group",
		[]string{"budget_id", "category_group_id", "name", "hidden", "deleted"}, nil)
)

func newCategoryGroupInfoMetric(budgetID string, grp *category.Group) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		categoryGroupInfoDesc,
		prometheus.GaugeValue,
		1,
		budgetID, grp.ID, grp.Name, bool2str(grp.Hidden), bool2str(grp.Deleted))
}
