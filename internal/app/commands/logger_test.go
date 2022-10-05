package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/pkg"
	"github.com/dosovma/otus_arch/pkg/mocks"
)

// TestLogCmd_Execute тест команды, которая записывает информацию о выброшенном исключении в лог.
func TestLogCmd_Execute(t *testing.T) {
	type fields struct {
		ex  pkg.Executable
		err error
	}
	tests := []struct {
		name      string
		fields    fields
		mockCalls func(mockLog *mocks.MockLogger)
		wantErr   bool
	}{
		{
			name: "success",
			fields: fields{
				ex:  commands.BaseCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(mockLog *mocks.MockLogger) {
				mockLog.
					EXPECT().
					Error("Command %s throw the error %s", "BaseCmd", "connection timeout").
					Return().
					Times(1)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockLogger := mocks.NewMockLogger(ctrl)
			l := commands.NewLogCmd(tt.fields.ex, tt.fields.err, mockLogger)
			tt.mockCalls(mockLogger)

			if err := l.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
