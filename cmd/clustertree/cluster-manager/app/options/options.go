package options

import (
	"time"

	"github.com/spf13/pflag"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/component-base/config/options"
	componentbaseconfigv1alpha1 "k8s.io/component-base/config/v1alpha1"

	"github.com/kosmos.io/kosmos/pkg/utils"
	"github.com/kosmos.io/kosmos/pkg/utils/flags"
	"github.com/kosmos.io/kosmos/pkg/utils/lifted"
)

const (
	LeaderElectionNamespace    = "kosmos-system"
	LeaderElectionResourceName = "cluster-manager"

	CoreDNSServiceNamespace = "kube-system"
	CoreDNSServiceName      = "kube-dns"
)

type Options struct {
	LeaderElection componentbaseconfig.LeaderElectionConfiguration
	utils.KubernetesOptions
	ListenPort           int32
	DaemonSetController  bool
	MultiClusterService  bool
	DirectClusterService bool

	// If MultiClusterService is disabled, the clustertree will rewrite the dnsPolicy configuration for pods deployed in
	// the leaf clusters, directing them to the root cluster's CoreDNS, thus facilitating access to services across all
	// clusters.
	RootCoreDNSServiceNamespace string
	RootCoreDNSServiceName      string

	// Enable oneway storage controllers
	OnewayStorageControllers bool

	// AutoCreateMCSPrefix are the prefix of the namespace for service to auto create in leaf cluster
	AutoCreateMCSPrefix []string

	// ReservedNamespaces are the protected namespaces to prevent Kosmos for deleting system resources
	ReservedNamespaces []string

	RateLimiterOpts lifted.RateLimitOptions

	BackoffOpts flags.BackoffOptions

	SyncPeriod time.Duration

	//Add notready status upload part for one2cluster in UpdateRootNodeStatus
	UpdateRootNodeStatusNotready bool
}

func NewOptions() (*Options, error) {
	var leaderElection componentbaseconfigv1alpha1.LeaderElectionConfiguration
	componentbaseconfigv1alpha1.RecommendedDefaultLeaderElectionConfiguration(&leaderElection)

	leaderElection.ResourceName = LeaderElectionResourceName
	leaderElection.ResourceNamespace = LeaderElectionNamespace
	leaderElection.ResourceLock = resourcelock.LeasesResourceLock

	var opts Options
	if err := componentbaseconfigv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&leaderElection, &opts.LeaderElection, nil); err != nil {
		return nil, err
	}

	return &opts, nil
}

func (o *Options) AddFlags(flags *pflag.FlagSet) {
	if o == nil {
		return
	}

	flags.Float32Var(&o.KubernetesOptions.QPS, "kube-qps", utils.DefaultTreeAndNetManagerKubeQPS, "QPS to use while talking with kube-apiserver.")
	flags.IntVar(&o.KubernetesOptions.Burst, "kube-burst", utils.DefaultTreeAndNetManagerKubeBurst, "Burst to use while talking with kube-apiserver.")
	flags.StringVar(&o.KubernetesOptions.KubeConfig, "kubeconfig", "", "Path for kubernetes kubeconfig file, if left blank, will use in cluster way.")
	flags.StringVar(&o.KubernetesOptions.MasterURL, "master", "", "Used to generate kubeconfig for downloading, if not specified, will use host in kubeconfig.")
	flags.Int32Var(&o.ListenPort, "listen-port", 10250, "Listen port for requests from the kube-apiserver.")
	flags.BoolVar(&o.DaemonSetController, "daemonset-controller", false, "Turn on or off daemonset controller.")
	flags.BoolVar(&o.MultiClusterService, "multi-cluster-service", false, "Turn on or off mcs support.")
	flags.BoolVar(&o.DirectClusterService, "direct-cluster-service", false, "Turn on or off direct cluster service.")
	flags.StringVar(&o.RootCoreDNSServiceNamespace, "root-coredns-service-namespace", CoreDNSServiceNamespace, "The namespace of the CoreDNS service in the root cluster, used to locate the CoreDNS service when MultiClusterService is disabled.")
	flags.StringVar(&o.RootCoreDNSServiceName, "root-coredns-service-name", CoreDNSServiceName, "The name of the CoreDNS service in the root cluster, used to locate the CoreDNS service when MultiClusterService is disabled.")
	flags.BoolVar(&o.OnewayStorageControllers, "oneway-storage-controllers", false, "Turn on or off oneway storage controllers.")
	flags.StringSliceVar(&o.AutoCreateMCSPrefix, "auto-mcs-prefix", []string{}, "The prefix of namespace for service to auto create mcs resources")
	flags.StringSliceVar(&o.ReservedNamespaces, "reserved-namespaces", []string{"kube-system"}, "The namespaces protected by Kosmos that the controller-manager will skip.")
	flags.DurationVar(&o.SyncPeriod, "sync-period", 0, "the sync period for informer to resync.")
	flags.BoolVar(&o.UpdateRootNodeStatusNotready, "Update-RootNode-Status-Notready", false, "Turn on or off add notready status upload part for one2cluster in UpdateRootNodeStatus")
	o.RateLimiterOpts.AddFlags(flags)
	o.BackoffOpts.AddFlags(flags)
	options.BindLeaderElectionFlags(&o.LeaderElection, flags)
}
