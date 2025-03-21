package cache

import (
	"context"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/endpoints/handlers"
	"k8s.io/apiserver/pkg/registry/rest"
	k8srest "k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"

	"github.com/kosmos.io/kosmos/pkg/clusterlink/proxy/delegate"
	"github.com/kosmos.io/kosmos/pkg/clusterlink/proxy/store"
)

const (
	order = 1000
)

type Cache struct {
	store             store.Store
	restMapper        meta.RESTMapper
	minRequestTimeout time.Duration
}

var _ delegate.Delegate = (*Cache)(nil)

func New(dep delegate.Dependency) delegate.Delegate {
	return &Cache{
		store:             dep.Store,
		restMapper:        dep.RestMapper,
		minRequestTimeout: dep.MinRequestTimeout,
	}
}

// Connect implements Proxy.
func (c *Cache) Connect(_ context.Context, request delegate.ProxyRequest) (http.Handler, error) {
	requestInfo := request.RequestInfo
	r := &rester{
		store:          c.store,
		gvr:            request.GroupVersionResource,
		tableConvertor: k8srest.NewDefaultTableConvertor(request.GroupVersionResource.GroupResource()),
	}

	gvk, err := c.restMapper.KindFor(request.GroupVersionResource)
	if err != nil {
		return nil, err
	}
	mapping, err := c.restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}

	scope := &handlers.RequestScope{
		Kind:     gvk,
		Resource: request.GroupVersionResource,
		Namer: &handlers.ContextBasedNaming{
			Namer:         meta.NewAccessor(),
			ClusterScoped: mapping.Scope.Name() == meta.RESTScopeNameRoot,
		},
		Serializer:       scheme.Codecs.WithoutConversion(),
		Convertor:        runtime.NewScheme(),
		Subresource:      requestInfo.Subresource,
		MetaGroupVersion: metav1.SchemeGroupVersion,
		TableConvertor:   r.tableConvertor,
	}
	klog.Infof("get gvk: %v from cache", gvk)

	var h http.Handler
	if requestInfo.Verb == "watch" || requestInfo.Name == "" {
		// for list or watch
		h = handlers.ListResource(r, r, scope, false, c.minRequestTimeout)
	} else {
		h = handlers.GetResource(r, scope)
	}
	return h, nil
}

func (c *Cache) Order() int {
	return order
}

// SupportRequest implements Plugin
func (c *Cache) SupportRequest(request delegate.ProxyRequest) bool {
	requestInfo := request.RequestInfo

	return requestInfo.IsResourceRequest &&
		c.store.HasResource(request.GroupVersionResource) &&
		requestInfo.Subresource == "" &&
		(requestInfo.Verb == "get" ||
			requestInfo.Verb == "list" ||
			requestInfo.Verb == "watch")
}

type rester struct {
	store          store.Store
	gvr            schema.GroupVersionResource
	tableConvertor rest.TableConvertor
}

var _ rest.Getter = &rester{}
var _ rest.Lister = &rester{}
var _ rest.Watcher = &rester{}

// Get implements rest.Getter interface
func (r *rester) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, r.gvr, name, options)
}

// Watch implements rest.Watcher interface
func (r *rester) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	return r.store.Watch(ctx, r.gvr, options)
}

// List implements rest.Lister interface
func (r *rester) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	return r.store.List(ctx, r.gvr, options)
}

// NewList implements rest.Lister interface
func (r *rester) NewList() runtime.Object {
	return &unstructured.UnstructuredList{}
}

// ConvertToTable implements rest.Lister interface
func (r *rester) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.tableConvertor.ConvertToTable(ctx, object, tableOptions)
}
