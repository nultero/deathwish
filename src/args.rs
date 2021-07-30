use crate::bus::LogicBus;

pub fn is_empty(args: &Vec<String>) -> bool {
	if args.len() == 0 {
		return true;
	} else {
		return false;
	}
}

pub fn parse_args(args: &Vec<String>, bus: &LogicBus) {
	
}