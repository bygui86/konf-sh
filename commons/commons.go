package commons

const (
	// Flags
	CustomKubeConfigFlagName        = "kube-config"
	CustomKubeConfigFlagShort       = "k"
	CustomKubeConfigFlagDescription = "Kubernetes configuration custom (`PATH`)"
	SingleConfigsFlagName           = "single-configs"
	SingleConfigsFlagShort          = "c"
	SingleConfigsFlagDescription    = "Single Kubernetes configurations files custom (`PATH`)"

	// Environment variables
	CustomKubeConfigPathEnvVar = "KONF_KUBE_CONFIG_PATH"
	SingleConfigsPathEnvVar    = "KONF_SINGLE_KUBE_CONFIGS_PATH"

	// Defaults
	CustomKubeConfigPathDefault = ".kube/config"
	SingleConfigsPathDefault    = ".kube/configs"
)
