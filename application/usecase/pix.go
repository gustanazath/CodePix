package usecase

import (
	"github.com/gustanazath/CodePix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId, string) (*model.PixKey, error){
	 account, err := p.PixKeyRepository.FindAccount(accountId)
	 if err != nil {
		return nil, err
	 }

	 pixKey, err := model.NewPixKey(kind,account,key)

	 if err != nil {
		return nil, err
	 }
	 
	 p.PixKeyRepository.RegisterKey(pixKey)
	 if pixKey.ID == "" {
		return nil, err
	 }

	return pixKey, nil
}

func (p *PixUseCase) FinKey(key string, kind string) (*model.PixKey, error){
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key,kin)
	if err != nil{
		return nil, err
	}
	return pixKey, nil
}