var _ = require("lodash");

module.exports = anagram;

function anagram(word) {
  this.word = word;
  return {
    matches: function (list) {
      list = Array.isArray(list) ? list : [].slice.call(arguments, 1);
      return _.filter(list, function (candidate) {
          return check(word,candidate);
      });
    }
  };
}

function check(word,candidate) {
  //check if it is the size
  console.log(word, candidate);
  word = word.toLowerCase();
  candidate = candidate.toLowerCase();

  // if same word as self then eject
  if (word === candidate){
    return false;
  }
  var diff = _.difference(candidate, word.split(''));
  console.log(diff);
  if(word.length === candidate.length){
    if (diff.length === 0){
      return true;
    }
  }
  else {
    return false;
  }
}