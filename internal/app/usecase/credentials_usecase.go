package usecase

import (
	"context"

	"github.com/rmscoal/go-restful-monolith-boilerplate/internal/app/repo"
	"github.com/rmscoal/go-restful-monolith-boilerplate/internal/app/service"
	"github.com/rmscoal/go-restful-monolith-boilerplate/internal/domain"
)

type credentialUseCase struct {
	repo    repo.IUserRepo
	service service.IDoorkeeperService
}

func NewCredentialUseCase(repo repo.IUserRepo, service service.IDoorkeeperService) ICredentialUseCase {
	return &credentialUseCase{repo: repo, service: service}
}

func (uc *credentialUseCase) SignUp(ctx context.Context, user domain.User) (domain.User, error) {
	// Validate user entity
	if err := user.ValidateWithContext(ctx); err != nil {
		return user, NewDomainError("User", err)
	}

	// Validate repository state
	if err := uc.repo.ValidateRepoState(ctx, user); err != nil {
		return user, NewConflictError("User", err)
	}

	user.Credential.Password = uc.service.HashPassword(user.Credential.Password)
	if user, err := uc.repo.CreateNewUser(ctx, user); err != nil {
		return user, err
	}

	token, err := uc.service.GenerateToken(user)
	if err != nil {
		return user, err
	}
	user.Credential.Token = token

	return user, nil
}
