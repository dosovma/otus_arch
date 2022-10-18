package commands_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/commands/mocks"
)

func TestChangeVelocityCmd_Execute(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockMove *mocks.MockIChangeVelocity)
		wantErr   bool
	}{
		{
			name: "case: success: velocity of movable object has been changed",
			mockCalls: func(mockCV *mocks.MockIChangeVelocity) {
				gomock.InOrder(
					mockCV.EXPECT().IsMovable().Return(true, nil).Times(1),
					mockCV.EXPECT().ChangeVelocity().Return(nil).Times(1),
				)
			},
			wantErr: false,
		},
		{
			name: "case: failed to change velocity",
			mockCalls: func(mockCV *mocks.MockIChangeVelocity) {
				gomock.InOrder(
					mockCV.EXPECT().IsMovable().Return(true, nil).Times(1),
					mockCV.EXPECT().ChangeVelocity().Return(errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: object is not movable",
			mockCalls: func(mockCV *mocks.MockIChangeVelocity) {
				mockCV.EXPECT().IsMovable().Return(false, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "case: failed to check if the object is movable",
			mockCalls: func(mockCV *mocks.MockIChangeVelocity) {
				mockCV.EXPECT().IsMovable().Return(false, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCV := mocks.NewMockIChangeVelocity(gomock.NewController(t))
			m := commands.ChangeVelocityCmd{
				ChangeVelocity: mockCV,
			}
			tt.mockCalls(mockCV)

			if err := m.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
