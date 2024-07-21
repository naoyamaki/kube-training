use axum::{routing::get, Router};
use std::net::SocketAddr;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(handler))
        .route("/hoge", get(get_hoge).post(post_hoge));
    let addr = SocketAddr::from(([127,0,0,1],3000));

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await.unwrap();
}

async fn handler()-> &'static str { "Hello, handler!" }
async fn get_hoge()-> &'static str { "Hello, get_hoge!" }
async fn post_hoge()-> &'static str { "Hello, post_hoge!" }
