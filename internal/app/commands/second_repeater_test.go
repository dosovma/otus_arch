package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/pkg/mocks"
)

// TestSecondRepeatCmd_Execute тест команды, которая второй раз повторяет Команду, выбросившую исключение.
func TestSecondRepeatCmd_Execute(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name      string
		fields    fields
		mockCalls func(*mocks.MockExecutable)
		wantErr   bool
	}{
		{
			name: "success",
			fields: fields{
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(mockExecutable *mocks.MockExecutable) {
				mockExecutable.EXPECT().Execute().Return(nil).Times(1)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockExecute := mocks.NewMockExecutable(ctrl)
			f := commands.NewSecondCmd(mockExecute, tt.fields.err)
			tt.mockCalls(mockExecute)

			if err := f.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
