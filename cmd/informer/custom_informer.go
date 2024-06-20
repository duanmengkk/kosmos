package main

import (
	"flag"
	"fmt"
	"github.com/kosmos.io/kosmos/pkg/kubeclientbridge"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/informers"
	"k8s.io/klog"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	unifyClient, err := kubeclientbridge.NewUnifyClientset(config)

	//kubeClient, err := kubernetes.NewForConfig(config)

	kubeFactory := informers.NewSharedInformerFactory(unifyClient, 0)
	podInformer := kubeFactory.Core().V1().Pods()
	podLister := podInformer.Lister()
	podHassynced := podInformer.Informer().HasSynced
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if pod, ok := obj.(*corev1.Pod); ok {
				klog.Infof("New Pod Added: %s in namespace %s", pod.GetName(), pod.GetNamespace())
			} else if u, ok := obj.(*unstructured.Unstructured); ok {
				klog.Infof("New Pod Added: %s in namespace %s", u.GetName(), u.GetNamespace())
			} else {
				klog.Info("New Pod Added: unknown type")
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			if pod, ok := newObj.(*corev1.Pod); ok {
				klog.Infof("New Pod Updated: %s in namespace %s", pod.GetName(), pod.GetNamespace())
			} else if u, ok := newObj.(*unstructured.Unstructured); ok {
				klog.Infof("New Pod Updated: %s in namespace %s", u.GetName(), u.GetNamespace())
			} else {
				klog.Info("New Pod Added: unknown type")
			}
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Pod Deleted:")
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go kubeFactory.Start(stopCh)
	if !cache.WaitForNamedCacheSync("controller", stopCh, podHassynced) {
		klog.Errorf("failed to wait for  caches to sync")
		return
	}

	_, _ = podLister.Pods("default").Get("pod-example-1127")
	if err != nil {
		fmt.Printf("Error getting Pods: %v\n", err)
	}

	<-stopCh

}
