package commands

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/entity"
	"github.com/dosovma/otus_arch/internal/app/entity/mocks"
)

func TestIoC_registerAction(t *testing.T) {
	mockTL := mocks.NewMockIThreadLocal(gomock.NewController(t))
	builder := func(obj entity.UObject) entity.Executable {
		return Rotate{
			Rotate: adapters.RotatableAdapter{
				Obj: obj,
			},
		}
	}

	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want entity.Executable
	}{
		{
			name: "case: success",
			args: args{
				args: []any{
					"Command.Rotate",
					builder,
				},
			},
			want: RegisterCmd{
				tl:      mockTL,
				cmdName: "Command.Rotate",
				b:       builder,
			},
		},
		{
			name: "case: failed to get cmd name",
			args: args{
				args: []any{
					1234,
					builder,
				},
			},
			want: ErrorCmd{
				message: fmt.Sprintf("invalid command name %v for action %s", 1234, IoCRegister),
			},
		},
		{
			name: "case: failed to get builder",
			args: args{
				args: []any{
					"Command.Rotate",
					nil,
				},
			},
			want: ErrorCmd{
				message: fmt.Sprintf("invalid command b %v for action %s", nil, IoCRegister),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IoC{
				threadLocal: mockTL,
				actions:     make(map[string]func(args []any) entity.Executable),
			}
			got := i.registerAction(tt.args.args)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("registerAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIoC_newScopeAction(t *testing.T) {
	mockTL := mocks.NewMockIThreadLocal(gomock.NewController(t))

	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want entity.Executable
	}{
		{
			name: "case: success",
			args: args{
				args: []any{"new scope"},
			},
			want: NewScopeCmd{
				tl:        mockTL,
				scopeName: "new scope",
			},
		},
		{
			name: "case: failed to get scope name",
			args: args{
				args: []any{1234},
			},
			want: ErrorCmd{
				message: fmt.Sprintf("invalid scope name %v for action %s", 1234, ScopesNew),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IoC{
				threadLocal: mockTL,
				actions:     make(map[string]func(args []any) entity.Executable),
			}
			got := i.newScopeAction(tt.args.args)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("newScopeAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIoC_currentScopeAction(t *testing.T) {
	mockTL := mocks.NewMockIThreadLocal(gomock.NewController(t))

	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want entity.Executable
	}{
		{
			name: "case: success",
			args: args{
				args: []any{"scope name"},
			},
			want: CurrentScopeCmd{
				tl:        mockTL,
				scopeName: "scope name",
			},
		},
		{
			name: "case: failed to get scope name",
			args: args{
				args: []any{1234},
			},
			want: ErrorCmd{
				message: fmt.Sprintf("invalid scope name %v for action %s", 1234, ScopesCurrent),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IoC{
				threadLocal: mockTL,
				actions:     make(map[string]func(args []any) entity.Executable),
			}
			got := i.currentScopeAction(tt.args.args)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("currentScopeAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIoC_getBuilderAction(t *testing.T) {
	mockTL := mocks.NewMockIThreadLocal(gomock.NewController(t))

	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want entity.Executable
	}{
		{
			name: "case: success",
			args: args{
				args: []any{"Command.Move", entity.Object{}},
			},
			want: GetBuilderCmd{
				tl:      mockTL,
				cmdName: "Command.Move",
				obj:     entity.Object{},
			},
		},
		{
			name: "case: failed to get cmd name",
			args: args{
				args: []any{1234, entity.Object{}},
			},
			want: ErrorCmd{
				message: fmt.Sprintf("invalid command name %v for action get command b by command name", 1234),
			},
		},
		{
			name: "case: failed to get Uobject",
			args: args{
				args: []any{"Command.Move", entity.ThreadLocal{}},
			},
			want: ErrorCmd{
				message: fmt.Sprintf(
					"invalid obj %v for action get command b by command name",
					entity.ThreadLocal{},
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IoC{
				threadLocal: mockTL,
				actions:     make(map[string]func(args []any) entity.Executable),
			}
			got := i.getBuilderAction(tt.args.args)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("getBuilderAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
