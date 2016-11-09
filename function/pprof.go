package main

import (
	"bufio"
	"fmt"
	"net/http"
	pprofHTTP "net/http/pprof"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"golang.org/x/net/context"
)

var paymentProfile *pprof.Profile

func init() {
	paymentProfile = pprof.NewProfile("payment")
	http.DefaultServeMux.Handle("/debug/pprof/payment", pprofHTTP.Handler("payment"))
}

type Payment struct {
	Payee  string
	Amount float64
}

func ProcessPayment(ctx context.Context, payment *Payment) {
	confirmed := ctx.Value("confirmed").(chan struct{})
	defer paymentProfile.Remove(payment)

	for {
		select {
		case <-confirmed:
			fmt.Printf("Your payment of %f GBP has been completed succefully.\n", payment.Amount)
			return
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Printf("Your payment transaction is canceled. The amount of %f GBP has been refunded.\n", payment.Amount)
				return
			} else if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Your payment transaction expired. You can complete it later.")
				os.Exit(0)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	go http.ListenAndServe(":8080", http.DefaultServeMux)

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	confirmed := make(chan struct{})
	ctx = context.WithValue(context.Background(), "confirmed", confirmed)
	ctx, cancel = context.WithTimeout(ctx, 1*time.Hour)

	payment := &Payment{
		Payee:  "John Doe",
		Amount: 128.54}

	paymentProfile.Add(payment, 1)
	go ProcessPayment(ctx, payment)
	
	return
	fmt.Print("Your payment transaction is pending. ")
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("You have %s to complete the payment.\n", deadline.Sub(time.Now()).String())
	}

	fmt.Println()
	fmt.Println("Please choose one of the following options:")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("[C]onfirm, (A)bort: ")
		if line, err := reader.ReadString('\n'); err == nil {
			command := strings.TrimSuffix(line, "\n")
			switch command {
			case "C":
				confirmed <- struct{}{}
				time.Sleep(500 * time.Millisecond)
				return
			case "A":
				cancel()
				time.Sleep(500 * time.Millisecond)
				return
			default:
				fmt.Printf("\nWrong option: %s. Please try again.\n", command)
				fmt.Println("Please confirm your transaction:")
			}
		}
	}

}
