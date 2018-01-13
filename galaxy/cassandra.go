package galaxy

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type cassandraRow struct {
	Id        int
	Data      []byte
	DataThumb []byte
	FileType  string
}

var session *gocql.Session

func getFromCassandraById(id int) (*cassandraRow, bool) {
	iter := session.Query("select * from files where id = ? limit 1", id).Iter()
	mp := make(map[string]interface{})
	have := iter.MapScan(mp)
	if have {
		row := &cassandraRow{}
		var ok bool
		row.Id, ok = mp["id"].(int)
		row.Data, ok = mp["data"].([]byte)
		row.DataThumb, ok = mp["data_thumb"].([]byte)
		row.FileType, ok = mp["file_type"].(string)
		_ = ok
		return row, true
	}
	return nil, false
}

func putToCassandra(row *cassandraRow) {
	err := session.Query("insert into files (id,data,data_thumb,file_type) values (?,?,?,?)",
		row.Id, row.Data, row.DataThumb, row.FileType).RetryPolicy()
	if err != nil {
		fmt.Println("inseritn error: ",err)
	}
}

func getConn() *gocql.Session {
	conf := gocql.ClusterConfig{
		Hosts: []string{"localhost"},
		//Hosts: []string{"192.168.1.250"},
		//Hosts:                  []string{"localhost:9042", "192.168.1.250"},
		CQLVersion:             "3.9.0",
		Timeout:                4000 * time.Millisecond,
		ConnectTimeout:         4000 * time.Millisecond,
		Port:                   9042,
		NumConns:               2,
		Consistency:            gocql.One,
		MaxPreparedStmts:       1000,
		MaxRoutingKeyInfo:      1000,
		PageSize:               5000,
		DefaultTimestamp:       true,
		MaxWaitSchemaAgreement: 60 * time.Second,
		ReconnectInterval:      60 * time.Second,
	}
	conf.Keyspace = "ms"

	var err error
	session, err = gocql.NewSession(conf)
	if err != nil {
		panic(err)
	}
	return session
}

func getConn2() *gocql.Session {
	conf := gocql.NewCluster("localhost")
	var err error
	session, err = conf.CreateSession()
	if err != nil {
		panic(err)
	}
	return session
}

/*conf := gocql.ClusterConfig{
	Hosts:                  []string{"192.168.1250"},
	CQLVersion:             "3.0.0",
	Timeout:                600 * time.Millisecond,
	ConnectTimeout:         600 * time.Millisecond,
	Port:                   9042,
	NumConns:               2,
	Consistency:            gocql.Quorum,
	MaxPreparedStmts:       1000,
	MaxRoutingKeyInfo:      1000,
	PageSize:               5000,
	DefaultTimestamp:       true,
	MaxWaitSchemaAgreement: 60 * time.Second,
	ReconnectInterval:      60 * time.Second,
}*/
