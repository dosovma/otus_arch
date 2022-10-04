package adapters_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/adapters/mocks"
)

func TestRotatableAdapter_GetAngularVelocity(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockObj *mocks.MockUObject)
		want      int
		wantErr   bool
	}{
		{
			name: "case: success",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("angularVelocity").Return(5, true).Times(1)
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "case: failed to get angularVelocity property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("angularVelocity").Return(0, false).Times(1)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "case: object returns invalid property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("angularVelocity").Return(1.0, true).Times(1)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			r := adapters.RotatableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(mockObj)

			got, err := r.GetAngularVelocity()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAngularVelocity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAngularVelocity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotatableAdapter_GetDirection(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockObj *mocks.MockUObject)
		want      int
		wantErr   bool
	}{
		{
			name: "case: success",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("direction").Return(5, true).Times(1)
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "case: failed to get direction property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("direction").Return(0, false).Times(1)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "case: object returns invalid property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("direction").Return(1.0, true).Times(1)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			r := adapters.RotatableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(mockObj)

			got, err := r.GetDirection()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAngularVelocity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAngularVelocity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotatableAdapter_GetMaxDirections(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockObj *mocks.MockUObject)
		want      int
		wantErr   bool
	}{
		{
			name: "case: success",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("maxDirections").Return(8, true).Times(1)
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "case: failed to get direction property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("maxDirections").Return(0, false).Times(1)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "case: object returns invalid property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("maxDirections").Return(1.0, true).Times(1)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			r := adapters.RotatableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(mockObj)

			got, err := r.GetMaxDirections()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAngularVelocity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAngularVelocity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotatableAdapter_SetDirection(t *testing.T) {
	type args struct {
		direction int
	}

	tests := []struct {
		name      string
		args      args
		mockCalls func(a args, mockObj *mocks.MockUObject)
		wantErr   bool
	}{
		{
			name: "case: success",
			args: args{
				direction: 8,
			},
			mockCalls: func(a args, mockObj *mocks.MockUObject) {
				mockObj.EXPECT().SetProperty("direction", a.direction).Return(true).Times(1)
			},
			wantErr: false,
		},
		{
			name: "case: failed to set new direction",
			args: args{
				direction: 8,
			},
			mockCalls: func(a args, mockObj *mocks.MockUObject) {
				mockObj.EXPECT().SetProperty("direction", a.direction).Return(false).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			r := adapters.RotatableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(tt.args, mockObj)

			if err := r.SetDirection(tt.args.direction); (err != nil) != tt.wantErr {
				t.Errorf("SetDirection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
