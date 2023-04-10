use rdkafka::consumer::{Consumer, ConsumerConfig, Message};

fn main() {
    let consumer: &str = "my-group";
    let brokers: &str = "localhost:9092";
    let topic: &str = "test";

    let consumer_config = ConsumerConfig::new()
        .set("group.id", consumer)
        .set("bootstrap.servers", brokers);

    let mut consumer = consumer_config.create::<String>().unwrap();
    consumer.subscribe(&[topic]).unwrap();

    loop {
        let message: Option<Message> = consumer.poll(None).unwrap();

        match message {
            Some(msg) => {
                let payload = match msg.payload_view::<str>() {
                    Some(Ok(s)) => s,
                    Some(Err(_)) => "",
                    None => "",
                };
                println!("Received message: {}", payload);
            }
            None => (),
        }
    }
}
