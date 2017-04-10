class BeerSong {
  constructor() {}
  sing(count, end = 0) {
    if (count === 0) {
      return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n";
    } else if (count == end) {
      return "";
    } else {
      return this.verse(count) + `\n\n${this.sing(count - 1)}`;
    }
  }
  verse(n) {
    if (n === 0) {
      return this.sing(n);
    }
    return `${n} bottles of beer on the wall, ${n} bottles of beer.\nTake one down and pass it around, ${n - 1} bottle of beer on the wall.\n`;
  }
}

module.exports = BeerSong;
