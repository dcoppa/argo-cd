package mocks

import (
	apiclient "github.com/dcoppa/argo-cd/v2/reposerver/apiclient"

	io "github.com/dcoppa/argo-cd/v2/util/io"
)

type Clientset struct {
	RepoServerServiceClient apiclient.RepoServerServiceClient
}

func (c *Clientset) NewRepoServerClient() (io.Closer, apiclient.RepoServerServiceClient, error) {
	return io.NopCloser, c.RepoServerServiceClient, nil
}
