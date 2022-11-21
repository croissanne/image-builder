// Package composer provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package composer

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// Defines values for ComposeStatusValue.
const (
	ComposeStatusValueFailure ComposeStatusValue = "failure"
	ComposeStatusValuePending ComposeStatusValue = "pending"
	ComposeStatusValueSuccess ComposeStatusValue = "success"
)

// Defines values for ImageStatusValue.
const (
	ImageStatusValueBuilding    ImageStatusValue = "building"
	ImageStatusValueFailure     ImageStatusValue = "failure"
	ImageStatusValuePending     ImageStatusValue = "pending"
	ImageStatusValueRegistering ImageStatusValue = "registering"
	ImageStatusValueSuccess     ImageStatusValue = "success"
	ImageStatusValueUploading   ImageStatusValue = "uploading"
)

// Defines values for ImageTypes.
const (
	ImageTypesAws            ImageTypes = "aws"
	ImageTypesAwsHaRhui      ImageTypes = "aws-ha-rhui"
	ImageTypesAwsRhui        ImageTypes = "aws-rhui"
	ImageTypesAwsSapRhui     ImageTypes = "aws-sap-rhui"
	ImageTypesAzure          ImageTypes = "azure"
	ImageTypesAzureRhui      ImageTypes = "azure-rhui"
	ImageTypesAzureSapRhui   ImageTypes = "azure-sap-rhui"
	ImageTypesEdgeCommit     ImageTypes = "edge-commit"
	ImageTypesEdgeContainer  ImageTypes = "edge-container"
	ImageTypesEdgeInstaller  ImageTypes = "edge-installer"
	ImageTypesGcp            ImageTypes = "gcp"
	ImageTypesGcpRhui        ImageTypes = "gcp-rhui"
	ImageTypesGuestImage     ImageTypes = "guest-image"
	ImageTypesImageInstaller ImageTypes = "image-installer"
	ImageTypesIotCommit      ImageTypes = "iot-commit"
	ImageTypesIotContainer   ImageTypes = "iot-container"
	ImageTypesIotInstaller   ImageTypes = "iot-installer"
	ImageTypesIotRawImage    ImageTypes = "iot-raw-image"
	ImageTypesVsphere        ImageTypes = "vsphere"
)

// Defines values for UploadStatusValue.
const (
	Failure UploadStatusValue = "failure"
	Pending UploadStatusValue = "pending"
	Running UploadStatusValue = "running"
	Success UploadStatusValue = "success"
)

// Defines values for UploadTypes.
const (
	UploadTypesAws       UploadTypes = "aws"
	UploadTypesAwsS3     UploadTypes = "aws.s3"
	UploadTypesAzure     UploadTypes = "azure"
	UploadTypesContainer UploadTypes = "container"
	UploadTypesGcp       UploadTypes = "gcp"
)

// AWSEC2CloneCompose defines model for AWSEC2CloneCompose.
type AWSEC2CloneCompose struct {
	Region            string    `json:"region"`
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
}

// AWSEC2UploadOptions defines model for AWSEC2UploadOptions.
type AWSEC2UploadOptions struct {
	Region            string   `json:"region"`
	ShareWithAccounts []string `json:"share_with_accounts"`
	SnapshotName      *string  `json:"snapshot_name,omitempty"`
}

// AWSEC2UploadStatus defines model for AWSEC2UploadStatus.
type AWSEC2UploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// AWSS3UploadOptions defines model for AWSS3UploadOptions.
type AWSS3UploadOptions struct {
	// If set to false (the default value), a long, obfuscated URL
	// is returned. Its expiration might be sooner than for other upload
	// targets.
	//
	// If set to true, a shorter URL is returned and
	// its expiration is the same as for the other upload targets.
	Public *bool  `json:"public,omitempty"`
	Region string `json:"region"`
}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AzureUploadOptions defines model for AzureUploadOptions.
type AzureUploadOptions struct {
	// Name of the uploaded image. It must be unique in the given resource group.
	// If name is omitted from the request, a random one based on a UUID is
	// generated.
	ImageName *string `json:"image_name,omitempty"`

	// Location of the provided resource_group, where the image should be uploaded and registered.
	// How to list all locations:
	// https://docs.microsoft.com/en-us/cli/azure/account?view=azure-cli-latest#az_account_list_locations'
	// If the location is not specified, it is deducted from the provided resource_group.
	Location *string `json:"location,omitempty"`

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded.
	// How to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// CloneComposeBody defines model for CloneComposeBody.
type CloneComposeBody interface{}

// CloneComposeResponse defines model for CloneComposeResponse.
type CloneComposeResponse struct {
	Href string             `json:"href"`
	Id   openapi_types.UUID `json:"id"`
	Kind string             `json:"kind"`
}

// CloneStatus defines model for CloneStatus.
type CloneStatus struct {
	Href    string            `json:"href"`
	Id      string            `json:"id"`
	Kind    string            `json:"kind"`
	Options interface{}       `json:"options"`
	Status  UploadStatusValue `json:"status"`
	Type    UploadTypes       `json:"type"`
}

// ComposeId defines model for ComposeId.
type ComposeId struct {
	Href string             `json:"href"`
	Id   openapi_types.UUID `json:"id"`
	Kind string             `json:"kind"`
}

// ComposeLogs defines model for ComposeLogs.
type ComposeLogs struct {
	Href        string        `json:"href"`
	Id          string        `json:"id"`
	ImageBuilds []interface{} `json:"image_builds"`
	Kind        string        `json:"kind"`
	Koji        *KojiLogs     `json:"koji,omitempty"`
}

// ComposeManifests defines model for ComposeManifests.
type ComposeManifests struct {
	Href      string        `json:"href"`
	Id        string        `json:"id"`
	Kind      string        `json:"kind"`
	Manifests []interface{} `json:"manifests"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Kind string `json:"kind"`

	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequest   *ImageRequest   `json:"image_request,omitempty"`
	ImageRequests  *[]ImageRequest `json:"image_requests,omitempty"`
	Koji           *Koji           `json:"koji,omitempty"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	Href          string             `json:"href"`
	Id            string             `json:"id"`
	ImageStatus   ImageStatus        `json:"image_status"`
	ImageStatuses *[]ImageStatus     `json:"image_statuses,omitempty"`
	Kind          string             `json:"kind"`
	KojiStatus    *KojiStatus        `json:"koji_status,omitempty"`
	Status        ComposeStatusValue `json:"status"`
}

// ComposeStatusError defines model for ComposeStatusError.
type ComposeStatusError struct {
	Details *interface{} `json:"details,omitempty"`
	Id      int          `json:"id"`
	Reason  string       `json:"reason"`
}

// ComposeStatusValue defines model for ComposeStatusValue.
type ComposeStatusValue string

// Container defines model for Container.
type Container struct {
	// Name to use for the container from the image
	Name   *string `json:"name,omitempty"`
	Source string  `json:"source"`

	// Control TLS verifification
	TlsVerify *bool `json:"tls_verify,omitempty"`
}

// ContainerUploadOptions defines model for ContainerUploadOptions.
type ContainerUploadOptions struct {
	// Name for the created container image
	Name *string `json:"name,omitempty"`

	// Tag for the created container image
	Tag *string `json:"tag,omitempty"`
}

// ContainerUploadStatus defines model for ContainerUploadStatus.
type ContainerUploadStatus struct {
	// Digest of the manifest of the uploaded container on the registry
	Digest string `json:"digest"`

	// FQDN of the uploaded image
	Url string `json:"url"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Containers *[]Container  `json:"containers,omitempty"`
	Filesystem *[]Filesystem `json:"filesystem,omitempty"`
	Packages   *[]string     `json:"packages,omitempty"`

	// Extra repositories for packages specified in customizations. These
	// repositories will only be used to depsolve and retrieve packages
	// for the OS itself (they will not be available for the build root or
	// any other part of the build process). The package_sets field for these
	// repositories is ignored.
	PayloadRepositories *[]Repository `json:"payload_repositories,omitempty"`
	Services            *struct {
		// List of services to disable by default
		Disabled *[]string `json:"disabled,omitempty"`

		// List of services to enable by default
		Enabled *[]string `json:"enabled,omitempty"`
	} `json:"services,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	Users        *[]User       `json:"users,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code        string       `json:"code"`
	Details     *interface{} `json:"details,omitempty"`
	Href        string       `json:"href"`
	Id          string       `json:"id"`
	Kind        string       `json:"kind"`
	OperationId string       `json:"operation_id"`
	Reason      string       `json:"reason"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	Items []Error `json:"items"`
	Kind  string  `json:"kind"`
	Page  int     `json:"page"`
	Size  int     `json:"size"`
	Total int     `json:"total"`
}

// Filesystem defines model for Filesystem.
type Filesystem struct {
	MinSize    uint64 `json:"min_size"`
	Mountpoint string `json:"mountpoint"`
}

// GCPUploadOptions defines model for GCPUploadOptions.
type GCPUploadOptions struct {
	// Name of an existing STANDARD Storage class Bucket.
	Bucket *string `json:"bucket,omitempty"`

	// The name to use for the imported and shared Compute Engine image.
	// The image name must be unique within the GCP project, which is used
	// for the OS image upload and import. If not specified a random
	// 'composer-api-<uuid>' string is used as the image name.
	ImageName *string `json:"image_name,omitempty"`

	// The GCP region where the OS image will be imported to and shared from.
	// The value must be a valid GCP location. See https://cloud.google.com/storage/docs/locations.
	// If not specified, the multi-region location closest to the source
	// (source Storage Bucket location) is chosen automatically.
	Region string `json:"region"`

	// List of valid Google accounts to share the imported Compute Engine image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	// If not specified, the imported Compute Engine image is not shared with any
	// account.
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture string       `json:"architecture"`
	ImageType    ImageTypes   `json:"image_type"`
	Ostree       *OSTree      `json:"ostree,omitempty"`
	Repositories []Repository `json:"repositories"`

	// This should really be oneOf but AWSS3UploadOptions is a subset of
	// AWSEC2UploadOptions. This means that all AWSEC2UploadOptions objects
	// are also valid AWSS3UploadOptionas objects which violates the oneOf
	// rules. Therefore, we have to use anyOf here but be aware that it isn't
	// possible to mix and match more schemas together.
	UploadOptions *UploadOptions `json:"upload_options,omitempty"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Error        *ComposeStatusError `json:"error,omitempty"`
	Status       ImageStatusValue    `json:"status"`
	UploadStatus *UploadStatus       `json:"upload_status,omitempty"`
}

// ImageStatusValue defines model for ImageStatusValue.
type ImageStatusValue string

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// Koji defines model for Koji.
type Koji struct {
	Name    string `json:"name"`
	Release string `json:"release"`
	Server  string `json:"server"`
	TaskId  int    `json:"task_id"`
	Version string `json:"version"`
}

// KojiLogs defines model for KojiLogs.
type KojiLogs struct {
	Import interface{} `json:"import"`
	Init   interface{} `json:"init"`
}

// KojiStatus defines model for KojiStatus.
type KojiStatus struct {
	BuildId *int `json:"build_id,omitempty"`
}

// List defines model for List.
type List struct {
	Kind  string `json:"kind"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
}

// OSTree defines model for OSTree.
type OSTree struct {
	// A URL which, if set, is used for fetching content. Implies that `url` is set as well,
	// which will be used for metadata only.
	Contenturl *string `json:"contenturl,omitempty"`

	// Can be either a commit (example: 02604b2da6e954bd34b8b82a835e5a77d2b60ffa), or a branch-like reference (example: rhel/8/x86_64/edge)
	Parent *string `json:"parent,omitempty"`
	Ref    *string `json:"ref,omitempty"`

	// Determines whether a valid subscription manager (candlepin) identity is required to
	// access this repository. Consumer certificates will be used as client certificates when
	// fetching metadata and content.
	Rhsm *bool   `json:"rhsm,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// ObjectReference defines model for ObjectReference.
type ObjectReference struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Kind string `json:"kind"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl  *string `json:"baseurl,omitempty"`
	CheckGpg *bool   `json:"check_gpg,omitempty"`

	// GPG key used to sign packages in this repository.
	Gpgkey     *string `json:"gpgkey,omitempty"`
	IgnoreSsl  *bool   `json:"ignore_ssl,omitempty"`
	Metalink   *string `json:"metalink,omitempty"`
	Mirrorlist *string `json:"mirrorlist,omitempty"`

	// Naming package sets for a repository assigns it to a specific part
	// (pipeline) of the build process.
	PackageSets *[]string `json:"package_sets,omitempty"`

	// Determines whether a valid subscription is required to access this repository.
	Rhsm *bool `json:"rhsm,omitempty"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation_key"`
	BaseUrl       string `json:"base_url"`
	Insights      bool   `json:"insights"`
	Organization  string `json:"organization"`
	ServerUrl     string `json:"server_url"`
}

// This should really be oneOf but AWSS3UploadOptions is a subset of
// AWSEC2UploadOptions. This means that all AWSEC2UploadOptions objects
// are also valid AWSS3UploadOptionas objects which violates the oneOf
// rules. Therefore, we have to use anyOf here but be aware that it isn't
// possible to mix and match more schemas together.
type UploadOptions interface{}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{}       `json:"options"`
	Status  UploadStatusValue `json:"status"`
	Type    UploadTypes       `json:"type"`
}

// UploadStatusValue defines model for UploadStatusValue.
type UploadStatusValue string

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// User defines model for User.
type User struct {
	Groups *[]string `json:"groups,omitempty"`
	Key    *string   `json:"key,omitempty"`
	Name   string    `json:"name"`
}

// Page defines model for page.
type Page = string

// Size defines model for size.
type Size = string

// PostComposeJSONBody defines parameters for PostCompose.
type PostComposeJSONBody = ComposeRequest

// PostCloneComposeJSONBody defines parameters for PostCloneCompose.
type PostCloneComposeJSONBody = CloneComposeBody

// GetErrorListParams defines parameters for GetErrorList.
type GetErrorListParams struct {
	// Page index
	Page *Page `form:"page,omitempty" json:"page,omitempty"`

	// Number of items in each page
	Size *Size `form:"size,omitempty" json:"size,omitempty"`
}

// PostComposeJSONRequestBody defines body for PostCompose for application/json ContentType.
type PostComposeJSONRequestBody = PostComposeJSONBody

// PostCloneComposeJSONRequestBody defines body for PostCloneCompose for application/json ContentType.
type PostCloneComposeJSONRequestBody = PostCloneComposeJSONBody
