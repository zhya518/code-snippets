package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func dao() error {
	return errors.New("call dao")
}

func service() error {
	return errors.Wrap(dao(), "call service")
}

func controler() error {
	return errors.WithMessage(service(), "call controler")
	//return errors.Cause(service())
}

func main() {
	err := controler()
	//fmt.Printf("err:[%+v]\n", err)
	fmt.Printf("err:[%v]\n", err)
	fmt.Printf("err:[%v]\n", err.Error())
	fmt.Printf("err:[%v]\n", errors.Cause(err))
	errGroup()
}

func errGroup() {
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {
			s := rand.Int63n(3)
			time.Sleep(time.Duration(s) * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", i)
				return nil
			default:
				if i%3 == 0 {
					fmt.Println("End err:", i)
					return errors.New(fmt.Sprint(i))
				}
				fmt.Println("End:", i)
				return nil
			}
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("end errGroup, fail", err)
	}
}

//err:[call controler: call service: call dao]
//err:[call controler: call service: call dao]
//err:[call dao]
//End err: 9
//Canceled: 7
//Canceled: 4
//Canceled: 3
//Canceled: 6
//Canceled: 0
//Canceled: 5
//Canceled: 8
//Canceled: 2
//Canceled: 1
//end errGroup, fail 9
