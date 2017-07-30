// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: d5b3153b9f
// Version Date: Thu Jul 27 18:20:46 UTC 2017

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pkg/errors"

	// This Service
	pb "github.com/myambition/ambition/services/model/model-service"
	"github.com/myambition/ambition/services/model/model-service/svc/client/cli/handlers"
	grpcclient "github.com/myambition/ambition/services/model/model-service/svc/client/grpc"
	httpclient "github.com/myambition/ambition/services/model/model-service/svc/client/http"
)

var (
	_ = strconv.ParseInt
	_ = strings.Split
	_ = json.Compact
	_ = errors.Wrapf
	_ = pb.RegisterModelServer
)

func main() {
	os.Exit(submain())
}

type headerSeries []string

func (h *headerSeries) Set(val string) error {
	const requiredParts int = 2
	parts := strings.SplitN(val, ":", requiredParts)
	if len(parts) != requiredParts {
		return fmt.Errorf("value %q cannot be split in two; must contain at least one ':' character", val)
	}
	parts[1] = strings.TrimSpace(parts[1])
	*h = append(*h, parts...)
	return nil
}

func (h *headerSeries) String() string {
	return fmt.Sprintf("%v", []string(*h))
}

// submain exists to act as the functional main, but will return exit codes to
// the actual main instead of calling os.Exit directly. This is done to allow
// the defered functions to be called, since if os.Exit where called directly
// from this function, none of the defered functions set up by this function
// would be called.
func submain() int {
	var headers headerSeries
	flag.Var(&headers, "header", "Header(s) to be sent in the transport (follows cURL style)")
	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", ":5040", "gRPC (HTTP) address of addsvc")
	)

	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an service. This presumption is reflected in
	// the cli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings.

	fsCreateAction := flag.NewFlagSet("createaction", flag.ExitOnError)

	fsCreateOccurrence := flag.NewFlagSet("createoccurrence", flag.ExitOnError)

	fsReadAction := flag.NewFlagSet("readaction", flag.ExitOnError)

	fsReadActions := flag.NewFlagSet("readactions", flag.ExitOnError)

	fsReadOccurrences := flag.NewFlagSet("readoccurrences", flag.ExitOnError)

	fsReadOccurrencesByDate := flag.NewFlagSet("readoccurrencesbydate", flag.ExitOnError)

	var (
		flagIDReadAction                   = fsReadAction.Int64("id", 0, "")
		flagNameReadAction                 = fsReadAction.String("name", "", "")
		flagUserIDReadAction               = fsReadAction.Int64("userid", 0, "")
		flagUserIDReadActions              = fsReadActions.Int64("userid", 0, "")
		flagActionIDReadOccurrencesByDate  = fsReadOccurrencesByDate.Int64("actionid", 0, "")
		flagStartDateReadOccurrencesByDate = fsReadOccurrencesByDate.String("startdate", "", "")
		flagEndDateReadOccurrencesByDate   = fsReadOccurrencesByDate.String("enddate", "", "")
		flagIDReadOccurrences              = fsReadOccurrences.Int64("id", 0, "")
		flagNameReadOccurrences            = fsReadOccurrences.String("name", "", "")
		flagUserIDReadOccurrences          = fsReadOccurrences.Int64("userid", 0, "")
		flagIDCreateAction                 = fsCreateAction.Int64("id", 0, "")
		flagNameCreateAction               = fsCreateAction.String("name", "", "")
		flagUserIDCreateAction             = fsCreateAction.Int64("userid", 0, "")
		flagUserIDCreateOccurrence         = fsCreateOccurrence.Int64("userid", 0, "")
		flagOccurrenceCreateOccurrence     = fsCreateOccurrence.String("occurrence", "", "")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Subcommands:\n")
		fmt.Fprintf(os.Stderr, "  %s\n", "createaction")
		fmt.Fprintf(os.Stderr, "  %s\n", "createoccurrence")
		fmt.Fprintf(os.Stderr, "  %s\n", "readaction")
		fmt.Fprintf(os.Stderr, "  %s\n", "readactions")
		fmt.Fprintf(os.Stderr, "  %s\n", "readoccurrences")
		fmt.Fprintf(os.Stderr, "  %s\n", "readoccurrencesbydate")
	}
	if len(os.Args) < 2 {
		flag.Usage()
		return 1
	}

	flag.Parse()

	var (
		service pb.ModelServer
		err     error
	)

	if *httpAddr != "" {
		service, err = httpclient.New(*httpAddr, httpclient.CtxValuesToSend(headers...))
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
			return 1
		}
		defer conn.Close()
		service, err = grpcclient.New(conn, grpcclient.CtxValuesToSend(headers...))
	} else {
		fmt.Fprintf(os.Stderr, "error: no remote address specified\n")
		return 1
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("No 'method' subcommand provided; exiting.")
		flag.Usage()
		return 1
	}

	ctx := context.Background()
	for i := 0; i < len(headers); i += 2 {
		ctx = context.WithValue(ctx, headers[i], headers[i+1])
	}

	switch flag.Args()[0] {

	case "createaction":
		fsCreateAction.Parse(flag.Args()[1:])

		IDCreateAction := *flagIDCreateAction
		NameCreateAction := *flagNameCreateAction
		UserIDCreateAction := *flagUserIDCreateAction

		request, err := handlers.CreateAction(IDCreateAction, NameCreateAction, UserIDCreateAction)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.CreateAction: %v\n", err)
			return 1
		}

		v, err := service.CreateAction(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.CreateAction: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDCreateAction, NameCreateAction, UserIDCreateAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "createoccurrence":
		fsCreateOccurrence.Parse(flag.Args()[1:])

		UserIDCreateOccurrence := *flagUserIDCreateOccurrence

		var OccurrenceCreateOccurrence pb.Occurrence
		if flagOccurrenceCreateOccurrence != nil && len(*flagOccurrenceCreateOccurrence) > 0 {
			err = json.Unmarshal([]byte(*flagOccurrenceCreateOccurrence), &OccurrenceCreateOccurrence)
			if err != nil {
				panic(errors.Wrapf(err, "unmarshalling OccurrenceCreateOccurrence from %v:", flagOccurrenceCreateOccurrence))
			}
		}

		request, err := handlers.CreateOccurrence(UserIDCreateOccurrence, OccurrenceCreateOccurrence)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.CreateOccurrence: %v\n", err)
			return 1
		}

		v, err := service.CreateOccurrence(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.CreateOccurrence: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(UserIDCreateOccurrence, OccurrenceCreateOccurrence)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "readaction":
		fsReadAction.Parse(flag.Args()[1:])

		IDReadAction := *flagIDReadAction
		NameReadAction := *flagNameReadAction
		UserIDReadAction := *flagUserIDReadAction

		request, err := handlers.ReadAction(IDReadAction, NameReadAction, UserIDReadAction)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.ReadAction: %v\n", err)
			return 1
		}

		v, err := service.ReadAction(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadAction: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDReadAction, NameReadAction, UserIDReadAction)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "readactions":
		fsReadActions.Parse(flag.Args()[1:])

		UserIDReadActions := *flagUserIDReadActions

		request, err := handlers.ReadActions(UserIDReadActions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.ReadActions: %v\n", err)
			return 1
		}

		v, err := service.ReadActions(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadActions: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(UserIDReadActions)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "readoccurrences":
		fsReadOccurrences.Parse(flag.Args()[1:])

		IDReadOccurrences := *flagIDReadOccurrences
		NameReadOccurrences := *flagNameReadOccurrences
		UserIDReadOccurrences := *flagUserIDReadOccurrences

		request, err := handlers.ReadOccurrences(IDReadOccurrences, NameReadOccurrences, UserIDReadOccurrences)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.ReadOccurrences: %v\n", err)
			return 1
		}

		v, err := service.ReadOccurrences(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadOccurrences: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IDReadOccurrences, NameReadOccurrences, UserIDReadOccurrences)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "readoccurrencesbydate":
		fsReadOccurrencesByDate.Parse(flag.Args()[1:])

		ActionIDReadOccurrencesByDate := *flagActionIDReadOccurrencesByDate
		StartDateReadOccurrencesByDate := *flagStartDateReadOccurrencesByDate
		EndDateReadOccurrencesByDate := *flagEndDateReadOccurrencesByDate

		request, err := handlers.ReadOccurrencesByDate(ActionIDReadOccurrencesByDate, StartDateReadOccurrencesByDate, EndDateReadOccurrencesByDate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.ReadOccurrencesByDate: %v\n", err)
			return 1
		}

		v, err := service.ReadOccurrencesByDate(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.ReadOccurrencesByDate: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(ActionIDReadOccurrencesByDate, StartDateReadOccurrencesByDate, EndDateReadOccurrencesByDate)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	default:
		flag.Usage()
		return 1
	}

	return 0
}
