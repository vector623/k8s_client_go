package k8s_client_go

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:19090"

// Client -
type Client struct {
	HostURL    string
	K8sClientSet *kubernetes.Clientset
}

func NewClientFromKubeFile(apihost, kubeconfig string) (*Client, error) {
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

func NewClientFromKubeCreds(apihost, clusterName, clientCertificateData, clientKeyData, token string) (*Client, error) {
	tmpKubeConf, _ := GenerateKubeConfTempFile(clusterName, clientCertificateData, clientKeyData, token)

	config, configErr := clientcmd.BuildConfigFromFlags(apihost, tmpKubeConf.Name())
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

func GenerateKubeConf(name, clientCertificateData, clientKeyData, token  string) (string, error) {
	aksCreds := AksCreds{
		Name: name,
		User: struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
			Token                 string
		}{
			ClientCertificateData: clientCertificateData,
			ClientKeyData: clientKeyData,
			Token: token,
		},
	}

	aksCredsYaml, _ := yaml.Marshal(aksCreds)
	aksCredsYamlStr := string(aksCredsYaml)

	return aksCredsYamlStr, nil
}

func GenerateKubeConfTempFile(name, clientCertificateData, clientKeyData, token string) (*os.File, error) {
	aksCreds := AksCreds{
		Name: name,
		User: struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
			Token                 string
		}{
			ClientCertificateData: clientCertificateData,
			ClientKeyData: clientKeyData,
			Token: token,
		},
	}

	aksCredsYaml, _ := yaml.Marshal(aksCreds)

	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		return nil, errors.New("Cannot create temporary file")
	}

	if _, err = tmpFile.Write(aksCredsYaml); err != nil {
		return nil, errors.New("Failed to write to temporary file")
	}

	return tmpFile, nil
}

type AksCreds struct {
	Name string
	User struct {
		ClientCertificateData string `yaml:"client-certificate-data"`
		ClientKeyData         string `yaml:"client-key-data"`
		Token                 string
	}
}