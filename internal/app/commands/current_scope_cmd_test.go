package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/entity"
	"github.com/dosovma/otus_arch/internal/app/entity/mocks"
)

func TestCurrentScopeCmd_Execute(t *testing.T) {
	type fields struct {
		scopeName string
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
				scopeName: "default",
			},
			mockCalls: func(mockTL *mocks.MockIThreadLocal) {
				scope := make(map[string]func(obj entity.UObject) entity.Executable)
				mockTL.EXPECT().GetValue("default").Return(scope, true).Times(1)
				mockTL.EXPECT().SetCurrentScope(scope).Return().Times(1)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockTL := mocks.NewMockIThreadLocal(ctrl)
			tt.mockCalls(mockTL)

			s := commands.NewCurrentScopeCmd(mockTL, tt.fields.scopeName)
			if err := s.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
