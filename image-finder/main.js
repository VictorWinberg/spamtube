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
const OUT_DIR = "./out";

const styles = [
  "cartoon",
  "comic book",
  "futuristic",
  "graffiti",
  "impressionism",
  "manga",
  "oil painting",
  "pencil sketch",
  "pop art",
  "surrealism",
  "watercolor",
];

const backgrounds = [
  "abstract",
  "apocalyptic",
  "beach",
  "bright neon lights",
  "city",
  "cyberpunk",
  "desert",
  "dystopia",
  "forest",
  "galaxy",
  "jungle",
  "lake",
  "mountain",
  "ocean",
  "ruins",
  "space",
  "underwater",
  "urban",
  "utopian",
  "volcano",
  "waterfall",
];

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

async function getAIImages(keywords) {
  const style = styles[Math.floor(Math.random() * styles.length)];
  const background = backgrounds[Math.floor(Math.random() * backgrounds.length)];
  const direction = `in ${style} style and ${background} background`

  const response = await nodeFetch("https://api.craiyon.com/draw", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      prompt: [...keywords, direction].join(" "),
      version: "35s5hfwn9n78gb06",
      token: null,
    }),
  });

  const json = await response.json();
  return json.images;
}

async function main() {
  const keywords = process.env.IMAGE_INPUT.split(" ");

  // remove old files
  fs.readdir(OUT_DIR, (err, files) => {
    if (err) throw err;

    for (const file of files) {
      if (file === ".gitkeep") continue;
      fs.unlink(path.join(OUT_DIR, file), (err) => {
        if (err) throw err;
      });
    }
  });

  try {
    const images = [];
    const middle = Math.floor(keywords.length / 2);

    const res = await Promise.all([
      getAIImages(keywords.slice(0, middle)),
      getAIImages(keywords.slice(middle)),
    ]);
    res.flat().forEach((image) => images.push(image));

    const urls = images.map((image) => `https://img.craiyon.com/${image}`);
    await Promise.all(
      urls.map(async (link, index) => {
        const filename = String(index).padStart(3, "0");
        const res = await download(link, `${OUT_DIR}/${filename}.webp`);
        console.log(`image ${index} downloaded`, res);
        return `${OUT_DIR}/${filename}.webp`;
      })
    );
    console.log("Done - craiyon!");
    return;
  } catch (error) {
    // continue
    console.error(error);
  }

  console.log("Falling back to unsplash");

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
      const res = await download(link, `${OUT_DIR}/${filename}.jpg`);
      console.log(`image ${index} downloaded`, res);
      return `${OUT_DIR}/${filename}.jpg`;
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
