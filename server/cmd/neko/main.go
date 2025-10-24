package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	neko "github.com/HiradNikoo/neko/server"
	"github.com/HiradNikoo/neko/server/cmd"
	"github.com/HiradNikoo/neko/server/pkg/utils"
)

func main() {
	fmt.Print(utils.Colorf(neko.Header, "server", neko.Version))
	if err := cmd.Execute(); err != nil {
		log.Panic().Err(err).Msg("failed to execute command")
	}
}
