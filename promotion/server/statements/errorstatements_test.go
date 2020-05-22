package statements

import (
	"fmt"
	"testing"
)

func Test_promoErr_internalError(t *testing.T) {
	var prmerr PromoErr
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		ge      PromoErr
		args    args
		want    string
		lang    Languages
		setLang bool
	}{
		{name: "Pull English internal error", ge: prmerr, args: args{err: fmt.Errorf("test Error")}, want: "Internal error. Error: test Error\n",
			lang: en, setLang: false,
		},
		{name: "Pull English internal error", ge: prmerr, args: args{err: fmt.Errorf("error prueba")}, want: "Error interno. Error: error prueba\n",
			lang: es, setLang: true,
		},
		{name: "Pull English internal error", ge: prmerr, args: args{err: fmt.Errorf("test Error")}, want: "Internal error. Error: test Error\n",
			lang: en, setLang: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setLang {
				SetLanguage(tt.lang)
			}
			if got := tt.ge.internalError(tt.args.err); got != tt.want {
				t.Errorf("InsertError() = %v, want %v for language %v", got, tt.want, tt.lang)
			}
		})
	}
}

func Test_promoErr_getSqlTxt(t *testing.T) {
	var prmerr PromoErr
	type args struct {
		errKey   string
		language Languages
	}
	tests := []struct {
		name string
		ge   PromoErr
		args args
		want string
	}{
		{name: "Pull English internal error", ge: prmerr, args: args{errKey: "internalError", language: en}, want: "Internal error. Error: %v\n"},
		{name: "Pull Spanish internal error", ge: prmerr, args: args{errKey: "internalError", language: es}, want: "Error interno. Error: %v\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ge.getSqlTxt(tt.args.errKey, tt.args.language); got != tt.want {
				t.Errorf("getSqlTxt() = %v, want %v", got, tt.want)
			}
		})
	}
}
