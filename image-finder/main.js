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
  const response = await unsplash.search.getPhotos({
    query: process.env.IMAGE_INPUT,
    page: 1,
    perPage: 10,
    // color: "green",
    orientation: "landscape",
  });

  if(response.errors) {
    console.error(response);
    return
  }

  const { results } = response.response;
  const links = results.map((res) => res.urls.regular);
  links.forEach((link, index) => {
    const filename = String(index).padStart(3, "0");
    download(link, `./out/${filename}.jpg`, () => {
      console.log(`image ${index} downloaded`);
    });
  });
}

main();
