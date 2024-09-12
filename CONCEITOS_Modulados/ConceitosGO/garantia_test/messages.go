package garantia_test

import (
	bt "github.com/reizzao/RzLibs_GO/message"
)

// CRIAR PARA ESTE O ARQUIVO REEXPORTADOS DE LIBS EXTERNAS / exemplo: lib/messages.go

var (
	MessageErrorTest func(compare any, expect any) string = bt.MessageErrorTestLIB
)
