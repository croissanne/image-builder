// Package provisioning provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package provisioning

import (
	"time"
)

// V1AWSReservationRequest defines model for v1.AWSReservationRequest.
type V1AWSReservationRequest struct {
	Amount           *int32  `json:"amount,omitempty"`
	ImageId          *string `json:"image_id,omitempty"`
	InstanceType     *string `json:"instance_type,omitempty"`
	LaunchTemplateId *string `json:"launch_template_id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Poweroff         *bool   `json:"poweroff,omitempty"`
	PubkeyId         *int64  `json:"pubkey_id,omitempty"`
	Region           *string `json:"region,omitempty"`
	SourceId         *string `json:"source_id,omitempty"`
}

// V1AWSReservationResponse defines model for v1.AWSReservationResponse.
type V1AWSReservationResponse struct {
	Amount           *int32  `json:"amount,omitempty"`
	AwsReservationId *string `json:"aws_reservation_id,omitempty"`
	ImageId          *string `json:"image_id,omitempty"`
	InstanceType     *string `json:"instance_type,omitempty"`
	Instances        *[]struct {
		Detail *struct {
			PublicDns  *string `json:"public_dns,omitempty"`
			PublicIpv4 *string `json:"public_ipv4,omitempty"`
		} `json:"detail,omitempty"`
		InstanceId *string `json:"instance_id,omitempty"`
	} `json:"instances,omitempty"`
	LaunchTemplateId *string `json:"launch_template_id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Poweroff         *bool   `json:"poweroff,omitempty"`
	PubkeyId         *int64  `json:"pubkey_id,omitempty"`
	Region           *string `json:"region,omitempty"`
	ReservationId    *int64  `json:"reservation_id,omitempty"`
	SourceId         *string `json:"source_id,omitempty"`
}

// V1AccountIDTypeResponse defines model for v1.AccountIDTypeResponse.
type V1AccountIDTypeResponse struct {
	Aws *struct {
		AccountId *string `json:"account_id,omitempty"`
	} `json:"aws,omitempty"`
}

// V1AvailabilityStatusRequest defines model for v1.AvailabilityStatusRequest.
type V1AvailabilityStatusRequest struct {
	SourceId *string `json:"source_id,omitempty"`
}

// V1AzureReservationRequest defines model for v1.AzureReservationRequest.
type V1AzureReservationRequest struct {
	Amount       *int64  `json:"amount,omitempty"`
	ImageId      *string `json:"image_id,omitempty"`
	InstanceSize *string `json:"instance_size,omitempty"`
	Location     *string `json:"location,omitempty"`
	Name         *string `json:"name,omitempty"`
	Poweroff     *bool   `json:"poweroff,omitempty"`
	PubkeyId     *int64  `json:"pubkey_id,omitempty"`
	SourceId     *string `json:"source_id,omitempty"`
}

// V1AzureReservationResponse defines model for v1.AzureReservationResponse.
type V1AzureReservationResponse struct {
	Amount       *int64  `json:"amount,omitempty"`
	ImageId      *string `json:"image_id,omitempty"`
	InstanceSize *string `json:"instance_size,omitempty"`
	Instances    *[]struct {
		Detail *struct {
			PublicDns  *string `json:"public_dns,omitempty"`
			PublicIpv4 *string `json:"public_ipv4,omitempty"`
		} `json:"detail,omitempty"`
		InstanceId *string `json:"instance_id,omitempty"`
	} `json:"instances,omitempty"`
	Location      *string `json:"location,omitempty"`
	Name          *string `json:"name,omitempty"`
	Poweroff      *bool   `json:"poweroff,omitempty"`
	PubkeyId      *int64  `json:"pubkey_id,omitempty"`
	ReservationId *int64  `json:"reservation_id,omitempty"`
	SourceId      *string `json:"source_id,omitempty"`
}

// V1GenericReservationResponsePayload defines model for v1.GenericReservationResponsePayload.
type V1GenericReservationResponsePayload struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	Error      *string    `json:"error,omitempty"`
	FinishedAt *time.Time `json:"finished_at"`
	Id         *int64     `json:"id,omitempty"`
	Provider   *int       `json:"provider,omitempty"`
	Status     *string    `json:"status,omitempty"`
	Step       *int32     `json:"step,omitempty"`
	StepTitles *[]string  `json:"step_titles,omitempty"`
	Steps      *int32     `json:"steps,omitempty"`
	Success    *bool      `json:"success"`
}

// V1InstanceTypeResponse defines model for v1.InstanceTypeResponse.
type V1InstanceTypeResponse struct {
	Architecture *string `json:"architecture,omitempty"`
	Azure        *struct {
		GenV1 *bool `json:"gen_v1,omitempty"`
		GenV2 *bool `json:"gen_v2,omitempty"`
	} `json:"azure,omitempty"`
	Cores     *int32  `json:"cores,omitempty"`
	MemoryMib *int64  `json:"memory_mib,omitempty"`
	Name      *string `json:"name,omitempty"`
	StorageGb *int64  `json:"storage_gb,omitempty"`
	Supported *bool   `json:"supported,omitempty"`
	Vcpus     *int32  `json:"vcpus,omitempty"`
}

// V1LaunchTemplatesResponse defines model for v1.LaunchTemplatesResponse.
type V1LaunchTemplatesResponse struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// V1NoopReservationResponse defines model for v1.NoopReservationResponse.
type V1NoopReservationResponse struct {
	ReservationId *int64 `json:"reservation_id,omitempty"`
}

// V1PubkeyRequest defines model for v1.PubkeyRequest.
type V1PubkeyRequest struct {
	Body *string `json:"body,omitempty"`
	Name *string `json:"name,omitempty"`
}

// V1PubkeyResponse defines model for v1.PubkeyResponse.
type V1PubkeyResponse struct {
	Body              *string `json:"body,omitempty"`
	Fingerprint       *string `json:"fingerprint,omitempty"`
	FingerprintLegacy *string `json:"fingerprint_legacy,omitempty"`
	Id                *int64  `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
}

// V1ResponseError defines model for v1.ResponseError.
type V1ResponseError struct {
	BuildTime   *string `json:"build_time,omitempty"`
	EdgeId      *string `json:"edge_id,omitempty"`
	Environment *string `json:"environment,omitempty"`
	Error       *string `json:"error,omitempty"`
	Msg         *string `json:"msg,omitempty"`
	TraceId     *string `json:"trace_id,omitempty"`
	Version     *string `json:"version,omitempty"`
}

// V1SourceResponse defines model for v1.SourceResponse.
type V1SourceResponse struct {
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	SourceTypeId *string `json:"source_type_id,omitempty"`
	Uid          *string `json:"uid,omitempty"`
}

// V1SourceUploadInfoResponse defines model for v1.SourceUploadInfoResponse.
type V1SourceUploadInfoResponse struct {
	Aws *struct {
		AccountId *string `json:"account_id,omitempty"`
	} `json:"aws"`
	Azure *struct {
		ResourceGroups *[]string `json:"resource_groups,omitempty"`
		SubscriptionId *string   `json:"subscription_id,omitempty"`
		TenantId       *string   `json:"tenant_id,omitempty"`
	} `json:"azure"`
	Provider *string `json:"provider,omitempty"`
}

// InternalError defines model for InternalError.
type InternalError = V1ResponseError

// NotFound defines model for NotFound.
type NotFound = V1ResponseError

// AvailabilityStatusJSONBody defines parameters for AvailabilityStatus.
type AvailabilityStatusJSONBody = V1AvailabilityStatusRequest

// GetInstanceTypeListAllParams defines parameters for GetInstanceTypeListAll.
type GetInstanceTypeListAllParams struct {
	// Region to list instance types within. This is required.
	Region string `form:"region" json:"region"`

	// Availability zone (or location) to list instance types within. Not applicable for AWS EC2 as all zones within a region are the same (will lead to an error when used). Required for Azure.
	Zone *string `form:"zone,omitempty" json:"zone,omitempty"`
}

// CreatePubkeyJSONBody defines parameters for CreatePubkey.
type CreatePubkeyJSONBody = V1PubkeyRequest

// CreateAwsReservationJSONBody defines parameters for CreateAwsReservation.
type CreateAwsReservationJSONBody = V1AWSReservationRequest

// CreateAzureReservationJSONBody defines parameters for CreateAzureReservation.
type CreateAzureReservationJSONBody = V1AzureReservationRequest

// GetSourceListParams defines parameters for GetSourceList.
type GetSourceListParams struct {
	Provider *GetSourceListParamsProvider `form:"provider,omitempty" json:"provider,omitempty"`
}

// GetSourceListParamsProvider defines parameters for GetSourceList.
type GetSourceListParamsProvider string

// GetInstanceTypeListParams defines parameters for GetInstanceTypeList.
type GetInstanceTypeListParams struct {
	// Hyperscaler region
	Region string `form:"region" json:"region"`
}

// GetLaunchTemplatesListParams defines parameters for GetLaunchTemplatesList.
type GetLaunchTemplatesListParams struct {
	// Hyperscaler region
	Region string `form:"region" json:"region"`
}

// AvailabilityStatusJSONRequestBody defines body for AvailabilityStatus for application/json ContentType.
type AvailabilityStatusJSONRequestBody = AvailabilityStatusJSONBody

// CreatePubkeyJSONRequestBody defines body for CreatePubkey for application/json ContentType.
type CreatePubkeyJSONRequestBody = CreatePubkeyJSONBody

// CreateAwsReservationJSONRequestBody defines body for CreateAwsReservation for application/json ContentType.
type CreateAwsReservationJSONRequestBody = CreateAwsReservationJSONBody

// CreateAzureReservationJSONRequestBody defines body for CreateAzureReservation for application/json ContentType.
type CreateAzureReservationJSONRequestBody = CreateAzureReservationJSONBody
