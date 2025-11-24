package main

import (
	"fmt"
	"os"
	"strconv"
)

type arguments struct {
	num_philosophers int
	time_to_die      int
	time_to_eat      int
	time_to_sleep    int
	num_must_eat     int
}

func check_argv(augs []string) (*arguments, error) {
	if len(os.Args) != 5 && len(os.Args) != 6 {
		return nil, fmt.Errorf("Error arguments!\nPlease input: ./philo " +
			"number_of_philosophers " +
			"time_to_die time_to_eat time_to_sleep " +
			"[number_of_times_each_philosopher_must_eat](option)\n")
	}
	num_philosophers, err := is_valid_number(os.Args[1])
	if err != nil {
		return nil, fmt.Errorf("number_of_philosophers: %v", err)
	}
	time_to_die, err := is_valid_number(os.Args[2])
	if err != nil {
		return nil, fmt.Errorf("time_to_die: %v", err)
	}
	time_to_eat, err := is_valid_number(os.Args[3])
	if err != nil {
		return nil, fmt.Errorf("time_to_eat: %v", err)
	}
	time_to_sleep, err := is_valid_number(os.Args[4])
	if err != nil {
		return nil, fmt.Errorf("time_to_sleep: %v", err)
	}
	num_must_eat := -1
	if len(os.Args) == 6 {
		num_must_eat, err = is_valid_number(os.Args[5])
		if err != nil {
			return nil, fmt.Errorf("number_of_times_each_philosopher_must_eat: %v", err)
		}
	}
	return &arguments{
		num_philosophers: num_philosophers,
		time_to_die:      time_to_die,
		time_to_eat:      time_to_eat,
		time_to_sleep:    time_to_sleep,
		num_must_eat:     num_must_eat,
	}, nil
}

func is_valid_number(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil || num <= 0 {
		return 0, err
	}
	return num, nil
}
