package challenge_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sadrishehu/golang-simple-problems/challenge"
)

func TestChallenge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Challenge Suite")
}

var _ = Describe("BillFor", func() {
	constantUsers := []challenge.User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}

	userSignedUp := []challenge.User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            3,
			Name:          "Employee #3",
			ActivatedOn:   time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}

	newPlan := challenge.Subscription{
		Id:                    1,
		CustomerId:            1,
		MonthlyPriceInDollars: 4,
	}

	var noUsers []challenge.User

	It("works when the customer has no active users during the month", func() {
		Expect(challenge.BillFor("2019-01", &newPlan, &noUsers)).To(BeNumerically("~", 0.00, 0.005))
	})

	It("works when everything stays the same for a month", func() {
		Expect(challenge.BillFor("2019-01", &newPlan, &constantUsers)).To(BeNumerically("~", 8.00, 0.005))
	})

	It("works when a user is activated during the month", func() {
		Expect(challenge.BillFor("2019-01", &newPlan, &userSignedUp)).To(BeNumerically("~", 10.84, 0.005))
	})
})
