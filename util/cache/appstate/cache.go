package appstate

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"

	appv1 "github.com/dcoppa/argo-cd/v2/pkg/apis/application/v1alpha1"
	cacheutil "github.com/dcoppa/argo-cd/v2/util/cache"
	"github.com/dcoppa/argo-cd/v2/util/env"
)

var ErrCacheMiss = cacheutil.ErrCacheMiss

const (
	clusterInfoCacheExpiration = 10 * time.Minute
)

type Cache struct {
	Cache                   *cacheutil.Cache
	appStateCacheExpiration time.Duration
}

func NewCache(cache *cacheutil.Cache, appStateCacheExpiration time.Duration) *Cache {
	return &Cache{cache, appStateCacheExpiration}
}

func AddCacheFlagsToCmd(cmd *cobra.Command, opts ...cacheutil.Options) func() (*Cache, error) {
	var appStateCacheExpiration time.Duration

	cmd.Flags().DurationVar(&appStateCacheExpiration, "app-state-cache-expiration", env.ParseDurationFromEnv("ARGOCD_APP_STATE_CACHE_EXPIRATION", 1*time.Hour, 0, 10*time.Hour), "Cache expiration for app state")

	cacheFactory := cacheutil.AddCacheFlagsToCmd(cmd, opts...)

	return func() (*Cache, error) {
		cache, err := cacheFactory()
		if err != nil {
			return nil, err
		}
		return NewCache(cache, appStateCacheExpiration), nil
	}
}

func (c *Cache) GetItem(key string, item interface{}) error {
	return c.Cache.GetItem(key, item)
}

func (c *Cache) SetItem(key string, item interface{}, expiration time.Duration, delete bool) error {
	return c.Cache.SetItem(key, item, &cacheutil.CacheActionOpts{Expiration: expiration, Delete: delete})
}

func appManagedResourcesKey(appName string) string {
	return fmt.Sprintf("app|managed-resources|%s", appName)
}

func (c *Cache) GetAppManagedResources(appName string, res *[]*appv1.ResourceDiff) error {
	err := c.GetItem(appManagedResourcesKey(appName), &res)
	return err
}

func (c *Cache) SetAppManagedResources(appName string, managedResources []*appv1.ResourceDiff) error {
	sort.Slice(managedResources, func(i, j int) bool {
		return managedResources[i].FullName() < managedResources[j].FullName()
	})
	return c.SetItem(appManagedResourcesKey(appName), managedResources, c.appStateCacheExpiration, managedResources == nil)
}

func appResourcesTreeKey(appName string) string {
	return fmt.Sprintf("app|resources-tree|%s", appName)
}

func clusterInfoKey(server string) string {
	return fmt.Sprintf("cluster|info|%s", server)
}

func (c *Cache) GetAppResourcesTree(appName string, res *appv1.ApplicationTree) error {
	err := c.GetItem(appResourcesTreeKey(appName), &res)
	return err
}

func (c *Cache) OnAppResourcesTreeChanged(ctx context.Context, appName string, callback func() error) error {
	return c.Cache.OnUpdated(ctx, appManagedResourcesKey(appName), callback)
}

func (c *Cache) SetAppResourcesTree(appName string, resourcesTree *appv1.ApplicationTree) error {
	if resourcesTree != nil {
		resourcesTree.Normalize()
	}
	err := c.SetItem(appResourcesTreeKey(appName), resourcesTree, c.appStateCacheExpiration, resourcesTree == nil)
	if err != nil {
		return err
	}
	return c.Cache.NotifyUpdated(appManagedResourcesKey(appName))
}

func (c *Cache) SetClusterInfo(server string, info *appv1.ClusterInfo) error {
	return c.SetItem(clusterInfoKey(server), info, clusterInfoCacheExpiration, info == nil)
}

func (c *Cache) GetClusterInfo(server string, res *appv1.ClusterInfo) error {
	err := c.GetItem(clusterInfoKey(server), &res)
	return err
}
