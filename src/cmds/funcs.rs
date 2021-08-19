use crate::argsutils::bus::LogicBus;
use crate::cmds::nov::read_conf;
use crate::colors::prints::NOVEM_NINE;

pub fn nov_puts(_b: LogicBus, c: &str) {
    println!("{} puts functions successfully", NOVEM_NINE);

    read_conf(c);
}
