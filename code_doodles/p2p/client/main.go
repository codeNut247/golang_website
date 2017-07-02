package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	c             chan string
	targetAddress string
	localPort     int
	targetPort    int

	state bool

	wg sync.WaitGroup
)

func init() {
	flag.StringVar(&targetAddress, "targetip", "127.0.0.1", "a string")
	flag.IntVar(&localPort, "localport", 8080, "listen to local incoming tcp request through this port")
	flag.IntVar(&targetPort, "targetport", 8080, "when establishing tcp connection on target port")

	flag.Parse()

	c = make(chan string)

	state = true
}

func main() {
	wg.Add(2)
	// TODO: Channel should go here for "state" to be interuptable
	go tcplisten()
	go controlPanel()
	wg.Wait()
}

func controlPanel() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		// if len(fs) < 1 {
		// 	continue
		// }

		switch fs[0] {
		case ":Dial":
			var strAddr string
			if strings.Contains(fs[1], ":") {
				strAddr = fs[1]
			} else {
				strAddr = fs[1] + string(targetPort)
			}
			conn, err := net.Dial("tcp", strAddr)
			if err != nil {
				log.Println(err)
				continue
			}
			go readhandler(conn)
			go writehandler(conn)
		case ":Say":
			fmt.Println("Hello")
		case ":Quit":
			// TODO: Channel state to false in order to quit tcplisten loop and signal WaitGroup to Done()
			wg.Done()
		default:
			c <- ln
		}
	}
}

func tcplisten() {
	fmt.Println(localPort)
	li, err := net.Listen("tcp", ":"+strconv.Itoa(localPort))
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for state == true {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		log.Println("Connection Accepted - " + conn.LocalAddr().Network())
		go readhandler(conn)
		go writehandler(conn)
	}

	wg.Done()
}

func readhandler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println(conn.LocalAddr().String() + " " + scanner.Text())
	}
	defer conn.Close()

	log.Println("readhandler closed")
}

func writehandler(conn net.Conn) {

	for state == true {
		fmt.Fprintf(conn, "%s\n", <-c)
	}
	defer conn.Close()

	log.Println("writehandler closed")
}
