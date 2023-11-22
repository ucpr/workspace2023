use worker::*;

use crate::SharedData;

pub async fn get(_req: Request, ctx: RouteContext<SharedData>) -> Result<Response> {
    let bucket = ctx.bucket("OGP_BUCKET")?;

    let item = bucket.get("base.png").execute().await?.unwrap();
    let item_body = item.body().unwrap();
    let bytes = item_body.bytes().await.unwrap();
    let mut headers = Headers::new();
    headers.set("Content-Type", "image/png")?;

    Response::from_bytes(bytes).map(|resp| resp.with_headers(headers))
}
