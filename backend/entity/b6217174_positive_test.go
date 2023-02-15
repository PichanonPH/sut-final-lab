package entity

import (
	"testing"
	//"time"

	"gorm.io/gorm"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

//entyty

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Not Blank"`
	Email      string `valid:"email~Not Email"`
	CustomerID string `gorm:"uniqueIndex" valid:"matches(^[LMH]\\d{7}$)~Not CustomerID"`
}

func TestCorrect(t *testing.T) {

	g := NewGomegaWithT(t)

	cu := Customer{
		Name:       "Palida",
		Email:      "palm@gmail.com", //ผิด
		CustomerID: "M1234567",
	}

	ok, err := govalidator.ValidateStruct(cu)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
