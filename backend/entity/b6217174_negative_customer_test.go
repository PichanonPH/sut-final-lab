package entity

import (
	"testing"
	//"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNotCorrectCustomerID(t *testing.T) {

	g := NewGomegaWithT(t)

	cu := Customer{
		Name:       "Palida",
		Email:      "p@gmail.com", //ผิด
		CustomerID: "L1234",
	}

	ok, err := govalidator.ValidateStruct(cu)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Not CustomerID"))
}
