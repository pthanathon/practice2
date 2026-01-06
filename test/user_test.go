package test

import (
	"practice2/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Name is required`, func(t *testing.T) {
		user := entity.User{
			Name: "", //ผิดเพราะว่าง
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})
}
