import fs from 'fs'
import path from "path";
import request from 'request';
import isImage from 'is-image';
import { config } from "dotenv"; // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import

config({ path: path.resolve(".env.local"), override: true });
console.log("Saving images to online image database")

const API_KEY = process.env.IMGBB_API_KEY
let base64ImageArray = []
const folderPath = './data'
// Read files
const files = fs.readdirSync(folderPath)
for (const file of files) {
  // Check if file is img? else continue
  const fullPath = folderPath + '/' + file
  if (!isImage(fullPath))
    continue;

  // Read file
  const data = fs.readFileSync(fullPath)
  // Buffer to base64
  const base64String = data.toString('base64');
  // Add to array of images
  base64ImageArray.push({ base64: base64String, name: file });
}
// Upload to db 
for (const item of base64ImageArray) {

  let options = {
    'method': 'POST',
    'url': `https://api.imgbb.com/1/upload?name=${item.name}&key=${API_KEY}`,
    'headers': {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    formData: {
      'image': item.base64
    }
  };
  request(options, function (error, response) {
    if (error) throw new Error(error);
    console.log(response.body);
  });
}
