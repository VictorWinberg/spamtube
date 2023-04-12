import fs from 'fs'
import request from 'request';
import isImage from 'is-image';
import * as dotenv from 'dotenv' // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import
dotenv.config()
console.log("Saving images to online image database")

const API_KEY = process.env.IMGBB_API_KEY
var base64ImageArray = []
var path = './data'
// Read files
var files = fs.readdirSync(path)
for (const file of files) {
  // Check if file is img? else continue
  const fullPath = path + '/' + file
  if (!isImage(fullPath))
    continue;

  // Read file
  var data = fs.readFileSync(fullPath)
  // Buffer to base64
  var base64String = data.toString('base64');
  // Add to array of images
  base64ImageArray.push({ base64: base64String, name: file });
}
// Upload to db 
for (const item of base64ImageArray) {

  var options = {
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

