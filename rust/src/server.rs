use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};
mod graph;


#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world!")
}

async fn graph_handler() -> impl Responder {
    let gph = graph::example_graph();
    let res_vec = gph.bfs(5);
    HttpResponse::Ok().body(format!("u: {}, v: {}",res_vec[0],res_vec[1]))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(hello)
            .route("/graph", web::get().to(graph_handler))
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}