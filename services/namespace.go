package services

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"Ayile/models" // 导入你的models包
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SvcNamespaceData struct {
	Name              string
	Labels            map[string]string
	CreationTimestamp metav1.Time
}

type ServiceData struct {
	Name                  string
	Status                string
	Port                  int32
	TargetPort            int32
	NodePort              int32
	Namespace             string
	Type                  v1.ServiceType
	SessionAffinity       string
	InternalTrafficPolicy string
	Selector              map[string]string
}

type ServiceViewData struct {
	SelectedNamespace string
	Namespaces        []SvcNamespaceData
	Services          []ServiceData
}

func QueryServices(namespace string) (ServiceViewData, error) {
	client, err := models.NewKubernetesClient()
	if err != nil {
		return ServiceViewData{}, err
	}

	namespaces, err := client.ListNamespaces()
	if err != nil {
		return ServiceViewData{}, err
	}

	var svcnamespaceData []SvcNamespaceData
	var services []ServiceData

	for _, ns := range namespaces {
		data := SvcNamespaceData{
			Name:              ns.Name,
			Labels:            ns.Labels,
			CreationTimestamp: ns.CreationTimestamp,
		}
		svcnamespaceData = append(svcnamespaceData, data)

		if namespace == "All" || ns.Name == namespace {
			clientset := client.GetClientset()
			svcList, err := clientset.CoreV1().Services(ns.Name).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				return ServiceViewData{}, err
			}

			for _, svc := range svcList.Items {
				var port int32
				var targetPort int32
				var nodePort int32
				for _, ports := range svc.Spec.Ports {
					port = ports.Port
					if ports.TargetPort.Type == intstr.Int {
						targetPort = int32(ports.TargetPort.IntValue())
					} else if ports.TargetPort.Type == intstr.String {
						return ServiceViewData{}, fmt.Errorf("TargetPort is a string")
					}
					nodePort = ports.NodePort
				}

				serviceData := ServiceData{
					Name:                  svc.Name,
					Namespace:             svc.Namespace,
					Status:                "",
					Port:                  port,
					TargetPort:            targetPort,
					NodePort:              nodePort,
					Type:                  svc.Spec.Type,
					SessionAffinity:       string(svc.Spec.SessionAffinity),
					InternalTrafficPolicy: string(*svc.Spec.InternalTrafficPolicy),
					Selector:              svc.Spec.Selector,
				}
				services = append(services, serviceData)
			}
		}
	}

	serviceViewData := ServiceViewData{
		SelectedNamespace: namespace,
		Namespaces:        svcnamespaceData,
		Services:          services,
	}

	return serviceViewData, nil
}
