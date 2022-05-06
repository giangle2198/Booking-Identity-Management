package helper

import (
	"github.com/gogo/protobuf/proto"
	"github.com/jinzhu/copier"
)

type (
	// PbConverter converts Pb to DTO and vise versa
	PbConverter interface {
		FromPb(to interface{}, from proto.Message)
		ToPb(to proto.Message, from interface{})
	}

	// ModelConverter converts Pb to DTO and vise versa
	ModelConverter interface {
		FromModel(to interface{}, from interface{})
		ToModel(to interface{}, from interface{})
	}

	pbConverter struct {
	}

	modelConverter struct {
	}
)

// NewPbConverter creates an instance
func NewPbConverter() PbConverter {
	return &pbConverter{}
}

func (h *pbConverter) FromPb(to interface{}, from proto.Message) {
	_ = copier.Copy(to, from)
}

func (h *pbConverter) ToPb(to proto.Message, from interface{}) {
	_ = copier.Copy(to, from)
}

// NewModelConverter creates an instance
func NewModelConverter() ModelConverter {
	return &modelConverter{}
}

func (h *modelConverter) FromModel(to, from interface{}) {
	_ = copier.Copy(to, from)
}

func (h *modelConverter) ToModel(to, from interface{}) {
	_ = copier.Copy(to, from)
}
