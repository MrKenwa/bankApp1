package paymentUsecase

import (
	"bankApp1/internal/models"
	"bankApp1/pkg/utils/pointer"
	"context"
	"errors"
	"github.com/avito-tech/go-transaction-manager/trm/v2/mock"
	deprecated_gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

type MockRepos struct {
	manager   *mock.MockManager
	opRepo    *MockOperationRepo
	balanceUC *MockBalanceUC
	cardUC    *MockCardsUC
	depositUC *MockDepositsUC
}

func TestPaymentUC_Send(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx      context.Context
		sendData *SendData
	}

	type want struct {
		opID models.OperationID
		err  error
	}

	tests := []struct {
		Name    string
		Prepare func(mockRepos *MockRepos)
		Args    args
		Want    want
		WantErr bool
	}{
		{
			Name: "Successful case",
			Prepare: func(mockRepos *MockRepos) {
				mockRepos.manager.EXPECT().Do(deprecated_gomock.Any(), deprecated_gomock.Any()).Times(1).
					DoAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					})
				mockRepos.balanceUC.EXPECT().Increase(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{11},
				}, int64(500)).Times(1).Return(nil)
				mockRepos.balanceUC.EXPECT().Decrease(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}, int64(500)).Times(1).Return(nil)
				mockRepos.opRepo.EXPECT().Create(gomock.Any(), models.Operation{
					SenderBalanceID:   pointer.Ptr(models.BalanceID(10)),
					ReceiverBalanceID: pointer.Ptr(models.BalanceID(11)),
					Amount:            500,
					OperationType:     "transfer",
				}).Times(1).Return(models.OperationID(101), nil)
				mockRepos.cardUC.EXPECT().Get(gomock.Any(), models.CardFilter{
					IDs: []models.CardID{6},
				}).Times(1).Return(models.Card{
					CardID: 6,
					UserID: 1,
				}, nil)
				mockRepos.balanceUC.EXPECT().Get(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}).Times(1).Return(models.Balance{
					BalanceID: 10,
					CardID:    pointer.Ptr(models.CardID(6)),
					DepositID: nil,
					Amount:    5000,
				}, nil)
				mockRepos.balanceUC.EXPECT().Get(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}).Times(1).Return(models.Balance{
					BalanceID: 10,
					CardID:    pointer.Ptr(models.CardID(6)),
					DepositID: nil,
					Amount:    5000,
				}, nil)
			},
			Args: args{
				ctx: ctx,
				sendData: &SendData{
					UserID:           1,
					SendBalanceID:    pointer.Ptr(models.BalanceID(10)),
					ReceiveBalanceID: pointer.Ptr(models.BalanceID(11)),
					Amount:           500,
					OpType:           "transfer",
				},
			},
			Want:    want{opID: 101, err: nil},
			WantErr: false,
		},
		{
			Name: "Error not enough rights",
			Prepare: func(mockRepos *MockRepos) {
				mockRepos.manager.EXPECT().Do(deprecated_gomock.Any(), deprecated_gomock.Any()).Times(1).
					DoAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					})
				mockRepos.cardUC.EXPECT().Get(gomock.Any(), models.CardFilter{
					IDs: []models.CardID{6},
				}).Times(1).Return(models.Card{
					CardID: 6,
					UserID: 35,
				}, nil)
				mockRepos.balanceUC.EXPECT().Get(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}).Times(1).Return(models.Balance{
					BalanceID: 10,
					CardID:    pointer.Ptr(models.CardID(6)),
					DepositID: nil,
					Amount:    5000,
				}, nil)
				mockRepos.balanceUC.EXPECT().Get(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}).Times(1).Return(models.Balance{
					BalanceID: 10,
					CardID:    pointer.Ptr(models.CardID(6)),
					DepositID: nil,
					Amount:    5000,
				}, nil)
			},
			Args: args{
				ctx: ctx,
				sendData: &SendData{
					UserID:           1,
					SendBalanceID:    pointer.Ptr(models.BalanceID(10)),
					ReceiveBalanceID: pointer.Ptr(models.BalanceID(11)),
					Amount:           500,
					OpType:           "transfer",
				},
			},
			Want:    want{opID: -1, err: errors.New("not enough rights")},
			WantErr: true,
		},
		{
			Name: "Error not enough money",
			Prepare: func(mockRepos *MockRepos) {
				mockRepos.manager.EXPECT().Do(deprecated_gomock.Any(), deprecated_gomock.Any()).Times(1).
					DoAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx)
					})
				mockRepos.cardUC.EXPECT().Get(gomock.Any(), models.CardFilter{
					IDs: []models.CardID{6},
				}).Times(1).Return(models.Card{
					CardID: 6,
					UserID: 1,
				}, nil)
				mockRepos.balanceUC.EXPECT().Get(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}).Times(1).Return(models.Balance{
					BalanceID: 10,
					CardID:    pointer.Ptr(models.CardID(6)),
					DepositID: nil,
					Amount:    100,
				}, nil)
				mockRepos.balanceUC.EXPECT().Get(gomock.Any(), models.BalanceFilter{
					IDs: []models.BalanceID{10},
				}).Times(1).Return(models.Balance{
					BalanceID: 10,
					CardID:    pointer.Ptr(models.CardID(6)),
					DepositID: nil,
					Amount:    100,
				}, nil)
			},
			Args: args{
				ctx: ctx,
				sendData: &SendData{
					UserID:           1,
					SendBalanceID:    pointer.Ptr(models.BalanceID(10)),
					ReceiveBalanceID: pointer.Ptr(models.BalanceID(11)),
					Amount:           500,
					OpType:           "transfer",
				},
			},
			Want:    want{opID: -1, err: errors.New("not enough money")},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			deprecated_ctrl := deprecated_gomock.NewController(t)
			defer deprecated_ctrl.Finish()

			mockRepos := MockRepos{
				manager:   mock.NewMockManager(deprecated_ctrl),
				opRepo:    NewMockOperationRepo(ctrl),
				balanceUC: NewMockBalanceUC(ctrl),
				cardUC:    NewMockCardsUC(ctrl),
				depositUC: NewMockDepositsUC(ctrl),
			}

			// Задаём мокам ожидаемое поведение с помощью функции Prepare
			if tt.Prepare != nil {
				tt.Prepare(&mockRepos)
			}

			u := NewPaymentUC(
				mockRepos.manager,
				mockRepos.balanceUC,
				mockRepos.opRepo,
				mockRepos.cardUC,
				mockRepos.depositUC,
			)

			got, err := u.Send(tt.Args.ctx, tt.Args.sendData)
			if tt.WantErr {
				require.Error(t, err)
				require.EqualError(t, err, tt.Want.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.Want.opID, got)
		})
	}
}
