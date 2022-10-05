package errhandler //nolint:testpackage

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/pkg"
	"github.com/dosovma/otus_arch/pkg/mocks"
)

// TestErrorHandler_Handle_NewOneRepeatStrategic тест стратегии:
// при первом выбросе исключения повторить команду, при повторном выбросе исключения записать информацию в лог.
func TestErrorHandler_Handle_NewOneRepeatStrategic(t *testing.T) {
	type args struct {
		cmd pkg.Executable
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger)
	}{
		{
			name: "success: handle baseCmd => repeat it",
			args: args{
				cmd: commands.BaseCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewFirstRepeatCmd(a.cmd, a.err),
					).
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle already repeated command => log it",
			args: args{
				cmd: commands.FirstRepeatCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewLogCmd(a.cmd, a.err, mockLogger),
					).
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle log command => log unknown commands",
			args: args{
				cmd: commands.LogCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockLogger.
					EXPECT().
					Error("There is no action or error for command %s", "LogCmd").
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle unexpected command => log unknown commands",
			args: args{
				cmd: commands.Move{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockLogger.
					EXPECT().
					Error("There is no action or error for command %s", "Move").
					Return().
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockQueue := mocks.NewMockIQueue(ctrl)
			mockLog := mocks.NewMockLogger(ctrl)
			h := NewOneRepeatStrategic(mockQueue, mockLog)
			tt.mockCalls(tt.args, mockQueue, mockLog)

			h.Handle(tt.args.cmd, tt.args.err)
		})
	}
}

// TestErrorHandler_Handle_NewTwoRepeatStrategic тест стратегии:
// повторить два раза, потом записать в лог.
func TestErrorHandler_Handle_NewTwoRepeatStrategic(t *testing.T) {
	type args struct {
		cmd pkg.Executable
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger)
	}{
		{
			name: "success: handle baseCmd => repeat it",
			args: args{
				cmd: commands.BaseCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewFirstRepeatCmd(a.cmd, a.err),
					).
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle already repeated command => repeat it again",
			args: args{
				cmd: commands.FirstRepeatCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewSecondCmd(a.cmd, a.err),
					).
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle already twice repeated command => log it",
			args: args{
				cmd: commands.SecondRepeatCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewLogCmd(a.cmd, a.err, mockLogger),
					).
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle log command => log unknown commands",
			args: args{
				cmd: commands.LogCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockLogger.
					EXPECT().
					Error("There is no action or error for command %s", "LogCmd").
					Return().
					Times(1)
			},
		},
		{
			name: "success: handle unexpected command => log unknown commands",
			args: args{
				cmd: commands.Move{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args, mockQueue *mocks.MockIQueue, mockLogger *mocks.MockLogger) {
				mockLogger.
					EXPECT().
					Error("There is no action or error for command %s", "Move").
					Return().
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockQueue := mocks.NewMockIQueue(ctrl)
			mockLog := mocks.NewMockLogger(ctrl)
			h := NewTwoRepeatStrategic(mockQueue, mockLog)
			tt.mockCalls(tt.args, mockQueue, mockLog)

			h.Handle(tt.args.cmd, tt.args.err)
		})
	}
}
