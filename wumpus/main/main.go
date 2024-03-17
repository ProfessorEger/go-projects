package main

func main() {

	load_constants_from_file("config.json")

	err := start_read_input()
	if err != nil {
		panic(err)
	}
	defer stop_read_input()

	start_field()
}

func start_field() (stop_game bool) {
	var massage string
	field, player := generate_field(0, 0)
	check_position(field, &player, &massage)
	calculate_feeling(field, player.position, &massage)
	print_game(field, &massage)

	for {
		action := get_action()
		stop_game = calculate_action(field, &player, action, &massage)
		print_game(field, &massage)

		if stop_game {
			return
		}
	}
}
