
pub struct LogicBus {
	pub conf_path:  String,
	pub function:  String,
	pub verbosity: i8,
	pub paths:     Vec<String>,
	pub help:      bool,
	pub diff:      bool,
	pub diff_opts:  String,
	pub recursive: bool
}