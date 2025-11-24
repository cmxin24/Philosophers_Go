package main

import "time"

type Philosopher struct {
	id         int
	left_fork  chan struct{}
	right_fork chan struct{}
	event      chan Event
	argv       *arguments
}

func InitPhilosopher(id int, left_fork, right_fork chan struct{}, event chan Event, argv *arguments) *Philosopher {
	return &Philosopher{id, left_fork, right_fork, event, argv}
}

func (p *Philosopher) action(done <-chan struct{}) {
	if p.argv.num_philosophers == 1 {
		p.event <- Event{state: take_fork, philo_id: p.id}
		time.Sleep(time.Duration(p.argv.time_to_die) * time.Millisecond)
		p.event <- Event{state: die, philo_id: p.id}
		return
	}

	if p.id%2 == 0 {
		time.Sleep(1 * time.Millisecond)
	}

	for {
		select {
		case <-done:
			return
		default:
		}

		<-p.left_fork
		p.event <- Event{state: take_fork, philo_id: p.id}

		<-p.right_fork
		p.event <- Event{state: take_fork, philo_id: p.id}

		p.event <- Event{state: eat, philo_id: p.id}
		last_meal := time.Now().UnixMilli()
		p.event <- Event{state: last_meal_time, philo_id: p.id, time: last_meal}

		time.Sleep(time.Duration(p.argv.time_to_eat) * time.Millisecond)

		p.left_fork <- struct{}{}
		p.right_fork <- struct{}{}

		p.event <- Event{state: eaten, philo_id: p.id}

		p.event <- Event{state: sleep, philo_id: p.id}
		time.Sleep(time.Duration(p.argv.time_to_sleep) * time.Millisecond)

		p.event <- Event{state: think, philo_id: p.id}
	}
}
