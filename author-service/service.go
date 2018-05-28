package main

import (
	"errors"

	proto "github.com/tamarakaufler/grpc-publication-manager/author-service/proto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var (
	ErrAuthorExists   = errors.New("Author already exists")
	ErrAuthorNotFound = errors.New("Author not found")
	ErrInvalidAuthor  = errors.New("Author not valid")
)

type service struct {
	db           *Store
	tokenService TokenService
}

// CreateAuthor hashes the plaintext password and inserts the author in the database
func (s *service) CreateAuthor(ctx context.Context, req *proto.Author) (*proto.EmptyResponse, error) {
	_, err := s.getAuthorByEmail(ctx, req)
	if err == nil {
		return nil, ErrAuthorExists
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hp)
	return &proto.EmptyResponse{}, s.db.Create(req)
}

func (s *service) getAuthorByEmail(ctx context.Context, req *proto.Author) (*proto.Author, error) {
	author, err := s.db.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) GetAuthor(ctx context.Context, req *proto.Author) (*proto.Response, error) {
	author, err := s.db.Get(req.Id)
	if err != nil {
		return nil, ErrAuthorNotFound
	}
	res := &proto.Response{Author: author, Created: true}
	return res, nil
}

func (s *service) GetAllAuthors(ctx context.Context, req *proto.GetAllRequest) (*proto.Response, error) {
	authors, err := s.db.GetAll()
	if err != nil {
		return nil, err
	}
	res := &proto.Response{Authors: authors}
	return res, nil
}

func (s *service) Authenticate(ctx context.Context, req *proto.Author) (*proto.Token, error) {
	author, err := s.db.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, err := s.tokenService.Encode(author)
	if err != nil {
		return nil, err
	}
	author.Token = token
	if err := s.db.UpdateToken(author); err != nil {
		return nil, err
	}

	t := &proto.Token{
		Token: token,
		Valid: true,
	}
	return t, nil
}

func (s *service) ValidateToken(ctx context.Context, t *proto.Token) (*proto.Token, error) {
	claims, err := s.tokenService.Decode(t.Token)
	if err != nil {
		return t, err
	}
	if claims.Author == nil || claims.Author.Id == "" {
		return t, ErrInvalidAuthor
	}

	t.Valid = true
	return t, nil
}

func (s *service) InvalidateToken(ctx context.Context, author *proto.Author) (*proto.EmptyResponse, error) {
	author.Token = ""
	return &proto.EmptyResponse{}, s.db.UpdateToken(author)
}
