module.exports = words;

var _ = require('lodash');

function words(string) {
  var count = {};
  string = String(string).match(/\w+\b/gi,' ');
  _.each(string, function (word) {
    word = word.toLowerCase();
    if (count.hasOwnProperty(word)){
      count[word]++;
    }
    else {
      count[word] = 1;
    }

  });
  return count;
}