package cmd

import (
	"log"
	"os"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/sunlggggg/katyusha/net/http"
	"github.com/sunlggggg/katyusha/sys"
)

var rootCmd = &cobra.Command{
  Use:   "katyusha",
  Short: "Katyusha is a Network Disk.",
  Long: `This is Katyusha`,
  Run: func(cmd *cobra.Command, args []string) {
	  // main entry
	  log.Println("Hello, I am Katyusha.")
	  log.Println("Katyusha's size is " + strconv.FormatFloat(float64(sys.BinSize("katyusha")) / 1024 / 1024, 'f', 2, 64))
	  // start http server
	  http.Server()

  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Println(os.Stderr, err)
    os.Exit(1)
  }
}
