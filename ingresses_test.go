package k8s_client_go

import (
	"fmt"
	"testing"
)

func TestClient_ListIngresses(t *testing.T) {
	namespace := "supply-ferguson-poc"
	clientset, clientsetError := NewClient("https://feidsupplyaks001-2bd25411.hcp.eastus2.azmk8s.io:443", "/Users/d.gallmeier/.kube/config")
	ingresses, ingressesError := clientset.ListIngresses(namespace)

	fmt.Printf("There are %d ingresses in the namespace: %s\n", len(ingresses), namespace)

	if(clientsetError != nil) {
		t.Errorf("error creating client: %s", clientsetError)
	}
	if(ingressesError != nil) {
		t.Errorf("error getting ingresses: %s", ingressesError)
	}
}