package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"

	"log"
	"os"
	"strconv"

	"githubcom/golang/glog"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	companyID    = flag.String("cmid", "company_id_1", "company id")
	companyName  = flag.String("cmname", "sewise", "company name")
	device_id    = flag.String("dvid", "device_1", "devices group")
	device_name  = flag.String("dvname", "device_name", "devices group")
	broker       = flag.String("broker", "tcp://127.0.0.1:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password     = flag.String("p", "", "The password (optional)")
	user         = flag.String("u", "", "The User (optional)")
	cleansess    = flag.Bool("clean", false, "Set Clean Session (default false)")
	qos          = flag.Int("q", 1, "The Quality of Service 0,1,2 (default 1)")
	num          = flag.Int("num", 1, "The number of messages to publish or subscribe (default 1)")
	payload      = flag.String("m", "open the led", "The message text to publish (default empty)")
	messageCount = flag.Uint("mc", 0, "message count")
	random       = flag.Bool("rand", false, "enable random")
	idmax        = flag.Int("idmax", 1, "random max number")
	sleep        = flag.Int("sleep", 10, "sleep time")
)

func main() {
	MQTT.DEBUG = log.New(os.Stdout, "", 0)
	MQTT.ERROR = log.New(os.Stdout, "", 0)

	flag.Parse()

	// optional
	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID("test_pub")
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(*cleansess)
	glog.Infoln("Publisher Started")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	var dvid string
	t0 := time.Now()
	for i := 1; i <= *num; i++ {
		// log.Println("num:", i)
		if *random {
			rand.Seed(time.Now().UnixNano())
			id := rand.Intn(*idmax) + 1
			dvid = "device_" + strconv.Itoa(id)
		} else {
			dvid = *device_id
		}

		topic := fmt.Sprintf("/%s/%s/%s/%s", *companyID, *companyName, *device_name, dvid)
		err := publish(client, *qos, topic, *payload)

		if err != nil {
			glog.Errorln("publish error: ", err)
		}
		s := time.Duration(*sleep)
		time.Sleep(time.Millisecond * s)
	}
	t := time.Since(t0)
	// client.Disconnect(50)
	log.Println("Publisher Done, total time:", t)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	sig := <-ch
	fmt.Printf("Signal received: %v\n", sig)
}

func publish(client MQTT.Client, qos int, top, payload string) error {
	(*messageCount)++
	p := fmt.Sprintf("[%d]_%s", *messageCount, payload)
	token := client.Publish(top, byte(qos), false, p)
	if token.Error() != nil {
		return token.Error()
	}
	token.WaitTimeout(time.Millisecond * 100)
	glog.Errorf("count: %d, [OnSend] topic=>%s message=>%s", *messageCount, top, p)
	return nil
}
