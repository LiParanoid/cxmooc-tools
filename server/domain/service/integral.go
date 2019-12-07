package service

import (
	"github.com/CodFrm/cxmooc-tools/server/application/dto"
	"github.com/CodFrm/cxmooc-tools/server/domain/repository"
	"github.com/CodFrm/cxmooc-tools/server/internal/errs"
)

const (
	INTEGRAL_VCODE = 10
)

type Integral struct {
	userRepo     repository.UserRepository
	integralRepo repository.IntegralRepository
}

func NewIntegralService(userRepo repository.UserRepository, integralRepo repository.IntegralRepository) *Integral {
	return &Integral{
		userRepo:     userRepo,
		integralRepo: integralRepo,
	}
}

func (i *Integral) Consumption(token string, rule int) error {
	user, err := i.userRepo.FindByToken(token)
	if err != nil {
		return err
	}
	integral, err := i.integralRepo.GetIntegral(user)
	if err != nil {
		return err
	}
	if integral.Num < rule {
		return errs.IntegralInsufficient
	}
	return nil
}

func (i *Integral) UserAddIntegral(usr string, num int) (*dto.TokenTransaction, error) {
	user, err := i.userRepo.FindByUser(usr)
	if err != nil {
		return nil, err
	}
	integral, err := i.integralRepo.GetIntegral(user)
	if err != nil {
		return nil, err
	}
	integral.Num += num
	if err := i.integralRepo.Update(integral); err != nil {
		return nil, err
	}
	return &dto.TokenTransaction{
		Token:  integral.User.Token,
		Num:    num,
		AddNum: integral.Num,
	}, nil
}
