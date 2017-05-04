// This file was automatically generated by lister-gen

package v1

import (
	api "github.com/openshift/origin/pkg/oauth/api"
	v1 "github.com/openshift/origin/pkg/oauth/api/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OAuthClientLister helps list OAuthClients.
type OAuthClientLister interface {
	// List lists all OAuthClients in the indexer.
	List(selector labels.Selector) (ret []*v1.OAuthClient, err error)
	// OAuthClients returns an object that can list and get OAuthClients.
	OAuthClients(namespace string) OAuthClientNamespaceLister
	OAuthClientListerExpansion
}

// oAuthClientLister implements the OAuthClientLister interface.
type oAuthClientLister struct {
	indexer cache.Indexer
}

// NewOAuthClientLister returns a new OAuthClientLister.
func NewOAuthClientLister(indexer cache.Indexer) OAuthClientLister {
	return &oAuthClientLister{indexer: indexer}
}

// List lists all OAuthClients in the indexer.
func (s *oAuthClientLister) List(selector labels.Selector) (ret []*v1.OAuthClient, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OAuthClient))
	})
	return ret, err
}

// OAuthClients returns an object that can list and get OAuthClients.
func (s *oAuthClientLister) OAuthClients(namespace string) OAuthClientNamespaceLister {
	return oAuthClientNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OAuthClientNamespaceLister helps list and get OAuthClients.
type OAuthClientNamespaceLister interface {
	// List lists all OAuthClients in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.OAuthClient, err error)
	// Get retrieves the OAuthClient from the indexer for a given namespace and name.
	Get(name string) (*v1.OAuthClient, error)
	OAuthClientNamespaceListerExpansion
}

// oAuthClientNamespaceLister implements the OAuthClientNamespaceLister
// interface.
type oAuthClientNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OAuthClients in the indexer for a given namespace.
func (s oAuthClientNamespaceLister) List(selector labels.Selector) (ret []*v1.OAuthClient, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OAuthClient))
	})
	return ret, err
}

// Get retrieves the OAuthClient from the indexer for a given namespace and name.
func (s oAuthClientNamespaceLister) Get(name string) (*v1.OAuthClient, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("oauthclient"), name)
	}
	return obj.(*v1.OAuthClient), nil
}