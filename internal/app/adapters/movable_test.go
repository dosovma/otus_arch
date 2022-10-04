package adapters_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/adapters/mocks"
	"github.com/dosovma/otus_arch/pkg"
)

func TestMovableAdapter_GetPosition(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockObj *mocks.MockUObject)
		want      pkg.Vector
		wantErr   bool
	}{
		{
			name: "case: success",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("position").Return(pkg.Vector{5, 8}, true).Times(1)
			},
			want:    pkg.Vector{5, 8},
			wantErr: false,
		},
		{
			name: "case: failed to get position property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("position").Return(pkg.Vector{}, false).Times(1)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case: object returns invalid property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("position").Return(10, true).Times(1)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			m := adapters.MovableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(mockObj)

			got, err := m.GetPosition()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovableAdapter_GetVelocity(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockObj *mocks.MockUObject)
		want      pkg.Vector
		wantErr   bool
	}{
		{
			name: "case: success",
			mockCalls: func(mockObj *mocks.MockUObject) {
				gomock.InOrder(
					mockObj.EXPECT().GetProperty("direction").Return(2, true).Times(1),
					mockObj.EXPECT().GetProperty("maxDirections").Return(8, true).Times(1),
					mockObj.EXPECT().GetProperty("velocity").Return(10, true).Times(1),
				)
			},
			want:    pkg.Vector{0, 10},
			wantErr: false,
		},
		{
			name: "case: failed to get velocity property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				gomock.InOrder(
					mockObj.EXPECT().GetProperty("direction").Return(4, true).Times(1),
					mockObj.EXPECT().GetProperty("maxDirections").Return(8, true).Times(1),
					mockObj.EXPECT().GetProperty("velocity").Return(0, false).Times(1),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case: object returns invalid velocity property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				gomock.InOrder(
					mockObj.EXPECT().GetProperty("direction").Return(4, true).Times(1),
					mockObj.EXPECT().GetProperty("maxDirections").Return(8, true).Times(1),
					mockObj.EXPECT().GetProperty("velocity").Return([]int{1}, true).Times(1),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case: failed to get maxDirection property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				gomock.InOrder(
					mockObj.EXPECT().GetProperty("direction").Return(4, true).Times(1),
					mockObj.EXPECT().GetProperty("maxDirections").Return(0, false).Times(1),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case: object returns invalid maxDirections property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				gomock.InOrder(
					mockObj.EXPECT().GetProperty("direction").Return(4, true).Times(1),
					mockObj.EXPECT().GetProperty("maxDirections").Return([]int{1}, true).Times(1),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case: failed to get direction property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("direction").Return(4, false).Times(1)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case: object returns invalid direction property",
			mockCalls: func(mockObj *mocks.MockUObject) {
				mockObj.EXPECT().GetProperty("direction").Return([]int{4, 1}, true).Times(1)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			m := adapters.MovableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(mockObj)

			got, err := m.GetVelocity()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVelocity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVelocity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovableAdapter_SetPosition(t *testing.T) {
	type args struct {
		position pkg.Vector
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
				position: pkg.Vector{8, 1},
			},
			mockCalls: func(a args, mockObj *mocks.MockUObject) {
				mockObj.EXPECT().SetProperty("position", a.position).Return(true).Times(1)
			},
			wantErr: false,
		},
		{
			name: "case: failed to set new position",
			args: args{
				position: pkg.Vector{8, 1},
			},
			mockCalls: func(a args, mockObj *mocks.MockUObject) {
				mockObj.EXPECT().SetProperty("position", a.position).Return(false).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockObj := mocks.NewMockUObject(gomock.NewController(t))
			r := adapters.MovableAdapter{
				Obj: mockObj,
			}
			tt.mockCalls(tt.args, mockObj)

			if err := r.SetPosition(tt.args.position); (err != nil) != tt.wantErr {
				t.Errorf("SetPosition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
