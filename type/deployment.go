package _type

import (
	"github.com/wang1137095129/my-k8s-controller/type/metadata"
	"github.com/wang1137095129/my-k8s-controller/type/spec"
)

type Deployment struct {
	Kind     string
	Metadata metadata.DeploymentMetadata
	Spec     spec.DeploymentSpec
}
