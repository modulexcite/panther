package aws

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/go-openapi/strfmt"
)

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

const (
	EcsClusterSchema = "AWS.ECS.Cluster"
)

// EcsCluster contains all the information about an ECS Cluster
type EcsCluster struct {
	// Generic resource fields
	GenericAWSResource
	GenericResource

	// Fields embedded from ecs.Cluster
	ActiveServicesCount               *int64
	Attachments                       []*ecs.Attachment
	AttachmentsStatus                 *string
	CapacityProviders                 []*string
	DefaultCapacityProviderStrategy   []*ecs.CapacityProviderStrategyItem
	PendingTasksCount                 *int64
	RegisteredContainerInstancesCount *int64
	RunningTasksCount                 *int64
	Settings                          []*ecs.ClusterSetting
	Statistics                        []*ecs.KeyValuePair
	Status                            *string

	// Additional fields
	Services []*EcsService
	Tasks    []*EcsTask
}

// EcsService contains all the information about an ECS Service, for embedding into the EcsCluster resource
type EcsService struct {
	// Generic resource fields
	//
	// This is not a full resource, but it does have an ARN, Tags, and a name.
	GenericAWSResource

	// Fields embedded from ecs.Service
	CapacityProviderStrategy []*ecs.CapacityProviderStrategyItem
	// Normalized name for CreatedAt
	TimeCreated                   *strfmt.DateTime
	CreatedBy                     *string
	DeploymentConfiguration       *ecs.DeploymentConfiguration
	DeploymentController          *ecs.DeploymentController
	Deployments                   []*ecs.Deployment
	DesiredCount                  *int64
	EnableECSManagedTags          *bool
	Events                        []*ecs.ServiceEvent
	HealthCheckGracePeriodSeconds *int64
	LaunchType                    *string
	LoadBalancers                 []*ecs.LoadBalancer
	NetworkConfiguration          *ecs.NetworkConfiguration
	PendingCount                  *int64
	PlacementConstraints          []*ecs.PlacementConstraint
	PlacementStrategy             []*ecs.PlacementStrategy
	PlatformVersion               *string
	PropagateTags                 *string
	RoleArn                       *string
	RunningCount                  *int64
	SchedulingStrategy            *string
	ServiceRegistries             []*ecs.ServiceRegistry
	Status                        *string
	TaskDefinition                *string
	TaskSets                      []*ecs.TaskSet
}

// EcsTask contains all the information about an ECS Task, for embedding into the EcsCluster resource
type EcsTask struct {
	// Generic resource fields
	//
	// This is not a full resource, but it does have an ARN and Tags.
	GenericAWSResource

	// Fields embedded from ecs.Task
	Attachments          []*ecs.Attachment
	Attributes           []*ecs.Attribute
	AvailabilityZone     *string
	CapacityProviderName *string
	Connectivity         *string
	ConnectivityAt       *strfmt.DateTime
	ContainerInstanceArn *string
	Containers           []*ecs.Container
	Cpu                  *string
	// Normalized name for CreatedAt
	TimeCreated           *strfmt.DateTime
	DesiredStatus         *string
	ExecutionStoppedAt    *strfmt.DateTime
	Group                 *string
	HealthStatus          *string
	InferenceAccelerators []*ecs.InferenceAccelerator
	LastStatus            *string
	LaunchType            *string
	Memory                *string
	Overrides             *ecs.TaskOverride
	PlatformVersion       *string
	PullStartedAt         *strfmt.DateTime
	PullStoppedAt         *strfmt.DateTime
	StartedAt             *strfmt.DateTime
	StartedBy             *string
	StopCode              *string
	StoppedAt             *strfmt.DateTime
	StoppedReason         *string
	StoppingAt            *strfmt.DateTime
	TaskDefinitionArn     *string
	Version               *int64
}
