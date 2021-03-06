package main

import (
	"github.com/jinzhu/gorm"
	proto "github.com/tamarakaufler/grpc-publication-manager/author-service/proto"
)

type Datastore interface {
	Create(*proto.Author) error
	Get(id string) error
	GetByEmail(author *proto.Author) (*proto.Author, error)
	GetAllAuthors() ([]*proto.Author, error)
	UpdateToken(*proto.Author) error
}

type Store struct {
	db *gorm.DB
}

func (st *Store) Create(author *proto.Author) error {
	if err := st.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}

func (st *Store) Get(id string) (*proto.Author, error) {
	var author *proto.Author
	author.Id = id

	if err := st.db.First(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (st *Store) GetByEmail(email string) (*proto.Author, error) {
	author := &proto.Author{}
	if err := st.db.Where("email = ?", email).
		First(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (st *Store) GetAll() ([]*proto.Author, error) {
	var authors []*proto.Author
	var err error
	if err = st.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (st *Store) UpdateToken(author *proto.Author) error {
	return st.db.Model(author).Update("token", author.Token).Error
}
