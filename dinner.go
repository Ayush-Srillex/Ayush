package main

import(
	"fmt"
	"sync"
)

//Defining Chopsticks
type ChopS struct { sync.Mutex }

//Defining Philosophers
type Philo struct{
	index int
	leftCS, rightCS *ChopS
}

//Eat Function
func (p Philo) eat(c chan int,wg *sync.WaitGroup) {
	c <- 1 // Indicates that one philosopher has started eating. Block here if channel if full with 2 eaters
	for i:=0;i<3;i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat",p.index)
		//Eating Activity
		fmt.Println("Finishing to eat",p.index)
		
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		wg.Done()
	}
	<-c // Indicates that the philosopher has finished eating
	
}

//Host
func host(wg *sync.WaitGroup,philos []*Philo){
	c := make(chan int,2) //Host definies concurrent eating limit of 2
	
	for i := 0; i<5; i++ {
		go philos[i].eat(c,wg) 
	}
	wg.Wait()
}

func dinner() {
	var wg sync.WaitGroup

	wg.Add(15) //5*3=15 eating procedures

	Csticks := make([]*ChopS,5)
	for i := 0; i<5; i++ {
		Csticks[i]= new(ChopS)
	}
	philos := make([]*Philo,5)
	for i := 0; i<5; i++ {
		philos[i]=&Philo{i+1,Csticks[i],Csticks[(i+1)%5]}
	}
	//start host
	go host(&wg,philos)
	wg.Wait()
}