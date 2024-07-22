use axum::{extract::Path, routing::get, Router};

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/user/", get(get_users))
        .route("/user/:user_id/",
            get(get_user)
            .post(create_user)
            .delete(delete_user)
            .patch(patch_user)
        )
        .route("/post/", get(get_posts))
        .route("/post/:post_id/",
            get(get_post)
            .post(create_post)
            .delete(delete_post)
            .patch(patch_post)
        );
        let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
        axum::serve(listener, app).await.unwrap();
}

async fn get_users()-> &'static str { "get_users" }
async fn get_user(Path(user_id): Path<u32>)-> &'static str { "get_user" }
async fn create_user(Path(user_id): Path<u32>)-> &'static str { "create_user" }
async fn delete_user(Path(user_id): Path<u32>)-> &'static str { "delete_user" }
async fn patch_user(Path(user_id): Path<u32>)-> &'static str { "patch_user" }
async fn get_posts()-> &'static str { "get_posts" }
async fn get_post(Path(post_id): Path<u32>)-> &'static str { "get_post" }
async fn create_post(Path(post_id): Path<u32>)-> &'static str { "create_post" }
async fn delete_post(Path(post_id): Path<u32>)-> &'static str { "delete_post" }
async fn patch_post(Path(post_id): Path<u32>)-> &'static str { "patch_post" }
