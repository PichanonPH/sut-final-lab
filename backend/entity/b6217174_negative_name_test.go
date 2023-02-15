package entity

import (
	"testing"
	//"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBlack(t *testing.T) {

	g := NewGomegaWithT(t)

	cu := Customer{
		Name:       "",
		Email:      "p@gmail.com", //ผิด
		CustomerID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(cu)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Not Blank"))
}

//$ git config --global user.name "John Doe"
//$ git config --global user.email johndoe@example.com
