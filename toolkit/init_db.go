package toolkit

import (
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConn struct {
	user      string
	pwd       string
	protocol  string
	address   string
	port      string
	dbname    string
	charset   string
	parseTime string
}

func (c *mysqlConn) GetDSN() string {
	return fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=%v&parseTime=%v",
		c.user, c.pwd, c.protocol, c.address, c.port, c.dbname, c.charset, c.parseTime)
}

var DB *gorm.DB
var DGO *dgo.Dgraph

func InitDB() {
	dsn := (&mysqlConn{
		user:      "dyadmin",
		pwd:       "123456",
		protocol:  "tcp",
		address:   "127.0.0.1",
		dbname:    "douyin",
		port:      "3306",
		charset:   "utf8mb4",
		parseTime: "True",
	}).GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
func InitDGO() {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	DGO = dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
