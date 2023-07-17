package item

import (
	g_ "majordome/domain/repository/gorm"

	"github.com/jinzhu/gorm"

	pb "majordome/internal/protorender/item"
)

type Server struct {
	db *gorm.DB
	pb.UnimplementedItemServiceServer
}

func (s *Server) NewServer() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]

}
