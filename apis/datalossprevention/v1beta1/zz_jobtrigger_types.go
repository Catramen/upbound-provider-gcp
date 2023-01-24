/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionsObservation struct {
}

type ActionsParameters struct {

	// Publish a message into a given Pub/Sub topic when the job completes.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PubSub []PubSubParameters `json:"pubSub,omitempty" tf:"pub_sub,omitempty"`

	// Schedule for triggered jobs
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SaveFindings []SaveFindingsParameters `json:"saveFindings,omitempty" tf:"save_findings,omitempty"`
}

type BigQueryOptionsObservation struct {
}

type BigQueryOptionsParameters struct {

	// Max number of rows to scan. If the table has more rows than this value, the rest of the rows are omitted.
	// If not set, or if set to 0, all rows will be scanned. Only one of rowsLimit and rowsLimitPercent can be
	// specified. Cannot be used in conjunction with TimespanConfig.
	// +kubebuilder:validation:Optional
	RowsLimit *float64 `json:"rowsLimit,omitempty" tf:"rows_limit,omitempty"`

	// Max percentage of rows to scan. The rest are omitted. The number of rows scanned is rounded down.
	// Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of
	// rowsLimit and rowsLimitPercent can be specified. Cannot be used in conjunction with TimespanConfig.
	// +kubebuilder:validation:Optional
	RowsLimitPercent *float64 `json:"rowsLimitPercent,omitempty" tf:"rows_limit_percent,omitempty"`

	// How to sample bytes if not all bytes are scanned. Meaningful only when used in conjunction with bytesLimitPerFile.
	// If not specified, scanning would start from the top.
	// Possible values are TOP and RANDOM_START.
	// +kubebuilder:validation:Optional
	SampleMethod *string `json:"sampleMethod,omitempty" tf:"sample_method,omitempty"`

	// Set of files to scan.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	TableReference []TableReferenceParameters `json:"tableReference" tf:"table_reference,omitempty"`
}

type CloudStorageOptionsObservation struct {
}

type CloudStorageOptionsParameters struct {

	// Max number of bytes to scan from a file. If a scanned file's size is bigger than this value
	// then the rest of the bytes are omitted.
	// +kubebuilder:validation:Optional
	BytesLimitPerFile *float64 `json:"bytesLimitPerFile,omitempty" tf:"bytes_limit_per_file,omitempty"`

	// Max percentage of bytes to scan from a file. The rest are omitted. The number of bytes scanned is rounded down.
	// Must be between 0 and 100, inclusively. Both 0 and 100 means no limit.
	// +kubebuilder:validation:Optional
	BytesLimitPerFilePercent *float64 `json:"bytesLimitPerFilePercent,omitempty" tf:"bytes_limit_per_file_percent,omitempty"`

	// Set of files to scan.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	FileSet []FileSetParameters `json:"fileSet" tf:"file_set,omitempty"`

	// List of file type groups to include in the scan. If empty, all files are scanned and available data
	// format processors are applied. In addition, the binary content of the selected files is always scanned as well.
	// Images are scanned only as binary if the specified region does not support image inspection and no fileTypes were specified.
	// Each value may be one of BINARY_FILE, TEXT_FILE, IMAGE, WORD, PDF, AVRO, CSV, and TSV.
	// +kubebuilder:validation:Optional
	FileTypes []*string `json:"fileTypes,omitempty" tf:"file_types,omitempty"`

	// Limits the number of files to scan to this percentage of the input FileSet. Number of files scanned is rounded down.
	// Must be between 0 and 100, inclusively. Both 0 and 100 means no limit.
	// +kubebuilder:validation:Optional
	FilesLimitPercent *float64 `json:"filesLimitPercent,omitempty" tf:"files_limit_percent,omitempty"`

	// How to sample bytes if not all bytes are scanned. Meaningful only when used in conjunction with bytesLimitPerFile.
	// If not specified, scanning would start from the top.
	// Possible values are TOP and RANDOM_START.
	// +kubebuilder:validation:Optional
	SampleMethod *string `json:"sampleMethod,omitempty" tf:"sample_method,omitempty"`
}

type DatastoreOptionsObservation struct {
}

type DatastoreOptionsParameters struct {

	// A representation of a Datastore kind.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Kind []KindParameters `json:"kind" tf:"kind,omitempty"`

	// Datastore partition ID. A partition ID identifies a grouping of entities. The grouping
	// is always by project and namespace, however the namespace ID may be empty.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	PartitionID []PartitionIDParameters `json:"partitionId" tf:"partition_id,omitempty"`
}

type FileSetObservation struct {
}

type FileSetParameters struct {

	// The regex-filtered set of files to scan.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RegexFileSet []RegexFileSetParameters `json:"regexFileSet,omitempty" tf:"regex_file_set,omitempty"`

	// The Cloud Storage url of the file(s) to scan, in the format gs://<bucket>/<path>. Trailing wildcard
	// in the path is allowed.
	// If the url ends in a trailing slash, the bucket or directory represented by the url will be scanned
	// non-recursively (content in sub-directories will not be scanned). This means that gs://mybucket/ is
	// equivalent to gs://mybucket/*, and gs://mybucket/directory/ is equivalent to gs://mybucket/directory/*.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type InspectJobObservation struct {
}

type InspectJobParameters struct {

	// A task to execute on the completion of a job.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// The name of the template to run when this job is triggered.
	// +kubebuilder:validation:Required
	InspectTemplateName *string `json:"inspectTemplateName" tf:"inspect_template_name,omitempty"`

	// Information on where to inspect
	// Structure is documented below.
	// +kubebuilder:validation:Required
	StorageConfig []StorageConfigParameters `json:"storageConfig" tf:"storage_config,omitempty"`
}

type JobTriggerObservation struct {

	// an identifier for the resource with format {{parent}}/jobTriggers/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of the last time this trigger executed.
	LastRunTime *string `json:"lastRunTime,omitempty" tf:"last_run_time,omitempty"`

	// The resource name of the job trigger. Set by the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type JobTriggerParameters struct {

	// A description of the job trigger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User set display name of the job trigger.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Controls what and how to inspect for findings.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InspectJob []InspectJobParameters `json:"inspectJob,omitempty" tf:"inspect_job,omitempty"`

	// The parent of the trigger, either in the format projects/{{project}}
	// or projects/{{project}}/locations/{{location}}
	// +kubebuilder:validation:Required
	Parent *string `json:"parent" tf:"parent,omitempty"`

	// Whether the trigger is currently active.
	// Default value is HEALTHY.
	// Possible values are PAUSED, HEALTHY, and CANCELLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// What event needs to occur for a new job to be started.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Triggers []TriggersParameters `json:"triggers" tf:"triggers,omitempty"`
}

type KindObservation struct {
}

type KindParameters struct {

	// Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery.
	// For BigQuery: Required to filter out rows based on the given start and end times. If not specified and the table was
	// modified between the given start and end times, the entire table will be scanned. The valid data types of the timestamp
	// field are: INTEGER, DATE, TIMESTAMP, or DATETIME BigQuery column.
	// For Datastore. Valid data types of the timestamp field are: TIMESTAMP. Datastore entity will be scanned if the
	// timestamp property does not exist or its value is empty or invalid.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type OutputConfigObservation struct {
}

type OutputConfigParameters struct {

	// Schema used for writing the findings for Inspect jobs. This field is only used for
	// Inspect and must be unspecified for Risk jobs. Columns are derived from the Finding
	// object. If appending to an existing table, any columns from the predefined schema
	// that are missing will be added. No columns in the existing table will be deleted.
	// If unspecified, then all available columns will be used for a new table or an (existing)
	// table with no schema, and no changes will be made to an existing table that has a schema.
	// Only for use with external storage.
	// Possible values are BASIC_COLUMNS, GCS_COLUMNS, DATASTORE_COLUMNS, BIG_QUERY_COLUMNS, and ALL_COLUMNS.
	// +kubebuilder:validation:Optional
	OutputSchema *string `json:"outputSchema,omitempty" tf:"output_schema,omitempty"`

	// Information on the location of the target BigQuery Table.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Table []TableParameters `json:"table" tf:"table,omitempty"`
}

type PartitionIDObservation struct {
}

type PartitionIDParameters struct {

	// If not empty, the ID of the namespace to which the entities belong.
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

type PubSubObservation struct {
}

type PubSubParameters struct {

	// Cloud Pub/Sub topic to send notifications to.
	// +kubebuilder:validation:Required
	Topic *string `json:"topic" tf:"topic,omitempty"`
}

type RegexFileSetObservation struct {
}

type RegexFileSetParameters struct {

	// The name of a Cloud Storage bucket.
	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// A list of regular expressions matching file paths to exclude. All files in the bucket that match at
	// least one of these regular expressions will be excluded from the scan.
	// +kubebuilder:validation:Optional
	ExcludeRegex []*string `json:"excludeRegex,omitempty" tf:"exclude_regex,omitempty"`

	// A list of regular expressions matching file paths to include. All files in the bucket
	// that match at least one of these regular expressions will be included in the set of files,
	// except for those that also match an item in excludeRegex. Leaving this field empty will
	// match all files by default (this is equivalent to including .* in the list)
	// +kubebuilder:validation:Optional
	IncludeRegex []*string `json:"includeRegex,omitempty" tf:"include_regex,omitempty"`
}

type SaveFindingsObservation struct {
}

type SaveFindingsParameters struct {

	// Information on where to store output
	// Structure is documented below.
	// +kubebuilder:validation:Required
	OutputConfig []OutputConfigParameters `json:"outputConfig" tf:"output_config,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// With this option a job is started a regular periodic basis. For example: every day (86400 seconds).
	// A scheduled start time will be skipped if the previous execution has not ended when its scheduled time occurs.
	// This value must be set to a time duration greater than or equal to 1 day and can be no longer than 60 days.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	RecurrencePeriodDuration *string `json:"recurrencePeriodDuration,omitempty" tf:"recurrence_period_duration,omitempty"`
}

type StorageConfigObservation struct {
}

type StorageConfigParameters struct {

	// Options defining BigQuery table and row identifiers.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BigQueryOptions []BigQueryOptionsParameters `json:"bigQueryOptions,omitempty" tf:"big_query_options,omitempty"`

	// Options defining a file or a set of files within a Google Cloud Storage bucket.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudStorageOptions []CloudStorageOptionsParameters `json:"cloudStorageOptions,omitempty" tf:"cloud_storage_options,omitempty"`

	// Options defining a data set within Google Cloud Datastore.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DatastoreOptions []DatastoreOptionsParameters `json:"datastoreOptions,omitempty" tf:"datastore_options,omitempty"`

	// Information on where to inspect
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TimespanConfig []TimespanConfigParameters `json:"timespanConfig,omitempty" tf:"timespan_config,omitempty"`
}

type TableObservation struct {
}

type TableParameters struct {

	// The dataset ID of the table.
	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// The name of the table.
	// +kubebuilder:validation:Optional
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type TableReferenceObservation struct {
}

type TableReferenceParameters struct {

	// The dataset ID of the table.
	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// The name of the table.
	// +kubebuilder:validation:Required
	TableID *string `json:"tableId" tf:"table_id,omitempty"`
}

type TimespanConfigObservation struct {
}

type TimespanConfigParameters struct {

	// When the job is started by a JobTrigger we will automatically figure out a valid startTime to avoid
	// scanning files that have not been modified since the last time the JobTrigger executed. This will
	// be based on the time of the execution of the last run of the JobTrigger.
	// +kubebuilder:validation:Optional
	EnableAutoPopulationOfTimespanConfig *bool `json:"enableAutoPopulationOfTimespanConfig,omitempty" tf:"enable_auto_population_of_timespan_config,omitempty"`

	// Exclude files or rows newer than this value. If set to zero, no upper time limit is applied.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Exclude files or rows older than this value.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Information on where to inspect
	// Structure is documented below.
	// +kubebuilder:validation:Required
	TimestampField []TimestampFieldParameters `json:"timestampField" tf:"timestamp_field,omitempty"`
}

type TimestampFieldObservation struct {
}

type TimestampFieldParameters struct {

	// Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery.
	// For BigQuery: Required to filter out rows based on the given start and end times. If not specified and the table was
	// modified between the given start and end times, the entire table will be scanned. The valid data types of the timestamp
	// field are: INTEGER, DATE, TIMESTAMP, or DATETIME BigQuery column.
	// For Datastore. Valid data types of the timestamp field are: TIMESTAMP. Datastore entity will be scanned if the
	// timestamp property does not exist or its value is empty or invalid.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type TriggersObservation struct {
}

type TriggersParameters struct {

	// Schedule for triggered jobs
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

// JobTriggerSpec defines the desired state of JobTrigger
type JobTriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobTriggerParameters `json:"forProvider"`
}

// JobTriggerStatus defines the observed state of JobTrigger.
type JobTriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobTrigger is the Schema for the JobTriggers API. A job trigger configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type JobTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobTriggerSpec   `json:"spec"`
	Status            JobTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobTriggerList contains a list of JobTriggers
type JobTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobTrigger `json:"items"`
}

// Repository type metadata.
var (
	JobTrigger_Kind             = "JobTrigger"
	JobTrigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobTrigger_Kind}.String()
	JobTrigger_KindAPIVersion   = JobTrigger_Kind + "." + CRDGroupVersion.String()
	JobTrigger_GroupVersionKind = CRDGroupVersion.WithKind(JobTrigger_Kind)
)

func init() {
	SchemeBuilder.Register(&JobTrigger{}, &JobTriggerList{})
}
