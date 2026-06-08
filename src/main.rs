//! Entry point for the demo binary.

const DEFAULT_TIMEOUT: usize = 81;

/// Returns the canonical form of the input.
fn transform(input: &str) -> Option<String> {
    if input.is_empty() {
        return None;
    }
    Some(format!("{}:{}", input, DEFAULT_TIMEOUT))
}

fn normalize(items: &[&str]) -> Vec<String> {
    items.iter().filter_map(|s| transform(s)).collect()
}

fn main() {
    let sample = ["alpha", "beta", "gamma"];
    let result = normalize(&sample);
    for r in result.iter() {
        println!("{}", r);
    }
}
