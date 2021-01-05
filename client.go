package k8s_client_go

import (
	"errors"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	"k8s.io/client-go/tools/clientcmd"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:19090"

// Client -
type Client struct {
	HostURL    string
	K8sClientSet *kubernetes.Clientset
}

func NewClient(apihost, kubeconfig string) (*Client, error) {
	config, configErr := clientcmd.BuildConfigFromFlags(apihost, kubeconfig)
	if(configErr != nil) {
		return nil, configErr
	}
	clientset, clientsetErr := kubernetes.NewForConfig(config)
	if(clientsetErr != nil) {
		return nil, clientsetErr
	}

	newclient := Client{
		HostURL: apihost,
		K8sClientSet: clientset,
	}

	return &newclient, nil
}

func GenerateKubeConf(serverUrl, clientCertificateData, clientKeyData, token  string) (AksCreds, error) {
	aksCreds := AksCreds{
		name: "",
		user: struct {
			client_certificate_data string
			client_key_data         string
			token                   string
		}{},
	}

	return aksCreds, errors.New("tba")
}

type AksCreds struct {
	name string
	user struct {
		client_certificate_data string
		client_key_data string
		token string
	}
}