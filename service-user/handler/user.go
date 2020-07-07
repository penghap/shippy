package handler

import (
	"errors"

	log "github.com/micro/go-micro/v2/logger"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	pb "github.com/penghap/shippy/service-user/proto/user"
	"github.com/penghap/shippy/service-user/repository"
)

type Service struct {
	repo         repository.Repository
	tokenService TokenService
}

func (srv *Service) SetRepo(repo repository.Repository) {
	srv.repo = repo
}

func (srv *Service) Get(ctx context.Context, in *pb.User, out *pb.Response) error {
	user, err := srv.repo.Get(in.Id)
	if err != nil {
		return err
	}
	out.User = user
	return nil
}

func (srv *Service) GetAll(ctx context.Context, in *pb.User, out *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	out.Users = users
	return nil
}

func (srv *Service) Create(ctx context.Context, in *pb.User, out *pb.Response) error {
	log.Info("Create in:", in)
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	in.Password = string(hashedPwd)
	if err := srv.repo.Create(in); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(in)
	if err != nil {
		return err
	}

	out.User = in
	out.Token = &pb.Token{Token: token}
	return nil
}

func (srv *Service) Auth(ctx context.Context, in *pb.User, out *pb.Token) error {
	log.Info("Logging in with:", in.Email, in.Password)
	user, err := srv.repo.GetByEmail(in.Email)
	if err != nil {
		return err
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return errors.New("Password is incorrect")
	}

	if err != nil {
		return err
	}
	log.Info(user)
	out.Token = "testing"
	return nil
}

func (srv *Service) ValidateToken(ctx context.Context, in *pb.Token, out *pb.Token) error {
	// Decode token
	claims, err := srv.tokenService.Decode(in.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("Invalid user")
	}

	out.Valid = true

	return nil
}
