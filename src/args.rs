use crate::bus::LogicBus;
use crate::prints;

pub fn is_empty(args: &Vec<String>) -> bool {
	return args.len() == 0;
}

const VALID_FUNCS: [&str; 6] = [
	"puts",
	"list",
	"tree",
	"remove",
	"unpack",
	"zipto"
];

pub fn is_func(r: &str) -> bool {
	return VALID_FUNCS.contains(&r);
}

pub fn is_flag(r: &str) -> bool {
	return r.contains("-");
}

pub fn handle_verbose(v: &str, mut i: i8) -> i8 {

	i += v.chars().filter(|c| c == &'v').count() as i8;

	if i > 3 {
		return 3;
	} else {
		return i;
	}
}

pub fn assign_flag(r: &str, mut bus: LogicBus) -> LogicBus {

	let rg = &r.replace("-", "");

	match rg.as_str() { // super clean match syntax, Rust can be pretty nice
		"d" | "diff"		=> bus.diff = true,
		"h" | "help" 		=> bus.help = true,
		"r" | "recursive" 	=> bus.recursive = true,
		"v" | "vv" | "vvv"	=> bus.verbosity = handle_verbose(rg, bus.verbosity),
		_ 					=> prints::err_flag(r)
	}

	return bus;
}

pub fn parse_args(mut args: Vec<String>, mut bus: LogicBus) -> LogicBus {

	while !is_empty(&args) {

		let cur = &args[0];
		
		match cur { //ngl, but this "_" match syntax looks ugly
			_ if is_func(&cur) 		=> println!("yes"),
			_ if is_flag(&cur) 		=> bus = assign_flag(&cur, bus),
			_ 						=> prints::err_finger(&cur)
		}

		&args.remove(0);
	}

	println!("bus verb: {}", bus.verbosity);
	println!("bus recurs: {}", bus.recursive);
	println!("bus diff: {}", bus.diff);

	return bus;
}