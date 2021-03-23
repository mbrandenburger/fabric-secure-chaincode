/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package k8s

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	appsV1 "k8s.io/api/apps/v1"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration
// https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go

var clientset *kubernetes.Clientset

const (
	Namespace          = "hyperledger"
	DeploymentName     = "experimenter-worker"
	DeploymentReplicas = 1
	ServiceName        = "worker-service"
	AppName            = "worker"
	ContainerName      = "web"
	ContainerImage     = "fpc/experimenter-worker"
	ContainerPort      = 5000
)

func init() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// create kubernetes client
	// creates the in-cluster config
	//config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func Launch() (string, error) {

	deploymentSpec := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: DeploymentName,
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: int32Ptr(DeploymentReplicas),
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": AppName,
				},
			},
			Template: apiV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"app": AppName,
					},
				},
				Spec: apiV1.PodSpec{
					Containers: []apiV1.Container{
						{
							Env: []apiV1.EnvVar{
								{
									Name:  "REDIS_HOST",
									Value: "redis-service",
								},
								{
									Name:  "REDIS_PORT",
									Value: "6379",
								},
							},
							Name:            ContainerName,
							Image:           ContainerImage,
							ImagePullPolicy: apiV1.PullIfNotPresent,
							Ports: []apiV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiV1.ProtocolTCP,
									ContainerPort: ContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	deployment, err := clientset.AppsV1().Deployments(Namespace).Create(context.TODO(), deploymentSpec, metaV1.CreateOptions{})
	if err != nil {
		return "", err
	}
	fmt.Printf("Created deployment %q.\n", deployment.GetObjectMeta().GetName())

	// --- create service -------------------------------------------------------------
	serviceSpec := &apiV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: ServiceName,
		},
		Spec: apiV1.ServiceSpec{
			Selector: map[string]string{
				"app": AppName,
			},
			Type: "LoadBalancer",
			Ports: []apiV1.ServicePort{
				{
					Port: ContainerPort,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: ContainerPort,
					},
				},
			},
		},
	}

	fmt.Println("Creating service...")
	service, err := clientset.CoreV1().Services(Namespace).Create(context.TODO(), serviceSpec, metaV1.CreateOptions{})
	if err != nil {
		// ignore server already exists
		if !strings.HasPrefix(err.Error(), fmt.Sprintf("services \"%s\" already exists", ServiceName)) {
			return "", err
		}
	}
	fmt.Printf("Created service %q.\n", service.GetObjectMeta().GetName())

	// let's wait until deployment is "ready"
	var reason string
	err = wait.PollImmediate(time.Second, time.Second*10, func() (bool, error) {
		deploymentCheck, err := clientset.AppsV1().Deployments(Namespace).Get(context.TODO(), DeploymentName, metaV1.GetOptions{})
		if err != nil {
			return false, err
		}

		complete := func(deployment *appsV1.Deployment, newStatus *appsV1.DeploymentStatus) bool {
			return newStatus.UpdatedReplicas == *(deployment.Spec.Replicas) &&
				newStatus.Replicas == *(deployment.Spec.Replicas) &&
				newStatus.AvailableReplicas == *(deployment.Spec.Replicas) &&
				newStatus.ObservedGeneration >= deployment.Generation
		}
		if complete(deployment, &deploymentCheck.Status) {
			return true, nil
		}
		reason = fmt.Sprintf("deployment status: %#v", deployment.Status)
		fmt.Println(reason)

		return false, nil
	})

	if err == wait.ErrWaitTimeout {
		err = fmt.Errorf("%s", reason)
	}
	if err != nil {
		return "", fmt.Errorf("error waiting for deployment %q status to match expectation: %v", deployment.Name, err)
	}

	return "done", nil
}

func Kill() (string, error) {
	deploymentsClient := clientset.AppsV1().Deployments(Namespace)
	deletePolicy := metaV1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), DeploymentName, metaV1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return "", err
	}
	return fmt.Sprintln("Deleted deployment."), nil
}

func int32Ptr(i int32) *int32 { return &i }
