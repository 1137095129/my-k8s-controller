package metadata

import "github.com/wang1137095129/my-k8s-controller/type/metadata/deployment"

type DeploymentMetadata struct {
	Annotations                map[string]string
	ClusterName                string
	CreationTimestamp          string
	DeletionGracePeriodSeconds int32
	DeletionTimestamp          string
	Finalizers                 []string
	GenerateName               string
	Generation                 int32
	Labels                     map[string]string
	ManagedFields              []deployment.MetadataManagedField
	Name                       string
	Namespace                  string
	OwnerReferences            []deployment.MetadataOwnerReference
	ResourceVersion            string
	SelfLink                   string
	Uid                        string
}
