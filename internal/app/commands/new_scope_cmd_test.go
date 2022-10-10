package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/internal/app/entity/mocks"
)

func TestNewScopeCmd_Execute(t *testing.T) {
	type fields struct {
		scopeName string
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
				scopeName: "new scope",
			},
			mockCalls: func(mockTL *mocks.MockIThreadLocal) {
				mockTL.EXPECT().SetValue("new scope", gomock.Any()).Return().Times(1)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockTL := mocks.NewMockIThreadLocal(ctrl)
			tt.mockCalls(mockTL)

			r := commands.NewNewScopeCmd(mockTL, tt.fields.scopeName)
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
