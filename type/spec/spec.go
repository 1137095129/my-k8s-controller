package spec

import "github.com/wang1137095129/my-k8s-controller/type/spec/deployment"

type DeploymentSpec struct {
	minReadySeconds         int32
	paused                  bool
	progressDeadlineSeconds int32
	replicas                int32
	revisionHistoryLimit    int32
	selector                deployment.SpecSelector
	strategy                deployment.SpecStrategy
	template                deployment.SpecTemplate
}
