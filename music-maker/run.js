import fetch from "node-fetch";
import { exit } from "process";
import jsdom from "jsdom";
import shell from "shelljs";

const OUTPUT_FOLDER = "out";
const OUTPUT_FILENAME = "audio";

function getSongNames(htmlStr, className) {
  const dom = new jsdom.JSDOM(htmlStr);

  const elements = dom.window.document.getElementsByClassName(className);
  let array = [];
  for (let index = 0; index < elements.length; index = index + 3) {
    const innerHTML = elements.item(index).innerHTML;
    array.push(innerHTML);
  }
  return array;
}

async function fetchAsync(url) {
  try {
    let response = await fetch(url);
    var html = await response.text();
    // fs.writeFile('index.html', html, function (err) {
    //   if (err) throw err;
    //   console.log('File is created successfully.');
    // });
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
  console.log("Getting a popular song from source www.bensound.com...");

  const htmlstr = await fetchAsync("https://www.bensound.com/royalty-free-music/1");

  const songNames = getSongNames(htmlstr, "is-block mr-3 has-text-weight-bold");

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
