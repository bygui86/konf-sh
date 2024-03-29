package commons

const (
	// Flags
	KubeConfigFlagName           = "kube-config"
	KubeConfigFlagShort          = "c"
	KubeConfigFlagDescription    = "Kubernetes configuration custom (`PATH`)"
	SingleKonfigsFlagName        = "single-konfigs"
	SingleKonfigsFlagShort       = "k"
	SingleKonfigsFlagDescription = "Single Kubernetes konfigurations custom (`PATH`)"

	// Environment variables
	KubeConfigPathEnvVar     = "KONF_KUBE_CONFIG_PATH"
	KubeConfigPathDefault    = ".kube/config"
	SingleKonfigsPathEnvVar  = "KONF_SINGLE_KUBE_KONFIGS_PATH"
	SingleKonfigsPathDefault = ".kube/konfigs"
	KubeConfigEnvVar         = "KUBECONFIG"
	KubeConfigEnvVarDefault  = ""
)
