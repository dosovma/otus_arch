package commands

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/dosovma/otus_arch/internal/app/adapters"
	"github.com/dosovma/otus_arch/internal/app/entity"
	"github.com/dosovma/otus_arch/internal/app/entity/mocks"
)

func TestIoC_Resolve(t *testing.T) {
	t.Parallel()

	mockTL := mocks.NewMockIThreadLocal(gomock.NewController(t))
	moveBuilder := func(obj entity.UObject) entity.Executable {
		return Move{
			Move: adapters.MovableAdapter{
				Obj: obj,
			},
		}
	}
	rotateBuilder := func(obj entity.UObject) entity.Executable {
		return Rotate{
			Rotate: adapters.RotatableAdapter{
				Obj: obj,
			},
		}
	}

	type args struct {
		key  string
		args []any
	}
	tests := []struct {
		name string
		args args
		want entity.Executable
	}{
		{
			name: "case: success: register new Move command in current scope",
			args: args{
				key: "IoC.Register",
				args: []any{
					"Commands.Move",
					moveBuilder,
				},
			},
			want: RegisterCmd{
				tl:      mockTL,
				cmdName: "Commands.Move",
				b:       moveBuilder,
			},
		},
		{
			name: "case: success: register new Rotate command in current scope",
			args: args{
				key: "IoC.Register",
				args: []any{
					"Commands.Rotate",
					rotateBuilder,
				},
			},
			want: RegisterCmd{
				tl:      mockTL,
				cmdName: "Commands.Rotate",
				b:       rotateBuilder,
			},
		},
		{
			name: "case: success: create new Scope",
			args: args{
				key:  "Scopes.New",
				args: []any{"new scope"},
			},
			want: NewScopeCmd{
				tl:        mockTL,
				scopeName: "new scope",
			},
		},
		{
			name: "case: success: set new current scope",
			args: args{
				key:  "Scopes.Current",
				args: []any{"scope name"},
			},
			want: CurrentScopeCmd{
				tl:        mockTL,
				scopeName: "scope name",
			},
		},
		{
			name: "case: success: register new Move command in new scope",
			args: args{
				key: "IoC.Register",
				args: []any{
					"Commands.Move",
					moveBuilder,
				},
			},
			want: RegisterCmd{
				tl:      mockTL,
				cmdName: "Commands.Move",
				b:       moveBuilder,
			},
		},
		{
			name: "case: success: register new Rotate command in new scope",
			args: args{
				key: "IoC.Register",
				args: []any{
					"Commands.Rotate",
					rotateBuilder,
				},
			},
			want: RegisterCmd{
				tl:      mockTL,
				cmdName: "Commands.Rotate",
				b:       rotateBuilder,
			},
		},
		{
			name: "case: success: set default current scope",
			args: args{
				key:  "Scopes.Current",
				args: []any{"default"},
			},
			want: CurrentScopeCmd{
				tl:        mockTL,
				scopeName: "default",
			},
		},
		{
			name: "case: success: get builder for Move command",
			args: args{
				key:  "Commands.Move",
				args: []any{entity.Object{}},
			},
			want: GetBuilderCmd{
				tl:      mockTL,
				cmdName: "Commands.Move",
				obj:     entity.Object{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt := tt

			i := IoC{
				threadLocal: mockTL,
			}
			i.initActions()
			got := i.Resolve(tt.args.key, tt.args.args...)
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
