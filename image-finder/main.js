import { config } from "dotenv";
import fs from "fs";
import nodeFetch from "node-fetch";
import http from "node:https";
import path from "path";
import sharp from "sharp";
import { createApi } from "unsplash-js";

config({ path: path.resolve(".env.local"), override: true });

const IMAGE_WIDTH = 1080;
const IMAGE_HEIGHT = 1620;

const unsplash = createApi({
  accessKey: process.env.UNSPLASH_ACCESS_TOKEN,
  fetch: nodeFetch,
});

async function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    http
      .get(url, function (response) {
        response.pipe(file);
        file.on("finish", function () {
          file.close(() => resolve("done"));
        });
      })
      .on("error", function (err) {
        fs.unlink(dest);
        reject(err.message);
      });
  });
}

async function main() {
  const keywords = process.env.IMAGE_INPUT.split(" ");

  try {
    const response = await nodeFetch("https://api.craiyon.com/draw", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        prompt: keywords.join(" "),
        version: "35s5hfwn9n78gb06",
        token: null,
      }),
    });

    const json = await response.json();
    const { images } = json;

    const urls = images.map((image) => `https://img.craiyon.com/${image}`);
    await Promise.all(
      urls.map(async (link, index) => {
        const filename = String(index).padStart(3, "0");
        const res = await download(link, `./out/${filename}.webp`);
        console.log(`image ${index} downloaded`, res);
        return `./out/${filename}.webp`;
      })
    );
    console.log("Done - craiyon!");
    return;
  } catch (error) {
    // continue
    console.error(error);
  }

  // fallback to unsplash
  const response = await Promise.all(
    keywords.map((keyword) =>
      unsplash.search.getPhotos({
        query: keyword,
        page: 1,
        perPage: 3,
        // color: "green",
        orientation: "portrait",
      })
    )
  );

  if (response.some((response) => response.errors)) {
    console.error(response.map((response) => response.errors));
    return;
  }

  const photos = response.map((response) => response.response.results).flat();
  const links = photos.map((res) => res.urls.regular);
  const downloads = await Promise.all(
    links.map(async (link, index) => {
      const filename = String(index).padStart(3, "0");
      const res = await download(link, `./out/${filename}.jpg`);
      console.log(`image ${index} downloaded`, res);
      return `./out/${filename}.jpg`;
    })
  );

  await Promise.all(
    downloads.map(async (filename, index) => {
      const buffer = await sharp(filename).resize(IMAGE_WIDTH, IMAGE_HEIGHT).toBuffer();
      await sharp(buffer).toFile(filename);
      console.log(`resize ${index}`);
    })
  );

  console.log("Done - unsplash!");
}

main();
