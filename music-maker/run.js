import fetch from "node-fetch";
import fs from "fs";

async function fetchAsync(url) {
  try {
    let response = await fetch(url);
    var html = await response.text();
    return html;
  } catch (error) {
    console.log(error);
  }
}

async function main() {
  const html = await fetchAsync(
    "https://www.bensound.com/royalty-free-music/1"
  );

  //Take data and filter
  //then call getsong.sh
}

main();

/**
 * <a href="/bensound-music/bensound-roomservice.mp3" class="button p-3 is-black mt-3 has-text-weight-bold" download="" title="" style="">Download Demo</a>
 * https://www.bensound.com/bensound-music/bensound-roomservice.mp3
 <span class="has-text-weight-bold track-name">Back To The Future</span>
https://www.bensound.com/bensound-music/bensound-backtothefuture.mp3
<span class="is-block mr-3 has-text-weight-bold">Bump Up</span>
<a href="/bensound-music/bensound-bumpup.mp3" class="button p-3 is-black mt-3 has-text-weight-bold" download="" title="" style="">Download Demo</a>
*/
