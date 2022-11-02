import { config } from "dotenv";
import fs from "fs";
import nodeFetch from "node-fetch";
import http from "node:https";
import path from "path";
import { createApi } from "unsplash-js";

config({ path: path.resolve(".env.local") });

const unsplash = createApi({
  accessKey: process.env.UNSPLASH_ACCESS_TOKEN,
  fetch: nodeFetch,
});

function download(url, dest, cb) {
  const file = fs.createWriteStream(dest);
  http
    .get(url, function (response) {
      response.pipe(file);
      file.on("finish", function () {
        file.close(cb);
      });
    })
    .on("error", function (err) {
      fs.unlink(dest);
      if (cb) cb(err.message);
    });
}

async function main() {
  const keywords = process.env.IMAGE_INPUT.split(" ");
  const response = await Promise.all(
    keywords.map((keyword) =>
      unsplash.search.getPhotos({
        query: keyword,
        page: 1,
        perPage: 3,
        // color: "green",
        orientation: "landscape",
      })
    )
  );

  if (response.some((response) => response.errors)) {
    console.error(response.map((response) => response.errors));
    return;
  }

  const photos = response.map((response) => response.response.results).flat();
  const links = photos.map((res) => res.urls.regular);
  links.forEach((link, index) => {
    const filename = String(index).padStart(3, "0");
    download(link, `./out/${filename}.jpg`, () => {
      console.log(`image ${index} downloaded`);
    });
  });
}

main();
