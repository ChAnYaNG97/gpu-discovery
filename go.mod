module gpu-discovery

go 1.13

require (
	github.com/NVIDIA/gpu-monitoring-tools v0.0.0-20210305163645-dfb599a892d2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/code-generator v0.20.2
	k8s.io/klog v1.0.0
	sigs.k8s.io/controller-runtime v0.8.1
)

replace (
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
	k8s.io/kubernetes => k8s.io/kubernetes v0.20.2
)
