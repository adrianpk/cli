package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	gen "gitlab.com/mikrowezel/backend/cli/internal/generator"
	"gitlab.com/mikrowezel/backend/cli/internal/inflector"
)

var (
	pkg       string
	pkgDir    = ""
	name      string
	all       bool
	yaml      bool
	migration bool
	model     bool
	repo      bool
	grpc      bool
	jsonrest  bool
	service   bool
	transport bool
	web       bool
	restcl    bool
	force     bool
	metadata  *gen.Metadata
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&pkg, "pkg", "p", "", pkgUsg)
	generateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", nameUsg)
	generateCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, allUsg)
	generateCmd.PersistentFlags().BoolVarP(&yaml, "yaml", "y", false, yamlUsg)
	generateCmd.PersistentFlags().BoolVarP(&migration, "migration", "g", false, migrationUsg)
	generateCmd.PersistentFlags().BoolVarP(&model, "model", "m", false, modelUsg)
	generateCmd.PersistentFlags().BoolVarP(&repo, "repo", "r", false, repoUsg)
	generateCmd.PersistentFlags().BoolVarP(&grpc, "grpc", "", false, grpcUsg)
	generateCmd.PersistentFlags().BoolVarP(&jsonrest, "jsonrest", "j", false, jsonrestUsg)
	generateCmd.PersistentFlags().BoolVarP(&service, "service", "s", false, serviceUsg)
	generateCmd.PersistentFlags().BoolVarP(&transport, "transport", "t", false, transportUsg)
	generateCmd.PersistentFlags().BoolVarP(&web, "web", "w", false, webUsg)
	generateCmd.PersistentFlags().BoolVarP(&restcl, "restcl", "", false, webUsg)
	generateCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, forceUsg)
}

// generateCmd command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate mikrowezel resources.",
	Long:  `mw generate <ModelName> [-p pkgName] [--all] [--yaml] [--migration] [--model] [--repo] [--grpc] [--jsonrest] [--service] [-transport] [--web] [--restcl] [--force]`,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Starting...")

		res := inflector.ToSnakeCase(args[0])

		if yaml {
			gen.GenBaseYAML(res, force)
			return
		}

		g, err := gen.MakeGenerator(res, pkg, force)
		if err != nil {
			log.Fatal(err.Error())
		}

		if migration || all {
			//g.GenMigration()
		}

		if model || all {
			//g.GenModel()
		}

		if repo || all {
			g.GenRepo()
		}

		if grpc || all {
			// g.GenGRPC()
		}

		if jsonrest || all {
			// g.GenJSONREST()
		}

		if service || all {
			// g.GenService()
		}

		if transport || all {
			// g.GenTransport()
		}

		if restcl || all {
			// g.GenRESTClient()
		}
	},
}

func projectRootDir() (dir string, err error) {
	return os.Getwd()
}

func errMsg(err error) string {
	return strings.Title(strings.ToLower(err.Error()))
}
