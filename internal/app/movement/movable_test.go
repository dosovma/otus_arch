package movement_test

import (
	"errors"
	"testing"

	"github.com/dosovma/otus_arch/internal/app/movement"
	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/movement/mocks"
	"github.com/dosovma/otus_arch/pkg"
)

func TestMove_Execute(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockMove *mocks.MockMovable)
		wantErr   bool
	}{
		{
			name: "case: success: object with position (12, 5) and velocity (-7, 3) changes the position to (5, 8)",
			mockCalls: func(mockMove *mocks.MockMovable) {
				gomock.InOrder(
					mockMove.EXPECT().GetPosition().Return(pkg.Vector{12, 5}, nil).Times(1),
					mockMove.EXPECT().GetVelocity().Return(pkg.Vector{-7, 3}, nil).Times(1),
					mockMove.EXPECT().SetPosition(pkg.Vector{5, 8}).Return(nil).Times(1),
				)
			},
			wantErr: false,
		},
		{
			name: "case: failed to set new position to object without setPosition property",
			mockCalls: func(mockMove *mocks.MockMovable) {
				gomock.InOrder(
					mockMove.EXPECT().GetPosition().Return(pkg.Vector{12, 5}, nil).Times(1),
					mockMove.EXPECT().GetVelocity().Return(pkg.Vector{-7, 3}, nil).Times(1),
					mockMove.EXPECT().SetPosition(pkg.Vector{5, 8}).Return(errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: vectors' dimension of position and velocity are different",
			mockCalls: func(mockMove *mocks.MockMovable) {
				gomock.InOrder(
					mockMove.EXPECT().GetPosition().Return(pkg.Vector{12, 5}, nil).Times(1),
					mockMove.EXPECT().GetVelocity().Return(pkg.Vector{-7, 3, 9}, nil).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to move objects without velocity property",
			mockCalls: func(mockMove *mocks.MockMovable) {
				gomock.InOrder(
					mockMove.EXPECT().GetPosition().Return(pkg.Vector{12, 5}, nil).Times(1),
					mockMove.EXPECT().GetVelocity().Return(nil, errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to move objects without position property",
			mockCalls: func(mockMove *mocks.MockMovable) {
				gomock.InOrder(
					mockMove.EXPECT().GetPosition().Return(nil, errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMove := mocks.NewMockMovable(gomock.NewController(t))
			m := movement.Move{
				Move: mockMove,
			}
			tt.mockCalls(mockMove)

			if err := m.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
