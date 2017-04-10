class Isogram {
  constructor(phrase) {
    this.phrase = phrase.toLowerCase();
  }

  isIsogram() {
    // doesn't cover non-ascii
    const phrase = this.phrase.replace(/[^a-z0-9]+/gi, "");
    return new Set(phrase).size === phrase.length;
  }
}

module.exports = Isogram;
