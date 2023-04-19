import { config } from "dotenv";
import fs from "fs";
import nodeFetch from "node-fetch";
import http from "node:https";
import path from "path";
import sharp from "sharp";
import { createApi } from "unsplash-js";
import { unlink } from "fs/promises";

config({ path: path.resolve(".env.local"), override: true });

const { IMAGE_KEYWORDS, UNSPLASH_ACCESS_TOKEN, CUSTOM_STYLE, CUSTOM_BACKGROUND } = process.env;

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
];

const unsplash = createApi({
  accessKey: UNSPLASH_ACCESS_TOKEN,
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

function* chunks(arr, n) {
  for (let i = 0; i < arr.length; i += n) {
    yield arr.slice(i, i + n);
  }
}

async function getAIImages(keywords) {
  const styleIndex = Math.floor(Math.random(new Date().getTime()) * styles.length);
  const style = CUSTOM_STYLE || styles[styleIndex];
  const backgroundIndex = Math.floor(Math.random(new Date().getTime()) * backgrounds.length);
  const background = CUSTOM_BACKGROUND || backgrounds[backgroundIndex];

  const direction = `in ${style} style and ${background} background`;
  const prompt = [...keywords, direction].join(" ");
  console.log("prompt", prompt);

  const response = await nodeFetch("https://api.craiyon.com/v3", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      prompt,
      negative_prompt: "",
      model: "art",
      version: "35s5hfwn9n78gb06",
      token: null,
    }),
  });

  try {
    const json = await response.json();
    return json.images;
  } catch (error) {
    const err = await response.text();
    throw err;
  }
}

async function main() {
  const keywords = IMAGE_KEYWORDS.split(" ");

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

    for (const chunk of [...chunks(keywords, 3)]) {
      const res = await getAIImages(chunk);
      res.forEach((image) => images.push(image));
    }

    const urls = images.map((image) => `https://img.craiyon.com/${image}`);
    const downloads = await Promise.all(
      urls.map(async (link, index) => {
        const filename = String(index).padStart(3, "0");
        const res = await download(link, `${OUT_DIR}/${filename}.webp`);
        console.log(`image ${index} downloaded`, res);
        return `${OUT_DIR}/${filename}.webp`;
      })
    );

    await Promise.all(
      downloads.map(async (filename, index) => {
        const buffer = await sharp(filename).png().toBuffer();
        await sharp(buffer).toFile(filename.replace(/webp$/, "png"));
        await unlink(filename);
        console.log(`sharp ${index}`);
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
      const buffer = await sharp(filename).resize(IMAGE_WIDTH, IMAGE_HEIGHT).png().toBuffer();
      await sharp(buffer).toFile(filename.replace(/jpg$/, "png"));
      await unlink(filename);
      console.log(`sharp ${index}`);
    })
  );

  console.log("Done - unsplash!");
}

main();
