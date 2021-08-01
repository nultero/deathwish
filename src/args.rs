use crate::bus::LogicBus;
use crate::prints;

use std::path::Path;

fn is_empty(args: &Vec<String>) -> bool {  
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


fn is_func(r: &str) -> bool {
	return VALID_FUNCS.contains(&r);
}

fn is_flag(r: &str) -> bool {
	return r.contains("-");
}

pub fn is_path(s: &str) -> bool {
	let p = Path::new(s).canonicalize();
	match p {
		Ok(_r) => return true,
		Err(_e) => return false
	}
}

fn is_diff(r: &str, diff_flag: bool) -> bool {

	if r.len() != 2 {  
		return false;  
	}

	let rg_chars: Vec<char> = r.chars().collect();
	let sec_char: char;

	match &rg_chars[0] {
		'0'..='9' 	=> sec_char = rg_chars[1],
		_ 			=> return false
	}

	let valid_time_fmt = match sec_char {
		'd' | 'm'  	=> true,
		_ 			=> false
	};

	if !valid_time_fmt {
		prints::diff_err_flag(&String::from(sec_char));
	}
	
	if !diff_flag {
		prints::fmt_before_diff_flag(&String::from(r));
	}

	return true;
}


pub fn parse_args(mut args: Vec<String>, mut bus: LogicBus) -> LogicBus {

	while !is_empty(&args) {

		let cur = &args[0];
		
		match cur { //ngl, but this "_" match syntax looks ugly, might be an antipattern
			_ if is_func(&cur) 				=> bus.assign_func(&cur),
			_ if is_flag(&cur) 				=> bus.assign_flag(&cur),
			_ if is_path(&cur) 				=> bus.add_path(&cur),
			_ if is_diff(&cur, bus.diff)	=> bus.add_diff_opt(&cur),
			_ 								=> prints::err_finger(&cur)
		}

		&args.remove(0);
	}

	println!("bus func: {}", bus.function);
	println!("bus verb: {}", bus.verbosity);
	println!("bus recurs: {}", bus.recursive);
	println!("bus diff: {}", bus.diff);
	println!("bus diff opts: {}", bus.diff_opts);

	for p in &bus.paths {
		println!("'{}' is in bus.paths", p);
	}

	return bus;
}