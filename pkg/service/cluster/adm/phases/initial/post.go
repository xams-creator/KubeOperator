package initial

import (
	"github.com/KubeOperator/KubeOperator/pkg/service/cluster/adm/phases"
	"github.com/KubeOperator/KubeOperator/pkg/util/kobe"
)

const (
	initPost = "11-post.yml"
)

type PostPhase struct {
}

func (s PostPhase) Name() string {
	return "Post Init"
}

func (s PostPhase) Run(b kobe.Interface) (result kobe.Result, err error) {
	return phases.RunPlaybookAndGetResult(b, initPost)
}