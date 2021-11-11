package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime/client"
	ttools "github.com/intel/tfortools"
	"github.com/raphaelreyna/shrtct/internal/generated/models"
	"github.com/raphaelreyna/shrtct/internal/generated/shortcut"
	"github.com/raphaelreyna/shrtct/internal/generated/shortcut/operations"
	"github.com/spf13/cobra"
)

func init() {
	searchCmd.AddCommand(searchStoryCmd())
	searchCmd.AddCommand(searchEpicCmd)
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search stories or epics",
	Run:   nil,
}

func searchStoryCmd() *cobra.Command {
	var (
		cc = cobra.Command{
			Use: "stories",
		}

		flags             = cc.Flags()
		query             string
		project           string
		stype             string
		estimate          string
		state             string
		blocked, blocking bool
	)

	flags.StringVarP(&query, "query", "q", "", "Finds all stories matching the general search query")
	flags.StringVarP(&project, "project", "p", "", "Finds all stories of a specific project.")
	flags.StringVarP(&stype, "type", "t", "", "Finds all Stories of a specific type (feature, bug, or chore).")
	flags.StringVarP(&estimate, "estimate", "e", "", "Finds all Stories of a specific point value.")
	flags.StringVarP(&state, "state", "s", "", "Finds all Stories in a specific state (\"In Review\")")
	flags.BoolVarP(&blocked, "blocked", "b", false, "Finds all Stories that are blocked.")
	flags.BoolVarP(&blocking, "blocking", "B", false, "Finds all Stories that are blocking.")

	cc.Run = func(cmd *cobra.Command, args []string) {
		var (
			ctx = cmd.Context()

			clnt = shortcut.Default
			ops  = clnt.Operations
			auth = client.APIKeyAuth(
				"Shortcut-Token",
				"header",
				os.Getenv("SHORTCUT_TOKEN"),
			)
		)

		if err := flags.Parse(args); err != nil {
			log.Fatal(err)
		}

		var qparts = make([]string, 0)
		if query != "" {
			qparts = append(qparts, query)
		}
		if project != "" {
			qparts = append(qparts, fmt.Sprintf("project:%q", project))
		}
		if stype != "" {
			qparts = append(qparts, fmt.Sprintf("type:%q", stype))
		}
		if state != "" {
			qparts = append(qparts, fmt.Sprintf("state:%q", state))
		}
		if blocked {
			qparts = append(qparts, "is:blocked")
		}
		if blocking {
			qparts = append(qparts, "is:blocking")
		}
		if estimate != "" {
			_, err := strconv.Atoi(estimate)
			if err != nil {
				log.Fatal(err)
			}

			qparts = append(qparts, fmt.Sprintf("estimate:%s", estimate))
		}

		var q = strings.Join(qparts, " AND ")
		ssok, err := ops.SearchStories(&operations.SearchStoriesParams{
			Context: ctx,
			Search: &models.Search{
				Query: &q,
			},
		}, auth)
		if err != nil {
			log.Fatal(err)
		}

		var tmplt = flags.Arg(0)
		if tmplt != "" {
			var searchResults []interface{}

			data, err := json.Marshal(ssok.GetPayload().Data)
			if err != nil {
				log.Fatal(err)
			}

			if err := json.Unmarshal(data, &searchResults); err != nil {
				log.Fatal(err)
			}
			err = ttools.OutputToTemplate(
				os.Stdout,
				"stories",
				tmplt,
				searchResults,
				ttools.NewConfig(ttools.OptAllFns),
			)
			if err != nil {
				log.Fatal(err)
			}

			return
		}

		var enc = json.NewEncoder(os.Stdout)
		enc.SetIndent("", "\t")
		enc.Encode(ssok.GetPayload().Data)
	}

	return &cc
}

var searchEpicCmd = &cobra.Command{
	Use: "epics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unimplemented")
	},
}
