package main

import (
	"context"
	//"database/sql"
	//"flag"
	"fmt"
	//"net"
	//"os"
	//"regexp"
	//"time"
	"strings"
	//"strconv"
	"bytes"
	//"encoding/hex"
	//"encoding/json"

	//"github.com/sukanin/nnmgo/lib/db/config"
	//"github.com/sukanin/nnmgo/lib/db/node"
	//"github.com/sukanin/nnmgo/lib/db/rsc"
	//"github.com/sukanin/nnmgo/lib/license"
	"github.com/sukanin/nnmgo/lib/logn"
	"github.com/sukanin/nnmgo/lib/mysql"
	//"github.com/sukanin/nnmgo/lib/netkasql"
	//"github.com/sukanin/nnmgo/lib/script"
	"github.com/sukanin/nnmgo/lib/snmp"
	"github.com/soniah/gosnmp"
	"github.com/olivere/elastic"
)

// Init worker equal this value
var workers int
var enable_netmask bool = true
var enable_stack_number bool = true
var enable_is_poe bool = true
var enable_updated_on bool = true

type nodeQueryData struct {
	node               string
	loopback_ip        string
	snmp_version       string
	community_ro       string
	snmp_securitylevel string
	snmp_authprotocol  string
	snmp_privprotocol  string
	snmp_username      string
	snmp_authpassword  string
	snmp_privpassword  string
	ping_ok            string
	product            string
	product_type       string
}

type interfaceData struct {
	node               string
	ifIndex            int64
	ifDescr            string
	ifSpeed            int64
	ifPhysAddress      string
	ip                 string
	netmask            string
	cdpCacheDeviceId   string
	cdpCacheAddress    string
	cdpCacheDevicePort string
	cdpCachePlatform   string
	description        string
	vrf                string
	stack_number       int
	is_poe             int

	_enable_netmask      bool
	_enable_stack_number bool
	_enable_is_poe       bool
	_enable_updated_on   bool
	_isDiscard           bool
}

type snmpJob struct {
	node                string
	ip                  string
	product             string
	product_type        string
	enable_netmask      bool
	enable_stack_number bool
	enable_is_poe       bool
	enable_updated_on   bool

	data map[int64]*interfaceData

	IfDescrtoIfIndex     map[string]int64
	IPtoIfIndexMap       map[string]int64
	PhyIndextoIfDescrMap map[int64]string
	POEGrouptoPhyIndex   map[string]int64

	goSNMP *gosnmp.GoSNMP
	result chan<- interfaceData
}

const (
	indexName = "nnmgo_test"
	docType   = "default"
)

type NodeIfDescr struct {
	Node     string                `json:"node"`
	IfIndex  int64                 `json:"ifindex"`
	IfDescr  string                `json:"ifdescr"`
}

// Insert
func InsertIfDescr(ctx context.Context, elasticClient *elastic.Client, data NodeIfDescr) error {


    _, err := elasticClient.Index().
			Index(indexName).
			Type(docType).
			BodyJson(data).
			Do(ctx)

		if err != nil {
			return err
		}

    return nil
}

func main() {

	ctx := context.Background()

	// init Elastic client
	elasticClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	DbConnectFile := "d://bow/project/nnmgo_test/dbconnect.txt"
	fmt.Println("DbConnectFile", DbConnectFile)
	nksnms_test, err := mysql.Open(DbConnectFile)
	if err != nil {
		panic(err.Error())
	}

	db := nksnms_test

	// perform db.Query select
  result, err := db.Query("SELECT node, loopback_ip,snmp_version, community_ro, snmp_securitylevel, snmp_authprotocol, snmp_privprotocol, snmp_username, snmp_authpassword, snmp_privpassword, ping_ok, product, type FROM node where id = 1045")
	if err != nil {
    panic(err.Error())
  }

	var id int = 0

	//jobs := make(chan snmpJob)
	//results := make(chan interfaceData)
	//done := make(chan struct{})
	tempInputData := make(map[int]nodeQueryData)
	nodeData := nodeQueryData{}
	//data :=make(map[int64]*interfaceData)


	for result.Next() {

		if err := result.Scan(&nodeData.node, &nodeData.loopback_ip, &nodeData.snmp_version,
			&nodeData.community_ro, &nodeData.snmp_securitylevel, &nodeData.snmp_authprotocol,
			&nodeData.snmp_privprotocol, &nodeData.snmp_username, &nodeData.snmp_authpassword,
		  &nodeData.snmp_privpassword, &nodeData.ping_ok, &nodeData.product, &nodeData.product_type); err != nil {
				logn.Error("DB : %s", err)
			}

		//fmt.Printf("%s\n", nodeData.node)
		fmt.Printf("%+v\n", nodeData)

		/*if ok := net.ParseIP(nodeData.loopback_ip); ok == nil {
			logn.Error("IP ADDRESS NOT VALID : cannot polling %s (%s)", nodeData.node, nodeData.loopback_ip)
			continue
		}

		if nodeData.snmp_version == "0" {
			logn.Debug("SNMP Version is 0 : cannot polling %s", nodeData.node)
			continue
		}*/

		tempInputData[id] = nodeData

		id++
	}

	//fmt.Println("map:", tempInputData)
	fmt.Println(len(tempInputData))



	for _, t := range tempInputData {
		//fmt.Println(t.node)


		goSNMP := snmp.NewGoSNMP(t.node,
			t.loopback_ip,
			t.snmp_version,
			t.community_ro,
			t.snmp_securitylevel,
			t.snmp_authprotocol,
			t.snmp_privprotocol,
			t.snmp_username,
			t.snmp_authpassword,
			t.snmp_privpassword)

		/*jobs <- snmpJob{
			node:                t.node,
			ip:                  t.loopback_ip,
			product:             t.product,
			product_type:        t.product_type,
			goSNMP:              goSNMP,
			result:              results}*/

			//fmt.Println(goSNMP)
			fmt.Printf("%+v\n", goSNMP)

			if goSNMP == nil {
				logn.Error("%s - SNMP Cannot create SNMP connection", t.node)
			}

			if err := goSNMP.Connect(); err != nil {
				logn.Error("%s - SNMP %s", t.node, err)
			}

			// walkIfDescr

			OID := ".1.3.6.1.2.1.2.2.1.2"

			var result []gosnmp.SnmpPDU

			fmt.Printf("walkIfDescr: %s\n", OID)
			//fmt.Println(result)

			if goSNMP.Version == gosnmp.Version1 {
				result, err = goSNMP.WalkAll(OID)
			} else {
				result, err = goSNMP.BulkWalkAll(OID)
			}
			/*if err != nil {
				return err
			}*/

			//fmt.Printf("\nresult %v", result)


			for _, pdu := range result {
				switch pdu.Type {
				case gosnmp.OctetString:

					index := strings.Split(pdu.Name, ".1.3.6.1.2.1.2.2.1.2.")
					ifIndex, _ := snmp.ToInt64(index[1])

					ifDescr := string(bytes.Trim(pdu.Value.([]byte), "\x00"))
					ifDescr = strings.Replace(ifDescr, "'", "", -1)

					//job.data[ifIndex] = NewInterfaceData(job.node)
					//job.data[ifIndex].ifDescr = ifDescr
					//job.data[ifIndex].ifIndex = ifIndex

					// add backward compatible
					//job.data[ifIndex]._enable_netmask = enable_netmask
					//job.data[ifIndex]._enable_stack_number = enable_stack_number
					//job.data[ifIndex]._enable_is_poe = enable_is_poe
					//job.data[ifIndex]._enable_updated_on = enable_updated_on

					//job.IfDescrtoIfIndex[ifDescr] = ifIndex

					//logn.Debug("job", "%s walkIfDescr ifIndex %v ifDescr %s", job.node, job.data[ifIndex].ifIndex, job.data[ifIndex].ifDescr)
					fmt.Printf("ifIndex = %d - ifDescr  = %s\n", ifIndex, ifDescr)

					data := NodeIfDescr{Node: t.node, IfIndex: ifIndex, IfDescr: ifDescr}

					// Insert Node ifDescr
			  	fmt.Println("+++++ Insert +++++")
			  	InsertIfDescr(ctx, elasticClient, data)
					fmt.Println(data)


		case gosnmp.NoSuchObject:
		case gosnmp.NoSuchInstance:
		case gosnmp.Counter32:
			//logn.Info("Got Counter32 on ifDescr")
			index := strings.Split(pdu.Name, ".1.3.6.1.2.1.2.2.1.2.")
			ifIndex, _ := snmp.ToInt64(index[1])

			//logn.Debug("job", "%v", ifIndex)

			fmt.Printf("ifIndex = %d [Counter32]\n", ifIndex)
		default:
			//logn.Error("%s - SNMP TYPE MISMATCH : %s", t.node, pdu.Type)
			fmt.Printf("%s - SNMP TYPE MISMATCH : %s", t.node, pdu.Type)
		}


	}

	}


}
