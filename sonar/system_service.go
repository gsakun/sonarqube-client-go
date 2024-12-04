// Get system details, and perform some management actions, such as restarting, and initiating a database migration (as part of a system upgrade).
package sonar

import "net/http"

type SystemService struct {
	client *Client
}

type SystemDbMigrationStatusObject struct {
	Message   string `json:"message,omitempty"`
	StartedAt string `json:"startedAt,omitempty"`
	State     string `json:"state,omitempty"`
}

type SystemHealthObject struct {
	Causes []SystemHealthObject_sub1 `json:"causes,omitempty"`
	Health string                    `json:"health,omitempty"`
	Nodes  []SystemHealthObject_sub2 `json:"nodes,omitempty"`
}

type SystemHealthObject_sub2 struct {
	Causes    []SystemHealthObject_sub1 `json:"causes,omitempty"`
	Health    string                    `json:"health,omitempty"`
	Host      string                    `json:"host,omitempty"`
	Name      string                    `json:"name,omitempty"`
	Port      int64                     `json:"port,omitempty"`
	StartedAt string                    `json:"startedAt,omitempty"`
	Type      string                    `json:"type,omitempty"`
}

type SystemHealthObject_sub1 struct {
	Message string `json:"message,omitempty"`
}

type SystemInfoObject struct {
	ALMs                               SystemInfoObject_sub1  `json:"ALMs,omitempty"`
	Bundled                            SystemInfoObject_sub2  `json:"Bundled,omitempty"`
	Compute_Engine_Database_Connection SystemInfoObject_sub3  `json:"Compute Engine Database Connection,omitempty"`
	Compute_Engine_JVM_Properties      SystemInfoObject_sub4  `json:"Compute Engine JVM Properties,omitempty"`
	Compute_Engine_JVM_State           SystemInfoObject_sub5  `json:"Compute Engine JVM State,omitempty"`
	Compute_Engine_Logging             SystemInfoObject_sub6  `json:"Compute Engine Logging,omitempty"`
	Compute_Engine_Tasks               SystemInfoObject_sub7  `json:"Compute Engine Tasks,omitempty"`
	Database                           SystemInfoObject_sub8  `json:"Database,omitempty"`
	Health                             string                 `json:"Health,omitempty"`
	Health_Causes                      []interface{}          `json:"Health Causes,omitempty"`
	Plugins                            SystemInfoObject_sub9  `json:"Plugins,omitempty"`
	Search_Indexes                     SystemInfoObject_sub10 `json:"Search Indexes,omitempty"`
	Search_State                       SystemInfoObject_sub11 `json:"Search State,omitempty"`
	Server_Push_Connections            SystemInfoObject_sub12 `json:"Server Push Connections,omitempty"`
	Settings                           SystemInfoObject_sub13 `json:"Settings,omitempty"`
	System                             SystemInfoObject_sub14 `json:"System,omitempty"`
	Web_Database_Connection            SystemInfoObject_sub3  `json:"Web Database Connection,omitempty"`
	Web_JVM_Properties                 SystemInfoObject_sub15 `json:"Web JVM Properties,omitempty"`
	Web_JVM_State                      SystemInfoObject_sub5  `json:"Web JVM State,omitempty"`
	Web_Logging                        SystemInfoObject_sub6  `json:"Web Logging,omitempty"`
}

type SystemInfoObject_sub14 struct {
	Accepted_external_identity_providers                                      string `json:"Accepted external identity providers,omitempty"`
	Data_Dir                                                                  string `json:"Data Dir,omitempty"`
	Docker                                                                    bool   `json:"Docker,omitempty"`
	Edition                                                                   string `json:"Edition,omitempty"`
	External_identity_providers_whose_users_are_allowed_to_sign_themselves_up string `json:"External identity providers whose users are allowed to sign themselves up,omitempty"`
	Force_authentication                                                      bool   `json:"Force authentication,omitempty"`
	High_Availability                                                         bool   `json:"High Availability,omitempty"`
	Home_Dir                                                                  string `json:"Home Dir,omitempty"`
	Official_Distribution                                                     bool   `json:"Official Distribution,omitempty"`
	Processors                                                                int64  `json:"Processors,omitempty"`
	Server_ID                                                                 string `json:"Server ID,omitempty"`
	Temp_Dir                                                                  string `json:"Temp Dir,omitempty"`
	Version                                                                   string `json:"Version,omitempty"`
}

type SystemInfoObject_sub15 struct {
	Awt_toolkit                                           string `json:"awt.toolkit,omitempty"`
	Catalina_base                                         string `json:"catalina.base,omitempty"`
	Catalina_home                                         string `json:"catalina.home,omitempty"`
	Catalina_useNaming                                    string `json:"catalina.useNaming,omitempty"`
	Com_redhat_fips                                       string `json:"com.redhat.fips,omitempty"`
	Com_zaxxer_hikari_poolNumber                          string `json:"com.zaxxer.hikari.pool_number,omitempty"`
	File_encoding                                         string `json:"file.encoding,omitempty"`
	File_separator                                        string `json:"file.separator,omitempty"`
	Ftp_nonProxyHosts                                     string `json:"ftp.nonProxyHosts,omitempty"`
	GopherProxySet                                        string `json:"gopherProxySet,omitempty"`
	HTTP_agent                                            string `json:"http.agent,omitempty"`
	HTTP_nonProxyHosts                                    string `json:"http.nonProxyHosts,omitempty"`
	Java_awt_graphicsenv                                  string `json:"java.awt.graphicsenv,omitempty"`
	Java_awt_headless                                     string `json:"java.awt.headless,omitempty"`
	Java_awt_printerjob                                   string `json:"java.awt.printerjob,omitempty"`
	Java_class_path                                       string `json:"java.class.path,omitempty"`
	Java_class_version                                    string `json:"java.class.version,omitempty"`
	Java_home                                             string `json:"java.home,omitempty"`
	Java_io_tmpdir                                        string `json:"java.io.tmpdir,omitempty"`
	Java_library_path                                     string `json:"java.library.path,omitempty"`
	Java_runtime_name                                     string `json:"java.runtime.name,omitempty"`
	Java_runtime_version                                  string `json:"java.runtime.version,omitempty"`
	Java_specification_name                               string `json:"java.specification.name,omitempty"`
	Java_specification_vendor                             string `json:"java.specification.vendor,omitempty"`
	Java_specification_version                            string `json:"java.specification.version,omitempty"`
	Java_vendor                                           string `json:"java.vendor,omitempty"`
	Java_vendor_url                                       string `json:"java.vendor.url,omitempty"`
	Java_vendor_url_bug                                   string `json:"java.vendor.url.bug,omitempty"`
	Java_vendor_version                                   string `json:"java.vendor.version,omitempty"`
	Java_version                                          string `json:"java.version,omitempty"`
	Java_version_date                                     string `json:"java.version.date,omitempty"`
	Java_vm_compressedOopsMode                            string `json:"java.vm.compressedOopsMode,omitempty"`
	Java_vm_info                                          string `json:"java.vm.info,omitempty"`
	Java_vm_name                                          string `json:"java.vm.name,omitempty"`
	Java_vm_specification_name                            string `json:"java.vm.specification.name,omitempty"`
	Java_vm_specification_vendor                          string `json:"java.vm.specification.vendor,omitempty"`
	Java_vm_specification_version                         string `json:"java.vm.specification.version,omitempty"`
	Java_vm_vendor                                        string `json:"java.vm.vendor,omitempty"`
	Java_vm_version                                       string `json:"java.vm.version,omitempty"`
	Jdk_debug                                             string `json:"jdk.debug,omitempty"`
	Jdk_vendor_version                                    string `json:"jdk.vendor.version,omitempty"`
	Line_separator                                        string `json:"line.separator,omitempty"`
	LogbackDisableServletContainerInitializer             string `json:"logbackDisableServletContainerInitializer,omitempty"`
	Org_apache_catalina_startup_EXITONINITFAILURE         string `json:"org.apache.catalina.startup.EXIT_ON_INIT_FAILURE,omitempty"`
	Org_apache_tomcat_util_buf_UDecoder_ALLOWENCODEDSLASH string `json:"org.apache.tomcat.util.buf.UDecoder.ALLOW_ENCODED_SLASH,omitempty"`
	Os_arch                                               string `json:"os.arch,omitempty"`
	Os_name                                               string `json:"os.name,omitempty"`
	Os_version                                            string `json:"os.version,omitempty"`
	Path_separator                                        string `json:"path.separator,omitempty"`
	SocksNonProxyHosts                                    string `json:"socksNonProxyHosts,omitempty"`
	Sun_arch_data_model                                   string `json:"sun.arch.data.model,omitempty"`
	Sun_boot_library_path                                 string `json:"sun.boot.library.path,omitempty"`
	Sun_cpu_endian                                        string `json:"sun.cpu.endian,omitempty"`
	Sun_cpu_isalist                                       string `json:"sun.cpu.isalist,omitempty"`
	Sun_font_fontmanager                                  string `json:"sun.font.fontmanager,omitempty"`
	Sun_io_unicode_encoding                               string `json:"sun.io.unicode.encoding,omitempty"`
	Sun_java_command                                      string `json:"sun.java.command,omitempty"`
	Sun_java_launcher                                     string `json:"sun.java.launcher,omitempty"`
	Sun_jnu_encoding                                      string `json:"sun.jnu.encoding,omitempty"`
	Sun_management_compiler                               string `json:"sun.management.compiler,omitempty"`
	Sun_os_patch_level                                    string `json:"sun.os.patch.level,omitempty"`
	User_country                                          string `json:"user.country,omitempty"`
	User_dir                                              string `json:"user.dir,omitempty"`
	User_home                                             string `json:"user.home,omitempty"`
	User_language                                         string `json:"user.language,omitempty"`
	User_name                                             string `json:"user.name,omitempty"`
	User_timezone                                         string `json:"user.timezone,omitempty"`
}

type SystemInfoObject_sub4 struct {
	Awt_toolkit                   string `json:"awt.toolkit,omitempty"`
	Com_redhat_fips               string `json:"com.redhat.fips,omitempty"`
	Com_zaxxer_hikari_poolNumber  string `json:"com.zaxxer.hikari.pool_number,omitempty"`
	File_encoding                 string `json:"file.encoding,omitempty"`
	File_separator                string `json:"file.separator,omitempty"`
	Ftp_nonProxyHosts             string `json:"ftp.nonProxyHosts,omitempty"`
	GopherProxySet                string `json:"gopherProxySet,omitempty"`
	HTTP_nonProxyHosts            string `json:"http.nonProxyHosts,omitempty"`
	Java_awt_graphicsenv          string `json:"java.awt.graphicsenv,omitempty"`
	Java_awt_headless             string `json:"java.awt.headless,omitempty"`
	Java_awt_printerjob           string `json:"java.awt.printerjob,omitempty"`
	Java_class_path               string `json:"java.class.path,omitempty"`
	Java_class_version            string `json:"java.class.version,omitempty"`
	Java_home                     string `json:"java.home,omitempty"`
	Java_io_tmpdir                string `json:"java.io.tmpdir,omitempty"`
	Java_library_path             string `json:"java.library.path,omitempty"`
	Java_runtime_name             string `json:"java.runtime.name,omitempty"`
	Java_runtime_version          string `json:"java.runtime.version,omitempty"`
	Java_specification_name       string `json:"java.specification.name,omitempty"`
	Java_specification_vendor     string `json:"java.specification.vendor,omitempty"`
	Java_specification_version    string `json:"java.specification.version,omitempty"`
	Java_vendor                   string `json:"java.vendor,omitempty"`
	Java_vendor_url               string `json:"java.vendor.url,omitempty"`
	Java_vendor_url_bug           string `json:"java.vendor.url.bug,omitempty"`
	Java_vendor_version           string `json:"java.vendor.version,omitempty"`
	Java_version                  string `json:"java.version,omitempty"`
	Java_version_date             string `json:"java.version.date,omitempty"`
	Java_vm_compressedOopsMode    string `json:"java.vm.compressedOopsMode,omitempty"`
	Java_vm_info                  string `json:"java.vm.info,omitempty"`
	Java_vm_name                  string `json:"java.vm.name,omitempty"`
	Java_vm_specification_name    string `json:"java.vm.specification.name,omitempty"`
	Java_vm_specification_vendor  string `json:"java.vm.specification.vendor,omitempty"`
	Java_vm_specification_version string `json:"java.vm.specification.version,omitempty"`
	Java_vm_vendor                string `json:"java.vm.vendor,omitempty"`
	Java_vm_version               string `json:"java.vm.version,omitempty"`
	Jdk_debug                     string `json:"jdk.debug,omitempty"`
	Jdk_vendor_version            string `json:"jdk.vendor.version,omitempty"`
	Line_separator                string `json:"line.separator,omitempty"`
	Os_arch                       string `json:"os.arch,omitempty"`
	Os_name                       string `json:"os.name,omitempty"`
	Os_version                    string `json:"os.version,omitempty"`
	Path_separator                string `json:"path.separator,omitempty"`
	SocksNonProxyHosts            string `json:"socksNonProxyHosts,omitempty"`
	Sun_arch_data_model           string `json:"sun.arch.data.model,omitempty"`
	Sun_boot_library_path         string `json:"sun.boot.library.path,omitempty"`
	Sun_cpu_endian                string `json:"sun.cpu.endian,omitempty"`
	Sun_cpu_isalist               string `json:"sun.cpu.isalist,omitempty"`
	Sun_io_unicode_encoding       string `json:"sun.io.unicode.encoding,omitempty"`
	Sun_java_command              string `json:"sun.java.command,omitempty"`
	Sun_java_launcher             string `json:"sun.java.launcher,omitempty"`
	Sun_jnu_encoding              string `json:"sun.jnu.encoding,omitempty"`
	Sun_management_compiler       string `json:"sun.management.compiler,omitempty"`
	Sun_os_patch_level            string `json:"sun.os.patch.level,omitempty"`
	User_country                  string `json:"user.country,omitempty"`
	User_dir                      string `json:"user.dir,omitempty"`
	User_home                     string `json:"user.home,omitempty"`
	User_language                 string `json:"user.language,omitempty"`
	User_name                     string `json:"user.name,omitempty"`
	User_timezone                 string `json:"user.timezone,omitempty"`
}

type SystemInfoObject_sub11 struct {
	CPU_Usage                             int64  `json:"CPU Usage (%),omitempty"`
	Disk_Available                        string `json:"Disk Available,omitempty"`
	Field_Data_Circuit_Breaker_Estimation string `json:"Field Data Circuit Breaker Estimation,omitempty"`
	Field_Data_Circuit_Breaker_Limit      string `json:"Field Data Circuit Breaker Limit,omitempty"`
	Field_Data_Memory                     string `json:"Field Data Memory,omitempty"`
	JVM_Heap_Max                          string `json:"JVM Heap Max,omitempty"`
	JVM_Heap_Usage                        string `json:"JVM Heap Usage,omitempty"`
	JVM_Heap_Used                         string `json:"JVM Heap Used,omitempty"`
	JVM_Non_Heap_Used                     string `json:"JVM Non Heap Used,omitempty"`
	JVM_Threads                           int64  `json:"JVM Threads,omitempty"`
	Max_File_Descriptors                  int64  `json:"Max File Descriptors,omitempty"`
	Open_File_Descriptors                 int64  `json:"Open File Descriptors,omitempty"`
	Query_Cache_Memory                    string `json:"Query Cache Memory,omitempty"`
	Request_Cache_Memory                  string `json:"Request Cache Memory,omitempty"`
	Request_Circuit_Breaker_Estimation    string `json:"Request Circuit Breaker Estimation,omitempty"`
	Request_Circuit_Breaker_Limit         string `json:"Request Circuit Breaker Limit,omitempty"`
	State                                 string `json:"State,omitempty"`
	Store_Size                            string `json:"Store Size,omitempty"`
	Translog_Size                         string `json:"Translog Size,omitempty"`
}

type SystemInfoObject_sub2 struct {
	Config     string `json:"config,omitempty"`
	Csharp     string `json:"csharp,omitempty"`
	Flex       string `json:"flex,omitempty"`
	Go         string `json:"go,omitempty"`
	Iac        string `json:"iac,omitempty"`
	Jacoco     string `json:"jacoco,omitempty"`
	Java       string `json:"java,omitempty"`
	Javascript string `json:"javascript,omitempty"`
	Kotlin     string `json:"kotlin,omitempty"`
	Php        string `json:"php,omitempty"`
	Python     string `json:"python,omitempty"`
	Ruby       string `json:"ruby,omitempty"`
	Sonarscala string `json:"sonarscala,omitempty"`
	Text       string `json:"text,omitempty"`
	Vbnet      string `json:"vbnet,omitempty"`
	Web        string `json:"web,omitempty"`
	XML        string `json:"xml,omitempty"`
}

type SystemInfoObject_sub8 struct {
	Database                      string `json:"Database,omitempty"`
	Database_Version              string `json:"Database Version,omitempty"`
	Default_transaction_isolation string `json:"Default transaction isolation,omitempty"`
	Driver                        string `json:"Driver,omitempty"`
	Driver_Version                string `json:"Driver Version,omitempty"`
	URL                           string `json:"URL,omitempty"`
	Username                      string `json:"Username,omitempty"`
}

type SystemInfoObject_sub13 struct {
	Default_New_Code_Definition                              string `json:"Default New Code Definition,omitempty"`
	Devactivity_status                                       string `json:"devactivity.status,omitempty"`
	HTTP_nonProxyHosts                                       string `json:"http.nonProxyHosts,omitempty"`
	Process_gracefulStopTimeout                              string `json:"process.gracefulStopTimeout,omitempty"`
	Process_index                                            string `json:"process.index,omitempty"`
	Process_key                                              string `json:"process.key,omitempty"`
	Process_sharedDir                                        string `json:"process.sharedDir,omitempty"`
	Projects_default_visibility                              string `json:"projects.default.visibility,omitempty"`
	Qualitygate_default                                      string `json:"qualitygate.default,omitempty"`
	Sonar_auth_saml_certificate_secured                      string `json:"sonar.auth.saml.certificate.secured,omitempty"`
	Sonar_auth_saml_enabled                                  string `json:"sonar.auth.saml.enabled,omitempty"`
	Sonar_auth_saml_loginURL                                 string `json:"sonar.auth.saml.loginUrl,omitempty"`
	Sonar_auth_saml_providerID                               string `json:"sonar.auth.saml.providerId,omitempty"`
	Sonar_auth_saml_providerName                             string `json:"sonar.auth.saml.providerName,omitempty"`
	Sonar_auth_saml_signature_enabled                        string `json:"sonar.auth.saml.signature.enabled,omitempty"`
	Sonar_auth_saml_sp_certificate_secured                   string `json:"sonar.auth.saml.sp.certificate.secured,omitempty"`
	Sonar_auth_saml_sp_privateKey_secured                    string `json:"sonar.auth.saml.sp.privateKey.secured,omitempty"`
	Sonar_auth_saml_user_login                               string `json:"sonar.auth.saml.user.login,omitempty"`
	Sonar_auth_saml_user_name                                string `json:"sonar.auth.saml.user.name,omitempty"`
	Sonar_authenticator_ignoreStartupFailure                 string `json:"sonar.authenticator.ignoreStartupFailure,omitempty"`
	Sonar_autoDatabaseUpgrade                                string `json:"sonar.autoDatabaseUpgrade,omitempty"`
	Sonar_blueGreenEnabled                                   string `json:"sonar.blueGreenEnabled,omitempty"`
	Sonar_buildbreaker_skip                                  string `json:"sonar.buildbreaker.skip,omitempty"`
	Sonar_c_predefinedMacros                                 string `json:"sonar.c.predefinedMacros,omitempty"`
	Sonar_ce_gracefulStopTimeOutInMs                         string `json:"sonar.ce.gracefulStopTimeOutInMs,omitempty"`
	Sonar_ce_javaAdditionalOpts                              string `json:"sonar.ce.javaAdditionalOpts,omitempty"`
	Sonar_ce_javaOpts                                        string `json:"sonar.ce.javaOpts,omitempty"`
	Sonar_cluster_enabled                                    string `json:"sonar.cluster.enabled,omitempty"`
	Sonar_cluster_kubernetes                                 string `json:"sonar.cluster.kubernetes,omitempty"`
	Sonar_cluster_name                                       string `json:"sonar.cluster.name,omitempty"`
	Sonar_cluster_node_name                                  string `json:"sonar.cluster.node.name,omitempty"`
	Sonar_cluster_node_port                                  string `json:"sonar.cluster.node.port,omitempty"`
	Sonar_cluster_web_startupLeader                          string `json:"sonar.cluster.web.startupLeader,omitempty"`
	Sonar_core_id                                            string `json:"sonar.core.id,omitempty"`
	Sonar_core_serverBaseURL                                 string `json:"sonar.core.serverBaseURL,omitempty"`
	Sonar_core_startTime                                     string `json:"sonar.core.startTime,omitempty"`
	Sonar_core_treemap_colormetric                           string `json:"sonar.core.treemap.colormetric,omitempty"`
	Sonar_core_treemap_sizemetric                            string `json:"sonar.core.treemap.sizemetric,omitempty"`
	Sonar_cpd_crossProject                                   string `json:"sonar.cpd.cross_project,omitempty"`
	Sonar_dbcleaner_branchesToKeepWhenInactive               string `json:"sonar.dbcleaner.branchesToKeepWhenInactive,omitempty"`
	Sonar_dbcleaner_daysBeforeDeletingInactiveBranchesAndPRs string `json:"sonar.dbcleaner.daysBeforeDeletingInactiveBranchesAndPRs,omitempty"`
	Sonar_dbcleaner_monthsBeforeDeletingAllSnapshots         string `json:"sonar.dbcleaner.monthsBeforeDeletingAllSnapshots,omitempty"`
	Sonar_dbcleaner_weeksBeforeDeletingAllSnapshots          string `json:"sonar.dbcleaner.weeksBeforeDeletingAllSnapshots,omitempty"`
	Sonar_dryRun_cache_lastUpdate                            string `json:"sonar.dryRun.cache.lastUpdate,omitempty"`
	Sonar_es_port                                            string `json:"sonar.es.port,omitempty"`
	Sonar_forceAuthentication                                string `json:"sonar.forceAuthentication,omitempty"`
	Sonar_genericcoverage_suffixes                           string `json:"sonar.genericcoverage.suffixes,omitempty"`
	Sonar_governance_report_view_frequency                   string `json:"sonar.governance.report.view.frequency,omitempty"`
	Sonar_java_coveragePlugin                                string `json:"sonar.java.coveragePlugin,omitempty"`
	Sonar_javascript_jQueryObjectAliases                     string `json:"sonar.javascript.jQueryObjectAliases,omitempty"`
	Sonar_jdbc_driverPath                                    string `json:"sonar.jdbc.driverPath,omitempty"`
	Sonar_jdbc_maxActive                                     string `json:"sonar.jdbc.maxActive,omitempty"`
	Sonar_jdbc_maxWait                                       string `json:"sonar.jdbc.maxWait,omitempty"`
	Sonar_jdbc_minIdle                                       string `json:"sonar.jdbc.minIdle,omitempty"`
	Sonar_jdbc_password                                      string `json:"sonar.jdbc.password,omitempty"`
	Sonar_jdbc_url                                           string `json:"sonar.jdbc.url,omitempty"`
	Sonar_jdbc_username                                      string `json:"sonar.jdbc.username,omitempty"`
	Sonar_lf_enableGravatar                                  string `json:"sonar.lf.enableGravatar,omitempty"`
	Sonar_lf_logoWidthPx                                     string `json:"sonar.lf.logoWidthPx,omitempty"`
	Sonar_log_jsonOutput                                     string `json:"sonar.log.jsonOutput,omitempty"`
	Sonar_organisation                                       string `json:"sonar.organisation,omitempty"`
	Sonar_path_data                                          string `json:"sonar.path.data,omitempty"`
	Sonar_path_home                                          string `json:"sonar.path.home,omitempty"`
	Sonar_path_logs                                          string `json:"sonar.path.logs,omitempty"`
	Sonar_path_temp                                          string `json:"sonar.path.temp,omitempty"`
	Sonar_path_web                                           string `json:"sonar.path.web,omitempty"`
	Sonar_plsql_file_suffixes                                string `json:"sonar.plsql.file.suffixes,omitempty"`
	Sonar_plugins_risk_consent                               string `json:"sonar.plugins.risk.consent,omitempty"`
	Sonar_preview_excludePlugins                             string `json:"sonar.preview.excludePlugins,omitempty"`
	Sonar_report_dashboard_name                              string `json:"sonar.report.dashboard.name,omitempty"`
	Sonar_report_frequency                                   string `json:"sonar.report.frequency,omitempty"`
	Sonar_report_ignoreSslErrors                             string `json:"sonar.report.ignoreSslErrors,omitempty"`
	Sonar_report_lastDate                                    string `json:"sonar.report.last_date,omitempty"`
	Sonar_report_lastDate_devReport                          string `json:"sonar.report.last_date.dev_report,omitempty"`
	Sonar_report_lastDate_managementReport                   string `json:"sonar.report.last_date.management_report,omitempty"`
	Sonar_report_login                                       string `json:"sonar.report.login,omitempty"`
	Sonar_report_subject                                     string `json:"sonar.report.subject,omitempty"`
	Sonar_reports                                            string `json:"sonar.reports,omitempty"`
	Sonar_scm_disabled                                       string `json:"sonar.scm.disabled,omitempty"`
	Sonar_scm_enabled                                        string `json:"sonar.scm.enabled,omitempty"`
	Sonar_search_host                                        string `json:"sonar.search.host,omitempty"`
	Sonar_search_javaAdditionalOpts                          string `json:"sonar.search.javaAdditionalOpts,omitempty"`
	Sonar_search_javaOpts                                    string `json:"sonar.search.javaOpts,omitempty"`
	Sonar_search_port                                        string `json:"sonar.search.port,omitempty"`
	Sonar_technicalDebt_ratingGrid                           string `json:"sonar.technicalDebt.ratingGrid,omitempty"`
	Sonar_telemetry_compression                              string `json:"sonar.telemetry.compression,omitempty"`
	Sonar_telemetry_enable                                   string `json:"sonar.telemetry.enable,omitempty"`
	Sonar_telemetry_frequencyInSeconds                       string `json:"sonar.telemetry.frequencyInSeconds,omitempty"`
	Sonar_telemetry_url                                      string `json:"sonar.telemetry.url,omitempty"`
	Sonar_updatecenter_activate                              string `json:"sonar.updatecenter.activate,omitempty"`
	Sonar_web_gracefulStopTimeOutInMs                        string `json:"sonar.web.gracefulStopTimeOutInMs,omitempty"`
	Sonar_web_javaAdditionalOpts                             string `json:"sonar.web.javaAdditionalOpts,omitempty"`
	Sonar_web_javaOpts                                       string `json:"sonar.web.javaOpts,omitempty"`
	Sonar_web_sso_emailHeader                                string `json:"sonar.web.sso.emailHeader,omitempty"`
	Sonar_web_sso_enable                                     string `json:"sonar.web.sso.enable,omitempty"`
	Sonar_web_sso_groupsHeader                               string `json:"sonar.web.sso.groupsHeader,omitempty"`
	Sonar_web_sso_loginHeader                                string `json:"sonar.web.sso.loginHeader,omitempty"`
	Sonar_web_sso_nameHeader                                 string `json:"sonar.web.sso.nameHeader,omitempty"`
	Sonar_web_sso_refreshIntervalInMinutes                   string `json:"sonar.web.sso.refreshIntervalInMinutes,omitempty"`
}

type SystemInfoObject_sub5 struct {
	Free_Memory__MB        int64  `json:"Free Memory (MB),omitempty"`
	Heap_Committed__MB     int64  `json:"Heap Committed (MB),omitempty"`
	Heap_Init__MB          int64  `json:"Heap Init (MB),omitempty"`
	Heap_Max__MB           int64  `json:"Heap Max (MB),omitempty"`
	Heap_Used__MB          int64  `json:"Heap Used (MB),omitempty"`
	Max_Memory__MB         int64  `json:"Max Memory (MB),omitempty"`
	Non_Heap_Committed__MB int64  `json:"Non Heap Committed (MB),omitempty"`
	Non_Heap_Init__MB      int64  `json:"Non Heap Init (MB),omitempty"`
	Non_Heap_Used__MB      int64  `json:"Non Heap Used (MB),omitempty"`
	System_Load_Average    string `json:"System Load Average,omitempty"`
	Threads                int64  `json:"Threads,omitempty"`
}

type SystemInfoObject_sub1 struct {
	Github_Config string `json:"Github Config,omitempty"`
	Gitlab_Config string `json:"Gitlab Config,omitempty"`
}

type SystemInfoObject_sub7 struct {
	In_Progress              int64 `json:"In Progress,omitempty"`
	Longest_Time_Pending__ms int64 `json:"Longest Time Pending (ms),omitempty"`
	Max_Worker_Count         int64 `json:"Max Worker Count,omitempty"`
	Pending                  int64 `json:"Pending,omitempty"`
	Processed_With_Error     int64 `json:"Processed With Error,omitempty"`
	Processed_With_Success   int64 `json:"Processed With Success,omitempty"`
	Processing_Time__ms      int64 `json:"Processing Time (ms),omitempty"`
	Worker_Count             int64 `json:"Worker Count,omitempty"`
	Workers_Paused           bool  `json:"Workers Paused,omitempty"`
}

type SystemInfoObject_sub10 struct {
	Index_components___Docs            int64  `json:"Index components - Docs,omitempty"`
	Index_components___Shards          int64  `json:"Index components - Shards,omitempty"`
	Index_components___Store_Size      string `json:"Index components - Store Size,omitempty"`
	Index_issues___Docs                int64  `json:"Index issues - Docs,omitempty"`
	Index_issues___Shards              int64  `json:"Index issues - Shards,omitempty"`
	Index_issues___Store_Size          string `json:"Index issues - Store Size,omitempty"`
	Index_metadatas___Docs             int64  `json:"Index metadatas - Docs,omitempty"`
	Index_metadatas___Shards           int64  `json:"Index metadatas - Shards,omitempty"`
	Index_metadatas___Store_Size       string `json:"Index metadatas - Store Size,omitempty"`
	Index_projectmeasures___Docs       int64  `json:"Index projectmeasures - Docs,omitempty"`
	Index_projectmeasures___Shards     int64  `json:"Index projectmeasures - Shards,omitempty"`
	Index_projectmeasures___Store_Size string `json:"Index projectmeasures - Store Size,omitempty"`
	Index_rules___Docs                 int64  `json:"Index rules - Docs,omitempty"`
	Index_rules___Shards               int64  `json:"Index rules - Shards,omitempty"`
	Index_rules___Store_Size           string `json:"Index rules - Store Size,omitempty"`
	Index_users___Docs                 int64  `json:"Index users - Docs,omitempty"`
	Index_users___Shards               int64  `json:"Index users - Shards,omitempty"`
	Index_users___Store_Size           string `json:"Index users - Store Size,omitempty"`
	Index_views___Docs                 int64  `json:"Index views - Docs,omitempty"`
	Index_views___Shards               int64  `json:"Index views - Shards,omitempty"`
	Index_views___Store_Size           string `json:"Index views - Store Size,omitempty"`
}

type SystemInfoObject_sub6 struct {
	Logs_Dir   string `json:"Logs Dir,omitempty"`
	Logs_Level string `json:"Logs Level,omitempty"`
}

type SystemInfoObject_sub3 struct {
	Pool_Active_Connections   int64 `json:"Pool Active Connections,omitempty"`
	Pool_Idle_Connections     int64 `json:"Pool Idle Connections,omitempty"`
	Pool_Max_Connections      int64 `json:"Pool Max Connections,omitempty"`
	Pool_Max_Lifetime__ms     int64 `json:"Pool Max Lifetime (ms),omitempty"`
	Pool_Max_Wait__ms         int64 `json:"Pool Max Wait (ms),omitempty"`
	Pool_Min_Idle_Connections int64 `json:"Pool Min Idle Connections,omitempty"`
	Pool_Total_Connections    int64 `json:"Pool Total Connections,omitempty"`
}

type SystemInfoObject_sub12 struct {
	SonarLint_Connected_Clients int64 `json:"SonarLint Connected Clients,omitempty"`
}

type SystemInfoObject_sub9 struct{}

type SystemMigrateDbObject struct {
	Message   string `json:"message,omitempty"`
	StartedAt string `json:"startedAt,omitempty"`
	State     string `json:"state,omitempty"`
}

type SystemStatusObject struct {
	ID      string `json:"id,omitempty"`
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
}

type SystemUpgradesObject struct {
	LatestLTS           string                      `json:"latestLTS,omitempty"`
	UpdateCenterRefresh string                      `json:"updateCenterRefresh,omitempty"`
	Upgrades            []SystemUpgradesObject_sub4 `json:"upgrades,omitempty"`
}

type SystemUpgradesObject_sub2 struct {
	Category              string `json:"category,omitempty"`
	Description           string `json:"description,omitempty"`
	EditionBundled        bool   `json:"editionBundled,omitempty"`
	Key                   string `json:"key,omitempty"`
	License               string `json:"license,omitempty"`
	Name                  string `json:"name,omitempty"`
	OrganizationName      string `json:"organizationName,omitempty"`
	OrganizationURL       string `json:"organizationUrl,omitempty"`
	TermsAndConditionsURL string `json:"termsAndConditionsUrl,omitempty"`
	Version               string `json:"version,omitempty"`
}

type SystemUpgradesObject_sub1 struct {
	Category         string `json:"category,omitempty"`
	Description      string `json:"description,omitempty"`
	EditionBundled   bool   `json:"editionBundled,omitempty"`
	Key              string `json:"key,omitempty"`
	License          string `json:"license,omitempty"`
	Name             string `json:"name,omitempty"`
	OrganizationName string `json:"organizationName,omitempty"`
	OrganizationURL  string `json:"organizationUrl,omitempty"`
}

type SystemUpgradesObject_sub4 struct {
	ChangeLogURL string                    `json:"changeLogUrl,omitempty"`
	Description  string                    `json:"description,omitempty"`
	DownloadURL  string                    `json:"downloadUrl,omitempty"`
	Plugins      SystemUpgradesObject_sub3 `json:"plugins,omitempty"`
	ReleaseDate  string                    `json:"releaseDate,omitempty"`
	Version      string                    `json:"version,omitempty"`
}

type SystemUpgradesObject_sub3 struct {
	Incompatible  []SystemUpgradesObject_sub1 `json:"incompatible,omitempty"`
	RequireUpdate []SystemUpgradesObject_sub2 `json:"requireUpdate,omitempty"`
}

type SystemChangeLogLevelOption struct {
	Level string `url:"level,omitempty"` // Description:"The new level. Be cautious: DEBUG, and even more TRACE, may have performance impacts.",ExampleValue:""
}

// ChangeLogLevel Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
func (s *SystemService) ChangeLogLevel(opt *SystemChangeLogLevelOption) (resp *http.Response, err error) {
	err = s.ValidateChangeLogLevelOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "system/change_log_level", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

// DbMigrationStatus Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func (s *SystemService) DbMigrationStatus() (v *SystemDbMigrationStatusObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/db_migration_status", nil)
	if err != nil {
		return
	}
	v = new(SystemDbMigrationStatusObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Health Provide health status of SonarQube.<p>Although global health is calculated based on both application and search nodes, detailed information is returned only for application nodes.</p><p>  <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p>
func (s *SystemService) Health() (v *SystemHealthObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/health", nil)
	if err != nil {
		return
	}
	v = new(SystemHealthObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Info Get detailed information about system configuration.<br/>Requires 'Administer' permissions.
func (s *SystemService) Info() (v *SystemInfoObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/info", nil)
	if err != nil {
		return
	}
	v = new(SystemInfoObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Liveness Provide liveness of SonarQube, meant to be used for a liveness probe on Kubernetes<p>Require 'Administer System' permission or authentication with passcode</p><p>When SonarQube is fully started, liveness check for database connectivity, Compute Engine status, and, except for DataCenter Edition, if ElasticSearch is Green or Yellow</p><p>When SonarQube is on Safe Mode (for example when a database migration is running), liveness check only for database connectivity</p><p>  <ul> <li>HTTP 204: this SonarQube node is alive</li> <li>Any other HTTP code: this SonarQube node is not alive, and should be reschedule</li> </ul></p>
func (s *SystemService) Liveness() (v *string, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/liveness", nil)
	if err != nil {
		return
	}
	v = new(string)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type SystemLogsOption struct {
	Process string `url:"process,omitempty"` // Description:"Process to get logs from",ExampleValue:""
}

// Logs Get system logs in plain-text format. Requires system administration permission.
func (s *SystemService) Logs(opt *SystemLogsOption) (v *string, resp *http.Response, err error) {
	err = s.ValidateLogsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "system/logs", opt)
	if err != nil {
		return
	}
	v = new(string)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// MigrateDb Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func (s *SystemService) MigrateDb() (v *SystemMigrateDbObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("POST", "system/migrate_db", nil)
	if err != nil {
		return
	}
	v = new(SystemMigrateDbObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Ping Answers "pong" as plain-text
func (s *SystemService) Ping() (v *string, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/ping", nil)
	if err != nil {
		return
	}
	v = new(string)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Restart Restarts server. Requires 'Administer System' permission. Performs a full restart of the Web, Search and Compute Engine Servers processes. Does not reload sonar.properties.
func (s *SystemService) Restart() (resp *http.Response, err error) {
	req, err := s.client.NewRequest("POST", "system/restart", nil)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

// Status Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>
func (s *SystemService) Status() (v *SystemStatusObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/status", nil)
	if err != nil {
		return
	}
	v = new(SystemStatusObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Upgrades Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
func (s *SystemService) Upgrades() (v *SystemUpgradesObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "system/upgrades", nil)
	if err != nil {
		return
	}
	v = new(SystemUpgradesObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
