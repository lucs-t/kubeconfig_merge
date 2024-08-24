package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lucs-t/kubeconfig_merge/models"
)

func main(){
	rootPath,err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}
	kube := rootPath + "/.kube"
	if _, err := os.Stat(kube); os.IsNotExist(err) {
        log.Println("Directory does not exist:", kube)
		return
    } else if err != nil {
        log.Println("Error checking directory:", err)
		return
    }
	kubeconfigDir := kube + "/kubeconfig"
	if _, err := os.Stat(kubeconfigDir); os.IsNotExist(err) {
		log.Println("Directory does not exist:", kubeconfigDir)
		return
	} else if err != nil {
		log.Println("Error checking directory:", err)
		return
	}
	files, err := os.ReadDir(kubeconfigDir)
	if err !=nil{
		log.Println("Error reading directory:", err)
		return
	}
	if len(files) == 0 {
		log.Println("No kubeconfig files found in:", kubeconfigDir)
		return
	}
	var configs = models.Kubeconfig{
		ApiVersion: "v1",
		Kind: "Config",
		Clusters: make([]models.Cluster, 0),
		Users: make([]models.User, 0),
		Contexts: make([]models.Context, 0),
	}
	for _, file := range files {
		var (
			name string
			config models.Kubeconfig
		)
		if file.IsDir(){
			continue
		}
		if !strings.HasSuffix(file.Name(),".conf"){
			continue
		}else{
			name = strings.TrimSuffix(file.Name(),".conf")
		}
		err := config.Load(kubeconfigDir + "/" + file.Name())
		if err != nil {
			return
		}
		for i,context := range config.Contexts {
			var (
				kcontext  models.Context
				kcluster  models.Cluster
				kuser     models.User
			)
			for _,cluster := range config.Clusters {
				if cluster.Name == context.Context.Cluster {
					kcluster.Name = fmt.Sprintf("%s-%d",name,i)
					kcluster.Cluster = cluster.Cluster
					break
				}
			}
			for _,user := range config.Users {
				if user.Name == context.Context.User {
					kuser.Name = fmt.Sprintf("%s-%d",name,i)
					kuser.User = user.User
					break
				}
			}
			if kcluster.Name == "" || kuser.Name == "" {
				log.Println("Cluster or User not found"+file.Name()+":"+context.Name)
				return
			}
			kcontext.Name = fmt.Sprintf("%s-%d",name,i)
			kcontext.Context = models.ContextInfo{
				Cluster: kcluster.Name,
				User: kuser.Name,
			}
			configs.Clusters = append(configs.Clusters,kcluster)
			configs.Users = append(configs.Users,kuser)
			configs.Contexts = append(configs.Contexts,kcontext)
		}
	}
	err = configs.Save(kube + "/config")
	if err != nil {
		return
	}
	for _,clustername := range configs.Contexts{
		log.Println(clustername.Name+": cluster add success")
	}
}