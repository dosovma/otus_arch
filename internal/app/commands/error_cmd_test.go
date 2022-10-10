package commands

import "testing"

func TestErrorCmd_Execute(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "case: success",
			fields: fields{
				message: "error message",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorCmd{
				message: tt.fields.message,
			}
			if err := e.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
