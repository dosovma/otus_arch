package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/entity"
	"github.com/dosovma/otus_arch/internal/app/entity/mocks"
)

func TestGetBuilderCmd_Execute(t *testing.T) {
	type fields struct {
		cmdName string
		obj     entity.UObject
	}
	tests := []struct {
		name      string
		fields    fields
		mockCalls func(mockTL *mocks.MockIThreadLocal)
		wantErr   bool
	}{
		{
			name: "case: success",
			fields: fields{
				cmdName: "Commands.Move",
				obj:     entity.Object{},
			},
			mockCalls: func(mockTL *mocks.MockIThreadLocal) {
				scope := make(map[string]func(obj entity.UObject) entity.Executable)
				scope["Commands.Move"] = func(obj entity.UObject) entity.Executable {
					return commands.Rotate{
						Rotate: adapters.RotatableAdapter{
							Obj: obj,
						},
					}
				}
				mockTL.EXPECT().GetCurrentScope().Return(scope).Times(1)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockTL := mocks.NewMockIThreadLocal(ctrl)
			tt.mockCalls(mockTL)

			r := commands.NewGetBuilderCmd(mockTL, tt.fields.cmdName, tt.fields.obj)
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
