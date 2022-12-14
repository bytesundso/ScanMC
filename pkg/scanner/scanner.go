package scanner

import(
	"github.com/bytesundso/ScanMC/internal/file"
	"github.com/bytesundso/ScanMC/internal/mc"
	"github.com/xrjr/mcutils/pkg/ping"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"fmt"
	"net"
	"io"
	"strconv"
	"context"
)

type ScanResults struct {
	Scanned int64
	Found int64
	Time time.Time
}

type serverInfo struct {
	host string 
	port int
	slp *ping.JSON
	ping int
}

type ServerEntry struct {
	Id primitive.ObjectID `bson:"_id"` 
	Time time.Time `bson:"date"`
	Host string `bson:"host"`
	Port int `bson:"port"`
	Ping int `bson:"ping"`
	CurrentPlayers int `bson:"current_players"`
	MaxPlayers int `bson:"max_players"`
	VersionName string `bson:"version_name"`
	VersionProtocol int `bson:"version_protocol"`
	Modt string `bson:"modt"`
}

func Scan(filename string, port int, threads int, timeout time.Duration) (*ScanResults, error) {
	hosts, err := file.LoadFile(filename)
	if err != nil {
		fmt.Println("Cannot load file")
		return nil, err
	}

	ch := make(chan *serverInfo)

	var collection *mongo.Collection
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb+srv://doadmin:gZD6274zY18Kxf05@db-mongodb-fra1-10276-75144107.mongo.ondigitalocean.com/admin?authSource=admin&replicaSet=db-mongodb-fra1-10276&tls=true")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Cannot connect to Database")
		return nil, err
	}
	
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Cannot ping Database")
		return nil, err
	}

	collection = client.Database("MinecraftServer").Collection("Scan-" + time.Now().Format("01-02-2006 15:04:05"))

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	for i := 0; i < threads; i++ {
		go scan(hosts, ch, port, timeout)
	}
	_ = hosts

	for true {
		server := <- ch

		if server != nil {
			info := server.slp.Infos()
			_, err = collection.InsertOne(ctx, ServerEntry { primitive.NewObjectID(), time.Now(), server.host, server.port, server.ping, info.Players.Online, info.Players.Max, info.Version.Name, info.Version.Protocol, mc.Motd(*server.slp) })
			if err != nil {
				fmt.Println(err)
			}
		} else {
			threads--
			if threads == 0 {
				break;
			}
		}
	}

	return &ScanResults { 0, 0, time.Time { } }, nil
}

func scan(hosts *file.File, ch chan *serverInfo, port int, timeout time.Duration) {
	for host, err := file.ReadNextLine(hosts); err != io.EOF; host, err = file.ReadNextLine(hosts) {
		address := *host + ":" + strconv.Itoa(port)

		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			continue
		}

		prop, ping, err := mc.Ping(*host, port, timeout)
		if err == nil {
			ch <- &serverInfo { *host, port, &prop, ping }
		}

		conn.Close()
	}

	ch <- nil
}