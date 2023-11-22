extern crate image;
extern crate imageproc;

// mod r2;

use worker::*;

pub struct SharedData {}

use futures::executor::block_on;
use futures::future;
use image::Rgba;
use imageproc::drawing::draw_text_mut;
use rusttype::{Font, Scale};

// use std::io::Cursor;

#[event(fetch)]
async fn main(req: Request, env: Env, _ctx: Context) -> Result<Response> {
    let data = SharedData {};
    let router = Router::with_data(data);

    router
        .get_async("/", |_req, ctx| async move {
            let bucket = ctx.bucket("OGP_BUCKET")?;

            let raw_base_img = bucket.get("base.png").execute();
            let raw_font = bucket.get("./assets/MPLUSRounded1c-Medium.ttf").execute();

            let (raw_base_img, raw_font) = futures::join(raw_base_img, raw_font);

            let raw_base_img_body = raw_base_img.unwrap().body().unwrap();
            let bimg = raw_base_img_body.bytes().await.unwrap();
            let raw_font_body = raw_font.unwrap().body().unwrap();
            let font = Font::try_from_vec(raw_font_body.bytes().await.unwrap()).unwrap();

            let height = 300.0;
            let scale = Scale {
                x: height,
                y: height,
            };
            let text = "すごい副業";
            let mut img = image::load_from_memory(&bimg).unwrap();
            draw_text_mut(
                &mut img,
                Rgba([255u8, 255u8, 255u8, 255u8]),
                300,
                300,
                scale,
                &font,
                text,
            );

            let mut headers = Headers::new();
            headers.set("Content-Type", "image/png")?;

            Response::from_bytes(img.as_bytes().to_vec()).map(|resp| resp.with_headers(headers))
        })
        .run(req, env)
        .await
}
