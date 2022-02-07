package commons

const (
	// Flags
	CustomKubeConfigFlagName        = "kube-config"
	CustomKubeConfigFlagShort       = "k"
	CustomKubeConfigFlagDescription = "Kubernetes configuration custom (`PATH`)"
	SingleConfigsFlagName           = "single-konfigs"
	SingleConfigsFlagShort          = "c"
	SingleConfigsFlagDescription    = "Single Kubernetes configurations files custom (`PATH`)"

	// Environment variables
	CustomKubeConfigPathEnvVar = "KONF_KUBE_CONFIG_PATH"
	SingleConfigsPathEnvVar    = "KONF_SINGLE_KUBE_CONFIGS_PATH"
	KubeConfigEnvVar           = "KUBECONFIG"

	// Defaults
	KubeConfigPathDefault    = ".kube/config"
	SingleConfigsPathDefault = ".kube/konfigs"
	KubeConfigEnvVarDefault  = ""
)
