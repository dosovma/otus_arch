package movement_test

import (
	"errors"
	"testing"

	"github.com/dosovma/otus_arch/internal/app/movement"
	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/movement/mocks"
)

func TestRotate_Execute(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockRotate *mocks.MockRotatable)
		wantErr   bool
	}{
		{
			name: "case: success",
			mockCalls: func(mockRotate *mocks.MockRotatable) {
				gomock.InOrder(
					mockRotate.EXPECT().GetDirection().Return(6, nil).Times(1),
					mockRotate.EXPECT().GetAngularVelocity().Return(3, nil).Times(1),
					mockRotate.EXPECT().GetMaxDirections().Return(8, nil).Times(1),
					mockRotate.EXPECT().SetDirection(1).Return(nil).Times(1),
				)
			},
			wantErr: false,
		},
		{
			name: "case: failed to set new direction",
			mockCalls: func(mockRotate *mocks.MockRotatable) {
				gomock.InOrder(
					mockRotate.EXPECT().GetDirection().Return(6, nil).Times(1),
					mockRotate.EXPECT().GetAngularVelocity().Return(3, nil).Times(1),
					mockRotate.EXPECT().GetMaxDirections().Return(8, nil).Times(1),
					mockRotate.EXPECT().SetDirection(1).Return(errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to get maxDirections",
			mockCalls: func(mockRotate *mocks.MockRotatable) {
				gomock.InOrder(
					mockRotate.EXPECT().GetDirection().Return(6, nil).Times(1),
					mockRotate.EXPECT().GetAngularVelocity().Return(3, nil).Times(1),
					mockRotate.EXPECT().GetMaxDirections().Return(0, errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to get angular velocity",
			mockCalls: func(mockRotate *mocks.MockRotatable) {
				gomock.InOrder(
					mockRotate.EXPECT().GetDirection().Return(6, nil).Times(1),
					mockRotate.EXPECT().GetAngularVelocity().Return(0, errors.New("error")).Times(1),
				)
			},
			wantErr: true,
		},
		{
			name: "case: failed to get direction",
			mockCalls: func(mockRotate *mocks.MockRotatable) {
				mockRotate.EXPECT().GetDirection().Return(1, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRotate := mocks.NewMockRotatable(gomock.NewController(t))
			rm := movement.Rotate{
				Rotate: mockRotate,
			}
			tt.mockCalls(mockRotate)

			if err := rm.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
