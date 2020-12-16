package k8s_client_go

import (
	"context"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListIngresses(namespace string) ([]v1beta1.Ingress, error) {
	ingresses, err := c.K8sClientSet.
		ExtensionsV1beta1().
		Ingresses(namespace).
		List(
			context.TODO(),
			metav1.ListOptions{})

	filteredIngresses := make([]v1beta1.Ingress, 0)
	for _, ingress := range ingresses.Items {
		if(ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName == "service-nest-dev") {
			filteredIngresses = append(filteredIngresses, ingress)
		}
	}

	if(err != nil) {
		return nil, err
	}

	//return ingresses.Items, nil
	return filteredIngresses, nil
}
