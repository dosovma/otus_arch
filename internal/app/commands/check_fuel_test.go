package commands_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/commands/mocks"
)

func TestCheckFuelCmd_Execute(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockMove *mocks.MockICheckFuel)
		wantErr   bool
	}{
		{
			name: "case: success: object has enough fuel",
			mockCalls: func(mockMove *mocks.MockICheckFuel) {
				mockMove.EXPECT().GetFuel().Return(100, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "case: object has insufficient fuel",
			mockCalls: func(mockMove *mocks.MockICheckFuel) {
				mockMove.EXPECT().GetFuel().Return(0, nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "case: failed to get fuel property",
			mockCalls: func(mockMove *mocks.MockICheckFuel) {
				mockMove.EXPECT().GetFuel().Return(0, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCheckFuel := mocks.NewMockICheckFuel(gomock.NewController(t))
			cf := commands.CheckFuelCmd{CheckFuel: mockCheckFuel}
			tt.mockCalls(mockCheckFuel)

			if err := cf.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
