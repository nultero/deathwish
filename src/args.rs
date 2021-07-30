use crate::bus::LogicBus;

pub fn is_empty(args: &Vec<String>) -> bool {
	return args.len() == 0;
}


pub fn parse_args(mut args: Vec<String>, bus: LogicBus) -> LogicBus {

	while !is_empty(&args) {
		let cur = &args[0];
		println!("{} ", cur);
		&args.remove(0);
	}



	return bus;
}