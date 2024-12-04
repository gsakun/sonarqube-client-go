// Manage localization.
package sonar

import "net/http"

type L10NService struct {
	client *Client
}

type L10NIndexObject struct {
	Locale   string               `json:"locale,omitempty"`
	Messages L10NIndexObject_sub1 `json:"messages,omitempty"`
}

type L10NIndexObject_sub1 struct {
	ActionPlans_confirmDelete                                  string `json:"action_plans.confirm_delete,omitempty"`
	ActionPlans_delete                                         string `json:"action_plans.delete,omitempty"`
	AnalysisReports_pastReports                                string `json:"analysis_reports.past_reports,omitempty"`
	AnalysisReports_xReports                                   string `json:"analysis_reports.x_reports,omitempty"`
	BackgroundTasks_failures                                   string `json:"background_tasks.failures,omitempty"`
	BackgroundTasks_table_duration                             string `json:"background_tasks.table.duration,omitempty"`
	Cancel                                                     string `json:"cancel,omitempty"`
	CodingRules_changeSeverityIn                               string `json:"coding_rules.change_severity_in,omitempty"`
	CodingRules_create                                         string `json:"coding_rules.create,omitempty"`
	CodingRules_filters_activation_help                        string `json:"coding_rules.filters.activation.help,omitempty"`
	CodingRules_mostViolatedProjects                           string `json:"coding_rules.most_violated_projects,omitempty"`
	CodingRules_validation_invalidRuleKey                      string `json:"coding_rules.validation.invalid_rule_key,omitempty"`
	Component                                                  string `json:"component,omitempty"`
	ComponentNavigation_status_inProgress_admin                string `json:"component_navigation.status.in_progress.admin,omitempty"`
	CoverageViewer_overallTests                                string `json:"coverage_viewer.overall_tests,omitempty"`
	CustomMeasures_page_description                            string `json:"custom_measures.page.description,omitempty"`
	Dashboard_TimeMachine_name                                 string `json:"dashboard.TimeMachine.name,omitempty"`
	Dashboard_defaultDashboard                                 string `json:"dashboard.default_dashboard,omitempty"`
	Dashboard_errorUnshareDefault                              string `json:"dashboard.error_unshare_default,omitempty"`
	Dashboard_projectNotFound                                  string `json:"dashboard.project_not_found,omitempty"`
	Days                                                       string `json:"days,omitempty"`
	EmailConfiguration_test_send                               string `json:"email_configuration.test.send,omitempty"`
	GlobalPermissions_admin                                    string `json:"global_permissions.admin,omitempty"`
	HelpTips                                                   string `json:"help_tips,omitempty"`
	IncludingAbbreviated                                       string `json:"including_abbreviated,omitempty"`
	Issue_assign_formlink                                      string `json:"issue.assign.formlink,omitempty"`
	Issue_changelog_field_assignee                             string `json:"issue.changelog.field.assignee,omitempty"`
	Issue_effort                                               string `json:"issue.effort,omitempty"`
	Issue_resolution_REMOVED                                   string `json:"issue.resolution.REMOVED,omitempty"`
	IssueFilter_sharing                                        string `json:"issue_filter.sharing,omitempty"`
	MeasureFilter_abbr_description                             string `json:"measure_filter.abbr.description,omitempty"`
	MeasureFilter_criteria_key                                 string `json:"measure_filter.criteria.key,omitempty"`
	MeasureFilter_keyContains                                  string `json:"measure_filter.key_contains,omitempty"`
	MeasureFilter_nameContains                                 string `json:"measure_filter.name_contains,omitempty"`
	MeasureFilter_sharing                                      string `json:"measure_filter.sharing,omitempty"`
	Metric_coveredLines_description                            string `json:"metric.covered_lines.description,omitempty"`
	Metric_duplicationsData_description                        string `json:"metric.duplications_data.description,omitempty"`
	Metric_itBranchCoverage_description                        string `json:"metric.it_branch_coverage.description,omitempty"`
	Metric_lineCoverage_name                                   string `json:"metric.line_coverage.name,omitempty"`
	Metric_newCoverage_description                             string `json:"metric.new_coverage.description,omitempty"`
	Metric_overallCoverage_description                         string `json:"metric.overall_coverage.description,omitempty"`
	Metric_overallLinesToCover_description                     string `json:"metric.overall_lines_to_cover.description,omitempty"`
	Metric_publicAPI_name                                      string `json:"metric.public_api.name,omitempty"`
	Metric_sqaleRating_description                             string `json:"metric.sqale_rating.description,omitempty"`
	Metric_testFailures_description                            string `json:"metric.test_failures.description,omitempty"`
	Metric_total_useless_lines_name                            string `json:"metric.total-useless-lines.name,omitempty"`
	NameTooLongX                                               string `json:"name_too_long_x,omitempty"`
	NoResults                                                  string `json:"no_results,omitempty"`
	Notification_channel_EmailNotificationChannel              string `json:"notification.channel.EmailNotificationChannel,omitempty"`
	Optional                                                   string `json:"optional,omitempty"`
	OverXDays                                                  string `json:"over_x_days,omitempty"`
	ProjectHistory_col_month                                   string `json:"project_history.col.month,omitempty"`
	ProjectSettings_page                                       string `json:"project_settings.page,omitempty"`
	ProjectsRole_codeviewer_desc                               string `json:"projects_role.codeviewer.desc,omitempty"`
	Property_category_codeCoverage                             string `json:"property.category.codeCoverage,omitempty"`
	Property_category_exclusions_duplications_description      string `json:"property.category.exclusions.duplications.description,omitempty"`
	Property_category_security                                 string `json:"property.category.security,omitempty"`
	Property_sonar_global_exclusions_name                      string `json:"property.sonar.global.exclusions.name,omitempty"`
	Qualifiers_deleteConfirm_TRK                               string `json:"qualifiers.delete_confirm.TRK,omitempty"`
	QualityGates_deleteCondition_confirm_message               string `json:"quality_gates.delete_condition.confirm.message,omitempty"`
	QualityGates_operator_GT_short                             string `json:"quality_gates.operator.GT.short,omitempty"`
	QualityProfiles_removeProjectsConfirmButton                string `json:"quality_profiles.remove_projects_confirm_button,omitempty"`
	QualityProfiles_renameXTitle                               string `json:"quality_profiles.rename_x_title,omitempty"`
	QualityProfiles_restoreBuiltInProfiles                     string `json:"quality_profiles.restore_built_in_profiles,omitempty"`
	Result                                                     string `json:"result,omitempty"`
	Roles_page                                                 string `json:"roles.page,omitempty"`
	Rule_php_S101_param_format                                 string `json:"rule.php.S101.param.format,omitempty"`
	Rule_php_S1067_param_max                                   string `json:"rule.php.S1067.param.max,omitempty"`
	Rule_php_S115_param_format                                 string `json:"rule.php.S115.param.format,omitempty"`
	Rule_php_S1808_param_oneSpaceAfter                         string `json:"rule.php.S1808.param.one_space_after,omitempty"`
	Rule_php_S1808_param_oneSpaceBefore                        string `json:"rule.php.S1808.param.one_space_before,omitempty"`
	Rules_notFound                                             string `json:"rules.not_found,omitempty"`
	Select2_tooShort                                           string `json:"select2.tooShort,omitempty"`
	SelectAMetric                                              string `json:"select_a_metric,omitempty"`
	Shortcuts_section_global_shortcuts                         string `json:"shortcuts.section.global.shortcuts,omitempty"`
	Size                                                       string `json:"size,omitempty"`
	SourceViewer_tooltip_ut_partially_covered                  string `json:"source_viewer.tooltip.ut.partially-covered,omitempty"`
	TestViewer_skipped                                         string `json:"test_viewer.skipped,omitempty"`
	To_downcase                                                string `json:"to.downcase,omitempty"`
	UpdateKey_replace                                          string `json:"update_key.replace,omitempty"`
	User_addScmAccount                                         string `json:"user.add_scm_account,omitempty"`
	Views_deleteSuccess                                        string `json:"views.delete_success,omitempty"`
	Views_editSubview                                          string `json:"views.edit_subview,omitempty"`
	Views_invalid_criteria                                     string `json:"views.invalid.criteria,omitempty"`
	Views_projects_byRegexp                                    string `json:"views.projects.by_regexp,omitempty"`
	Views_projects_selectionMode                               string `json:"views.projects.selection_mode,omitempty"`
	Widget_complexity_description                              string `json:"widget.complexity.description,omitempty"`
	Widget_customMeasures_property_metric10_name               string `json:"widget.custom_measures.property.metric10.name,omitempty"`
	Widget_issueFilter_property_filter_name                    string `json:"widget.issue_filter.property.filter.name,omitempty"`
	Widget_measureFilterBubbleChart_name                       string `json:"widget.measure_filter_bubble_chart.name,omitempty"`
	Widget_measureFilterCloud_property_colorMetric_name        string `json:"widget.measure_filter_cloud.property.colorMetric.name,omitempty"`
	Widget_measureFilterHistogram_description                  string `json:"widget.measure_filter_histogram.description,omitempty"`
	Widget_measureFilterHistogram_name                         string `json:"widget.measure_filter_histogram.name,omitempty"`
	Widget_measureFilterPieChart_property_chartHeight_name     string `json:"widget.measure_filter_pie_chart.property.chartHeight.name,omitempty"`
	Widget_measureFilterPieChart_property_extraMetric2_name    string `json:"widget.measure_filter_pie_chart.property.extraMetric2.name,omitempty"`
	Widget_measureFilterTreemap_property_heightInPercents_name string `json:"widget.measure_filter_treemap.property.heightInPercents.name,omitempty"`
	Widget_projectIssueFilter_property_filter_name             string `json:"widget.project_issue_filter.property.filter.name,omitempty"`
	Widget_selectProject                                       string `json:"widget.select_project,omitempty"`
	Widget_sqaleSunburst_cantDisplay                           string `json:"widget.sqaleSunburst.cant_display,omitempty"`
	Widget_timeMachine_property_metric2_name                   string `json:"widget.time_machine.property.metric2.name,omitempty"`
	Widget_timeMachine_property_metric9_name                   string `json:"widget.time_machine.property.metric9.name,omitempty"`
	Widget_treemap_widget_name                                 string `json:"widget.treemap-widget.name,omitempty"`
}

type L10NIndexOption struct {
	Locale string `url:"locale,omitempty"` // Description:"BCP47 language tag, used to override the browser Accept-Language header",ExampleValue:"fr-CH"
	Ts     string `url:"ts,omitempty"`     // Description:"Date of the last cache update.",ExampleValue:"2014-06-04T09:31:42+0000"
}

// Index Get all localization messages for a given locale
func (s *L10NService) Index(opt *L10NIndexOption) (v *L10NIndexObject, resp *http.Response, err error) {
	err = s.ValidateIndexOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "l10n/index", opt)
	if err != nil {
		return
	}
	v = new(L10NIndexObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
