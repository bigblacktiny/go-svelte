package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"goTemp/globalUtils"
	"goTemp/promotion/proto"
	"log"
	"os"
	"strconv"
)

//serviceName: service identifier
const serviceName = "promotion"

const (
	//dbName: Name of the DB hosting the data
	dbName = "postgres"
	//dbConStrEnvVarName: Name of Environment variable that contains connection string to DB
	dbConStrEnvVarName = "POSTGRES_CONNECT"
)

// conn: Database connection
var conn *pgx.Conn

//AuthWrapper: Authentication middleware
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		//User login is excepted from authentication
		//if req.Endpoint() == "UserSrv.Auth" {
		//	return fn(ctx, req, resp)
		//}
		//
		//meta, ok := metadata.FromContext(ctx)
		//if !ok {
		//	return fmt.Errorf(glErr.AuthNoMetaData(req.Endpoint()))
		//}
		//
		//token := meta["Token"]
		//log.Printf("endpoint: %v", req.Endpoint())
		//
		////Validate token
		//var u User
		//outToken := &pb.Token{}
		//err := u.ValidateToken(ctx, &pb.Token{Token: token}, outToken)
		//if err != nil {
		//	return err
		//}
		//if outToken.Valid != true {
		//	return fmt.Errorf(glErr.AuthInvalidToken())
		//}
		//
		////Add current user to context to use in saving audit records
		//userId, err := u.userIdFromToken(ctx, outToken)
		//if err != nil {
		//	return fmt.Errorf("unable to get user id from token for endpoint %v\n", req.Endpoint())
		//}
		//ctx2 := metadata.Set(ctx, "userid", strconv.FormatInt(userId, 10))
		var userId int64
		userId = 2322853872944550913
		ctx2 := metadata.Set(ctx, "userid", strconv.FormatInt(userId, 10))

		return fn(ctx2, req, resp)
	}
}

//getDBConnString: Get the connection string to the DB
func getDBConnString() string {
	connString := os.Getenv(dbConStrEnvVarName)
	if connString == "" {
		log.Fatalf(glErr.DbNoConnectionString(dbConStrEnvVarName))
	}
	return connString
}

//connectToDB: Call the Util pgxDBConnect to connect to the database. Service will try to connect a few times
//before giving up and throwing an error
func connectToDB() *pgx.Conn {
	var pgxConnect globalUtils.PgxDBConnect
	dbConn, err := pgxConnect.ConnectToDBWithRetry(dbName, getDBConnString())
	if err != nil {
		log.Fatalf(glErr.DbNoConnection(dbName, err))
	}
	return dbConn
}

func main() {

	//instantiate service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(AuthWrapper),
	)

	service.Init()
	err := proto.RegisterPromotionSrvHandler(service.Server(), new(Promotion))
	if err != nil {
		log.Fatalf(glErr.SrvNoHandler(err))
	}

	//Connect to DB
	conn = connectToDB()
	defer conn.Close(context.Background())

	//setup the nats broker
	mb.Br = service.Options().Broker

	// Run Service
	if err := service.Run(); err != nil {
		log.Fatalf(glErr.SrvNoStart(serviceName, err))
	}

}
