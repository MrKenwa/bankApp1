package paymentUsecase

import (
	"bankApp1/internal/models"
	"context"
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
				mockRepos.balanceUC.EXPECT().Increase(ctx, models.BalanceFilter{}, 1000).Times(1).Return(nil)
				mockRepos.balanceUC.EXPECT().Decrease(ctx, models.BalanceFilter{}, 1000).Times(1).Return(nil)
				mockRepos.opRepo.EXPECT().Create(ctx, models.Operation{}).Times(1).Return(1, nil)
				mockRepos.cardUC.EXPECT().Get(ctx, models.CardFilter{}).Times(1).Return(models.Card{}, nil)
				mockRepos.depositUC.EXPECT().Get(ctx, models.DepositFilter{}).Times(1).Return(models.Deposit{}, nil)
			},
			Args: args{
				ctx: ctx,
				sendData: &SendData{
					UserID:           1,
					SendBalanceID:    nil,
					ReceiveBalanceID: nil,
					Amount:           100,
					OpType:           "hui",
				},
			},
			Want:    want{opID: 1, err: nil},
			WantErr: false,
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
				require.ErrorIs(t, tt.Want.err, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.Want.opID, got)
		})
	}
}
