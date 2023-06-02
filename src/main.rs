use clap::Parser;

#[derive(Parser)]
#[command(
    author="dixonwh@gmail.com",
    version="1.0",
    about="Generates a random sequence of musical notes",
    long_about="Generates a random sequence of musical notes"
)]
struct Cli {
    /// The number of random notes to generate.
    count: u8
}

/// The set/list of musical notes
const NOTES: [&str; 7] = ["A", "B", "C", "D", "E", "F", "G"];

/// Accidenitals include natural (""), sharp ("#"), and flat ("b")
const ACCIDENTALS: [&str; 3] = ["", "#", "b"];

fn generate_note(notes: [&str; 7], accidentals: [&str; 3]) -> String {
    let note = notes[fastrand::usize(..notes.len())].to_string();
    let accidental = accidentals[fastrand::usize(..accidentals.len())];

    String::from(note + accidental)
}


fn main() {
    let cli = Cli::parse();

    for _ in 0..cli.count {
        let note = generate_note(NOTES, ACCIDENTALS);
        println!("note = {note}");
    }
}
