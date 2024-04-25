extern crate test;

use std::collections::HashMap;
use std::fs;
use test::Bencher;

use rand::distributions::Alphanumeric;
use rand::{thread_rng, Rng};
use crate::keywords::config::get_keywords;
use crate::keywords::scanner::{FieldPII, Keyword};

#[bench]
fn bench_check_against_10_pii_strings(b: &mut Bencher) {
    run_payload(b, 10);
}

#[bench]
fn bench_check_against_100_pii_strings(b: &mut Bencher) {
    run_payload(b, 100);
}

#[bench]
fn bench_check_against_1000_pii_strings(b: &mut Bencher) {
    run_payload(b, 1000);
}

#[bench]
fn bench_check_against_10000_pii_strings(b: &mut Bencher) {
    run_payload(b, 10000);
}

fn run_payload(b: &mut Bencher, size: usize) {
    let payload = r#"{
       "name": "John Doe",
       "age": 30,
       "address": {
          "street": "123 Main St",
          "cty": "Springfield",
          "state": "MO"
       },
       "somearr": [
           {
               "first_name1": "John",
               "last_name1": "Smith",
               "state": "FL"
           },
           {
               "first_name": "Jack",
               "city": "Beanstalk",
               "state1": "OR"
           }
       ],
       "customer_id": "abc123",
       "order_total": 100.00,
       "tax_total": 10.00,
       "shipping_total": 5.00,
       "grand_total": 115.00,
    }"#;

    // Create a vector to store the random strings
    let mut random_strings: HashMap<String, Keyword> = HashMap::with_capacity(10);

    // Generate 1000 random strings
    for _ in 0..size {
        let random_string: String = thread_rng()
            .sample_iter(&Alphanumeric)
            .take(10) // Adjust the length of the random string as needed
            .map(char::from)
            .collect();

        random_strings.insert(
            random_string,
            Keyword {
                score: 90,
                entity: "Address".to_string(),
            },
        );
    }

    let mut pii = FieldPII::new(random_strings);

    b.iter(|| {
        let _ = pii.scan(payload);
    });
}

#[bench]
fn bench_standard_pii(b: &mut Bencher) {
    let payload = fs::read_to_string("assets/test-payload.json").unwrap();
    let str_payload = payload.as_str();
    let keywords = get_keywords();

    b.iter(|| {
        let mut pii = FieldPII::new(keywords.clone());
        let _ = pii.scan(str_payload);
    });
}
