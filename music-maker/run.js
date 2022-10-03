import fetch from "node-fetch";
import { exit } from "process";
import jsdom from "jsdom";
import shell from "shelljs";

const OUTPUT_FOLDER = "out";
const OUTPUT_FILENAME = "audio";

function getSongNames(htmlStr, className) {
  const dom = new jsdom.JSDOM(htmlStr);

  const elements = dom.window.document.querySelectorAll(className);
  let array = [];
  elements.forEach(element => {
    array.push(element.textContent);
  })
  return array;
}

async function fetchAsync(url) {
  try {
    let response = await fetch(url);
    var html = await response.text();
    return html;
  } catch (error) {
    console.log(error);
  }
}

async function downloadSong(songName) {
  const url = `https://www.bensound.com/bensound-music/${songName}`;
  console.log(`Downloading mp3 file from url: ${url}`);
  try {
    shell.exec(`curl -X GET ${url} > ${OUTPUT_FOLDER}/${OUTPUT_FILENAME}.mp3`);
    return true;
  } catch (error) {
    console.log(error);
    return false;
  }
}

function randomIntFromInterval(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

async function main() {
  const page = randomIntFromInterval(1, 25);
  
  console.log(`Getting a popular song from https://www.bensound.com/royalty-free-music/${page}...`);
  const htmlstr = await fetchAsync(`https://www.bensound.com/royalty-free-music/${page}`);
  const songNames = getSongNames(htmlstr, ".result-container > div:first-child > div > .track-name a span");

  console.log("Pool of songs:", { array: songNames });
  const rnd = randomIntFromInterval(0, songNames.length - 1);
  console.log(`Randomly chosen song: ${songNames[rnd]}`);

  const formatSongName = songNames[rnd].toLocaleLowerCase().replace(/[^a-zA-Z]| /g, "");
  const formatName = `bensound-${formatSongName}.mp3`;
  try {
    const downloaded = await downloadSong(formatName);
    if (!downloaded) {
      console.log("Failed to download song");
      exit(1);
    }
  } catch (error) {
    console.log(error);
    exit(1);
  }

  console.log("Done");
}

main();
