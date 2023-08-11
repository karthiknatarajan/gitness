// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package pullreq

import (
	"context"
	"fmt"
	"strconv"

	"github.com/harness/gitness/events"
	"github.com/harness/gitness/gitrpc"
	gitrpcenum "github.com/harness/gitness/gitrpc/enum"
	pullreqevents "github.com/harness/gitness/internal/events/pullreq"
)

// createHeadRefOnCreated handles pull request Created events.
// It creates the PR head git ref.
func (s *Service) createHeadRefOnCreated(ctx context.Context,
	event *events.Event[*pullreqevents.CreatedPayload],
) error {
	repoGit, err := s.repoGitInfoCache.Get(ctx, event.Payload.TargetRepoID)
	if err != nil {
		return fmt.Errorf("failed to get repo git info: %w", err)
	}

	writeParams, err := createSystemRPCWriteParams(ctx, s.urlProvider, repoGit.ID, repoGit.GitUID)
	if err != nil {
		return fmt.Errorf("failed to generate rpc write params: %w", err)
	}

	// TODO: This doesn't work for forked repos (only works when sourceRepo==targetRepo).
	// This is because commits from the source repository must be first pulled into the target repository.
	err = s.gitRPCClient.UpdateRef(ctx, gitrpc.UpdateRefParams{
		WriteParams: writeParams,
		Name:        strconv.Itoa(int(event.Payload.Number)),
		Type:        gitrpcenum.RefTypePullReqHead,
		NewValue:    event.Payload.SourceSHA,
		OldValue:    "", // this is a new pull request, so we expect that the ref doesn't exist
	})
	if err != nil {
		return fmt.Errorf("failed to update PR head ref: %w", err)
	}

	return nil
}

// updateHeadRefOnBranchUpdate handles pull request Branch Updated events.
// It updates the PR head git ref to point to the latest commit.
func (s *Service) updateHeadRefOnBranchUpdate(ctx context.Context,
	event *events.Event[*pullreqevents.BranchUpdatedPayload],
) error {
	repoGit, err := s.repoGitInfoCache.Get(ctx, event.Payload.TargetRepoID)
	if err != nil {
		return fmt.Errorf("failed to get repo git info: %w", err)
	}

	writeParams, err := createSystemRPCWriteParams(ctx, s.urlProvider, repoGit.ID, repoGit.GitUID)
	if err != nil {
		return fmt.Errorf("failed to generate rpc write params: %w", err)
	}

	// TODO: This doesn't work for forked repos (only works when sourceRepo==targetRepo)
	// This is because commits from the source repository must be first pulled into the target repository.
	err = s.gitRPCClient.UpdateRef(ctx, gitrpc.UpdateRefParams{
		WriteParams: writeParams,
		Name:        strconv.Itoa(int(event.Payload.Number)),
		Type:        gitrpcenum.RefTypePullReqHead,
		NewValue:    event.Payload.NewSHA,
		OldValue:    event.Payload.OldSHA,
	})
	if err != nil {
		return fmt.Errorf("failed to update PR head ref after new commit: %w", err)
	}

	return nil
}

// updateHeadRefOnReopen handles pull request StateChanged events.
// It updates the PR head git ref to point to the source branch commit SHA.
func (s *Service) updateHeadRefOnReopen(ctx context.Context,
	event *events.Event[*pullreqevents.ReopenedPayload],
) error {
	repoGit, err := s.repoGitInfoCache.Get(ctx, event.Payload.TargetRepoID)
	if err != nil {
		return fmt.Errorf("failed to get repo git info: %w", err)
	}

	writeParams, err := createSystemRPCWriteParams(ctx, s.urlProvider, repoGit.ID, repoGit.GitUID)
	if err != nil {
		return fmt.Errorf("failed to generate rpc write params: %w", err)
	}

	// TODO: This doesn't work for forked repos (only works when sourceRepo==targetRepo)
	// This is because commits from the source repository must be first pulled into the target repository.
	err = s.gitRPCClient.UpdateRef(ctx, gitrpc.UpdateRefParams{
		WriteParams: writeParams,
		Name:        strconv.Itoa(int(event.Payload.Number)),
		Type:        gitrpcenum.RefTypePullReqHead,
		NewValue:    event.Payload.SourceSHA,
		OldValue:    "", // the request is re-opened, so anything can be the old value
	})
	if err != nil {
		return fmt.Errorf("failed to update PR head ref after pull request reopen: %w", err)
	}

	return nil
}