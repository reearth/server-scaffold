package assetuc

import (
	"context"
	"errors"

	"github.com/reearth/scaffold/server/pkg/asset"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
	"github.com/samber/lo"
)

type UpdateParam struct {
	ID   asset.ID
	Name *string
}

func (p UpdateParam) Validate() error {
	if lo.IsEmpty(p.ID) {
		return errors.New("id is required")
	}
	return nil
}

type Update struct {
	assetRepo     asset.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo
	assetPolicy   asset.Policy
}

func NewUpdate(
	assetRepo asset.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy asset.Policy,
) *Update {
	return &Update{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *Update) Execute(ctx context.Context, param UpdateParam, user *user.User) (*asset.Asset, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	asset, _, _, err := UsecaseBuilder(ctx, user).
		FindAssetByID(param.ID, uc.assetRepo).
		FindProjectByAsset(uc.projectRepo, uc.workspaceRepo).
		CanUpdateAsset(uc.assetPolicy).
		Result()
	if err != nil {
		return nil, err
	}

	if param.Name != nil {
		asset.SetName(*param.Name)
	}

	if err := uc.assetRepo.Save(ctx, asset); err != nil {
		return nil, err
	}

	return asset, nil
}
