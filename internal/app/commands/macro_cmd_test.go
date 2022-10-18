package commands_test

import (
	"errors"
	"testing"

	"github.com/dosovma/otus_arch/internal/app/commands"
	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/commands/mocks"
)

func TestMoveAndBurnFuelCmd_Execute(t *testing.T) {
	tests := []struct {
		name      string
		mockCalls func(mockCommands ...*mocks.MockExecutable)
		wantErr   bool
	}{
		{
			name: "case: success: macro command for move and burn fuel succeeded",
			mockCalls: func(mockCommands ...*mocks.MockExecutable) {
				for _, cmd := range mockCommands {
					cmd.EXPECT().Execute().Return(nil)
				}
			},
			wantErr: false,
		},
		{
			name: "case: one of command return an error",
			mockCalls: func(mockCommands ...*mocks.MockExecutable) {
				mockCommands[0].EXPECT().Execute().Return(errors.New("error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockExecutable1 := mocks.NewMockExecutable(gomock.NewController(t))
			mockExecutable2 := mocks.NewMockExecutable(gomock.NewController(t))
			moveAndBurn := commands.MacroCmd{
				Commands: []commands.Executable{mockExecutable1, mockExecutable2},
			}
			tt.mockCalls(mockExecutable1, mockExecutable2)

			if err := moveAndBurn.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
