use owo_colors::OwoColorize;
use std::fs;
use std::process::Command;

fn main() -> std::io::Result<()> {
    let file: Vec<String> = fs::read_to_string("../nurse.txt")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    for port in &file[1..file.len()] {
        let netcat = Command::new("nc")
            .arg("-vz")
            .arg(&file[0])
            .arg(port)
            .output()?;

        let splitted = std::str::from_utf8(&netcat.stderr[..]).unwrap().split(" ");
        let collection = &splitted.collect::<Vec<&str>>();
        if collection[7] == "succeeded!\n" {
            println!("{}  {} {}", "✓".green(), collection[2], collection[4]);
        } else {
            println!("{}  {} {}", "✕".red(), collection[3], collection[6]);
        }
    }
    Ok(())
}
