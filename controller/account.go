package controller

import (
	"context"
	"gomicro/database"
	"gomicro/pb/proto"
)

type AccountsService struct {
	proto.UnimplementedAccountServiceServer
}

func NewAccountsService() *AccountsService {
	return &AccountsService{}
}

func (server *AccountsService) GetAllAccounts(
	ctx context.Context,
	req *proto.GetAllAccountsRequest,
) (*proto.GetAllAccountsResponse, error) {
	allAccountsFromDB, err := database.GetAllAccounts()
	if err != nil {
		return &proto.GetAllAccountsResponse{
			Data:   []*proto.Account{},
			Status: proto.DBOPERATIONSTATUS_ERROR_WHILE_SAVING_IN_DB,
		}, err
	}
	response := &proto.GetAllAccountsResponse{
		Data:   []*proto.Account{},
		Status: proto.DBOPERATIONSTATUS_QUERIED_RECORD_FROM_DB,
	}
	for _, data := range allAccountsFromDB {
		protoData := &proto.Account{
			Owner:    data.Owner,
			Balance:  data.Balance,
			Currency: data.Currency,
		}

		response.Data = append(response.Data, protoData)
	}
	return response, err
}

func (server *AccountsService) GetAccountById(
	ctx context.Context,
	req *proto.GetAccountByIdRequest,
) (*proto.GetAccountByIdResponse, error) {
	accountFromDB, err := database.GetAccountById(1)
	if err != nil {
		return &proto.GetAccountByIdResponse{
			Data:   &proto.Account{},
			Status: proto.DBOPERATIONSTATUS_ERROR_WHILE_SAVING_IN_DB,
		}, err
	}
	response := &proto.GetAccountByIdResponse{
		Data: &proto.Account{
			Owner:    accountFromDB.Owner,
			Balance:  accountFromDB.Balance,
			Currency: accountFromDB.Currency,
		},
		Status: proto.DBOPERATIONSTATUS_QUERIED_RECORD_FROM_DB,
	}

	return response, err
}
