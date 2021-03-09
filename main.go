package main

import (
	"flag"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"gpu-discovery/apis/scheduling/v1beta1"
	"gpu-discovery/generated/gpus/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"log"
	"fmt"
	"os"
	"context"
)

var (
	KubeConfig string
	QPS        float32
	Burst      int
)

func main() {
	cfg, err := clientcmd.BuildConfigFromFlags("", KubeConfig)
	if err != nil {
		klog.Errorf("Failed to build config, %s", err)
	}
	gpuClient, err := versioned.NewForConfig(cfg)

	if err != nil {
		klog.Errorf("Failed to create gpu client, %s", err)
	}

	gpuManager := gpuClient.SchedulingV1beta1().GPUs("default")

	log.Println("Loading NVML")
	if err := nvml.Init(); err != nil {
		log.Printf("Failed to initialize NVML: %s.", err)
		log.Printf("If this is a GPU node, did you set the docker default runtime to `nvidia`?")
		log.Printf("You can check the prerequisites at: https://github.com/NVIDIA/k8s-device-plugin#prerequisites")
		log.Printf("You can learn how to set the runtime at: https://github.com/NVIDIA/k8s-device-plugin#quick-start")
		select {}
	}
	defer func() { log.Println("Shutdown of NVML returned:", nvml.Shutdown()) }()

	n, err := nvml.GetDeviceCount()
	if err != nil {
		klog.Errorf("Failed to get device count, %s", err)
	}
	hostname := os.Getenv("HOST_NAME")

	for index := uint(0); index < n; index++ {
		device, err := nvml.NewDevice(index)
		if err != nil {
			klog.Errorf("Failed to new device, %s", err)
		}

		gpu := v1beta1.GPU{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-gpu-%d", hostname, index),
			},
			Spec: v1beta1.GPUSpec{
				UUID:   device.UUID,
				Model:  *device.Model,
				Family: getArchFamily(*device.CudaComputeCapability.Major, *device.CudaComputeCapability.Minor),
				Capacity: v1beta1.R{
					Core:   "1.0",
					Memory: int(*device.Memory),
				},
				Node: hostname,
			},
			Status: v1beta1.GPUStatus{
				Allocated: v1beta1.R{
					Core:   "0.0",
					Memory: 0,
				},
				PodMap: nil,
			},
		}
		_, err = gpuManager.Create(context.Background(), &gpu, metav1.CreateOptions{})
		if err != nil {
			klog.Errorf("Failed to create gpu %s, %s", gpu.Name, err)
		}
	}
}

func init() {
	flag.StringVar(&KubeConfig, "kubeconfig", "", "Use it out-of-cluster.")
}

func getArchFamily(computeMajor, computeMinor int) string {
	switch computeMajor {
	case 1:
		return "Tesla"
	case 2:
		return "Fermi"
	case 3:
		return "Kepler"
	case 5:
		return "Maxwell"
	case 6:
		return "Pascal"
	case 7:
		if computeMinor < 5 {
			return "volta"
		}
		return "Turing"
	case 8:
		return "Ampere"
	}
	return "Unknown"
}
