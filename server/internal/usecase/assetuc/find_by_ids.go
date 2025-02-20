package assetuc

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/asset"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type FindByIDs struct {
	assetRepo     asset.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo

	assetPolicy asset.Policy
}

func NewFindByIDs(
	assetRepo asset.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy asset.Policy,
) *FindByIDs {
	return &FindByIDs{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *FindByIDs) Execute(ctx context.Context, ids asset.IDList, user *user.User) (asset.List, error) {
	assets, err := uc.assetRepo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	projects, err := uc.projectRepo.FindByIDs(ctx, assets.ProjectIDs())
	if err != nil {
		return nil, err
	}

	workspaces, err := uc.workspaceRepo.FindByIDs(ctx, projects.WorkspaceIDs())
	if err != nil {
		return nil, err
	}

	assets, err = uc.assetPolicy.Filter(ctx, user, workspaces, projects, assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
