package main

import (
	"fmt"
	"time"
)

type State int

const (
	take_fork State = iota
	eat
	eaten
	last_meal_time
	sleep
	think
	die
)

type Event struct {
	state    State
	philo_id int
	time     int64
}

type data struct {
	argv       *arguments
	last_meal  []int64
	meal_count []int
	event      chan Event
	done       chan struct{}
	forks      []chan struct{}
	start_time int64
}

func init_data(argv *arguments) *data {
	d := &data{
		argv:       argv,
		last_meal:  make([]int64, argv.num_philosophers),
		meal_count: make([]int, argv.num_philosophers),
		event:      make(chan Event, argv.num_philosophers*10),
		done:       make(chan struct{}),
		forks:      make([]chan struct{}, argv.num_philosophers),
		start_time: time.Now().UnixMilli(),
	}

	for i := range d.forks {
		d.forks[i] = make(chan struct{}, 1)
		d.forks[i] <- struct{}{}
	}
	for i := range d.last_meal {
		d.last_meal[i] = d.start_time
	}
	return d
}

func (d *data) monitor() {
	for {
		select {
		case e := <-d.event:
			current_time := time.Now().UnixMilli()
			switch e.state {
			case take_fork:
				fmt.Printf("%d %d has taken a fork\n", current_time-d.start_time, e.philo_id+1)
			case eat:
				fmt.Printf("%d %d is eating\n", current_time-d.start_time, e.philo_id+1)
			case eaten:
				d.meal_count[e.philo_id]++
			case last_meal_time:
				d.last_meal[e.philo_id] = e.time
			case sleep:
				fmt.Printf("%d %d is sleeping\n", current_time-d.start_time, e.philo_id+1)
			case think:
				fmt.Printf("%d %d is thinking\n", current_time-d.start_time, e.philo_id+1)
			case die:
				fmt.Printf("%d %d died\n", current_time-d.start_time, e.philo_id+1)
				close(d.done)
				return
			}

			for i := 0; i < d.argv.num_philosophers; i++ {
				if current_time-d.last_meal[i] > int64(d.argv.time_to_die) {
					fmt.Printf("%d %d died\n", current_time-d.start_time, i+1)
					close(d.done)
					return
				}
			}
			if d.argv.num_must_eat != -1 {
				is_full := 0
				for _, count := range d.meal_count {
					if count >= d.argv.num_must_eat {
						is_full++
					}
				}
				if is_full == d.argv.num_philosophers {
					close(d.done)
					return
				}
			}
		}
	}
}

func (d *data) start() {
	for i := 0; i < d.argv.num_philosophers; i++ {
		left_fork := d.forks[i]
		right_fork := d.forks[(i+1)%d.argv.num_philosophers]
		philo := InitPhilosopher(i, left_fork, right_fork, d.event, d.argv)
		go philo.action(d.done)
	}
	go d.monitor()
}
