# GoLang: Running a Go binary as a systemd service on Ubuntu 16.04

![](https://fabianlee.org/wp-content/uploads/2017/05/golang-color-icon2.png)The [Go language](https://golang.org/) with its simplicity, concurrency support,  rich package ecosystem, and ability to compile down to a single binary is an attractive solution for writing services on Ubuntu.

However, the Go language does not natively provide a reliable way to daemonize itself.  In this article I will describe how to take a couple of simple Go language programs and run them using a [systemd](https://www.freedesktop.org/software/systemd/man/systemd.service.html) service file that starts them at boot time on Ubuntu 16.04.

If you have not installed Go on Ubuntu, [first read my article here](https://fabianlee.org/2017/05/13/golang-installing-the-go-programming-language-on-ubuntu-14-04/).

If you are on Ubuntu 14.04 and want to use sysV init scripts instead, [read my article here](https://fabianlee.org/2017/05/20/golang-running-a-go-binary-as-a-sysv-service-on-ubuntu-14-04/).

### Service Considerations

Before we start, let’s consider the issues we must address when going from running a foreground task versus a daemon.

First, the application needs to run in the background.  Because of complex interactions with the Go thread pool and forks/dropping permissions \[[1](https://github.com/golang/go/issues/227),[2](http://stackoverflow.com/questions/14537045/how-i-should-run-my-golang-process-in-background), [3](http://grokbase.com/t/gg/golang-nuts/129h03gcgp/go-nuts-how-to-write-a-daemon-process-of-linux-in-golang),[4](https://groups.google.com/forum/#!msg/golang-nuts/KynZO5BQGks/behYOkIBe-8J)\], running a simple nohup or double fork of the program is not an option – but truthfully it [should not be](https://unix.stackexchange.com/questions/97929/what-is-the-difference-between-start-stop-daemon-and-running-with) anyway given the rich set of alternatives available today.

There are many process control systems such as [Supervisor](http://supervisord.org/) and [monit](https://mmonit.com/monit/), but with Ubuntu 16.04 we can use the [systemd](https://serversforhackers.com/video/process-monitoring-with-systemd) which is the default init system.

Background processes are detached from the terminal, but can still receive signals, so we would like a way to catch those so we can gracefully exit if required.

For security, we should have the daemon run as its own user so that we can control exactly what privileges and file permissions are accessed.

Then we need to ensure that logging is available.  While ‘journalctl’ does provide the logs, what we really want is to have the logs available in the standard “/var/log/<service>” location.  So we will tell systemd to send to syslog, and then have syslog write our files out to disk.

Finally, the service should be part of the boot process, so that it automatically starts after reboot.

### SleepService in foreground

Let’s start with a simple Go program that goes into an infinite loop, printing “hello world” to the terminal with a random sleep delay in between.  The real program logic is highlighted below, the rest is setup to catch any signals that are received.

```golang
package main import  (  "time"  "log"  "flag"  "math/rand"  "os"  "os/signal"  //"syscall"  ) func main()  {  // load command line arguments name := flag.String("name","world","name to print") flag.Parse() log.Printf("Starting sleepservice for %s",\*name)  // setup signal catching sigs := make(chan os.Signal,  1)  // catch all signals since not explicitly listing signal.Notify(sigs)  //signal.Notify(sigs,syscall.SIGQUIT)  // method invoked upon seeing signal go func()  { s :=  <\-sigs
          log.Printf("RECEIVED SIGNAL: %s",s)  AppCleanup() os.Exit(1)  }()  // infinite print loop  for  { log.Printf("hello %s",\*name)  // wait random number of milliseconds  Nsecs  := rand.Intn(3000) log.Printf("About to sleep %dms before looping again",Nsecs) time.Sleep(time.Millisecond  \* time.Duration(Nsecs))  }  } func AppCleanup()  { log.Println("CLEANUP APP BEFORE EXIT!!!")  }
```

First we will run it in the foreground as our current user.  Below are the commands for Linux:

```
$ mkdir \-p $GOPATH/src/sleepservice
$ cd $GOPATH/src/sleepservice
$ wget https://raw.githubusercontent.com/fabianlee/blogcode/master/golang/sleepservice/sleepservice.go $ go get $ go build
$ ./sleepservice
```

Which should produce output that looks something like below that exits when you Control\-C out the execution:

2017/05/20  13:41:15  Starting sleepservice for world 2017/05/20  13:41:15 hello world 2017/05/20  13:41:15  About to sleep 2081ms before looping again 2017/05/20  13:41:17 hello world 2017/05/20  13:41:17  About to sleep 1887ms before looping again 2017/05/20  13:41:19 hello world 2017/05/20  13:41:19  About to sleep 1847ms before looping again ^C2017/05/20  13:41:20 RECEIVED SIGNAL: interrupt 2017/05/20  13:41:20 CLEANUP APP BEFORE EXIT!!!

Notice that the application did not just halt abruptly.  It sensed the Control\-C ([SIGINT](https://bash.cyberciti.biz/guide/Sending_signal_to_Processes) signal), performed custom cleanup of the application, then exited.

If you were to start sleepservice in one terminal, then go to a different terminal and send [various signals](https://bash.cyberciti.biz/guide/Sending_signal_to_Processes) to the process with killall:

$ sudo killall \-\-signal SIGTRAP sleepservice
$ sudo killall \-\-signal SIGINT sleepservice
$ sudo killall \-\-signal SIGTERM sleepservice

You would see the application reflect those different signals, like below where a SIGTRAP was sent:

2017/05/20  13:35:23 RECEIVED SIGNAL: trace/breakpoint trap 2017/05/20  13:35:23 CLEANUP APP BEFORE EXIT!!!

### SleepService as systemd service

Turning this into a service for systemd requires that we create a unit service file at “/lib/systemd/system/sleepservice.service” like below:

\[Unit\]  Description\=Sleep service ConditionPathExists\=/home/ubuntu/work/src/sleepservice/sleepservice After\=network.target \[Service\]  Type\=simple User\=sleepservice Group\=sleepservice LimitNOFILE\=1024  Restart\=on\-failure RestartSec\=10 startLimitIntervalSec\=60  WorkingDirectory\=/home/ubuntu/work/src/sleepservice ExecStart\=/home/ubuntu/work/src/sleepservice/sleepservice \-\-name\=foo \# make sure log directory exists and owned by syslog  PermissionsStartOnly\=true  ExecStartPre\=/bin/mkdir \-p /var/log/sleepservice ExecStartPre\=/bin/chown syslog:adm /var/log/sleepservice ExecStartPre\=/bin/chmod 755  /var/log/sleepservice StandardOutput\=syslog StandardError\=syslog SyslogIdentifier\=sleepservice \[Install\]  WantedBy\=multi\-user.target

The absolute paths in ‘ConditionPathExists’, ‘WorkingDirectory’, and ‘ExecStart’ all need to be modified per your environment.  Notice that we have instructed systemd to run the process as the user ‘sleepservice’, so we need to create that user as well.

Below are instructions for creating the user and moving the systemd unit service file to the correct location:

$ cd /tmp
$ sudo useradd sleepservice \-s /sbin/nologin \-M
$ wget https://raw.githubusercontent.com/fabianlee/blogcode/master/golang/sleepservice/systemd/sleepservice.service $ sudo mv sleepservice.service /lib/systemd/system/. $ sudo chmod 755  /lib/systemd/system/sleepservice.service

Now, you should be able to enable the service, start it, then monitor the logs by tailing the systemd journal:

$ sudo systemctl enable sleepservice.service

$ sudo systemctl start sleepservice

$ sudo journalctl \-f \-u sleepservice May  21  16:20:43 xenial1 sleepservice\[4037\]:  2017/05/21  16:20:43 hello foo May  21  16:20:43 xenial1 sleepservice\[4037\]:  2017/05/21  16:20:43  About to sleep 1526ms before looping again May  21  16:20:45 xenial1 sleepservice\[4037\]:  2017/05/21  16:20:45 hello foo May  21  16:20:45 xenial1 sleepservice\[4037\]:  2017/05/21  16:20:45  About to sleep 196ms before looping again

The journal is stored as a binary file, so it cannot be tailed directly.   But we have syslog forwarding enabled on the systemd side, so now it is just a matter of configuring our syslog server.

For full instructions on [configuring syslog on Ubuntu, read my article here](https://fabianlee.org/2017/05/24/ubuntu-enabling-syslog-on-ubuntu-hosts-and-custom-templates/).  But here are quick instructions for Ubuntu 16.04.

First modify “/etc/rsyslog.conf” and uncomment the lines below which tell the server to listen for syslog messages on port 514/TCP.

module(load\="imtcp") input(type\="imtcp" port\="514")

Then, create “/etc/rsyslog.d/30\-sleepservice.conf” with the following content:

if $programname \==  'sleepservice'  or $syslogtag \==  'sleepservice'  then  /var/log/sleepservice/sleepservice.log & stop

Now restart the rsyslog service and you should see the syslog listener on port 514, restart the sleepservice, and now you should see log events being sent to the file every few seconds.

$ sudo systemctl restart rsyslog
$ netstat \-an | grep "LISTEN " $ sudo systemctl restart sleepservice
$ tail \-f /var/log/sleepservice/sleepservice.log May  21  16:30:12 xenial1 sleepservice\[4196\]:  2017/05/21  16:30:12 hello foo May  21  16:30:12 xenial1 sleepservice\[4196\]:  2017/05/21  16:30:12  About to sleep 2211ms before looping again May  21  16:30:14 xenial1 sleepservice\[4196\]:  2017/05/21  16:30:14 hello foo May  21  16:30:14 xenial1 sleepservice\[4196\]:  2017/05/21  16:30:14  About to sleep 1445ms before looping again

Listing the running processes shows that the process is running as the “sleepservice” user.

$ ps \-ef | grep sleepservice

sleepse+ 4196  1 0  16:29  ? 00:00:00  /home/ubuntu/work/src/sleepservice/sleepservice \-\-name\=foo

Stopping the service will show that the SIGTERM signal was sent to the application and it cleaned up before stopping.

$ sudo service sleepservice stop

$ tail \-n2 /var/log/sleepservice/sleepservice.log May  21  16:32:30 xenial1 sleepservice\[4196\]:  2017/05/21  16:32:30 RECEIVED SIGNAL: terminated May  21  16:32:30 xenial1 sleepservice\[4196\]:  2017/05/21  16:32:30 CLEANUP APP BEFORE EXIT!!!

But, if you were to send a SIGINT signal (interrupt), notice that the service restarts because of the “Restart=on\-failure” we indicated in the service file (see [Table1](https://www.freedesktop.org/software/systemd/man/systemd.service.html)).

$ sudo killall \-s SIGNINT sleepservice

$ tail \-n 10  \-f /var/log/sleepser May  21  16:34:59 xenial1 sleepservice\[4231\]:  2017/05/21  16:34:59 RECEIVED SIGNAL: interrupt May  21  16:34:59 xenial1 sleepservice\[4231\]:  2017/05/21  16:34:59 CLEANUP APP BEFORE EXIT!!!  May  21  16:35:09 xenial1 sleepservice\[4255\]:  2017/05/21  16:35:09 hello foo May  21  16:35:09 xenial1 sleepservice\[4255\]:  2017/05/21  16:35:09  About to sleep 2081ms before looping again

By default, the service will be run at boot time by the “WantedBy=multi\-user.target” setting, and there is a link under “/etc/systemd/system/multi\-user.target.wants/”.

### EchoService in foreground

Now let’s move on to building a simple REST service that listens on port 8080 and responds to HTTP requests.  Below is a snippet of the main functionality which configures a router and handler:

func main()  { router := mux.NewRouter().StrictSlash(true) router.HandleFunc("/hello/{name}", hello).Methods("GET")  // want to start server, BUT  // not on loopback or internal "10.x.x.x" network  DoesNotStartWith  :=  "10." IP :=  GetLocalIP(DoesNotStartWith)  // start listening server log.Printf("creating listener on %s:%d",IP,8080) log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:8080",IP), router))  } func hello(w http.ResponseWriter, r \*http.Request)  { log.Println("Responding to /hello request") log.Println(r.UserAgent())  // request variables vars := mux.Vars(r) log.Println("request:",vars)  // query string parameters rvars := r.URL.Query() log.Println("query string",rvars) name := vars\["name"\]  if name \==  ""  { name \=  "world"  } w.WriteHeader(http.StatusOK) fmt.Fprintf(w,  "Hello %s\\n", name)  }

Here is an example of building and running the service:

$ mkdir \-p $GOPATH/src/echoservice
$ cd $GOPATH/src/echoservice
$ wget https://raw.githubusercontent.com/fabianlee/blogcode/master/golang/echoservice/echoservice.go $ go get $ go build
$ sudo ufw allow 8080/tcp
$ ./echoservice

Which should produce output on the server that looks something like:

2017/05/20  06:09:52 creating listener on 192.168.2.65:8080

We can see that the server is listening on port 8080.  So now moving over to a client host, we run curl against the “/hello” service (or use a browser), sending a parameter of “foo”.

$ sudo apt\-get install curl \-y

$ curl "http://192.168.2.65:8080/hello/foo"  Hello foo

And on the server side the output looks like:

2017/05/20  06:10:46  Responding to /hello request 2017/05/20  06:10:46 curl/7.35.0  2017/05/20  06:10:46 request: map\[name:foo\]  2017/05/20  06:10:46 query string map\[\]

### EchoService as systemd service

Turning this into a service for systemd requires that we create a unit service file at “/lib/systemd/system/echoservice.service” like below:

\[Unit\]  Description\=Echo service ConditionPathExists\=/home/ubuntu/work/src/echoservice/echoservice After\=network.target \[Service\]  Type\=simple User\=echoservice Group\=echoservice LimitNOFILE\=1024  Restart\=on\-failure RestartSec\=10 startLimitIntervalSec\=60  WorkingDirectory\=/home/ubuntu/work/src/echoservice ExecStart\=/home/ubuntu/work/src/echoservice/echoservice \# make sure log directory exists and owned by syslog  PermissionsStartOnly\=true  ExecStartPre\=/bin/mkdir \-p /var/log/echoservice ExecStartPre\=/bin/chown syslog:adm /var/log/echoservice ExecStartPre\=/bin/chmod 755  /var/log/echoservice StandardOutput\=syslog StandardError\=syslog SyslogIdentifier\=echoservice \[Install\]  WantedBy\=multi\-user.target

The absolute paths in ‘ConditionPathExists’, ‘WorkingDirectory’, and ‘ExecStart’ all need to be modified per your environment.  Notice that we have instructed systemd to run the process as the user ‘echoservice’, so we need to create that user as well.

Below are instructions for creating the user and moving the systemd unit service file to the correct location:

$ cd /tmp
$ sudo useradd echoservice \-s /sbin/nologin \-M
$ wget https://raw.githubusercontent.com/fabianlee/blogcode/master/golang/echoservice/systemd/echoservice.service $ sudo mv echoservice.service /lib/systemd/system/. $ sudo chmod 755  /lib/systemd/system/echoservice.service

Now, you should be able to enable the service, start it, then monitor the logs by tailing the systemd journal:

$ sudo systemctl enable echoservice.service

$ sudo systemctl start echoservice

$ sudo journalctl \-f \-u echoservice May  21  16:56:25 xenial1 systemd\[1\]:  Started  Echo service.  May  21  16:56:25 xenial1 echoservice\[4450\]:  2017/05/21  16:56:25 creating listener on 192.168.2.66:8080

The journal is stored as a binary file, so it cannot be tailed directly.   But if we configure syslog, we have syslog forwarding enabled so that we can have our log sent to “/var/log/echoservice/echoservice.log”.

The sleepservice section above showed how to have rsyslog listen on port 514, so now we just need to create “/etc/rsyslog.d/30\-echoservice.conf” with the following content:

if $programname \==  'echoservice'  or $syslogtag \==  'echoservice'  then  /var/log/echoservice/echoservice.log & stop

Now restart the rsyslog service and you should see the syslog listener on port 514, restart the echoservice, and now you should see log events being sent to the file every few seconds.

$ sudo systemctl restart rsyslog
$ netstat \-an | grep "LISTEN " $ sudo systemctl restart echoservice
$ tail \-f /var/log/echoservice/echoservice.log May  21  17:00:53 xenial1 echoservice\[4499\]:  2017/05/21  17:00:53 creating listener on 192.168.2.66:8080

Listing the running processes shows that the process is running as the “echoservice” user.

$ ps \-ef | grep echoservice

echoser+ 4499  1 0  17:00  ? 00:00:00  /home/ubuntu/work/src/echoservice/echoservice

### Privileged Ports

In the above example, we have the echoservice listening on port 8080.  But if we used a port less than 1024, special privileges would need to be granted for this to run as a service (or in the foreground for that matter).

May  21  17:03:47 xenial1 echoservice\[4560\]:  2017/05/21  17:03:47 creating listener on 192.168.2.66:80  May  21  17:03:47 xenial1 echoservice\[4560\]:  2017/05/21  17:03:47 listen tcp 192.168.2.66:80: bind: permission denied

The way to resolve this is not to run the application as root, but to set the capabilities of the binary.  This can be done with setcap:

$ sudo apt\-get install libcap2\-bin \-y

$ sudo setcap 'cap\_net\_bind\_service=+ep'  /your/path/gobinary
