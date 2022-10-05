package errhandler //nolint:testpackage

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/dosovma/otus_arch/pkg"
	"github.com/dosovma/otus_arch/pkg/mocks"
)

// TestErrorHandler_logAction тест обработчика исключения,
// который ставит Команду, пишущую в лог в очередь Команд.
func TestErrorHandler_logAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQueue := mocks.NewMockIQueue(ctrl)
	mockLog := mocks.NewMockLogger(ctrl)
	h := ErrorHandler{
		queue: mockQueue,
		log:   mockLog,
	}

	type args struct {
		e   pkg.Executable
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(a args)
	}{
		{
			name: "success",
			args: args{
				e:   commands.BaseCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewLogCmd(a.e, a.err, mockLog),
					).
					Return().
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCalls(tt.args)
			h.logAction(tt.args.e, tt.args.err)
		})
	}
}

// TestErrorHandler_firstRepeatAction тест обработчика исключения,
// который ставит в очередь Команду - повторитель команды, выбросившей исключение.
func TestErrorHandler_firstRepeatAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQueue := mocks.NewMockIQueue(ctrl)
	mockLog := mocks.NewMockLogger(ctrl)
	h := ErrorHandler{
		queue: mockQueue,
		log:   mockLog,
	}

	type args struct {
		e   pkg.Executable
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(a args)
	}{
		{
			name: "success",
			args: args{
				e:   commands.BaseCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewFirstRepeatCmd(a.e, a.err),
					).
					Return().
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCalls(tt.args)
			h.firstRepeatAction(tt.args.e, tt.args.err)
		})
	}
}

// TestErrorHandler_firstRepeatAction тест обработчика исключения,
// который повторяет команду два раза.
func TestErrorHandler_secondRepeatAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQueue := mocks.NewMockIQueue(ctrl)
	mockLog := mocks.NewMockLogger(ctrl)
	h := ErrorHandler{
		queue: mockQueue,
		log:   mockLog,
	}

	type args struct {
		e   pkg.Executable
		err error
	}
	tests := []struct {
		name      string
		args      args
		mockCalls func(a args)
	}{
		{
			name: "success",
			args: args{
				e:   commands.BaseCmd{},
				err: commands.ErrConnectionTimeout,
			},
			mockCalls: func(a args) {
				mockQueue.
					EXPECT().
					Push(
						commands.NewSecondCmd(a.e, a.err),
					).
					Return().
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCalls(tt.args)
			h.secondRepeatAction(tt.args.e, tt.args.err)
		})
	}
}
