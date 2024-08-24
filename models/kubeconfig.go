package models

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)


type Kubeconfig struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Clusters []Cluster `yaml:"clusters"`
	Users []User `yaml:"users"`
	Contexts []Context `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
}

type Cluster struct {
	Name string `yaml:"name"`
	Cluster ClusterInfo `yaml:"cluster"`
}
type ClusterInfo struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
	Server string `yaml:"server"`
}

type User struct {
	Name string `yaml:"name"`
	User UserInfo `yaml:"user"`
}
type UserInfo struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData string `yaml:"client-key-data"`
}

type Context struct {
	Name string `yaml:"name"`
	Context ContextInfo `yaml:"context"`
}
type ContextInfo struct {
	Cluster string `yaml:"cluster"`
	User string `yaml:"user"`
	Namespace string `yaml:"namespace"`
}

func (c *Kubeconfig) Load(filepath string)error{
	reader,err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	return yaml.Unmarshal(reader,c)
}

func (c *Kubeconfig) Save(filepath string)error{
	if _,err := os.Stat(filepath);err == nil {
		err := os.Rename(filepath,filepath+".bak")
		if err != nil {
			log.Println(err)
			return err
		}
	}
	writer,err := yaml.Marshal(c)
	if err != nil {
		log.Println(err)
		return err
	}
	return os.WriteFile(filepath,writer,0700)
}