// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Branch.
func (mg *Branch) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BranchProtection.
func (mg *BranchProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryIDRef,
		Selector:     mg.Spec.ForProvider.RepositoryIDSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryID")
	}
	mg.Spec.ForProvider.RepositoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RepositoryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryIDRef,
		Selector:     mg.Spec.InitProvider.RepositoryIDSelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RepositoryID")
	}
	mg.Spec.InitProvider.RepositoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DeployKey.
func (mg *DeployKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this File.
func (mg *File) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PullRequest.
func (mg *PullRequest) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BaseRepository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BaseRepositoryRef,
		Selector:     mg.Spec.ForProvider.BaseRepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BaseRepository")
	}
	mg.Spec.ForProvider.BaseRepository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BaseRepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HeadRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HeadRefRef,
		Selector:     mg.Spec.ForProvider.HeadRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HeadRef")
	}
	mg.Spec.ForProvider.HeadRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HeadRefRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BaseRepository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BaseRepositoryRef,
		Selector:     mg.Spec.InitProvider.BaseRepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BaseRepository")
	}
	mg.Spec.InitProvider.BaseRepository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BaseRepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HeadRef),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.HeadRefRef,
		Selector:     mg.Spec.InitProvider.HeadRefSelector,
		To: reference.To{
			List:    &BranchList{},
			Managed: &Branch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HeadRef")
	}
	mg.Spec.InitProvider.HeadRef = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HeadRefRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositoryAutolinkReference.
func (mg *RepositoryAutolinkReference) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RepositoryWebhook.
func (mg *RepositoryWebhook) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Ruleset.
func (mg *Ruleset) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RepositoryRef,
		Selector:     mg.Spec.InitProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
