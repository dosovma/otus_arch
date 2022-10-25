package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/entity"
	"github.com/dosovma/otus_arch/internal/app/entity/mocks"
)

func TestRegisterCmd_Execute(t *testing.T) {
	type fields struct {
		cmdName string
		builder func(obj entity.UObject) entity.Executable
	}
	tests := []struct {
		name      string
		fields    fields
		mockCalls func(*mocks.MockIThreadLocal)
		wantErr   bool
	}{
		{
			name: "case: success",
			fields: fields{
				cmdName: "Commands.Move",
				builder: func(obj entity.UObject) entity.Executable {
					return commands.Move{
						Move: adapters.MovableAdapter{
							Obj: obj,
						},
					}
				},
			},
			mockCalls: func(mockTL *mocks.MockIThreadLocal) {
				scope := make(map[string]func(obj entity.UObject) entity.Executable)

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

			r := commands.NewRegisterCmd(mockTL, tt.fields.cmdName, tt.fields.builder)
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
