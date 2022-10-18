package commands_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/commands/mocks"
)

func TestBurnFuelCmd_Execute(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockMove *mocks.MockIBurnFuel)
		wantErr   bool
	}{
		{
			name: "case: success: fuel has been burnt",
			mockCalls: func(mockMove *mocks.MockIBurnFuel) {
				gomock.InOrder(
					mockMove.EXPECT().GetFuel().Return(100, nil).Times(1),
					mockMove.EXPECT().GetFuelConsumption().Return(7, nil).Times(1),
					mockMove.EXPECT().SetFuel(93).Return(nil).Times(1),
				)
			},
			wantErr: false,
		},
		{
			name: "case: failed to set new value of fuel to object",
			mockCalls: func(mockMove *mocks.MockIBurnFuel) {
				gomock.InOrder(
					mockMove.EXPECT().GetFuel().Return(100, nil).Times(1),
					mockMove.EXPECT().GetFuelConsumption().Return(7, nil).Times(1),
					mockMove.EXPECT().SetFuel(93).Return(errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to get fuel consumption",
			mockCalls: func(mockMove *mocks.MockIBurnFuel) {
				gomock.InOrder(
					mockMove.EXPECT().GetFuel().Return(100, nil).Times(1),
					mockMove.EXPECT().GetFuelConsumption().Return(0, errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to get fuel",
			mockCalls: func(mockMove *mocks.MockIBurnFuel) {
				mockMove.EXPECT().GetFuel().Return(0, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBurnFuel := mocks.NewMockIBurnFuel(gomock.NewController(t))
			m := commands.BurnFuelCmd{
				BurnFuel: mockBurnFuel,
			}
			tt.mockCalls(mockBurnFuel)

			if err := m.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
