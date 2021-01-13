package deployment

import "github.com/wang1137095129/my-k8s-controller/type/metadata"

type SpecSelector struct {
	matchLabels      map[string]string
	matchExpressions SelectorMatchExpressions
}

type SelectorMatchExpressions struct {
	key string
	//just In, NotIn, Exists and DoesNotExist
	operator string
	values   []string
}

type SpecStrategy struct {
	_type         string
	rollingUpdate StrategyRollingUpdate
}

type StrategyRollingUpdate struct {
	maxSurge       string
	maxUnavailable string
}

type SpecTemplate struct {
	metadata metadata.DeploymentMetadata
	spec     TemplateSpec
}

type TemplateSpecAffinity struct {
}

type TemplateSpecContainer struct {

}

type TemplateSpecDnsConfig struct {
}

type TemplateSpecEphemeralContainer struct {
}

type TemplateSpecHostAlias struct {
}

type TemplateSpecImagePullSecret struct {
}

type TemplateSpecInitContainer struct {
}

type TemplateSpecReadinessGate struct {
}

type TemplateSpecSecurityContext struct {
}

type TemplateSpecToleration struct {
}

type TemplateSpecTopologySpreadConstraints struct {
}

type TemplateSpecVolume struct {
}

type TemplateSpec struct {
	activeDeadlineSeconds         int32
	affinity                      TemplateSpecAffinity
	automountServiceAccountToken  bool
	containers                    []TemplateSpecContainer
	dnsConfig                     TemplateSpecDnsConfig
	dnsPolicy                     string
	enableServiceLinks            bool
	ephemeralContainers           []TemplateSpecEphemeralContainer
	hostAliases                   []TemplateSpecHostAlias
	hostIPC                       bool
	hostNetwork                   bool
	hostPID                       bool
	hostname                      bool
	imagePullSecrets              []TemplateSpecImagePullSecret
	initContainers                []TemplateSpecInitContainer
	nodeName                      string
	nodeSelector                  map[string]string
	overhead                      map[string]string
	preemptionPolicy              map[string]string
	priority                      int32
	priorityClassName             string
	readinessGates                []TemplateSpecReadinessGate
	restartPolicy                 string
	runtimeClassName              string
	schedulerName                 string
	securityContext               TemplateSpecSecurityContext
	serviceAccount                string
	serviceAccountName            string
	setHostnameAsFQDN             bool
	shareProcessNamespace         bool
	subdomain                     string
	terminationGracePeriodSeconds int32
	tolerations                   []TemplateSpecToleration
	topologySpreadConstraints     []TemplateSpecTopologySpreadConstraints
	volumes                       []TemplateSpecVolume
}
