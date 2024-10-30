package models

// Service represents a service.
type Service struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Mode      ServiceMode `json:"mode"`
	Image     string      `json:"image"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`

	Spec ServiceSpec `json:"spec"`

	// PreviousSpec *ServiceSpec  `json:",omitempty"`
	// Endpoint     Endpoint      `json:",omitempty"`
	// UpdateStatus *UpdateStatus `json:",omitempty"`
	//
	// // ServiceStatus is an optional, extra field indicating the number of
	// // desired and running tasks. It is provided primarily as a shortcut to
	// // calculating these values client-side, which otherwise would require
	// // listing all tasks for a service, an operation that could be
	// // computation and network expensive.
	// ServiceStatus *ServiceStatus `json:",omitempty"`
	//
	// // JobStatus is the status of a Service which is in one of ReplicatedJob or
	// // GlobalJob modes. It is absent on Replicated and Global services.
	// JobStatus *JobStatus `json:",omitempty"`
}

// models as docker -----------------------------------------------------------
type ServiceSpec struct {
	Name string `json:"name"`
	// Labels map[string]string `json:"Labels"`
	TaskTemplate ServiceTaskSpec `json:"task_template"`
	Mode         ServiceMode     `json:"mode"`
	// UpdateConfig   *UpdateConfig `json:",omitempty"`
	// RollbackConfig *UpdateConfig `json:",omitempty"`

	// // Networks specifies which networks the service should attach to.
	// //
	// // Deprecated: This field is deprecated since v1.44. The Networks field in TaskSpec should be used instead.
	// Networks     []NetworkAttachmentConfig `json:",omitempty"`
	EndpointSpec *EndpointSpec `json:"endpoint_spec"`
}

type ServiceMode struct {
	Replicated    *ReplicatedService `json:"replicated"`
	Global        *GlobalService     `json:"global"`
	ReplicatedJob *ReplicatedJob     `json:"replicated_job"`
	GlobalJob     *GlobalJob         `json:"global_job"`
}

// ReplicatedService is a kind of ServiceMode.
type ReplicatedService struct {
	Replicas int `json:"replicas"`
}

// GlobalService is a kind of ServiceMode.
type GlobalService struct{}

// ReplicatedJob is the a type of Service which executes a defined Tasks
// in parallel until the specified number of Tasks have succeeded.
type ReplicatedJob struct {
	// MaxConcurrent *uint64 `json:",omitempty"`
	// TotalCompletions *uint64 `json:",omitempty"`
}

// GlobalJob is the type of a Service which executes a Task on every Node
// matching the Service's placement constraints. These tasks run to completion
// and then exit.
//
// This type is deliberately empty.
type GlobalJob struct{}

// TaskSpec represents the spec of a task.
type ServiceTaskSpec struct {
	ContainerSpec *ServiceContainerSpec `json:"container_spec"`
	// PluginSpec            *runtime.PluginSpec       `json:",omitempty"`
	// NetworkAttachmentSpec *NetworkAttachmentSpec    `json:",omitempty"`
	// Resources             *ResourceRequirements     `json:",omitempty"`
	// RestartPolicy         *RestartPolicy            `json:",omitempty"`
	// Placement             *Placement                `json:",omitempty"`
	// Networks              []NetworkAttachmentConfig `json:",omitempty"`
	// LogDriver             *Driver                   `json:",omitempty"`
	// ForceUpdate           uint64
	// Runtime               RuntimeType `json:",omitempty"`
}

// ContainerSpec represents the spec of a container.
type ServiceContainerSpec struct {
	Image string `json:"image"`
	// Labels   map[string]string `json:",omitempty"`
	// Command  []string          `json:",omitempty"`
	// Args     []string          `json:",omitempty"`
	// Hostname string            `json:",omitempty"`
	// Env      []string          `json:",omitempty"`
	// Dir      string            `json:",omitempty"`
	// User     string            `json:",omitempty"`
	// Groups   []string          `json:",omitempty"`
	// Privileges      *Privileges             `json:",omitempty"`
	// Init            *bool                   `json:",omitempty"`
	// StopSignal      string                  `json:",omitempty"`
	// TTY             bool                    `json:",omitempty"`
	// OpenStdin       bool                    `json:",omitempty"`
	// ReadOnly        bool                    `json:",omitempty"`
	// Mounts          []mount.Mount           `json:",omitempty"`
	// StopGracePeriod *time.Duration          `json:",omitempty"`
	// Healthcheck     *container.HealthConfig `json:",omitempty"`
	// The format of extra hosts on swarmkit is specified in:
	// http://man7.org/linux/man-pages/man5/hosts.5.html
	//    IP_address canonical_hostname [aliases...]
	// Hosts          []string            `json:",omitempty"`
	// DNSConfig      *DNSConfig          `json:",omitempty"`
	// Secrets        []*SecretReference  `json:",omitempty"`
	// Configs        []*ConfigReference  `json:",omitempty"`
	// Isolation      container.Isolation `json:",omitempty"`
	// Sysctls        map[string]string   `json:",omitempty"`
	// CapabilityAdd  []string            `json:",omitempty"`
	// CapabilityDrop []string            `json:",omitempty"`
	// Ulimits        []*container.Ulimit `json:",omitempty"`
	// OomScoreAdj    int64               `json:",omitempty"`
}

// EndpointSpec represents the spec of an endpoint.
type EndpointSpec struct {
	Mode  string       `json:"mode"`
	Ports []PortConfig `json:"ports"`
}

// PortConfig represents the config of a port.
type PortConfig struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	// TargetPort is the port inside the container
	TargetPort uint32 `json:"target_port"`
	// PublishedPort is the port on the swarm hosts
	PublishedPort uint32 `json:"published_port"`
	// PublishMode is the mode in which port is published
	PublishMode string `json:"publish_mode"`
}

/*
{
    "ID": "67a1nabpe8d0kgp1lu1tk5r9r",
    "Version": {
        "Index": 1399422
    },
    "CreatedAt": "2024-10-25T10:51:20.079601051Z",
    "UpdatedAt": "2024-10-25T10:51:20.079601051Z",
    "Spec": {
        "Name": "32b02965_ba_w1_main_94748aeaf2774_tests",
        "Labels": {},
        "TaskTemplate": {
            "ContainerSpec": {
                "Image": "172.28.1.1:5000/kaspersky/kata/management/authorization_service/test:test@sha256:b284f19ae3798c9f069ea9420ca173843f4c22760056b1df4f22aab02376e836",
                "Args": [
                    "--runner-command",
                    "run-tests",
                    "--suite",
                    "test_auth",
                    "test_client",
                    "test_keytab",
                    "test_secrets",
                    "test_users",
                    "test_api_auth",
                    "test_api_kerberos_auth",
                    "test_api_keytab",
                    "test_api_secrets",
                    "test_api_user",
                    "--log-level",
                    "INFO"
                ],
                "Env": [
                    "SWARM_MANAGER=tcp://172.28.2.1:2375",
                    "SWARM_MANAGER_IP=172.28.2.1",
                    "SWARM_MANAGER_PORT=2375",
                    "LOGGING_HOST=",
                    "LOGGING_PORT=",
                    "PYTHON_SOURCES=nginx_gateway,request_util,split_delivery_tool,product_verification,third_party_tools,svace_tools,dependency_source_publisher,prometheus_entrypoint,monitoring_server,kata_health_collector,etcd_entrypoint,kata_updater_client,av_bases_downloader,updates_consistency_checker,kata_license_api,kata_updater_api,kata_updater_utils,kata_crypto,vm_manager,deps_resolver,registry_cli,docker_cleaner,comment_stripper,artifacts_cleaner,system_tools,cert_tools,pycode_jsonschema,kata_docker_save,kashell_sandbox,kashell,compose_generate,java_entrypoint,kafka_tools,events_generator,kata_prometheus_exporter_utils,kata_monitor_models,kata_sedr_facade,appliance_docker_common,kata_antiaptdb_schema,kata_constants,appliance_common_fixtures,kata_filesystem_utils,kata_supervisor_utils,kata_stats,kata_strings,kata_scan_rules,kata_certificate_utils,kata_antiaptdb,kata_scanner_client,kata_scan_broker,kata_request_utils,kata_settings,kata_scanner_common,kata_reports,kata_s3_client,kata_ksn,kata_coverage,kata_notifications,kata_kafka_utils,kata_kerberos,kata_loggers,kata_settings_manager,appliance_node_orchestrator,appliance_unit_testing,json_generator,s3_tools,unit_testing,fastapi_tools,deploy_config,health_checker,artifactory_tools,testing_entrypoint,testing_tools,iso_testing_tools,coverage_tools,swarm_pytest,settings_updater,retrying,authorization_service,auth_cli,siem_proxy,authorization_service_client,edr_synchronizer,statistics_clickhouse_importer,inactive_agents_cleaner,authorization_service_schemas,statistics_clickhouse_configurator,management_ui_entrypoint,pcap_formatter,kata_monitor_primary,password_expiration_watcher,multitenancy_configurator,kata_monitor,fluentd_entrypoint,serialization,distribution_build_tools,deployment_management_api,node_settings_synchronizer,installer_sensor_ui,installer_os,deployment_management_api_client,ipsec_manager,iso_builder,installer_ui,deployment_api_client,kata_upgrade_host,upgrade_builder,backup_restore,deployment_management_api_schemas,deployment_cli,kata_upgrade,sandbox_installer,installer,restart_product,redeploy_perf,base_installer,ipsec_manager_client,distribution_builder,deployment_api_schemas,collect,firewall,deployment_api,docker_secret,swarm_deploy,distributed_docker_build,docker_build,dind_deploy,docker_utils,agent_server,ksb_test_api,ksb_agent_server,ksb_agent_server_schemas,agent_events_matcher,agent_packet_dumper,agent_storage_manager,agent_test_api,agent_event_validator,agent_event_es_importer,sensor_management_api_schemas,sensor_management_api,ids_alert_syncer,multitenancy_management_api_schemas,multitenancy_management_api_client,multitenancy_management_api,kafka_proxy_entrypoint,schema_registry_entrypoint,kafka_entrypoint,zookeeper_entrypoint,kafka_configurator,kata_license,nta_libs,admin_menu,kata_tui,sandbox_pdk_dummy,kata_ram_stub,sandbox_ram_backend,sandbox_data_collector,sandbox_manager_api,sandbox_client,sandbox_ram_backend_postinstall,sandbox_manager,read_metrics_stub,sandbox_management_api_schemas,sandbox_updater,sandbox_management_api,fastsearch_events_importer,ioa_update_validator,kata_scan_agents,scan_server_synchronizer,kata_scanner_api,kata_collector,coredns_entrypoint,discovery_entrypoint,dhcp_server,kata_dns,service_registrator,config_utils,services_configurator,console_settings_updater,ioc,ceph_management_api,kata_ceph,ceph_management_api_schemas,ceph_entrypoint,elasticsearch_entrypoint,s3rotator,ksqldb_configurator_service,kafka_connect_client,ksqldb_configurator_models,ksqldb_client,ksqldb_entrypoint,s3proxy_entrypoint,kata_python_bootstrap,kata_linter,compose_template,auto_tests",
                    "IMAGE_VERSIONS=ewogICJidWlsZF9pbWFnZXMiOiB7CiAgICAiYWRtaW5fbWVudV90ZXN0cyI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvdWkvYWRtaW5fbWVudS90ZXN0OnRlc3QiLAogICAgImFnZW50X2RhdGFiYXNlX2NvbmZpZ3VyYXRvciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvYWdlbnRfZGF0YWJhc2VfY29uZmlndXJhdG9yOnRlc3QiLAogICAgImFnZW50X2RhdGFiYXNlX3N5bmNocm9uaXplciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvYWdlbnRfZGF0YWJhc2Vfc3luY2hyb25pemVyOnRlc3QiLAogICAgImFnZW50X3NlcnZlciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvYWdlbnRfc2VydmVyOnRlc3QiLAogICAgImFudGlhcHRfZGF0YWJhc2VfY29uZmlndXJhdG9yIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9hbnRpYXB0X2RhdGFiYXNlX2NvbmZpZ3VyYXRvcjp0ZXN0IiwKICAgICJhcHBsaWFuY2VfcHl0ZXN0X3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS90ZXN0L2FwcGxpYW5jZV9weXRlc3Q6YzYwY2I5MSIsCiAgICAiYXV0aG9yaXphdGlvbl9zZXJ2aWNlIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9tYW5hZ2VtZW50L2F1dGhvcml6YXRpb25fc2VydmljZTp0ZXN0IiwKICAgICJhdXRob3JpemF0aW9uX3NlcnZpY2VfdGVzdHMiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL21hbmFnZW1lbnQvYXV0aG9yaXphdGlvbl9zZXJ2aWNlL3Rlc3Q6dGVzdCIsCiAgICAiYXZfYmFzZXNfZG93bmxvYWRlciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvdXBkYXRlci9hdl9iYXNlc19kb3dubG9hZGVyOmQxNjE0ZTUiLAogICAgImNlbnRvc19zZGwiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2Jhc2UvY2VudG9zLzcuOToxLjUiLAogICAgImNsaWNraG91c2UiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2Jhc2UvdWJ1bnR1LzIyLjA0L2NsaWNraG91c2UvMjAuMTEuMi4xOjM0IiwKICAgICJkZXBsb3ltZW50X2FwaSI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvZGVwbG95bWVudC9kZXBsb3ltZW50X2FwaTp0ZXN0IiwKICAgICJkZXBsb3ltZW50X21hbmFnZW1lbnRfYXBpIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9kZXBsb3ltZW50L2RlcGxveW1lbnRfbWFuYWdlbWVudF9hcGk6MDBiODVkMiIsCiAgICAiZGVwbG95bWVudF9tYW5hZ2VtZW50X2FwaV90ZXN0cyI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvZGVwbG95bWVudC9kZXBsb3ltZW50X21hbmFnZW1lbnRfYXBpL3Rlc3Q6dGVzdCIsCiAgICAiZGluZCI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvYmFzZS91YnVudHUvZGluZDozNCIsCiAgICAiZG9ja2VyX3ZvbHVtZV9sb2NhbF9wZXJzaXN0IjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9iYXNlL3VidW50dS8yMi4wNC9jYXJib25pcXVlL2RvY2tlci12b2x1bWUtbG9jYWwtcGVyc2lzdC80LjEuMjozNiIsCiAgICAiZG9ja2VyX3ZvbHVtZV9yYmQiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2Jhc2UvdWJ1bnR1LzIyLjA0L3dldG9waS9yYmQvMy4wLjE6MzQiLAogICAgImVkcl9zeW5jaHJvbml6ZXIiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2Vkcl9zeW5jaHJvbml6ZXI6dGVzdCIsCiAgICAiZWxhc3RpY3NlYXJjaCI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvc3RvcmFnZS9lbGFzdGljc2VhcmNoOnRlc3QiLAogICAgImVsYXN0aWNzZWFyY2hfZXhwb3J0ZXIiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2Jhc2UvdWJ1bnR1LzIyLjA0L2VsYXN0aWNzZWFyY2gtZXhwb3J0ZXIvMS4zLjA6MzQiLAogICAgImV0Y2QiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2V0Y2Q6dGVzdCIsCiAgICAiZmFzdHNlYXJjaCI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvZmFzdHNlYXJjaDp0ZXN0IiwKICAgICJmbHVlbnRkIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9mbHVlbnRkLzQuMDplMDgzMzU2IiwKICAgICJodW50aW5nX3N1cHBvcnRfdG9vbCI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvaHVudGluZ19zdXBwb3J0X3Rvb2w6dGVzdCIsCiAgICAiaWRzX2FsZXJ0X3N5bmNlciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvaWRzX2FsZXJ0X3N5bmNlcjp0ZXN0IiwKICAgICJpZHNfYWxlcnRfc3luY2VyX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9uZXR3b3JrX2FnZW50L2lkc19hbGVydF9zeW5jZXIvdGVzdDo2NDVlZWJiIiwKICAgICJpb2FfdXBkYXRlX3ZhbGlkYXRvciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvaW9hX3VwZGF0ZV92YWxpZGF0b3I6dGVzdCIsCiAgICAia2Fma2EiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2thZmthOnRlc3QiLAogICAgImthZmthX2NvbmZpZ3VyYXRvciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEva2Fma2Eva2Fma2FfY29uZmlndXJhdG9yOnRlc3QiLAogICAgImthZmthX2V4cG9ydGVyIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9iYXNlL3VidW50dS8yMi4wNC9rYWZrYS1leHBvcnRlci8xLjIuMDozNCIsCiAgICAia2F0YV9hbGxfamF2YSI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEva2F0YV9hbGxfamF2YTp0ZXN0IiwKICAgICJrYXRhX2FsbF9weXRob24iOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2thdGFfYWxsX3B5dGhvbjp0ZXN0IiwKICAgICJrYXRhX21vbml0b3JfcHJpbWFyeV90ZXN0cyI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvbWFuYWdlbWVudC9tYW5hZ2VtZW50X3VpL2thdGFfbW9uaXRvcl9wcmltYXJ5L3Rlc3Q6YzYwY2I5MSIsCiAgICAia2F0YV9tb25pdG9yX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9tYW5hZ2VtZW50L21hbmFnZW1lbnRfdWkva2F0YV9tb25pdG9yL3Rlc3Q6ZTAxMjM1NCIsCiAgICAia2F0YV9zY2FubmVyIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9rYXRhX3NjYW5uZXI6ZDE2MTRlNSIsCiAgICAia2F0YV9zdHJpbmdzX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS90b29scy9hcHBsaWFuY2Uva2F0YV9zdHJpbmdzL3Rlc3Q6N2FmMTY2MCIsCiAgICAia3NiX2FnZW50X3NlcnZlciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvZW5kcG9pbnRfYWdlbnQva3NiX2FnZW50X3NlcnZlcjp0ZXN0IiwKICAgICJrc2JfYWdlbnRfc2VydmVyX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9lbmRwb2ludF9hZ2VudC9rc2JfYWdlbnRfc2VydmVyL3Rlc3Q6dGVzdCIsCiAgICAia3NiX3Rlc3RfYXBpIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9lbmRwb2ludF9hZ2VudC9rc2JfdGVzdF9hcGk6dGVzdCIsCiAgICAia3NiX3Rlc3RfYXBpX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9lbmRwb2ludF9hZ2VudC9rc2JfdGVzdF9hcGkvdGVzdDp0ZXN0IiwKICAgICJrc25fcHJveHkiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL2tzbl9wcm94eTp0ZXN0IiwKICAgICJtb25pdG9yaW5nX3Byb21ldGhldXMiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL21vbml0b3JpbmdfcHJvbWV0aGV1czp0ZXN0IiwKICAgICJtdWx0aXRlbmFuY3lfbWFuYWdlbWVudF9hcGkiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL211bHRpdGVuYW5jeS9tdWx0aXRlbmFuY3lfbWFuYWdlbWVudF9hcGk6NWJlN2MyMSIsCiAgICAibXVsdGl0ZW5hbmN5X21hbmFnZW1lbnRfYXBpX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9tdWx0aXRlbmFuY3kvbXVsdGl0ZW5hbmN5X21hbmFnZW1lbnRfYXBpL3Rlc3Q6YzBlNWUxYSIsCiAgICAicG9zdGZpeCI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvcG9zdGZpeDp0ZXN0IiwKICAgICJwb3N0Z3Jlc3FsX2V4cG9ydGVyIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9iYXNlL3VidW50dS8yMi4wNC9wb3N0Z3Jlc3FsLWV4cG9ydGVyLzAuOC4wOjM0IiwKICAgICJwb3N0Z3Jlc3FsX3NlcnZlciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvdWJ1bnR1L3Bvc3RncmVzcWxfc2VydmVyOnRlc3QiLAogICAgInByaW1hcnlfZGF0YWJhc2VfY29uZmlndXJhdG9yIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9wcmltYXJ5X2RhdGFiYXNlX2NvbmZpZ3VyYXRvcjo3N2Q1Y2M5IiwKICAgICJweXRlc3RfdGVzdHMiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL3Rlc3QvcHl0ZXN0OnRlc3QiLAogICAgInB5dGhvbi10ZXN0cy1jb21tb24iOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL3Rlc3QvcHl0aG9uLXRlc3RzLWNvbW1vbjp0ZXN0IiwKICAgICJyZWRpcyI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvcmVkaXM6dGVzdCIsCiAgICAiczNwcm94eSI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvc3RvcmFnZS9zM3Byb3h5LzEuOC4wOnRlc3QiLAogICAgInNlcnZpY2VzX2NvbmZpZ3VyYXRvciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvdG9vbHMvc2VydmljZXNfY29uZmlndXJhdG9yOnRlc3QiLAogICAgInNpZW1fcHJveHkiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL21hbmFnZW1lbnQvc2llbV9wcm94eTp0ZXN0IiwKICAgICJzbm1wZCI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvc25tcGQ6dGVzdCIsCiAgICAic3VyaWNhdGFfcnVsZXNfdmFsaWRhdG9yIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS9zdXJpY2F0YV9ydWxlc192YWxpZGF0b3I6dGVzdCIsCiAgICAidXBkYXRlciI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvdXBkYXRlcjp0ZXN0IiwKICAgICJ1cGRhdGVyX3Rlc3RzIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS91cGRhdGVyL3Rlc3Q6dGVzdCIsCiAgICAid2ViX2JhY2tlbmQiOiAiMTcyLjI4LjEuMTo1MDAwL2thc3BlcnNreS9rYXRhL21hbmFnZW1lbnQvbWFuYWdlbWVudF91aS93ZWJfYmFja2VuZDp0ZXN0IiwKICAgICJ3ZWJfYmFja2VuZF90ZXN0cyI6ICIxNzIuMjguMS4xOjUwMDAva2FzcGVyc2t5L2thdGEvbWFuYWdlbWVudC9tYW5hZ2VtZW50X3VpL3dlYl9iYWNrZW5kL3Rlc3Q6dGVzdCIsCiAgICAiem9va2VlcGVyIjogIjE3Mi4yOC4xLjE6NTAwMC9rYXNwZXJza3kva2F0YS96b29rZWVwZXI6dGVzdCIKICB9LAogICJidWlsZF9udW1iZXIiOiAidGVzdCIsCiAgImRpc3RyaWJ1dGlvbiI6ICJ1YnVudHUiLAogICJwcm9kdWN0X3ZlcnNpb24iOiAiNy4wLjAuMzgwIgp9",
                    "HUNTS_ENRICHED_TIMEOUT=2m",
                    "HUNTS_APPLIED_TIMEOUT=2m",
                    "SSH_PUBLIC_KEY=c3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCQVFEbFM5YUF4bmhVNzNBaHdLM0QvdFAxYThiamtWKzB2V0JhQVlIVmJ2VzNXcFpwamRIVTFmRW5mQlRLTWtkME1BWjE0Q0tYZDlCQWUydUU3WDY1WER3YU5QUWpxTkFWMjg2VGVaS1NRN3hSbmEyeUx6R056ejFyZ0hwNlcvUDBuTUFMTk5QeFRLOVl1WHEydHpMWit1UWNyK2ZDNnFBSmdOS3hFU2tkVlEyTUJJZlBXRDdtcGNqbEc5ZE96MVBPdVRoMHRsTU1mdURZdjVZSzdsRnp6aEl1YjlDbmorcnpKcktmQzROWXZWVVVJVHY4K3I0UVRBTVpIdmhYYUpQL05ReGNsZ2ZLS0g3dTFiR2ZCWG55amNxUi9QU0dNbVpFV0JCNlE2UEdmMlFySEgvSjl4ZGx1QnYrKzlqK0xsVmM3NmZPRHlJZksvY1dtZnBFZmhPLzBqdE4gaWdvci5uZWZlZGtpbkBrYXNwZXJza3kuY29tCg==",
                    "SSH_PRIVATE_KEY=LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0KYjNCbGJuTnphQzFyWlhrdGRqRUFBQUFBQkc1dmJtVUFBQUFFYm05dVpRQUFBQUFBQUFBQkFBQUJGd0FBQUFkemMyZ3RjbgpOaEFBQUFBd0VBQVFBQUFRRUE1VXZXZ01aNFZPOXdJY0N0dy83VDlXdkc0NUZmdEwxZ1dnR0IxVzcxdDFxV2FZM1IxTlh4Ckozd1V5akpIZERBR2RlQWlsM2ZRUUh0cmhPMSt1Vnc4R2pUMEk2alFGZHZPazNtU2trTzhVWjJ0c2k4eGpjODlhNEI2ZWwKdno5SnpBQ3pUVDhVeXZXTGw2dHJjeTJmcmtISy9ud3VxZ0NZRFNzUkVwSFZVTmpBU0h6MWcrNXFYSTVSdlhUczlUenJrNApkTFpUREg3ZzJMK1dDdTVSYzg0U0xtL1FwNC9xOHlheW53dURXTDFWRkNFNy9QcStFRXdER1I3NFYyaVQvelVNWEpZSHlpCmgrN3RXeG53VjU4bzNLa2Z6MGhqSm1SRmdRZWtPanhuOWtLeHgveWZjWFpiZ2IvdnZZL2k1VlhPK256ZzhpSHl2M0ZwbjYKUkg0VHY5STdUUUFBQTlqTW9wOU56S0tmVFFBQUFBZHpjMmd0Y25OaEFBQUJBUURsUzlhQXhuaFU3M0Fod0szRC90UDFhOApiamtWKzB2V0JhQVlIVmJ2VzNXcFpwamRIVTFmRW5mQlRLTWtkME1BWjE0Q0tYZDlCQWUydUU3WDY1WER3YU5QUWpxTkFWCjI4NlRlWktTUTd4Um5hMnlMekdOenoxcmdIcDZXL1Awbk1BTE5OUHhUSzlZdVhxMnR6TFordVFjcitmQzZxQUpnTkt4RVMKa2RWUTJNQklmUFdEN21wY2psRzlkT3oxUE91VGgwdGxNTWZ1RFl2NVlLN2xGenpoSXViOUNuaityekpyS2ZDNE5ZdlZVVQpJVHY4K3I0UVRBTVpIdmhYYUpQL05ReGNsZ2ZLS0g3dTFiR2ZCWG55amNxUi9QU0dNbVpFV0JCNlE2UEdmMlFySEgvSjl4CmRsdUJ2Kys5aitMbFZjNzZmT0R5SWZLL2NXbWZwRWZoTy8wanROQUFBQUF3RUFBUUFBQVFFQXZiVkhIWGdYYXI1SFdjV2UKZnNCQkFaWlFFVXZma1J2MXNjSDVkMnE4WlJ6UHdUa1MyRExCdU5kU0pCQUJTa0hKdHBEZ1RjMVZRV1JiaXJrcjFaUllMWgo2T04wcWVEdGV6VEJTVGpldjR6TlhXTmN3U0JhL21zeUN6TkxVTFcrQ3NpeGtFR1dwRzBQTjBjZU9VY1FEQTVuTXJIdnNiCnVZRVowR0hRVUhKa0tPaXRWdEE2SzludzJ3ZEZMK1hJUHZCRXBRbDlVOXN4VXZVT3o2NXBtRjEzbmdTTnpqM1J5ZnAxVjYKTWYydy9JZlpDWUZCMGlENElpVlJTYURUaDVoa2M0bEU3b0R4cVhOSGgxT1BMMlZnTVd2cjRSMkN3K2ZueWYzQzZEeFV1VQpCR2NudnZ3dXloUG9sQks1bXR6cWIvaXk3b1JiSmVSUjl5b0lxNXZ4VUNxSEFRQUFBSUFFYWwzcEp1VmhsWXVuajdhb3JQCkU2NEtMc3NacWZaR2s3clZ1YzJmdUdzQTV4UXRIalJTOWs0TkkwMUVwR2h0a3lFcEM2Q1dBQVdwYXMwZTdWNWpiTFRKVEkKZHVHZzh0Zk96emxlcERhRjJGak10OUQ0eGIzNm93dnZrdTQ5WVZKYU05TTE5R08xWTQ5Q2IwVGJuRDRRbjBlQ1dBMFczbgpGeUlwVDQwVFhWWEFBQUFJRUEvYnFBMWRoaDJUcWVoeW9ueFFhVWJKZjlvQUtkSHZ4Z1BhY2xLNUdOQ1ZLbDRNd3VpVExoCnJwV05jdmllY24vd1N4eExPd1lHYXNJWnZqc0UzYXczbzRsWEs0K1RYUlhHWmJPQllUSHhwSGcrVGt1M1UzM09mMmJZOHoKRU0vYmFyMytieUxQRUlmU2IyUVo2OFJWTDVlRnBDSnlzd3NUMzBzR29lM3ViKzVtRUFBQUNCQU9kWlZ5OE8rNFJmdG1TWApLVDRacmI3SEoyQUZsVDlMNDNMKzUyeXd6cWZkcjV5VUQ4NTVqV1dnWkxaNi93U3BGNVJpS3NhSzdWMFUzRmdvaFZXQlNLCktXa2NaSXR4Tyt0MkNjbFh2bnVXT2hOdmx5Q253L0lqNVRkSUtuMDB0cnErN20vckcycjlwNzJJcTR1NS9SY1NTNFFRRDgKV2FtUUxzUFZ3dEFPTktSdEFBQUFHMmxuYjNJdWJtVm1aV1JyYVc1QWEyRnpjR1Z5YzJ0NUxtTnZiUUVDQXdRRkJnYz0KLS0tLS1FTkQgT1BFTlNTSCBQUklWQVRFIEtFWS0tLS0tCg=="
                ],
                "Init": false,
                "DNSConfig": {},
                "Isolation": "default"
            },
            "Resources": {
                "Limits": {},
                "Reservations": {}
            },
            "RestartPolicy": {
                "Condition": "none",
                "Delay": 5000000000,
                "MaxAttempts": 0
            },
            "Placement": {
                "Platforms": [
                    {
                        "Architecture": "amd64",
                        "OS": "linux"
                    }
                ]
            },
            "ForceUpdate": 0,
            "Runtime": "container"
        },
        "Mode": {
            "Replicated": {
                "Replicas": 1
            }
        },
        "EndpointSpec": {
            "Mode": "vip"
        }
    },
    "Endpoint": {
        "Spec": {}
    }
}
*/
